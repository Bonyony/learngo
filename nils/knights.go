package nils

// this was a fun lesson check
// I added the drop method to test my knowledge

import "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if i == nil {
		return
	}
	c.leftHand = i
	fmt.Printf("%v picked up the %v\n", c.name, i.name)
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v cannot take any more items\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gave %v to %v\n", c.name, to.leftHand.name, to.name)
}

func (c *character) drop(i *item) {
	if i == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to drop\n", c.name)
		return
	}
	if c.leftHand != i {
		fmt.Printf("%v does not have that\n", c.name)
		return
	}
	c.leftHand = nil
	fmt.Printf("%v dropped the %v\n", c.name, i.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v has nothing in his hands!", c.name)
	}
	return fmt.Sprintf("%v has a %v", c.name, c.leftHand.name)
}

func Knights() {
	sword := item{"Sword"}
	shield := item{"Shield"}

	lancelot := character{"Sir Lancelot", nil}
	grumpert := character{"Grumpert", nil}

	lancelot.pickup(&sword)
	lancelot.give(&grumpert)

	fmt.Println(lancelot)
	fmt.Println(grumpert)

	grumpert.drop(&sword)
	fmt.Println(grumpert)
	grumpert.pickup(&shield)
	grumpert.drop(&sword)
}
