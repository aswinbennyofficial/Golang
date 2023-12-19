package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"math"
)

func printBalance(client *ethclient.Client) {
	var accountID string
	fmt.Println("Enter Account address")
	fmt.Scan(&accountID)

	account := common.HexToAddress(accountID)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Convert balance from wei to ether
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("Balance is:", ethValue,"ETH")
}

func main() {

	// loading environment variable
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	INFURIA_KEY := os.Getenv("INFURIA_KEY")
	INFURIA_URL := "https://mainnet.infura.io/v3/" + INFURIA_KEY

	// establish connection
	client, err := ethclient.Dial(INFURIA_URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	printBalance(client)
}
