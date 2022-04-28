package pb2proto

type Desc interface {
	// decode information from protobuf.Message to protodec.Desc
	Build(inputSource interface{}) (Desc, error)
	// analyze protodec information and return .proto format string
	Analyze() string
	// return protodec element depth
	Depth() int64
	// analyze protobuf.Options and return .proto format strings
	AnalyzeOptions() []string
}
