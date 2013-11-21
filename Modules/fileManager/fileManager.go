package fileManager    

import "github.com/teixeiras/raspGo/Modules/structs"

import (
    "fmt"
    "path/filepath"
)

type Module struct{
	structs.GenericModuleStruct
}

func (module  Module) GetName() string {
	return "File Manager";
}

func (module  Module) IsVisible() bool {
	return true;
}

func (module  Module) ListPublicOptions() []string  {
	return []string{"File Manager"};
}

func listFile() {
    files, _ := filepath.Glob("*")
    fmt.Println(files) // contains a list of all files in the current directory
}