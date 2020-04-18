package api

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	// Cmd is the exported cobra command which starts the webhook handler service
	Cmd = &cobra.Command{
		Use:   "api",
		Short: "Runs the web service",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Testing"))
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Printf("Starting server...")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Server did not shut down cleanly: %+v", err)
	}
	log.Printf("Shutting down server...")
}
