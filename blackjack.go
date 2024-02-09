package main

import (
	"fmt"
	"math/rand"
)

// Suit represents a card suit
type Suit uint8

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

// Rank represents a card rank
type Rank uint8

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card represents a playing card
type Card struct {
	Suit Suit
	Rank Rank
}

// String returns a string representation of the card
func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

// Deck represents a deck of cards
type Deck struct {
	cards []Card
}

// NewDeck creates a new deck of cards
func NewDeck() *Deck {
	deck := &Deck{}
	for s := Spades; s <= Clubs; s++ {
		for r := Ace; r <= King; r++ {
			deck.cards = append(deck.cards, Card{Suit: s, Rank: r})
		}
	}
	return deck
}

// Shuffle shuffles the deck using the Fisher-Yates shuffle algorithm
func (d *Deck) Shuffle() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(d.cards)
	for i := n - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// DrawCard draws a card from the deck
func (d *Deck) DrawCard() Card {
	if len(d.cards) == 0 {
		panic("Deck is empty")
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

// Hand represents a collection of cards
type Hand interface {
	Cards() []Card
	AddCard(card Card)
	Value() int // can be calculated differently based on different poker games
}

// BlackjackHand implements the Hand interface for Blackjack
type BlackjackHand struct {
	cards []Card
	soft bool // True if hand contains an Ace with value 11
}

// Cards returns the cards in the hand
func (h *BlackjackHand) Cards() []Card {
	return h.cards
}

// AddCard adds a card to the hand
func (h *BlackjackHand) AddCard(card Card) {
	h.cards = append(h.cards, card)
	if card.Rank == Ace {
		if h.Value() <= 10 {
			h.soft = true // Treat Ace as 11 unless it busts the hand

      // @TODO: improvement is to factor in 5 cards win. So sometimes, we want Ace to be 1 even tho total hand is <= 10.
		}
	}
}

// Value calculates the hand's value, considering Ace as 1 or 11
func (h *BlackjackHand) Value() int {
	value := 0
	for _, card := range h.cards {
		switch card.Rank {
		case Ace:
			if h.soft && value + 11 <= 21 {
				value += 11
			} else {
				value += 1
			}
		case Jack, Queen, King:
			value += 10
		default:
			value += int(card.Rank)
		}
	}
	return value
}

// PlayBlackjack simulates a Blackjack game against the dealer
func PlayBlackjack(deck *Deck) {
	player := BlackjackHand{}
	dealer := BlackjackHand{}

	// Deal initial cards
	player.AddCard(deck.DrawCard())
	player.AddCard(deck.DrawCard())
	dealer.AddCard(deck.DrawCard())
	dealer.AddCard(deck.DrawCard())

	fmt.Println("Your cards:", player.Cards())
	fmt.Println("Dealer shows:", dealer.Cards()[1]) // Don't reveal dealer's hole card

	// Player turn
	for {
  	action := ""
  	fmt.Println("Hit (h) or Stand (s)?")
  	fmt.Scanf("%s", &action)
  
  	if action == "h" {
  		player.AddCard(deck.DrawCard())
  		fmt.Println("Your cards:", player.Cards())
  		if player.Value() > 21 {
  			fmt.Println("Bust!")
  			break
  		}
  	} else if action == "s" {
  		break
  	} else {
  		fmt.Println("Invalid action. Please enter h or s.")
  	}
  }
  
  // Dealer turn
  fmt.Println("Dealer's cards:")
  for _, card := range dealer.Cards() {
  	fmt.Println(card)
  }
  
  for dealer.Value() < 17 {
  	dealer.AddCard(deck.DrawCard())
  	fmt.Println("Dealer hits:", dealer.Cards()[len(dealer.Cards())-1])
  	if dealer.Value() > 21 {
  		fmt.Println("Dealer busts!")
  		break
  	}
  }
  
  // Determine winner
  if player.Value() > 21 || (dealer.Value() <= 21 && dealer.Value() > player.Value()) {
  	fmt.Println("Dealer wins!")
  } else if dealer.Value() > 21 || (player.Value() <= 21 && player.Value() > dealer.Value()) {
  	fmt.Println("Player wins!")
  } else {
  	fmt.Println("Push!")
  }
}

// Improvements:
// Multiple players
// Doubling down
// Splitting pairs
// Insurance
// Different bet amounts
// More complex dealer strategy
