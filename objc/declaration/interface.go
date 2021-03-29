package declaration

import (
	"fmt"
	"regexp"
)

type Interface struct {
	name string
	superName string
}

func (i Interface) IsNode() {}

func (i Interface) Name() string {
	return i.name
}

func (i Interface) SuperName() string {
	return i.superName
}

func parseInterface(input string) (Node,error)  {
	re, err := regexp.Compile(`^@interface\s(?P<name>\w+)(?:\s:\s(?P<super>\w+))?$`)
	if err != nil {
		return nil, err
	}
	match := re.FindStringSubmatch(input)
	if len(match) == 0 {
		return nil, fmt.Errorf("not an interface")
	}
	i := Interface{
		name: match[re.SubexpIndex("name")],
		superName: match[re.SubexpIndex("super")],
	}
	return i, nil
}