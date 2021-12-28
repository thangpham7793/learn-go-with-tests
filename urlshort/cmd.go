package urlshort

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

var PORT = 8080

func Run() {
	yamlPath := flag.String("yml", "", "path to yaml file with url mapping")
	flag.Parse()

	content, err := ioutil.ReadFile(*yamlPath)
	if err != nil {
		log.Fatal(err)
	}

	urlMap, err := newUrlMapper(content)
	if err != nil {
		log.Fatal(err)
	}

	handler := YAMLHandler(urlMap, http.HandlerFunc(DefaultHandler))

	log.Printf("Starting server on port %d", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), handler)
}

func newUrlMapper(yml []byte) (map[string]string, error) {
	content := []map[string]string{}
	err := yaml.Unmarshal(yml, &content)

	if err != nil {
		return nil, err
	}

	urlMap := make(map[string]string)
	for _, m := range content {
		urlMap[m["path"]] = m["url"]
	}

	return urlMap, nil
}

func YAMLHandler(urlMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if redirect, ok := urlMap[r.URL.Path]; ok {
			http.Redirect(w, r, redirect, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Please use a different shortened url"))
}
