package main

import (
	"fmt"
)

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item) {
	// Adds an item to a players inventory
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s!\n", p.Name, item.Name)
}

func (p *Player) DropItem(itemName string) {
	// Removes an item from the invefmt.Println("Feeling gontory by name
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...) // remove item
			fmt.Printf("%s dropped %s.\n", p.Name, itemName)
			return
		}
	}

	fmt.Printf("%s does not have %s in inventory.\n", p.Name, itemName)
}

func (p *Player) UseItem(itemName string) {
	// Depending on the item - it will print different things
	for _, item := range p.Inventory {
		if item.Name == itemName {
			fmt.Println("Using ", item.Name)
			return
		} else {
			fmt.Println("No item to use")
			return
		}
	}
}

func main() {
	player1 := Player{
		Name: "dominator",
	}

	pickAxe := Item{
		Name: "pick axe",
		Type: "chops only wood",
	}

	axe := Item{
		Name: "axe",
		Type: "chops everything icluding heads",
	}

	player1.PickUpItem(pickAxe)
	player1.PickUpItem(axe)
	fmt.Printf("Here is player1 %+v\n", player1)

	player1.UseItem("pick axe")
	player1.DropItem("pick axe")
	fmt.Printf("Here is player1 %+v\n", player1)

}
