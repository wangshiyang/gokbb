package main

import (
	"net/http"
	"fmt"
	"shawn/gokbb/common/setting"
	"shawn/gokbb/routers"
	"shawn/gokbb/common/logging"
	"os"
	"os/signal"
	"context"
	"time"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logging.Info("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("Shutdown Server......")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logging.Error("Server Shutdown error:", err)
	}

	logging.Info("Server exiting")
}
