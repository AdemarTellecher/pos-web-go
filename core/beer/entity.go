package beer

// as tagh json define como os dados vão ser transformados em json
type Beer struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

type BeerType int

const (
	TypeAle   = 1
	TypeLager = 2
	TypeMalt  = 3
	TypeStout = 4
	TypeSoft  = 5
)

//desta forma a função String pertence ao tipo e pode ser usada da seguinte forma:
// var x TypeAle
// fmt.Println(x.String())

func (t BeerType) TypeBeer() string {
	switch t {
	case TypeAle:
		return "Beer Ale"
	case TypeLager:
		return "Beer Lager"
	case TypeMalt:
		return "Beer Malt"
	case TypeStout:
		return "Beer Stout"
	case TypeSoft:
		return "Beer Soft"
	}
	return "Unknown" //Retorna desconhecida
}

type BeerStyle int

// usando desta forma o compilador vai automaticamente definir os ids sequencialmente
// Lembrando que, caso tenha que adicionar mais modelos, colocar no final da fila.
const (
	StyleAmber = iota + 1
	StyleBlonde
	StyleBrown
	StyleCream
	StyleDark
	StylePale
	StyleStrong
	StyleWheat
	StyleRed
	StyleIPA
	StyleLime
	StylePilsner
	StyleGolden
	StyleFruit
	StyleHoney
)

func (t BeerStyle) StyleBeer() string {
	switch t {
	case StyleAmber:
		return "Amber"
	case StyleBrown:
		return "Brown"
	case StyleCream:
		return "Cream"
	case StyleDark:
		return "Dark"
	case StylePale:
		return "Pale"
	case StyleStrong:
		return "Strong"
	case StyleWheat:
		return "Wheat"
	case StyleRed:
		return "Red"
	case StyleIPA:
		return "India Pale Ale"
	case StyleLime:
		return "Lime"
	case StylePilsner:
		return "Pilsner"
	case StyleGolden:
		return "Golden"
	case StyleFruit:
		return "Fruit"
	case StyleHoney:
		return "Honey"
	}
	return "Unknown" // Retorna desconhecida
}
