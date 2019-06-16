package main
import (
	"fmt"
)
func names() (one string, two string) {
	one = "one"
	two = "two"
	return
}
/* names() can also be defined as
func names() (string, string) {
	return "one","two"
}
 */
func main() {
	notone, nottwo := names()
	fmt.Println(notone,nottwo)
}