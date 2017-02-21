package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdder interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdder
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Meat struct {
	Ingredient IngredientAdder
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdder is needed on the Ingredient field of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

type Onion struct {
	Ingredient IngredientAdder
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdder is needed on the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}
