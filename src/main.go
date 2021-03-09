package main

import (
	hr "github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"os"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type View struct {
	db *gorm.DB
}

type Router struct {
	view *View
	hr *hr.Router
}

func main() {
	// config
	wd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	t_file, _ := ioutil.ReadFile(filepath.Join(wd,"./static/config.json"))
	var inte_data map[string]interface{}
	if err := json.Unmarshal(t_file, &inte_data); err != nil {
		log.Fatal(err)
	}
	// connect db
	db, _ := gorm.Open("sqlite3", filepath.Join(wd,inte_data["db_path"].(string)))
	defer db.Close()
	// assign router
	hr_r := hr.New()
	view := &View{db:db}
	view.init()
	router := Router{view:view,hr:hr_r}
	router.init()
	// static files
	//http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	hr_r.ServeFiles("/static/*filepath", http.Dir(filepath.Join(wd,"./static")))
	log.Fatal(http.ListenAndServe(":"+inte_data["port"].(string), hr_r))
}
