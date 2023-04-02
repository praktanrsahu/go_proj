package InterfaceAnimal

import "fmt"

type Animal interface {
	Breathe(val string) (string, error)
}

type Cat struct {
	sound string
}

type Dog struct {
	sound string
}

func (c Cat) Breath(val string) (string, error) {
	if val != "15" {
		return "", fmt.Errorf("Breath provided- %v , Please provide appropriate. ", val)
	}
	val = "I breath " + val + " times per minute"
	return val, nil
}

func (c Cat) Sound() (string, error) {
	if c.sound != "Meow!!" {
		return "", fmt.Errorf("sound provided for Cat: %v , Please provide appropriate. ", c.sound)
	}
	c.sound = "My sound is " + c.sound
	return c.sound, nil
}

func (d Dog) Breath(val string) (string, error) {
	if val != "25" {
		return "", fmt.Errorf("Breath provided- %v , Please provide appropriate. ", val)
	}
	val = "I breath " + val + " times per minute"
	return val, nil
}

func (d Dog) Sound() (string, error) {
	if d.sound != "Bhow!!" {
		return "", fmt.Errorf("sound provided for Dog: %v , Please provide appropriate. ", d.sound)
	}
	d.sound = "My sound is " + d.sound
	return d.sound, nil
}
