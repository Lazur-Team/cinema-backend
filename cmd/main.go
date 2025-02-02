package main

import (
	"cinema/internal/config"
	"log/slog"
	"os"
)

func main() {
	// define config
	cfg := config.MustLoad()

	// define logger
	log := setupLogger(cfg.AppEnv)
	log.Info("Starting server", slog.Any("cfg", cfg.AppUrl))

	// define postgres
	// storage := pgsql.New(cfg)

	// define router
	// routes := router.New(storage, log)

	// run server
	// server := http.Server{
	// 	Addr:    cfg.AppUrl + ":" + cfg.AppPort,
	// 	Handler: routes,
	// }
	// if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 	log.Error("failed to start server", err)
	// }

	// log.Error("Server stopped")

	// // Graceful shutdown
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// sig := <-quit
	// log.Info("Shutting down server...", slog.Any("Signal", sig))

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Error("Server forced to shutdown", slog.Any("error", err))
	// } else {
	// 	log.Info("Server gracefully stopped")
	// }
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}
