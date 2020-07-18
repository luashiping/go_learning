package polymorphism

import (
	"fmt"
	"testing"
)

// 自定义code类型
type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphsim(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgrammer(goProg)
	writeFirstProgrammer(javaProg)
}
