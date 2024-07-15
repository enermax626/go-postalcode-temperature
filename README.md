# Sistema de Consulta de Clima por CEP

## Descrição

Este sistema em Go recebe um CEP, identifica a cidade correspondente e retorna o clima atual, incluindo a temperatura em graus Celsius, Fahrenheit e Kelvin. O sistema está publicado no Google Cloud Run.


## Como Rodar o Programa

### Acessar Diretamente pelo Endpoint

Você pode acessar o programa diretamente pelo endpoint publicado no Google Cloud Run:

```
https://<seu-endpoint>.run.app
```

### Rodar com Docker

1. **Clone o repositório:**

   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd <NOME_DO_REPOSITORIO>
   ```

2. **Compile e rode a aplicação usando Docker:**

   ```bash
   docker build -t go-postalcode-temperature .
   docker run -p 8080:8080 go-postalcode-temperature
   ```

3. **Acesse a aplicação localmente:**

   Abra o navegador e vá para `http://localhost:8080`