package main

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MongoPrimaryShardConnectionString string `envconfig:"mongo_hot_connection_string" required:"true"`
	MongoDBHotShardConnectionString   string `envconfig:"mongo_primary_connection_string" required:"true"`
	AMQPConnectionString              string `envconfig:"amqp_connection_string" required:"true"`
	LookupServiceConnectionString     string `envconfig:"lookup_service_connection_string" required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
