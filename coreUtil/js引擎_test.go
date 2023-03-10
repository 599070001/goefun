package E

import (
	. "github.com/599070001/goefun/core"
	"testing"
)

func TestRunJs(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := RunJs("function get(p){return p}", "get", "aaaaaaa"+E到文本(i))
		E调试输出(i, str)
	}

}
