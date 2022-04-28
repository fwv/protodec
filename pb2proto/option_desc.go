package pb2proto

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type OptionDesc struct {
	descriptor        protoreflect.FieldDescriptor
	value             protoreflect.Value
	depth             int64
	KStrs             []string
	KStrsWithFullName []string
	VStrs             []string
	Customizes        []bool
}

func NewOptionDes(depth int64, value protoreflect.Value) *OptionDesc {
	return &OptionDesc{
		value:             value,
		depth:             depth,
		KStrs:             []string{},
		KStrsWithFullName: []string{},
		VStrs:             []string{},
		Customizes:        []bool{},
	}
}

func (v *OptionDesc) Build(inputSource interface{}) (Desc, error) {
	fieldDescPB, ok := inputSource.(protoreflect.FieldDescriptor)
	if !ok {
		return v, fmt.Errorf("failed to convert inputsource to FieldDescriptor when build OptionDesc")
	}
	v.descriptor = fieldDescPB
	valuePB := v.value
	keystrWithFullName := string(fieldDescPB.FullName())
	keystr := string(fieldDescPB.Name())
	if fieldDescPB.IsList() {
		valuePBList := valuePB.List()
		for i := 0; i < valuePBList.Len(); i++ {
			valstr := v.AnalyzeValue(valuePBList.Get(i))
			v.KStrs = append(v.KStrs, keystr)
			v.KStrsWithFullName = append(v.KStrsWithFullName, keystrWithFullName)
			v.Customizes = append(v.Customizes, v.IsCustomize(keystrWithFullName))
			v.VStrs = append(v.VStrs, valstr)
		}
	} else {
		valstr := v.AnalyzeValue(valuePB)
		v.KStrs = append(v.KStrs, keystr)
		v.KStrsWithFullName = append(v.KStrsWithFullName, keystrWithFullName)
		v.Customizes = append(v.Customizes, v.IsCustomize(keystrWithFullName))
		v.VStrs = append(v.VStrs, valstr)
	}
	return v, nil
}

func (v *OptionDesc) Analyze() string {
	return ""
}

func (v *OptionDesc) AnalyzeValue(val protoreflect.Value) string {
	fieldDescPB := v.descriptor
	if fieldDescPB == nil {
		return ""
	}
	if fieldDescPB.Kind().String() == "string" {
		return fmt.Sprintf("\"%v\"", val.Interface())
	}
	return fmt.Sprintf("%v", val.Interface())
}

func (v *OptionDesc) Depth() int64 {
	return v.depth
}

func (v *OptionDesc) AnalyzeOptions() []string {
	return nil
}

func (v *OptionDesc) IsCustomize(fullName string) bool {
	return !strings.Contains(fullName, "google.protobuf")
}
