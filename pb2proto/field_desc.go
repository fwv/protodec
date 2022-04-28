package pb2proto

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldDesc struct {
	descriptor  protoreflect.FieldDescriptor
	depth       int64
	cardinality string
	fieldType   string
	name        string
	no          int32
	isMap       bool
	keyType     string
	valType     string
	options     []Desc
}

func NewFieldDesc(depth int64) *FieldDesc {
	return &FieldDesc{
		depth:       depth,
		cardinality: "",
		fieldType:   "",
		name:        "",
		no:          0,
		isMap:       false,
		keyType:     "",
		valType:     "",
		options:     []Desc{},
	}
}

func (f *FieldDesc) Build(inputSource interface{}) (Desc, error) {
	fieldDescPB, ok := inputSource.(protoreflect.FieldDescriptor)
	if !ok {
		return f, fmt.Errorf("faile to convert input source to FieldDescriptor")
	}
	f.descriptor = fieldDescPB
	f.cardinality = fieldDescPB.Cardinality().String()
	if fieldDescPB.IsMap() {
		f.isMap = true
		f.keyType = f.AnalyzeFieldType(fieldDescPB.MapKey())
		f.valType = f.AnalyzeFieldType(fieldDescPB.MapValue())
	}
	f.fieldType = f.AnalyzeFieldType(fieldDescPB)
	f.name = string(fieldDescPB.Name())
	f.no = int32(fieldDescPB.Number())
	// options
	if fieldDescPB.Options() != nil {
		fieldDescPB.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			optDesc, err := NewOptionDes(f.Depth()+1, v).Build(fd)
			if err != nil {
				return false
			}
			f.options = append(f.options, optDesc)
			return true
		})
	}
	return f, nil
}

func (f *FieldDesc) AnalyzeFieldType(fieldDesc protoreflect.FieldDescriptor) string {
	if fieldDesc == nil {
		return "unknown_type"
	}
	fieldType := fieldDesc.Kind().String()
	if fieldType == "message" {
		parentName := string(fieldDesc.Parent().FullName())
		msgName := string(fieldDesc.Message().FullName())
		return trimFieldName(parentName, msgName)
	} else if fieldType == "enum" {
		parentName := string(fieldDesc.Parent().FullName())
		enumName := string(fieldDesc.Enum().FullName())
		return trimFieldName(parentName, enumName)
	}
	return fieldType
}

func (f *FieldDesc) Analyze() string {
	alignStr := ""
	for i := 0; i < int(f.depth); i++ {
		alignStr += alignUnit
	}
	outStr := alignStr
	var cardinality, ftype, name, no string
	// syntax compatible
	if !f.isMap {
		if syntax == "proto3" {
			if f.cardinality != "optional" {
				cardinality = f.cardinality
			}
		} else if syntax == "proto2" {
			cardinality = f.cardinality
		}
	}
	if f.isMap {
		ftype = fmt.Sprintf("map<%s, %s>", f.keyType, f.valType)
	} else {
		ftype = f.fieldType
	}
	name = f.name
	no = fmt.Sprintf("%v", f.no)
	// options
	optionStrs := ""
	if len(f.options) > 0 {
		optionStrs += "["
		analyzeStrs := f.AnalyzeOptions()
		for i := 0; i < len(analyzeStrs); i++ {
			optionStrs += analyzeStrs[i]
			if i != len(analyzeStrs)-1 {
				optionStrs += ","
			}
		}
		optionStrs += "]"
	}
	if cardinality != "" {
		outStr += fmt.Sprintf("%s %s %s = %s%s;", cardinality, ftype, name, no, optionStrs)
	} else {
		outStr += fmt.Sprintf("%s %s = %s%s;", ftype, name, no, optionStrs)
	}
	return outStr
}

func (f *FieldDesc) Depth() int64 {
	return f.depth
}

func (f *FieldDesc) AnalyzeOptions() []string {
	optStrs := make([]string, 0)
	for _, desc := range f.options {
		optDesc, ok := desc.(*OptionDesc)
		if !ok {
			fmt.Println("failed to convert to OptionDesc")
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

func trimFieldName(parentName, fullName string) string {
	parentPaths := strings.Split(parentName, ".")
	paths := strings.Split(fullName, ".")
	name := ""
	for i := 0; i < len(paths); i++ {
		if i < len(parentPaths) && paths[i] == parentPaths[i] {
			continue
		}
		name += paths[i]
		if i != len(paths)-1 {
			name += "."
		}
	}
	return name
}
