package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化默认为默认零值""

	s = "hello"
	t.Log(len(s))
	// s[1] = '3'         // string是不可变的byte slice

	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	// s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s)) // 是byte数

	s = "中"
	t.Log(len(s)) // 是byte数
	// t.Logf("%x", s[0])
	// t.Logf("%x", s[1])
	// t.Logf("%x", s[2])
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0]) // 编码
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	b := []rune(s)
	t.Logf("%x", b[0])
	for _, c := range s { // range和字符串配合，迭代输出的是rune而不是byte
		t.Logf("%[1]c %[1]x", c)
	}
}
