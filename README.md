## Elasticsearch Setup

Crie os seguintes indices no Elasticsearch:

```
PUT /proposicoes
PUT /deputados
```

Gerar a API Key: http://127.0.0.1:5601/app/management/security/api_keys/

Os certificados do elasticsearch são automaticamente gerados pelo docker-compose em um volume compartilhado.
A maneira mais simples de rodar o projeto por enquanto é utilizando o docker-compose.

```bash
docker-compose run app bash
go run main.go
```
