package exchange

import (
	"fmt"
	"time"
)

// Ticker represents market data for a given trading pair.
type Ticker struct {
	Pair      string
	Price     float64
	Volume    float64
	Timestamp time.Time
}

// Order represents an instruction to buy or sell.
type Order struct {
	ID        string
	Pair      string
	Side      string // "buy" or "sell"
	Type      string // "market" or "limit"
	Price     float64
	Amount    float64
	Status    string // "open", "filled", "cancelled"
	Timestamp time.Time
}

// Balance represents the available balance for a specific asset.
type Balance struct {
	Asset     string
	Available float64
	Locked    float64
}

// Exchange defines the interface that any exchange integration must implement.
type Exchange interface {
	GetTicker(pair string) (*Ticker, error)
	PlaceOrder(pair, side, orderType string, price, amount float64) (*Order, error)
	GetBalance(asset string) (*Balance, error)
}

// MockExchange is a dummy implementation of the Exchange interface for testing.
type MockExchange struct{}

// NewMockExchange creates a new instance of MockExchange.
func NewMockExchange() *MockExchange {
	return &MockExchange{}
}

// GetTicker returns a dummy ticker.
func (m *MockExchange) GetTicker(pair string) (*Ticker, error) {
	fmt.Printf("MockExchange: Getting ticker for %s\n", pair)
	return &Ticker{
		Pair:      pair,
		Price:     50000.0, // Fixed dummy price
		Volume:    10.0,
		Timestamp: time.Now(),
	}, nil
}

// PlaceOrder returns a dummy filled order.
func (m *MockExchange) PlaceOrder(pair, side, orderType string, price, amount float64) (*Order, error) {
	fmt.Printf("MockExchange: Placing %s %s order for %f %s at %f\n", orderType, side, amount, pair, price)
	return &Order{
		ID:        "mock-order-id-123",
		Pair:      pair,
		Side:      side,
		Type:      orderType,
		Price:     price,
		Amount:    amount,
		Status:    "filled",
		Timestamp: time.Now(),
	}, nil
}

// GetBalance returns a dummy balance.
func (m *MockExchange) GetBalance(asset string) (*Balance, error) {
	fmt.Printf("MockExchange: Getting balance for %s\n", asset)
	return &Balance{
		Asset:     asset,
		Available: 1000.0,
		Locked:    0.0,
	}, nil
}
