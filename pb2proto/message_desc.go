package pb2proto

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type MessageDesc struct {
	descriptor protoreflect.MessageDescriptor
	depth      int64
	name       string
	fields     []Desc
	options    []Desc
	messages   []Desc
	enums      []Desc
}

func NewMessageDesc(depth int64) *MessageDesc {
	return &MessageDesc{
		depth:    depth,
		name:     "",
		fields:   make([]Desc, 0),
		options:  []Desc{},
		messages: []Desc{},
		enums:    []Desc{}}
}

func (m *MessageDesc) Build(inputSource interface{}) (Desc, error) {
	msgDescPB, ok := inputSource.(protoreflect.MessageDescriptor)
	if !ok {
		return m, fmt.Errorf("failed to convert input source to MessageDescriptor")
	}
	m.descriptor = msgDescPB
	m.name = string(msgDescPB.Name())
	// fields
	for i := 0; i < msgDescPB.Fields().Len(); i++ {
		fieldDescPB := msgDescPB.Fields().Get(i)
		field, err := NewFieldDesc(m.depth + 1).Build(fieldDescPB)
		if err != nil {
			return m, fmt.Errorf("failed to build field desc")
		}
		m.fields = append(m.fields, field)
	}
	// nested message
	for i := 0; i < msgDescPB.Messages().Len(); i++ {
		nestedMsgDesc := msgDescPB.Messages().Get(i)
		if nestedMsgDesc.IsMapEntry() {
			continue
		}
		nestedMsg, err := NewMessageDesc(m.depth + 1).Build(nestedMsgDesc)
		if err != nil {
			return nil, fmt.Errorf("failed to build nested message")
		}
		m.messages = append(m.messages, nestedMsg)
	}
	// enumrations
	for i := 0; i < msgDescPB.Enums().Len(); i++ {
		enumMsgDescPB := msgDescPB.Enums().Get(i)
		enumMsg, err := NewEnumDesc(m.depth + 1).Build(enumMsgDescPB)
		if err != nil {
			return nil, fmt.Errorf("failed to build enumeration")
		}
		m.enums = append(m.enums, enumMsg)
	}
	// message options
	if msgDescPB.Options() != nil {
		msgDescPB.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			optDesc, err := NewOptionDes(m.Depth()+1, v).Build(fd)
			if err != nil {
				return false
			}
			m.options = append(m.options, optDesc)
			return true
		})
	}
	return m, nil
}

func (m *MessageDesc) Analyze() string {
	alignStr := ""
	for i := 0; i < int(m.depth); i++ {
		alignStr += alignUnit
	}
	start := alignStr + "message " + m.name + " {" + "\n"
	optAlignStr := ""
	for i := 0; i < int(m.depth+1); i++ {
		optAlignStr += alignUnit
	}
	for _, optStr := range m.AnalyzeOptions() {
		start += optAlignStr + optStr + "\n"
	}
	for _, enumMsg := range m.enums {
		enumMsgStr := enumMsg.Analyze()
		start += enumMsgStr + "\n"
	}
	for _, nestedMsg := range m.messages {
		nestedMsgStr := nestedMsg.Analyze()
		start += nestedMsgStr + "\n"
	}
	for _, field := range m.fields {
		fieldStr := field.Analyze()
		start += fieldStr + "\n"
	}
	end := alignStr + "}\n"
	return start + end
}

func (m *MessageDesc) Depth() int64 {
	return m.depth
}

func (m *MessageDesc) AnalyzeOptions() []string {
	optStrs := make([]string, 0)
	for _, desc := range m.options {
		optDesc, ok := desc.(*OptionDesc)
		if !ok {
			fmt.Println("failed to convert to OptionDesc")
			return nil
		}
		for i := 0; i < len(optDesc.KStrs); i++ {
			kstr := optDesc.KStrsWithFullName[i]
			vstr := optDesc.VStrs[i]
			optStr := fmt.Sprintf("option(%s) = %s;", kstr, vstr)
			optStrs = append(optStrs, optStr)
		}
	}
	return optStrs
}
