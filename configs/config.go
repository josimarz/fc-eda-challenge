package configs

import "github.com/spf13/viper"

type Config struct {
	Port            string `mapstructure:"PORT"`
	WalletCoreDSN   string `mapstructure:"WALLET_CORE_DSN"`
	TransactionsDSN string `mapstructure:"TRANSACTIONS_DSN"`
	KafkaDSN        string `mapstructure:"KAFKA_DSN"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
