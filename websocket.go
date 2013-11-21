package main
import "github.com/teixeiras/raspGo/Modules/structs"
import (
	"github.com/gorilla/websocket"
	"net/http"
    "encoding/json"
    "log"
    "reflect"
)

type Connection struct {
    // The websocket connection.
    ws *websocket.Conn

    // Buffered channel of outbound messages.
    send chan []byte
}

func Call(params ... interface{}) (result []reflect.Value, err error) {
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = in
    return
}

func route (response map[string]interface{}) {
    moduleName, ok := response["module"].(string);
    if !ok {
        return;
    }
    actionName, ok := response["action"].(string)
    if !ok {
        return;
    }

    log.Println("Route search");
    log.Printf("Module: %s", moduleName);
    log.Printf("Action: %s", actionName);

    if moduleExist(moduleName) == true {
        
        log.Println("Module Found");

        var module structs.GenericModuleStruct = getModule(moduleName);
        
        if reflect.ValueOf(&module).MethodByName(actionName).IsValid() {
        
            log.Println("Method Found");
            
            args, _ := response["args"].(map[string]interface{});
            value,_ := Call(args);
            reflect.ValueOf(&module).MethodByName(actionName).Call(value);   
        }
    }   
}


func (c * Connection) reader() {
    for {
        _, message, err := c.ws.ReadMessage();
        log.Printf("Message received: %s", message);
        if err != nil {
            break;
        }
        var objmap map[string]interface{};
        err = json.Unmarshal(message, &objmap);
        route(objmap);
    }
    c.ws.Close();
}

func (c * Connection) sendMessage(module structs.Module, request structs.Request) (bool) {
    x := make(map[string]interface{})
    x["module"] = module.Id;
    x["action"] = request.Action;
    x["args"] = request.Arguments;
    c.ws.WriteJSON(x);
    return true;
}
func wsHandler(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

    if _, ok := err.(websocket.HandshakeError); ok {
        http.Error(w, "Not a websocket handshake", 400);
        return
    } else if err != nil {
        http.Error(w, "Not a websocket handshake", 400);
        return
    }

    c := &Connection{send: make(chan []byte, 256), ws: ws}

    defer func() { ws.Close() }()
    c.reader()
}

