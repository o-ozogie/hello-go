package request

type Request interface {
	Validate() bool
}
