package modconfig

import (
	"github.com/hashicorp/hcl/v2"
)

type ModTreeItemBase struct {
	// TODO NEEDED???
	ModTreeItemRemain hcl.Body `hcl:",remain" json:"-"`

	Mod   *Mod       `cty:"mod" json:"-"`
	Paths []NodePath `column:"path,jsonb" json:"-"`

	fullName string
	parents  []ModTreeItem
	children []ModTreeItem
}

// AddParent implements ModTreeItem
func (b *ModTreeItemBase) AddParent(parent ModTreeItem) error {
	b.parents = append(b.parents, parent)
	return nil
}

// GetParents implements ModTreeItem
func (b *ModTreeItemBase) GetParents() []ModTreeItem {
	return b.parents
}

// GetChildren implements ModTreeItem
func (b *ModTreeItemBase) GetChildren() []ModTreeItem {
	return b.children
}
func (b *ModTreeItemBase) GetPaths() []NodePath {
	// lazy load
	if len(b.Paths) == 0 {
		b.SetPaths()
	}
	return b.Paths
}

// SetPaths implements ModTreeItem
func (b *ModTreeItemBase) SetPaths() {
	for _, parent := range b.parents {
		for _, parentPath := range parent.GetPaths() {
			b.Paths = append(b.Paths, append(parentPath, b.fullName))
		}
	}
}
func (b *ModTreeItemBase) GetMod() *Mod {
	return b.Mod
}

// GetModTreeItemBase implements ModTreeItem
func (b *ModTreeItemBase) GetModTreeItemBase() *ModTreeItemBase {
	return b
}