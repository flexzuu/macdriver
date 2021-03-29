package declaration

import "fmt"

type Node interface {
	IsNode()
}

func Parse(input string) (Node, error)  {
	res, err := parseInterface(input)
	if err == nil {
		return res, nil
	}
	res, err = parseMethod(input)
	if err == nil {
		return res, nil
	}
	res, err = parseProperty(input)
	if err == nil {
		return res, nil
	}
	return nil, fmt.Errorf("could not parse")
}