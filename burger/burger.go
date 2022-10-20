package burger

type Burger struct {
	ingredients []Ingredient
}

func New(ingredients []Ingredient) *Burger {
	return &Burger{
		ingredients: ingredients,
	}
}
