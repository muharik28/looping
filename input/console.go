package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Console struct {
	reader *bufio.Reader
}

func NewConsole() *Console {

	return &Console{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *Console) ReadInt(message string) (int, error) {

	fmt.Print(message)

	val, err := c.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	val = strings.TrimSpace(val)

	return strconv.Atoi(val)
}
