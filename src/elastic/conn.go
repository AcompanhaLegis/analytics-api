package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	"os"
)

func GetClient() (*elasticsearch.Client, error) {
	godotenv.Load(".env")
	cert, err := os.ReadFile(os.Getenv("APP_CERTS_DIR") + "/ca.crt")
	if err != nil {
		logger.Fatalf("Error reading certificate: %s", err)
		return nil, err
	}
	cfg := elasticsearch.Config{
		CACert: cert,
		Addresses: []string{
			os.Getenv("ES_HOST"),
		},
		Username: os.Getenv("ELASTIC_USERNAME"),
		Password: os.Getenv("ELASTIC_PASSWORD"),
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Fatalf("Error creating the client: %s", err)
		return nil, err
	}
	return es, nil
}
