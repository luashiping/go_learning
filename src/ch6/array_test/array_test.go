package array_test

import "testing"

// 数组初始化
func TestAarrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

// 数组遍历
func TestAarryTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestAarraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	t.Log(arr3[3])
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)

	// 最后面两个元素
	arr3_sec2 := arr3[len(arr3)-2:]
	t.Log(arr3_sec2)

	// 倒数第一个元素
	arr3_sec3 := arr3[len(arr3)-1:]
	t.Log(arr3_sec3)
}
