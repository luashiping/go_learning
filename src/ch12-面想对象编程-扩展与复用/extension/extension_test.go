package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	// p *Pet
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

// func (d *Dog) Speak() {
// 	fmt.Println("Wang!")
// 	// d.p.Speak()
// }

// func (d *Dog) SpeakTo(host string) {
// 	// d.p.SpeakTo(host)
// 	d.Speak()
// 	fmt.Println(host)
// }

func TestDog(t *testing.T) {
	dog := new(Dog)
	// dog.Speak()
	dog.SpeakTo("Chao")

}
