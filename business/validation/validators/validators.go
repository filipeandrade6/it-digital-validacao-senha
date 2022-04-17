package validators

type void struct{}

var v void // sugar for struct{}

type validater interface {
	Check(s string) bool
	SetNext(validater)
}
