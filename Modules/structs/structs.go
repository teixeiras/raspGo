package structs

type GenericModuleStruct interface {
    GetName() string
    ListPublicOptions() []string
    IsVisible() bool
}

type Module struct {
    Id string
    ModuleObject GenericModuleStruct
}

type Request struct {
    Action string
    Arguments map[string]interface{}
}

type Page struct {
    Modules []Module
    Title string
    Url string
}



