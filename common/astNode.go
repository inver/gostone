package common

type AstNode struct {
	Type AstType
}

type ValueNode interface {
	GetValue()
	HasValue()
	Size()
}

type BooleanAstNode struct {
	Value bool
}

func (node *BooleanAstNode) GetValue() interface{} {
	return node.Value
}

type LongAstNode struct {
	value int32
}

type DoubleAstNode struct {
	value float32
}

type StringAstNode struct {
	value string
}
