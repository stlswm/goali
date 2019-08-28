package request

type Request interface {
	ExportJsonParams() []byte
}
