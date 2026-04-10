package config

// Config holds the overall configuration for the trading bot.
type Config struct {
	Exchange ExchangeConfig `json:"exchange"`
	Strategy StrategyConfig `json:"strategy"`
	Trading  TradingConfig  `json:"trading"`
}

// ExchangeConfig holds API credentials and exchange specific settings.
type ExchangeConfig struct {
	Name      string `json:"name"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

// StrategyConfig holds parameters for the trading strategy.
type StrategyConfig struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

// TradingConfig holds general trading parameters like pairs and risk limits.
type TradingConfig struct {
	Pairs          []string `json:"pairs"`
	MaxPosition    float64  `json:"max_position"`
	RiskPercentage float64  `json:"risk_percentage"`
}

// LoadConfig creates a basic configuration for testing purposes.
func LoadConfig() *Config {
	return &Config{
		Exchange: ExchangeConfig{
			Name:      "MockExchange",
			APIKey:    "test-key",
			APISecret: "test-secret",
		},
		Strategy: StrategyConfig{
			Name: "DummyStrategy",
			Parameters: map[string]string{
				"period": "14",
			},
		},
		Trading: TradingConfig{
			Pairs:          []string{"BTC-USD", "ETH-USD"},
			MaxPosition:    1.0,
			RiskPercentage: 0.05,
		},
	}
}
