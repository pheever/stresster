package module

import "reflect"

//IdentifiedTypes are filled at runtime after all the plugins are loaded
var IdentifiedTypes = map[string]reflect.Type{}

//LoadAll modules from directory
func LoadAll(directory string) {

}

//LoadSingle module
func LoadSingle(filename string) {

}
