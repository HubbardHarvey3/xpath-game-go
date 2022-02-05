package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"os"
)

var name string



func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to Xpath!")
	fmt.Println("What is your Name?")
	fmt.Scan(&name)

	// Go to the desert
	fmt.Printf("Welcome to the desert %v. \n", name)
	fmt.Printf(" 𝞇 𝞇 𝞇𝞇𝞇     𝞇   𝞇𝞇 𝞇 \n")
	fmt.Printf(" 𝞇           𝞇  𝞇𝞇          𝞇\n")
	fmt.Printf(" 𝞇 𝞇 𝞇          𝞇                 𝞇      𝞇   \n")
	fmt.Println("There are riddles you need to solve.")
	var answer string
	var ans int
	var x int = rand.Intn(10)
	var y int = rand.Intn(8)
	fmt.Printf("What is %v + %x?", x, y)
	fmt.Scan(&answer)
	ans, _ = strconv.Atoi(answer)
	if ans == (x+y) {
		fmt.Println("YOU DID IT!!! You can move on to the next round")
	} else {
		fmt.Println("Slowly, you die of thrist and your world slowly dims dark as you thrist to death.")
		os.Exit(1)
	}

	// the Forest Level
	fmt.Printf("Welcome to the forest %v.\n", name)
	fmt.Printf("🌳  🌳🌳🌳\n")
	fmt.Printf("🌳🌳     🌲🌴 🌳\n")
	fmt.Printf("🌳🌳🌳🌳🌴🌴\n")
	fmt.Printf("🌳🌳   🌴 🌳🌳🌳\n")
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(20)
	y = rand.Intn(10)
	fmt.Printf("What is %v + %x? \n", x, y)
	fmt.Scan(&answer)
	ans, _ = strconv.Atoi(answer)
	if ans == (x+y) {
		fmt.Println("YOU DID IT!!! You can move out of the forest and onto the next round")
	} else {
		fmt.Println("Slowly, the trees close in and you become disoriented before succombing to your hunger.")
		os.Exit(1)
	}

	// the Swamp
	fmt.Printf("Welcome to the swamp %v. Where only GATORS get out alive\n", name)
	fmt.Printf("🐊🐊    \n")
	fmt.Printf("🐸\n")
	fmt.Printf("🌩️  ⛈️    ⛈️    \n")
	fmt.Printf("🐊🐊    \n")
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(30)
	y = rand.Intn(10)
	fmt.Printf("What is %v + %x? \n", x, y)
	fmt.Scan(&answer)
	ans, _ = strconv.Atoi(answer)
	if ans == (x+y) {
		fmt.Println("YOU DID IT!!! YOU BEAT THE GAME!!!!! ")
		fmt.Println("You did great, you should play again.")
	} else {
		fmt.Println("Your last misstep spells certain doom as you flounder in the water and gators creep in to devour you.")
		os.Exit(1)
	}
}