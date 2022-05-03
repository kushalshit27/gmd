package pkg

import (
	"fmt"
	"os"
)

func ExitGDM() {
	fmt.Println("Bye Bye")
	os.Exit(3)
}
