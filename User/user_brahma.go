package User

import "sample/mocktest/InterfaceAnimal"

type Brahma struct {
	typebrahma InterfaceAnimal.Animal
}

func (b *Brahma) Use(myval string) (string, error) {
	val, err := b.typebrahma.Breathe(myval)
	if err != nil {
		return "", err
	}
	return val, nil
}
