package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/abatilo/contract-ai/pkg/api/v1/nlp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
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
	log.Printf("Connecting to nlpservice...")
	conn, err := grpc.Dial("dns:///nlpservice:9090", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		questions := []*nlp.Question{
			{
				Body: "Who are the tenants?",
			},
			{
				Body: "When was the agreement made?",
			},
			{
				Body: "Where is the landlord located?",
			},
			{
				Body: "What city is the landlord in?",
			},
			{
				Body: "What state is the landlord in?",
			},
			{
				Body: "What kind of agreement is this?",
			},
		}

		answersCh := make(chan *nlp.AskQuestionResponse, len(questions))

		for _, question := range questions {
			c := nlp.NewNLPServiceClient(conn)
			go func(q *nlp.Question, ch chan<- *nlp.AskQuestionResponse) {
				resp, err := c.AskQuestion(ctx, &nlp.AskQuestionRequest{
					Corpus:   &nlp.Corpus{Body: corpus},
					Question: q,
				})
				if err != nil {
					log.Printf("Couldn't AskQuestions: %+v", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				ch <- resp
			}(question, answersCh)
		}

		answers := []*nlp.AskQuestionResponse{}
		for i := 0; i < len(questions); i++ {
			answer := <-answersCh
			answers = append(answers, answer)
		}
		close(answersCh)

		answersBytes, err := json.Marshal(answers)
		if err != nil {
			log.Printf("Couldn't marshal answers slice: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(answersBytes)
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

const corpus = `I. THE PARTIES. This Residential Lease Agreement (“Agreement”) made this
December 27th, 2020 is between:
Landlord: ABC Landlord Inc. with a mailing address of 123 Main Street, City of Santa
Monica, State of California ("Landlord"), AND
Tenant(s): Jeffrey and Martha Johnson (“Tenant”).
Landlord and Tenant are each referred to herein as a “Party” and, collectively, as the
"Parties."
NOW, THEREFORE, FOR AND IN CONSIDERATION of the mutual promises and
agreements contained herein, the Tenant agrees to lease the Premises from the
Landlord under the following terms and conditions:`
