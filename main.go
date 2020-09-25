package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"bytes"
	admission "k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func printJson(data []byte) {
	var out bytes.Buffer
	var js string
	err := json.Indent(&out, []byte(data), "", "  ")
	if err != nil {
		log.Printf("json decode error: %s\n", err.Error())
		js = string(data)
	} else {
		js = out.String()
	}
	log.Println(js)
}

type Handler struct {}
func (h* Handler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request %s\n", r.RequestURI)
	data, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Println(err)
	} else {
		printJson(data)
		var review admission.AdmissionReview
		err = json.Unmarshal(data, &review)
		if err != nil {
			log.Println(err)
		}
		log.Println(review.Request.UID)

		resp := admission.AdmissionReview{
			TypeMeta: metav1.TypeMeta {
				Kind: "AdmissionReview",
				APIVersion: "admission.k8s.io/v1" },
			Response: &admission.AdmissionResponse {
				UID: review.Request.UID,
				Allowed: true,
			},
		}
		respb, err := json.Marshal(&resp)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(respb))
		w.Write(respb)
	}
}

func main() {
	log.Println("start")

	server := &http.Server{
		Addr:    ":443",
		Handler: &Handler{},
	}
	err := server.ListenAndServeTLS("/hook.crt", "/hook.key")
	if err != nil {
		panic(err)
	}
}
