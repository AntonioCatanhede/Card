package main

// var card string // eh possivel inicializar fora da funcao sem atribuir um valor

// func main() {

// 	var card string = "Ace of Spades" // Go
// 	card2 := "Ace of Clubs"           // operador ':=' utilizado apenas em inicializacao de variaveis
// 	fmt.Println(card2)
// 	card2 = "Ace of Hearts" // depois, usa-se o '=' parra assignment
// 	card3 = "Ace of Diamonds"
// 	fmt.Println(card)
// 	fmt.Println(card2)
// 	fmt.Println(card3)

// }

// func main() {
// 	// card := newCard()
// 	fmt.Println(runtime.NumCPU())
// 	fmt.Println("hello world")
// 	// fmt.Println(newCard()) // eh possivel imprimir direto

// }

func main() {
	// cards := []string{"Ace of Diamonds", newCard()} //um exemplo de slice

	// cards1 := deck{"Ace of Diamonds", newCard()} // msm coisa que a linha aimca, mas usando o type definido por mim
	cards := newDeckFromFile("my_cards")
	// fmt.Println(cards.toString())
	cards.shuffle()
	cards.printDeck()

	// fmt.Println(cards[0], " and ", cards[1])

	// cards = append(cards, "Six of Spades")

	// fmt.Println(cards[0], ", ", cards[1], ", ", cards[2]) // forma burra

	// baralho := newDeck()

	// baralho.printDeck()

	// for i, card := range cards { //for's !!!
	// 	fmt.Println(i, card)
	// }

	// cards1.printDeck() // usando funcao definida no arquivo deck.go
	// que eh igual ao codigo comentado acima

	// hand, remainingDeck := deal(cards, 5)

	// hand.printDeck()
	// remainingDeck.printDeck()

	// deck.printDeck(remainingDeck)
	// deck.printDeck(hand)

}

// sintaxe de funcao que retorna string
