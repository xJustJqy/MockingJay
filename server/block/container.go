package block

import (
	"github.com/xJustJqy/MockingJay/server/block/cube"
	"github.com/xJustJqy/MockingJay/server/item"
	"github.com/xJustJqy/MockingJay/server/item/inventory"
	"github.com/xJustJqy/MockingJay/server/world"
)

// ContainerViewer represents a viewer that is able to view a container and its inventory.
type ContainerViewer interface {
	// ViewSlotChange views a change of a single slot in the inventory, in which the item was changed to the
	// new item passed.
	ViewSlotChange(slot int, newItem item.Stack)
}

// ContainerOpener represents an entity that is able to open a container.
type ContainerOpener interface {
	// OpenBlockContainer opens a block container at the position passed.
	OpenBlockContainer(pos cube.Pos)
}

// Container represents a container of items, typically a block such as a chest. Containers may have their
// inventory opened by viewers.
type Container interface {
	AddViewer(v ContainerViewer, w *world.World, pos cube.Pos)
	RemoveViewer(v ContainerViewer, w *world.World, pos cube.Pos)
	Inventory() *inventory.Inventory
}
