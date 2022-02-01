package view

import (
	"github.com/atotto/clipboard"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run(tree *tview.TreeView) {
	app := tview.NewApplication()
	if err := app.SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}
}

func BuildTree(oyaml interface{}) *tview.TreeView {
	rootDir := "."
	root := tview.NewTreeNode(rootDir) //.SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	traverse(root, oyaml)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference != nil {
			clipboard.WriteAll(reference.(string))
		}
		children := node.GetChildren()
		if len(children) != 0 {
			node.SetExpanded(!node.IsExpanded())
		}
	})

	return tree
}

func makeSimpleNode(v interface{}, pref string, ref interface{}, selectable bool, color tcell.Color) *tview.TreeNode {
	var res string
	switch data := v.(type) {
	case string:
		res = data
	case int:
		res = string(data)
	default:
		res = "-unknown-"
	}

	n := tview.NewTreeNode(pref + res)
	if ref != nil {
		n.SetReference(ref)
	}
	if selectable {
		n.SetSelectable(selectable)
	}
	if selectable {
		n.SetSelectable(selectable)
	}
	if color != 0 {
		n.SetColor(color)
	}
	return n
}

func traverse(node *tview.TreeNode, oyaml interface{}) *tview.TreeNode {
	switch data := oyaml.(type) {
	case []interface{}:
		for _, v := range data {
			traverse(node, v)
		}
	case map[interface{}]interface{}:
		for k, v := range data {
			key := makeSimpleNode(k, "", nil, false, tcell.ColorGreen)
			node.AddChild(traverse(key, v))
		}
	case string:
		node.AddChild(makeSimpleNode(data, "$ ", data, true, 0))
	}

	return node
}
