package methods

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func HelloWorld(name string) (string, error) {
	name = strings.Trim(name, " ")
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("hello, %s!", name), nil
}
