package tree

import (
	"github.com/mngoutham/GOExamples"
	"testing"
)

func TestDepthNonEmpty(t *testing.T) {
	root := GOExamples.TreeNode{}
	root.Left = &GOExamples.TreeNode{Val: 1}
	root.Left.Left = &GOExamples.TreeNode{Val: 2}

	depth := GOExamples.MaxDepth(&root)
	if depth != 3 {
		t.Errorf("Invalid tree depth, got %d", depth)
	}
}

func TestDepthEmpty(t *testing.T) {
	depth := GOExamples.MaxDepth(nil)
	if depth != 0 {
		t.Errorf("Invalid tree depth, got %d", depth)
	}
}
