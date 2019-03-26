package testmod

import (
	"errors"
	"fmt"
)

func Hello(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!", name), nil
	default:
		return "", errors.New("unknow language")
	}
}
