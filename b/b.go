package b

import (
	"fmt"

	"github.com/dmitryax/prj1/a"
)

func PublicB1() {
	fmt.Println("Called PublicB1")
	a.Public1()
}
