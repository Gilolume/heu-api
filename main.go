package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// [START imports]
	"github.com/gorilla/mux"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	// [END imports])
)

func main() {
	log.Println("Starting heu-api")
	readConfigFile(configFile)
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET", "POST")
	r.HandleFunc("/test", test).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":80", r)

}

func index(w http.ResponseWriter, r *http.Request) {
	s := IndexResponse{}

	switch r.Method {
	case "GET":
		s = IndexResponse{Success: 1, Message: "Hello ! Index test with GET method"}
	case "POST":
		s = IndexResponse{Success: 1, Message: "Hello ! Index test with POST method"}
	default:
		s = IndexResponse{Success: 1, Message: "Hello ! Index test with not allowed method"}
	}

	response, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(response)
}

func test(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	filename := "./discours_macron.mp3"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Send the contents of the audio file with the encoding and
	// and sample rate information to be transcripted.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "fr-FR",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	// Print the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
}
