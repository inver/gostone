package infrastructure

type TestUnit struct {
	Name  string
	Cases []Case
}
type Case struct {
	Input          string
	ExpectedResult string
}
