package pb2proto

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type EnumDesc struct {
	descriptor protoreflect.EnumDescriptor
	depth      int64
	name       string
	values     []Desc
	options    []Desc
}

func NewEnumDesc(depth int64) *EnumDesc {
	return &EnumDesc{
		depth:   depth,
		name:    "",
		values:  []Desc{},
		options: []Desc{},
	}
}

func (e *EnumDesc) Build(inputSource interface{}) (Desc, error) {
	enumDescPB, ok := inputSource.(protoreflect.EnumDescriptor)
	if !ok {
		return nil, fmt.Errorf("failed to convert input source to Enum Descriptor")
	}
	e.descriptor = enumDescPB
	e.name = string(enumDescPB.Name())
	enumDescPB.Values()
	// values
	for i := 0; i < enumDescPB.Values().Len(); i++ {
		valueDescPB := enumDescPB.Values().Get(i)
		val, err := NewEnumValueDesc(e.depth + 1).Build(valueDescPB)
		if err != nil {
			return e, fmt.Errorf("failed to build enum value desc")
		}
		e.values = append(e.values, val)
	}
	// enum options
	if enumDescPB.Options() != nil {
		enumDescPB.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
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

func (e *EnumDesc) Analyze() string {
	alignStr := ""
	for i := 0; i < int(e.depth); i++ {
		alignStr += alignUnit
	}
	start := alignStr + "enum " + e.name + " {" + "\n"
	optAlignStr := ""
	for i := 0; i < int(e.depth+1); i++ {
		optAlignStr += alignUnit
	}
	for _, optStr := range e.AnalyzeOptions() {
		start += optAlignStr + optStr + "\n"
	}
	for _, valueMsg := range e.values {
		valueMsgStr := valueMsg.Analyze()
		start += valueMsgStr + "\n"
	}
	end := alignStr + "}\n"
	return start + end
}

func (e *EnumDesc) Depth() int64 {
	return e.depth
}

func (e *EnumDesc) AnalyzeOptions() []string {
	optStrs := make([]string, 0)
	for _, desc := range e.options {
		optDesc, ok := desc.(*OptionDesc)
		if !ok {
			fmt.Println("failed to build enum value desc")
		}
		for i := 0; i < len(optDesc.KStrs); i++ {
			kstr := optDesc.KStrsWithFullName[i]
			vstr := optDesc.VStrs[i]
			optStr := fmt.Sprintf("option %s = %s;", kstr, vstr)
			optStrs = append(optStrs, optStr)
		}
	}
	return optStrs
}
