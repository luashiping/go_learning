package flexible_reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	// t.Log(a == b) // map只能和nil比较
	t.Log("a==b?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	// t.Log(s1 == s2) // slice只能和nil比较
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {

	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.

	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct types.")
	}
	// Elem() 获取指针指向的值
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct typed.")
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}

	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	// settings := map[string]interface{}{}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
