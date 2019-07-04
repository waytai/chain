package chain

import ("crypto/sha256"
		"fmt"
		"strings"
	)


func main() {
	for i := 0; i < 100; i++ {
		data := "test@example.com"+string(i)
		c := getSha256Code(data)
		var index = strings.Index(c, "000")
		if index == 0 {
			fmt.Println(c,i)
			break;
		}
	}
}


func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	data := h.Sum(nil)
	hex :=fmt.Sprintf("%x", data)
	return hex
}
