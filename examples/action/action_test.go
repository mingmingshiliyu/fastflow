package action

import (
	"reflect"
	"testing"
)

type Action interface {
	Name() string
}

type ActionA struct {
}

func (a ActionA) Name() string {
	//TODO implement me
	panic("implement me")
}

type ActionB struct {
}

func (a ActionB) Name() string {
	//TODO implement me
	panic("implement me")
}

var a map[string]Action
var structMap map[string]reflect.Type

func TestPut(t *testing.T) {
	b := "actionA"
	registerStruct("actionA", ActionA{}, structMap)
	if b == "actionA" {

	}

}

// 注册结构体到map中
func registerStruct(name string, instance interface{}, structMap map[string]reflect.Type) {
	structType := reflect.TypeOf(instance)
	structMap[name] = structType
}
