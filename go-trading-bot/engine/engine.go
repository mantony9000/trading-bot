package engine

import (
	"fmt"
	"time"

	"go-trading-bot/config"
	"go-trading-bot/exchange"
	"go-trading-bot/strategy"
)

// Engine manages the trading loop, tying together exchange and strategy.
type Engine struct {
	Config   *config.Config
	Exchange exchange.Exchange
	Strategy strategy.Strategy
	running  bool
}

// NewEngine creates a new trading engine instance.
func NewEngine(cfg *config.Config, exch exchange.Exchange, strat strategy.Strategy) *Engine {
	return &Engine{
		Config:   cfg,
		Exchange: exch,
		Strategy: strat,
		running:  false,
	}
}

// Start begins the main trading loop.
func (e *Engine) Start() {
	fmt.Println("Starting trading engine...")
	e.running = true

	// Simulate a simple trading loop for demonstration.
	// In a real bot, you'd likely use websockets or a ticker.
	for i := 0; i < 3; i++ {
		if !e.running {
			break
		}

		fmt.Printf("\n--- Tick %d ---\n", i+1)
		e.tick()

		time.Sleep(1 * time.Second) // Wait between ticks
	}

	fmt.Println("Trading engine stopped.")
}

// Stop halts the trading loop.
func (e *Engine) Stop() {
	e.running = false
}

// tick represents a single iteration of the trading loop.
func (e *Engine) tick() {
	for _, pair := range e.Config.Trading.Pairs {
		// 1. Fetch Market Data
		ticker, err := e.Exchange.GetTicker(pair)
		if err != nil {
			fmt.Printf("Error getting ticker for %s: %v\n", pair, err)
			continue
		}

		// 2. Evaluate Strategy
		signal, err := e.Strategy.CalculateSignal(ticker)
		if err != nil {
			fmt.Printf("Error calculating signal for %s: %v\n", pair, err)
			continue
		}

		fmt.Printf("Signal for %s: %s\n", pair, signal)

		// 3. Execute Trade (if signal dictates)
		if signal == strategy.SignalBuy {
			// Example order execution logic
			_, err := e.Exchange.PlaceOrder(pair, "buy", "market", ticker.Price, 0.01)
			if err != nil {
				fmt.Printf("Error placing buy order: %v\n", err)
			}
		} else if signal == strategy.SignalSell {
			_, err := e.Exchange.PlaceOrder(pair, "sell", "market", ticker.Price, 0.01)
			if err != nil {
				fmt.Printf("Error placing sell order: %v\n", err)
			}
		}
	}
}
