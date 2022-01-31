## Setup Local
### Subindo o ambiente de dev
   - Garanta ter o docker e docker-compose instalados em sua máquina
     - `docker -v` e `docker-compose -v`
   - Utilize o comando `docker-compose up --build` para inicializar o servidor http
     - a API deve começar a ouvir em `http://localhost:8080/`

### Comandos cURL para testes
####  List Book
```bash
curl --location --request GET 'http://localhost:8080/api/v1/books/1'
```
####  Create Book
```bash
curl --location --request POST 'http://localhost:8080/api/v1/books/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "O Poder do Hábito",
    "author": "Charles Duhigg"
}'
```
####  Delete Book
```bash
curl --location --request DELETE 'http://localhost:8080/api/v1/books/1'
```
####  Update Book
```bash
curl --location --request PUT 'http://localhost:8080/api/v1/books/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Mais Rápido e Melhor"
}'
```
#### List All Books
```bash
curl --location --request GET 'http://localhost:8080/api/v1/books/'
```