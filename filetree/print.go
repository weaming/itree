package filetree

import (
	"fmt"
	"strings"
)

const (
	SPACE           = "   "
	HORIZONTAL_LINE = "─"
	VERTICAL_LINE   = "│"
	T_PREFIX        = "├"
	END_PREFIX      = "└"
)

func PrintFileNode(node *FileNode, prefix []string, depth, level int) {
	for i, x := range node.Children {
		if i == 0 {
			// first
			prefix = append(prefix, T_PREFIX)
		}

		if i+1 == len(node.Children) {
			// last
			prefix = append(prefix[:len(prefix)-1], END_PREFIX)
		}

		fmt.Printf("%v%v %v %v\n", strings.Join(prefix, ""), strings.Repeat(HORIZONTAL_LINE, 2), x.Name, x.TotalSize)

		if x.Type == TYPE_DIR && depth < level {
			prefix = append(prefix[:len(prefix)-1], VERTICAL_LINE)
			prefix = append(prefix, SPACE)
			PrintFileNode(x, prefix, depth+1, level)
			prefix = append(prefix[:len(prefix)-2], T_PREFIX)
		}
	}
}
