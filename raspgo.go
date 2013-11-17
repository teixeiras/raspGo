package main


import (
    "html/template"
	"fmt"
	"net/http"
	"time"

)
import "github.com/teixeiras/raspGo/Modules/raspiConfig"
import "github.com/teixeiras/raspGo/Modules/readproc"



func log(message string) {
	t := time.Now();
	fmt.Printf("%s: %s\n", t, message);

}

type Page struct {
    Title string
    Body  []byte
}


func index(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/"):];
    p := &Page{Title: title};
    
    t, _ := template.ParseFiles("index.html");
    t.Execute(w, p)};



func main() {
    readproc.Openproc();
    _ = raspiConfig.Expand_file_system();
	log("Server has started");
	http.HandleFunc("/", index);

	http.ListenAndServe(":9999", nil);
}
