package tree

import (
	"encoding/csv"
	"log"
	"os"
	"testing"

	"github.com/eugeneradionov/BE-test/entities"
)

var TestData = []*entities.Msg{
	{ID: `11`, ParentID: ``},
	{ID: `12`, ParentID: `11`},
	{ID: `13`, ParentID: `12`},
	{ID: `14`, ParentID: `13`},
	{ID: `15`, ParentID: `14`},
	{ID: `16`, ParentID: `15`},
	{ID: `17`, ParentID: `15`},

	{ID: `21`, ParentID: ``},
	{ID: `22`, ParentID: `21`},
	{ID: `23`, ParentID: `21`},
	{ID: `24`, ParentID: `21`},
	{ID: `25`, ParentID: `21`},
	{ID: `26`, ParentID: `21`},
	{ID: `27`, ParentID: `21`},

	{ID: `31`, ParentID: ``},
	{ID: `32`, ParentID: `31`},
	{ID: `33`, ParentID: `32`},
	{ID: `34`, ParentID: `33`},
	{ID: `35`, ParentID: `34`},
	{ID: `36`, ParentID: `35`},
	{ID: `37`, ParentID: `36`},
	{ID: `38`, ParentID: `37`},
	{ID: `39`, ParentID: `37`},
	{ID: `40`, ParentID: `37`},
	{ID: `41`, ParentID: `37`},
	{ID: `42`, ParentID: `37`},
	{ID: `43`, ParentID: `37`},
	{ID: `44`, ParentID: `37`},

	{ID: `51`, ParentID: ``},
	{ID: `52`, ParentID: `51`},
	{ID: `53`, ParentID: ``},
	{ID: `54`, ParentID: `53`},
	{ID: `55`, ParentID: ``},
	{ID: `56`, ParentID: `55`},
	{ID: `57`, ParentID: ``},
}

func TestNewTree(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := NewTree()
		for _, msg := range TestData {
			tree.Add(msg)
		}
		if !validTree(t, tree, TestData) {
			t.Fail()
		}
	}
}

func TestNewTreeLarge(t *testing.T) {
	tests := readCsvTests()
	for i := 0; i < 100; i++ {
		tree := NewTree()
		for _, msg := range tests {
			tree.Add(msg)
		}
		if !validTree(t, tree, tests) {
			t.Fail()
		}
	}
}

func readCsvTests() []*entities.Msg {
	f, err := os.Open(`test.csv`)
	if err != nil {
		log.Panic(`cant open file`)
	}
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		log.Panic(`cant read file`)
	}
	tests := make([]*entities.Msg, 0, len(lines))
	for _, l := range lines {
		tests = append(tests, &entities.Msg{
			ID:       l[0],
			ParentID: l[1],
		})
	}
	return tests
}

func validTree(t *testing.T, tree *Tree, tests []*entities.Msg) bool {
	// prepare a map of all id to msgMap all msgs are inserted
	msgMap := map[string]bool{}
	for _, msg := range tests {
		msgMap[msg.ID] = true
	}

	for _, n := range tree.children {
		if n.ParentID != `` {
			t.Error(`top level parent is not an empty string `)
			return false

		}
		if !validNode(t, msgMap, n) {
			t.Error(`node is not valid`)
			return false
		}
	}
	if len(msgMap) > 0 {
		t.Error(`missing msgs in tree`)
		return false
	}
	return true
}

func validNode(t *testing.T, msgMap map[string]bool, node *node) bool {
	delete(msgMap, node.ID)
	for _, n := range node.children {
		if node.ID != n.ParentID {
			t.Error(`child is not under proper parent`)
			return false
		}
		if !validNode(t, msgMap, n) {
			return false
		}
	}
	return true
}
