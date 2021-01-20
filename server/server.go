package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/database"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/routers"
)

var (
	runserver = flag.Bool("runserver", false, "Run the server")
	migrate   = flag.Bool("migrate", false, "Migrate the database and tables")
)

func main() {
	flag.Parse()

	// initiate and run server based on cmd
	if *runserver {
		Server()
	} else if *migrate {
		//migrat the tables
		Migrate()
	} else {
		fmt.Println("Please provide second cmd")
	}
}

//
// Migrate the models
//
func Migrate() {
	database.MigrateModels()
}

//
// Initiate the gin server
//
func Server() {
	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)

	//init router
	// Force log's color
	gin.ForceConsoleColor()

	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	//
	// Gin Helmet
	// Security middlewares for Gin (gin-gonic/gin) inspired by the popular helmet middleware package for Node JS express and koa.

	// Default returns a number of handlers that are advised to use for basic HTTP(s) protection
	// NoSniff(), DNSPrefetchControl(), FrameGuard(), SetHSTS(true), IENoOpen(), XSSFilter()
	//
	router.Use(helmet.Default())

	//
	// Add routers
	//
	//
	routers.Routes(router) // add routters from auth

	//
	//
	// Use middlewares
	//

	srv := &http.Server{
		Addr:    ":9091",
		Handler: router,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Println("Server listening on port ", 9091)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
