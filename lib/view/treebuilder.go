package view

import (
	"context"
	"fmt"

	"github.com/atotto/clipboard"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run(tree *tview.TreeView, ctx context.Context) {
	app := tview.NewApplication()
	go func() {
		<-ctx.Done()
		app.Stop()
	}()
	if err := app.EnableMouse(true).SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}
}

func BuildTree(oyaml interface{}, cancel context.CancelFunc) *tview.TreeView {
	rootDir := "."
	root := tview.NewTreeNode(rootDir).Expand() //.SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			cancel()
		}
	})
	traverse(root, oyaml)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference != nil {
			clipboard.WriteAll(fmt.Sprint(reference))
			cancel()
		}
		children := node.GetChildren()
		if len(children) != 0 {
			node.SetExpanded(!node.IsExpanded())
		}
	})

	return tree
}

func makeSimpleNode(v interface{}, pref string, ref interface{}, selectable bool, color tcell.Color) *tview.TreeNode {
	n := tview.NewTreeNode(pref + fmt.Sprint(v))
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
	n.Collapse()
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
	case string, int:
		node.AddChild(makeSimpleNode(data, "$ ", data, true, 0))
	}

	return node
}
