package console

import (
	"bufio"
	conv "github.com/KaymeKaydex/algorithms-basic-functions.git/conversion"
	"os"
	"strings"
)

func (Console) ReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
func (Console) ReadStringArray() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	slice := strings.Fields(scanner.Text())
	return slice
}
func (c *Console) ReadIntArray() ([]int, error) {
	return conv.SliceAtoi(c.ReadStringArray())
}
