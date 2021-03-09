package main

import (
	"net/http"
	"encoding/json"
	"os/exec"
	"bytes"
	"log"
)

type Route struct {
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (ro *Router) init() {
	routes := Routes{
		{"GET","/add_article/",ro.add_article},
		{"POST","/add_sentence/",ro.add_sentence},
		{"GET","/add_link/",ro.add_link},
		{"GET","/get_title/",ro.get_title},
		{"GET","/get_articles/",ro.get_articles},
		{"GET","/id2cate/",ro.id2cate},
	}
	for _, route := range routes {
		ro.hr.HandlerFunc(route.Method,route.Pattern,route.HandlerFunc)
	}
}
func (ro *Router) add_article(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	obj := ro.view.add_article(args["doi"][0],args["title"][0],args["category"][0])
	respondWithJSON(w,http.StatusOK,obj)
}
func (ro *Router) add_sentence(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	obj := ro.view.add_sentence(r.Form["art_id"][0],r.Form["syndrome"][0],r.Form["part"][0])
	respondWithJSON(w,http.StatusOK,obj)
}
func (ro *Router) add_link(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	obj := ro.view.add_link(args["sente_id"][0],args["bacteria_name"][0],args["trend_mark"][0])
	respondWithJSON(w,http.StatusOK,obj)
}
func (ro *Router) get_title(w http.ResponseWriter, r *http.Request) {
	//args := r.URL.Query()
	cmd := exec.Command("curl","-LH","\"Accept:application/json\"","http://dx.doi.org/10.1038/nrd842")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		log.Fatal(err)
	}
	//obj := ro.view.add_link(args["sente_id"][0],args["bacteria_name"][0],args["trend_mark"][0])
	respondWithJSON(w,http.StatusOK,out)
}
func (ro *Router) get_articles(w http.ResponseWriter, r *http.Request) {
	//args := r.URL.Query()
	obj := ro.view.get_articles(Article{})
	respondWithJSON(w,http.StatusOK,obj)
}
func (ro *Router) id2cate(w http.ResponseWriter, r *http.Request) {
	obj := ro.view.id2cate()
	respondWithJSON(w,http.StatusOK,obj)
}
/*
func (ro *Router) get(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	emails := ro.view.get_email(args["name"][0])
	respondWithJSON(w,http.StatusOK,emails)
}

func (ro *Router) post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := ro.view.add_email(r.Form["name"][0],r.Form["email"][0])
	respondWithJSON(w,http.StatusOK,u)
}
*/

func respondWithJSON(w http.ResponseWriter,code int, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // allow cross access
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
