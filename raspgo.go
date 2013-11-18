package main


import ( 
    "html/template"
	"fmt"
	"net/http"
	"time"

)
import "github.com/teixeiras/raspGo/Modules/raspiConfig"
import "github.com/teixeiras/raspGo/Modules/osRequests"



func log(message string) {
	t := time.Now();
	fmt.Printf("%s: %s\n", t, message);

}
type ModuleStruct interface {

}

type Module struct {
    Name string
    ModuleObject ModuleStruct
}

type Page struct {
    Modules []Module
    Title string
    Url string
}

var modules []Module;
    
func initModules() {
    modules = append(modules, Module{Name:"OS Operations", ModuleObject:nil});
    modules = append(modules, Module{Name:"Raspberry Operations", ModuleObject:nil});
    modules = append(modules, Module{Name:"File Manager Opearions", ModuleObject:nil});

}
func index(responseWriter http.ResponseWriter, request *http.Request) {
    title := request.URL.Path[len("/"):];
    page := &Page{Modules : modules ,Title : title, Url : request.Host};

    template, _ := template.ParseFiles("template/index.html");
    template.Execute(responseWriter, page)};



func main() {
    initModules();
    osRequests.Openproc();
    _ = raspiConfig.Expand_file_system();
	log("Server has started");
	http.HandleFunc("/", index);
    http.HandleFunc("/ws", wsHandler)
    http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("template/"))))

	http.ListenAndServe(":9999", nil);
}
