package main


import ( 
    "html/template"
	"net/http"
    "log"

)
import "github.com/teixeiras/raspGo/Modules/raspiConfig"
import "github.com/teixeiras/raspGo/Modules/osRequests"
import "github.com/teixeiras/raspGo/Modules/fileManager"
import "github.com/teixeiras/raspGo/Modules/structs"


var modules []structs.Module;
func getModule(id string) structs.GenericModuleStruct {
    for _, moduleObj := range modules {
        if (moduleObj.Id == id) {
            return moduleObj.ModuleObject;
        }
    }
    return nil;
}

func moduleExist(id string) bool {
    for _, moduleObj := range modules {
        if (moduleObj.Id == id) {
            return true;
        }
    }
    return false;
}

func initModules() {
    modules = append(modules, structs.Module{Id: "os.operations", ModuleObject: nil});
    modules = append(modules, structs.Module{Id: "raspberry.operations", ModuleObject: nil});
    
    var manager fileManager.Module;
    modules = append(modules, structs.Module{Id: "file.manager", ModuleObject: manager});

}

func index(responseWriter http.ResponseWriter, request *http.Request) {
    title := request.URL.Path[len("/"):];
    page := &structs.Page{Modules : modules ,Title : title, Url : request.Host};

    template, _ := template.ParseFiles("template/index.html");
    template.Execute(responseWriter, page)
}



func main() {
    initModules();
    osRequests.Openproc();
    _ = raspiConfig.Expand_file_system();
	log.Println("Server has started");
	http.HandleFunc("/", index);
    http.HandleFunc("/ws", wsHandler)
    http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("template/"))))

	http.ListenAndServe(":9999", nil);
}
