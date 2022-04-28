package pb2proto

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type EnumValueDesc struct {
	depth   int64
	name    string
	no      string
	options []Desc
}

func NewEnumValueDesc(depth int64) *EnumValueDesc {
	return &EnumValueDesc{
		depth:   depth,
		name:    "",
		no:      "",
		options: []Desc{},
	}
}

func (e *EnumValueDesc) Build(inputSource interface{}) (Desc, error) {
	enumValueDescPB, ok := inputSource.(protoreflect.EnumValueDescriptor)
	if !ok {
		return nil, fmt.Errorf("faile to convert input source to Enum Value Descriptor")
	}
	e.name = string(enumValueDescPB.Name())
	e.no = fmt.Sprintf("%d", enumValueDescPB.Number())
	// options
	if enumValueDescPB.Options() != nil {
		enumValueDescPB.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			optDesc, err := NewOptionDes(e.Depth()+1, v).Build(fd)
			if err != nil {
				return false
			}
			e.options = append(e.options, optDesc)
			return true
		})
	}
	return e, nil
}

func (e *EnumValueDesc) Analyze() string {
	alignStr := ""
	for i := 0; i < int(e.depth); i++ {
		alignStr += alignUnit
	}
	outStr := alignStr
	var name, no string
	name = e.name
	no = e.no
	// options
	optionStrs := ""
	if len(e.options) > 0 {
		optionStrs += "["
		analyzeStrs := e.AnalyzeOptions()
		for i := 0; i < len(analyzeStrs); i++ {
			optionStrs += analyzeStrs[i]
			if i != len(analyzeStrs)-1 {
				optionStrs += ","
			}
		}
		optionStrs += "]"
	}
	outStr += fmt.Sprintf("%s = %s%s;", name, no, optionStrs)
	return outStr
}

func (e *EnumValueDesc) Depth() int64 {
	return e.depth
}

func (e *EnumValueDesc) AnalyzeOptions() []string {
	optStrs := make([]string, 0)
	for _, desc := range e.options {
		optDesc, ok := desc.(*OptionDesc)
		if !ok {
			fmt.Println("failed to convert inpusource to OptionDesc")
			return nil
		}
		for i := 0; i < len(optDesc.KStrs); i++ {
			optStr := ""
			kstrWithFullName := optDesc.KStrsWithFullName[i]
			kstr := optDesc.KStrs[i]
			vstr := optDesc.VStrs[i]
			isCus := optDesc.Customizes[i]
			if isCus {
				optStr = fmt.Sprintf("(%s) = %s", kstrWithFullName, vstr)
			} else {
				optStr = fmt.Sprintf("%s = %s", kstr, vstr)
			}
			optStrs = append(optStrs, optStr)
		}
	}
	return optStrs
}
