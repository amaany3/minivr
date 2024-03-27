package publisherserver

import (
	"net/http"
	"os"

	handler "github.com/amaany3/minivr/server/internal/handler/publisher"
	"github.com/amaany3/minivr/server/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "publisher-server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		handler.NewHealthHandler(r)
	})

	logger.Notice("starting publisher server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("failed to start publisher server: %v", err)
		os.Exit(1)
	}
}
