package server

import (
	"avito-task/internal/config"
	"avito-task/internal/handler"
	"avito-task/internal/repository"
	"avito-task/internal/service"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg *config.Config
	// db     *sql.DB
	router *gin.Engine
}

func NewApp(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Init() {
	gin.SetMode(a.cfg.Mode)
	a.router = gin.New()
	a.router.Use(gin.Recovery())

	a.setComponents() // register router
}

// Run runs application.
func (a *App) Run(ctx context.Context) {
	// Config custom server, run , wait graceful shutdown
	server := &http.Server{
		Addr:           a.cfg.Port,
		Handler:        a.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Starting web server on port: ", a.cfg.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Application forced to shutdown: ", err)
	}
	log.Println("Application exiting")

}

// SetComponents sets components.
func (a *App) setComponents() {

	apiVersion := a.router.Group("/v1")

	db, err := repository.NewPostgres(a.cfg.DBConf)
	if err != nil {
		log.Println(err)
		return
	}

	err = repository.CreateTables(db)
	if err != nil {
		log.Println(err)
		return
	}

	accountRepo := repository.NewAccountRepo(db)
	currencyRepo := repository.NewCurrencyRepo(db)

	userRepo := repository.NewUserService(db)

	accountService := service.NewAccountService(accountRepo, currencyRepo)
	userService := service.NewUserService(userRepo)

	currencyService := service.NewCurrencyService(currencyRepo)

	handler.SetEndpoints(apiVersion, accountService, userService, currencyService)
}
