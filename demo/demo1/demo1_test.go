package demo1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDemo1(t *testing.T) {
	container := 42

	containerType := reflect.TypeOf(container)
	fmt.Println(containerType)

	containerValue := reflect.ValueOf(container)
	fmt.Println(containerValue)
}
