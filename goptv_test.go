package goptv

import (
	"fmt"
	"testing"
	"time"
)

func Test_IntV(t *testing.T) {
	var intp = 34
	fmt.Println(IntV(&intp))

}
func Test_StringV(t *testing.T) {
	var strp = "test"
	fmt.Println(StringV(&strp))
}
func Test_TimeV(t *testing.T) {
	var tt = time.Now().UTC()
	fmt.Println(TimeV(&tt))
}
func Test_BufV(t *testing.T) {
	var intp = 34
	fmt.Println(BufV(&intp))
	var strp = "test"
	fmt.Println(BufV(&strp))
}
