package declaration

import (
	"fmt"
	"regexp"
	"strings"
)

type Method struct {
	baseName string
	returnType string
	args []MethodArgs
}

func (m Method) IsNode() {}

func (m Method) Name() string {
	name := strings.Builder{}
	name.WriteString(m.baseName)
	for _, arg := range m.args {
		name.WriteString(arg.methodNameAddon)
		name.WriteString(":")
	}
	return name.String()
}

func (m Method) ReturnType() string {
	return strings.TrimSuffix(m.returnType, " *")
}

func (m Method) ReturnIsPointer() bool {
	return strings.HasSuffix(m.returnType, " *")
}

func (m Method) ReturnsVoid() bool {
	return m.returnType == "void"
}

type MethodArgs struct {
	name string
	methodNameAddon string
	argType string
}

func (arg MethodArgs) Name() string {
	return arg.name
}

func (arg MethodArgs) Type() string {
	return strings.TrimSuffix(arg.argType, " *")
}

func (arg MethodArgs) IsPointer() bool {
	return strings.HasSuffix(arg.argType, " *")
}

func parseMethod(input string) (Node,error)  {
	re, err := regexp.Compile(`^[-+]\s\((?P<returnType>\w+\s?\*?)\)(?P<baseName>\w+)(?P<args>:(?:.|\s)*)?;`)
	if err != nil {
		return nil, err
	}
	match := re.FindStringSubmatch(input)
	if len(match) == 0 {
		return nil, fmt.Errorf("not a method")
	}
	m := Method{
		returnType: match[re.SubexpIndex("returnType")],
		baseName: match[re.SubexpIndex("baseName")],
	}
	argsIndex := re.SubexpIndex("args")
	if argsIndex == -1 {
		// no args
		return m, nil
	}
	argsString := match[argsIndex]

	re, err = regexp.Compile(`(?:(?P<methodNameAddon>\w+)?:\((?P<argType>\w+\s?\*?)\)(?P<argName>\w+)\s*)`)
	if err != nil {
		return nil, err
	}

	argMatches := re.FindAllStringSubmatch(argsString, -1)

	for _, argMatch := range argMatches {
		methodNameAddonIdx := re.SubexpIndex("methodNameAddon")
		argTypeIdx := re.SubexpIndex("argType")
		argNameIdx := re.SubexpIndex("argName")

		m.args = append(m.args, MethodArgs{
			name:            argMatch[argNameIdx],
			methodNameAddon: argMatch[methodNameAddonIdx],
			argType:         argMatch[argTypeIdx],
		})
	}

	return m, nil
}