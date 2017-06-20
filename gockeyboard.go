package gockeyboard

/*
#cgo CFLAGS: -I./Lib
#cgo LDFLAGS: -v -L. -lkeyboard
#include <keyboard.h>
*/
import "C"
import (
	"fmt"
)

// StartGocKeyboard initialize the ncurses keyboard through C pseudo import
func StartGocKeyboard() {

	C.InitKeyboard()

	fmt.Printf("\nEnter: (Press 'q' to quit)")
	for {
		key := C.GetCharacter()
		fmt.Printf("%c", key)

		if key == 'q' {
			fmt.Printf("Bye!")
			break
		}
	}

	C.CloseKeyboard()
}
