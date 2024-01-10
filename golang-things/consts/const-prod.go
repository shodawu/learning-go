// +build prod

package consts

import "fmt"

func init() {
	Ver = "Production"
	fmt.Println(Ver)
}
