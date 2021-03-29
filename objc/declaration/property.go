package declaration

import (
	"fmt"
	"regexp"
	"strings"
)

type Property struct {
	name string
	propertyType string
}

func (p Property) IsNode() {}

func (p Property) Name() string {
	return strings.TrimPrefix(p.name, "*")
}

func (p Property) Type() string {
	return p.propertyType
}

func (p Property) IsPointer() bool {
	return strings.HasPrefix(p.name, "*")
}


func parseProperty(input string) (Node,error)  {
	re, err := regexp.Compile(`^@property(?:\(.*\))?\s(?P<propertyType>\w+)\s(?P<name>\*?\w+);`)
	if err != nil {
		return nil, err
	}
	match := re.FindStringSubmatch(input)
	if len(match) == 0 {
		return nil, fmt.Errorf("not an interface")
	}
	i := Property{
		propertyType: match[re.SubexpIndex("propertyType")],
		name: match[re.SubexpIndex("name")],
	}
	return i, nil
}