package main
import (
	"github.com/ohisama/key"
	"time"
)
func main() {
	var c byte;
	for {
		time.Sleep(200 * time.Millisecond)
		c = key.Scan()
		if c > 0 {
			print(string(c))
		}
	}
}
