# temperature
Desafio  fullcycle

Para configurar a API weather, edite o arquivo config.toml e inclua sua chave de API.
```
[weather]
key="your api key"
```
Para rodar o desafio basta executar o seguinte comando:
```
docker compose up -d --build
```
```
 curl -X GET http://localhost:8080/v1/temperature?cep=00000-000
 Response: 200
 {
	"StatusCode": 200,
	"message": "success",
	"data": {
		"state": "Rio+de+Janeiro",
		"temp_c": 22.4,
		"temp_f": 72.32,
		"temp_k": 295.55
	}
 }
  Response: 404
 {
	"StatusCode": 404,
	"message": "zipcode not found."
  }
 ```

## Test
```
 docker compose run app go test test/status_code_test.go
```