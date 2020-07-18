package interface_test

import "testing"

// 定义接口
type Programmer interface {
	WriteHelloWorld() string
}

// 定义接口的实现
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
