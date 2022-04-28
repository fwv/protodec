package pb2proto

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	syntax string
)

const (
	alignUnit string = "  "
)

type FileDesc struct {
	descriptor  protoreflect.FileDescriptor
	syntax      string
	pkg         string
	fileName    string
	dependencys []string
	options     []Desc
	enums       []Desc
	messages    []Desc
}

func NewFileDesc() *FileDesc {
	return &FileDesc{
		syntax:      "",
		pkg:         "",
		fileName:    "",
		dependencys: make([]string, 0),
		options:     []Desc{},
		enums:       []Desc{},
		messages:    make([]Desc, 0),
	}
}

func (f *FileDesc) Build(inputSource interface{}) (Desc, error) {
	fileDescPB, ok := inputSource.(protoreflect.FileDescriptor)
	if !ok || fileDescPB == nil {
		return f, fmt.Errorf("faile to convert input source to FileDescriptor")
	}
	f.descriptor = fileDescPB
	f.syntax = fileDescPB.Syntax().String()
	syntax = f.syntax
	f.pkg = string(fileDescPB.Package().Name())
	// imports
	for i := 0; i < fileDescPB.Imports().Len(); i++ {
		dep := fileDescPB.Imports().Get(i)
		f.dependencys = append(f.dependencys, dep.Path())
	}
	// messages
	for i := 0; i < fileDescPB.Messages().Len(); i++ {
		messageDescPB := fileDescPB.Messages().Get(i)
		message, err := NewMessageDesc(f.Depth() + 1).Build(messageDescPB)
		if err != nil {
			return f, fmt.Errorf("failed to build message desc")
		}
		f.messages = append(f.messages, message)
	}
	// enumrations
	for i := 0; i < fileDescPB.Enums().Len(); i++ {
		enumMsgDescPB := fileDescPB.Enums().Get(i)
		enumMsg, err := NewEnumDesc(f.Depth() + 1).Build(enumMsgDescPB)
		if err != nil {
			return nil, fmt.Errorf("failed to build enumeration")
		}
		f.enums = append(f.enums, enumMsg)
	}
	// options
	if fileDescPB != nil {
		fileDescPB.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
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

func (f *FileDesc) Analyze() string {
	syntaxStr := fmt.Sprintf("syntax = \"%s\";", f.syntax)
	pkgStr := fmt.Sprintf("package %s;", f.pkg)
	optStrs := ""
	for _, opt := range f.AnalyzeOptions() {
		optStrs += opt
		optStrs += "\n"
	}
	depStrs := ""
	for _, dep := range f.dependencys {
		depStr := fmt.Sprintf("import \"%s\";", dep)
		depStr += "\n"
		depStrs += depStr
	}
	enumStrs := ""
	for _, enumMsg := range f.enums {
		enumMsgStr := enumMsg.Analyze()
		enumStrs += enumMsgStr + "\n"
	}
	msgStrs := ""
	for _, message := range f.messages {
		msgStr := message.Analyze()
		msgStrs += msgStr
		msgStrs += "\n"
	}
	outStr := syntaxStr + "\n" +
		pkgStr + "\n" +
		optStrs +
		depStrs + "\n" +
		enumStrs +
		msgStrs
	return outStr
}

func (f *FileDesc) Depth() int64 {
	return -1
}

func (f *FileDesc) AnalyzeOptions() []string {
	optStrs := make([]string, 0)
	for _, desc := range f.options {
		optDesc, ok := desc.(*OptionDesc)
		if !ok {
			fmt.Println("faild to convert to OptionDesc")
			return optStrs
		}
		for i := 0; i < len(optDesc.KStrs); i++ {
			kstr := optDesc.KStrs[i]
			vstr := optDesc.VStrs[i]
			optStr := fmt.Sprintf("option %s = %s;", kstr, vstr)
			optStrs = append(optStrs, optStr)
		}
	}
	return optStrs
}

func Analyze(descriptor protoreflect.FileDescriptor) (string, error) {
	if descriptor == nil {
		return "", fmt.Errorf("file descriptor is nil")
	}
	fileDesc, err := NewFileDesc().Build(descriptor)
	if err != nil {
		return "", err
	}
	return fileDesc.Analyze(), nil
}
