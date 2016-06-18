package goptv

import (
	"fmt"
	"testing"
)

func Test_IntV(t *testing.T) {
	var intp = 34
	fmt.Println(IntV(&intp) + 1)

}
func Test_StringV(t *testing.T) {
	var strp = "test"
	fmt.Println(StringV(&strp))
}
func Test_BufV(t *testing.T) {
	var intp = 34
	fmt.Println(BufV(&intp))
	var strp = "test"
	fmt.Println(BufV(&strp))
}
