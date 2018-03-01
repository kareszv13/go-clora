package wrapper

/*
#include "test_loragw_spi.h"
#cgo LDFLAGS: -L. -llib
*/
import "C"

func main() {
	C.test()
}
