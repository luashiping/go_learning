package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

// 结构体定义
type Employee struct {
	Id   string
	Name string
	Age  int
}

// 通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

// func (e Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回引用/指针，相当于e:=&Employee{}
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e) // 加上取值符等于下面的语句, 变成指针类型
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	// e := &Employee{"0", "Bob", 20}
	// t.Log(e.String())

	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())

}
