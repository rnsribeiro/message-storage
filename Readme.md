# 📩 Message Storage API

Uma API REST desenvolvida em **Go (Golang)** utilizando o framework **Fiber** para interagir com o smart contract `MessageStorage` na **Polygon Mainnet**.

Este contrato permite:
- Armazenar mensagens com IDs únicos.
- Recuperar mensagens por ID.
- Verificar o total de mensagens armazenadas.

A conexão com a blockchain é feita via **Alchemy**, com configuração segura através de um arquivo `.env`.

---

## 🔗 Endpoints da API

- `POST /set-message`: Armazena uma nova mensagem no contrato.
- `GET /get-message/:id`: Recupera uma mensagem pelo seu ID.
- `GET /total-messages`: Retorna o número total de mensagens.

---

## 🗂 Estrutura do Projeto

```
message-storage/
├── abi/                        # ABI compilada do contrato
│   └── MessageStorage.abi
├── contracts/
│   └── MessageStorage.sol     # Contrato Solidity
├── message/
│   └── message_storage.go     # Go bindings gerados via abigen
├── .env                       # Variáveis de ambiente (não commitado)
├── .env.example               # Exemplo de variáveis de ambiente
├── .gitignore                 # Ignora arquivos sensíveis
├── go.mod                     # Dependências Go
├── main.go                    # Implementação do servidor Fiber
└── requests.http              # Requisições para teste da API
```

---

## ✅ Pré-requisitos

- [Node.js](https://nodejs.org/)
- [Go 1.18+](https://golang.org/)
- Conta no [Alchemy](https://www.alchemy.com/)
- Endereço do contrato na Polygon Mainnet
- Chave privada de uma carteira Polygon com saldo suficiente de **POL**
- [Git](https://git-scm.com/)
- Ferramenta REST como:
  - VS Code com extensão **REST Client**
  - IDEs JetBrains com suporte nativo a requisições HTTP

---

## ⚙️ Instruções de Setup

### 1. Clone o repositório

```bash
git clone https://github.com/rnsribeiro/message-storage.git
cd message-storage
```

---

### 2. Compile o contrato Solidity

```bash
npm install -g solc
mkdir abi
solcjs contracts/MessageStorage.sol --abi -o abi
```

Isso gerará:  
`abi/MessageStorage_sol_MessageStorage.abi`

---

### 3. Gere o pacote Go com `abigen`

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
mkdir message
abigen --abi abi/MessageStorage_sol_MessageStorage.abi --pkg message --type MessageStorage --out message/message_storage.go
```

---

### 4. Inicialize o projeto Go

```bash
go mod init github.com/rnsribeiro/message-storage
go get github.com/gofiber/fiber/v2
go get github.com/ethereum/go-ethereum
go get github.com/joho/godotenv
```

---

### 5. Configure o `.env`

Crie o arquivo `.env` com o seguinte conteúdo:

```env
PRIVATE_KEY=SEU_PRIVATE_KEY
ALCHEMY_URL=https://polygon-mainnet.g.alchemy.com/v2/SUA_API_KEY
CONTRACT_ADDRESS=ENDERECO_DO_CONTRATO
```

Garanta que `.env` esteja no `.gitignore`:

```bash
echo ".env" >> .gitignore
```

---

### 6. Execute o servidor

```bash
go mod tidy
go run main.go
```

A API estará disponível em:  
📍 `http://localhost:3000`

---

## 🔬 Testando a API

### Usando `requests.http` (com REST Client ou JetBrains)

#### POST /set-message

```http
POST http://localhost:3000/set-message
Content-Type: application/json

{
  "message": "Hello, Polygon! This is a test message."
}
```

**Resposta esperada:**

```json
{
  "message": "Message sent",
  "txHash": "0x..."
}
```

---

#### GET /get-message/:id

```http
GET http://localhost:3000/get-message/1
Content-Type: application/json
```

**Resposta esperada:**

```json
{
  "id": 1,
  "message": "Hello, Polygon! This is a test message."
}
```

---

#### GET /total-messages

```http
GET http://localhost:3000/total-messages
Content-Type: application/json
```

**Resposta esperada:**

```json
{
  "totalMessages": "1"
}
```

---

## 💻 Testes alternativos com `curl`

```bash
# Armazenar mensagem
curl -X POST http://localhost:3000/set-message -H "Content-Type: application/json" -d '{"message": "Hello, Polygon! This is a test message."}'

# Buscar mensagem por ID
curl http://localhost:3000/get-message/1

# Total de mensagens
curl http://localhost:3000/total-messages
```

---

## 🔐 Smart Contract (Solidity)

Arquivo: `contracts/MessageStorage.sol`  
Versão: `^0.8.2 <0.9.0`

### Principais funções:

```solidity
function setMessage(string memory _message) public returns (uint256);
function getMessage(uint256 _id) public view returns (string memory);
function getTotalMessages() public view returns (uint256);
```

---

## ⚠️ Considerações de Segurança

- **Chave privada:** Nunca commit no repositório! Use um sistema seguro (ex: AWS KMS) em produção.
- **Custo de Gas:** A função `setMessage` consome POL. Tenha saldo suficiente.
- **Controle de acesso:** Por padrão, qualquer um pode chamar `setMessage`. Para restringir, adicione `onlyOwner` (necessário reimplementar o contrato).
- **.env seguro:** Verifique se está no `.gitignore`.

---

## 🧪 Dicas adicionais

- **Testnet:** Use a **Polygon Amoy Testnet** antes da Mainnet para evitar custos.
- **Escalabilidade:** Para grandes volumes, considere armazenar a mensagem off-chain (ex: IPFS) e salvar apenas o hash on-chain.
- **Extensibilidade:** Pode adicionar autenticação ou novos endpoints (como busca por múltiplos IDs).

---

## ❓ Suporte

Abra uma issue no [repositório do GitHub](https://github.com/rnsribeiro/message-storage) ou entre em contato com o mantenedor.

---
