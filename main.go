package main

import (
	"fmt"
	"log"
	"net/http"
	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/bitbucket"
)

func main(){
	hook := bitbucket.New(&bitbucket.Config{
		UUID : "none",	
	})
	http.HandleFunc("/webhook", func(rw http.ResponseWriter, r *http.Request) {
		hook.ParsePayload(rw, r)
		hook.RegisterEvents(func(payload interface{}, header webhooks.Header) {
			fmt.Fprintf(rw, "pyload: %v, header: %v",payload, header)
		}, bitbucket.RepoPushEvent)
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}