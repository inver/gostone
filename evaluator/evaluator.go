package evaluator

import "github.com/inver/gostone/common"

type Evaluator struct {
}

func (*Evaluator) Process(node common.AstNode, context map[string]EvalNode) (string, error) {

	return "", nil
}
