package six910api

import (
	"fmt"
	"testing"
)

func TestHeaders_DeepCopy(t *testing.T) {
	var hd Headers
	hd.Set("test1", "tone")
	hd.Set("test2", "ttwo")

	cp := hd.DeepCopy()

	var tf = false
	var i = 0
	for k, v := range cp.headers {
		fmt.Print("k = ", k)
		fmt.Println(" v = ", v)
		if i == 0 && (k != "test1" || v != "tone") {
			tf = true
		} else if i == 1 && (k != "test2" || v != "ttwo") {
			tf = true
		}
		i++
	}
	fmt.Println("tf: ", tf)
	if tf {
		t.Fail()
	}
}
