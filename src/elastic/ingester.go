package elastic

import (
	"acompanhalegis-analytics-api/src/dadosabertos"
	"bytes"
	"encoding/json"
)

var client, _ = GetClient()

func IngestProposicoes(proposicoes []dadosabertos.Proposicao) error {
	for _, proposicao := range proposicoes {
		err := IngestProposicao(proposicao)
		if err != nil {
			return err
		}
	}

	return nil
}

func IngestProposicao(proposicao dadosabertos.Proposicao) error {
	data, err := json.Marshal(proposicao)
	if err != nil {
		logger.Errorf("Error marshalling proposicao: %v", err)
		return err
	}
	_, err = client.Index("proposicoes", bytes.NewReader(data))
	if err != nil {
		logger.Errorf("Error indexing proposicao: %v - %v", err, proposicao)
		return err
	}

	return nil
}
