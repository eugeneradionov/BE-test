package tree

import (
	"sync"

	"github.com/eugeneradionov/BE-test/entities"
)

type Tree struct {
	mu *sync.Mutex

	children    []*node
	childrenMap map[string]*node
}

func NewTree() *Tree {
	return &Tree{
		mu:          &sync.Mutex{},
		children:    make([]*node, 0),
		childrenMap: make(map[string]*node),
	}
}

func (t *Tree) Add(msg *entities.Msg) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if msg.ParentID == "" {
		newNode := newNode(msg)
		t.children = append(t.children, newNode)
		t.childrenMap[msg.ID] = newNode

		return
	}

	node, ok := t.childrenMap[msg.ParentID]
	if ok {
		newNode := newNode(msg)
		node.children = append(node.children, newNode)
		t.childrenMap[newNode.ID] = newNode

		return
	}

	// add new parent with children if there is no one found
	parent := newNode(&entities.Msg{
		ID:       msg.ParentID,
		ParentID: "",
	})

	t.children = append(t.children, parent)
	t.childrenMap[parent.ID] = parent

	newNode := newNode(msg)
	parent.children = append(parent.children, newNode)
}
