package console

import (
	"bufio"
	"os"
	"strings"
	"github.com/KaymeKaydex/algorithms-console-io.git/conversion"
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
func (Console) ReadIntArray() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	slice := strings.Fields(scanner.Text())
	return SliceAtoi(slice)
}