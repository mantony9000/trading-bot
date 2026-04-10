package main

import (
	"fmt"

	"go-trading-bot/config"
	"go-trading-bot/engine"
	"go-trading-bot/exchange"
	"go-trading-bot/strategy"
)

func main() {
	fmt.Println("Initializing Go Trading Bot...")

	// 1. Load Configuration
	cfg := config.LoadConfig()
	fmt.Printf("Loaded config for exchange: %s\n", cfg.Exchange.Name)

	// 2. Initialize Exchange
	var exch exchange.Exchange
	if cfg.Exchange.Name == "MockExchange" {
		exch = exchange.NewMockExchange()
	} else {
		// Initialize real exchange here
		panic("Unsupported exchange")
	}

	// 3. Initialize Strategy
	var strat strategy.Strategy
	if cfg.Strategy.Name == "DummyStrategy" {
		strat = strategy.NewDummyStrategy()
	} else {
		// Initialize real strategy here
		panic("Unsupported strategy")
	}

	// 4. Initialize and Start Engine
	eng := engine.NewEngine(cfg, exch, strat)
	eng.Start()

	fmt.Println("Bot execution completed.")
}
