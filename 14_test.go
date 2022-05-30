package l1_test

import (
	"log"
	"reflect"
	"testing"
)

// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel из
// переменной типа interface{}.

func checkType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

func checkType2(v interface{}) string {
	switch v.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case float64:
		return "float64"
	case []int:
		return "int slice"
	case nil:
		return "nil"
	case bool:
		return "bool"
	case chan interface{}:
		return "channel"
	default:
		return "unknown"
	}
}

func Test_checkType(t *testing.T) {
	log.Println(checkType(1))
	log.Println(checkType("asd"))
	log.Println(checkType(nil))
	log.Println(checkType([]int{1, 2, 3}))
	log.Println(checkType(make(chan interface{})))

	log.Println(checkType2(1))
	log.Println(checkType2("asd"))
	log.Println(checkType2(nil))
	log.Println(checkType2([]int{1, 2, 3}))
	log.Println(checkType2(make(chan interface{})))

}
