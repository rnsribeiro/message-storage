Message Storage API
This project provides a Go-based REST API using the Fiber framework to interact with a MessageStorage smart contract deployed on the Polygon Mainnet. The contract allows storing messages with unique IDs, retrieving messages by ID, and checking the total number of stored messages. The API connects to the blockchain via Alchemy and uses a .env file for secure configuration.
The repository is structured to be clear and easy to use, with detailed instructions for setup, compilation, and testing. The API exposes three endpoints:

POST /set-message: Stores a new message in the contract.
GET /get-message/:id: Retrieves a message by its ID.
GET /total-messages: Returns the total number of messages stored.

Project Structure
message-storage/
├── message/
│ └── message_storage.go # Generated Go package for contract interaction
├── output/
│ └── MessageStorage_sol_MessageStorage.abi # Compiled ABI
├── .env # Environment variables (not committed)
├── .gitignore # Excludes sensitive files
├── go.mod # Go module dependencies
├── main.go # Fiber server implementation
├── MessageStorage.sol # Solidity smart contract
└── requests.http # HTTP requests for testing

Prerequisites
Before setting up the project, ensure you have the following:

Node.js: For compiling the Solidity contract with solcjs.
Go: Version 1.18 or higher for running the server and abigen.
Alchemy Account: An API key for the Polygon Mainnet (sign up at Alchemy).
Contract Address: The address of the deployed MessageStorage contract on Polygon Mainnet.
Private Key: A Polygon wallet’s private key with sufficient MATIC for gas fees.
Git: For cloning and managing the repository.
REST Client: A tool like VS Code’s REST Client extension or JetBrains IDEs to run requests.http.

Setup Instructions
Follow these steps to set up and run the project.

1. Clone the Repository
   Clone the repository to your local machine:
   git clone https://github.com/rnsribeiro/message-storage.git
   cd message-storage

2. Compile the Smart Contract
   The MessageStorage.sol contract must be compiled to generate its ABI for Go integration.

Install solc:Install the Solidity compiler globally:
npm install -g solc

Compile the Contract:The contract is located in MessageStorage.sol. Compile it to generate the ABI:
mkdir output
solcjs MessageStorage.sol --abi -o output

This creates output/MessageStorage_sol_MessageStorage.abi.

3. Generate the Go Package with abigen
   Use abigen to create a Go package for interacting with the contract.

Install abigen:
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

Generate the Go Package:Run the following command to create the Go package in the message directory:
mkdir message
abigen --abi output/MessageStorage_sol_MessageStorage.abi --pkg message --type MessageStorage --out message/message_storage.go

This generates message/message_storage.go, which contains the contract’s Go bindings.

4. Initialize the Go Project
   Set up the Go module and install dependencies.

Initialize the Go Module:
go mod init github.com/rnsribeiro/message-storage

Install Dependencies:Install the required Go packages:
go get github.com/gofiber/fiber/v2
go get github.com/ethereum/go-ethereum
go get github.com/joho/godotenv

Set Up the .env File:Create a .env file in the project root with the following content:
ALCHEMY_API_KEY=your_alchemy_api_key
CONTRACT_ADDRESS=0xYourContractAddress
PRIVATE_KEY=your_private_key

Replace your_alchemy_api_key with your Alchemy API key for Polygon Mainnet.
Replace 0xYourContractAddress with the deployed contract’s address.
Replace your_private_key with your wallet’s private key (without 0x prefix).
Security Note: Ensure .env is listed in .gitignore to prevent committing sensitive data:echo ".env" >> .gitignore

5. Run the Server
   The main.go file implements the Fiber server, which exposes the API endpoints.

Verify Configuration:Ensure the .env file is correctly configured and message/message_storage.go exists.

Run the Server:
go mod tidy
go run main.go

The server will start at http://localhost:3000.

6. Test the API
   The requests.http file contains sample HTTP requests to test the API. Use a tool like VS Code’s REST Client or a JetBrains IDE to execute them.

Install REST Client:

For VS Code, install the REST Client extension.
For JetBrains IDEs, use the built-in HTTP Client.

Test Requests:The requests.http file includes:

POST /set-message: Stores a new message.
POST http://localhost:3000/set-message
Content-Type: application/json

{
"message": "Hello, Polygon! This is a test message."
}

Expected response:
{
"message": "Message sent",
"txHash": "0x..."
}

GET /get-message/:id: Retrieves a message by ID (e.g., ID 1).
GET http://localhost:3000/get-message/1
Content-Type: application/json

Expected response:
{
"id": 1,
"message": "Hello, Polygon! This is a test message."
}

GET /total-messages: Returns the total number of messages.
GET http://localhost:3000/total-messages
Content-Type: application/json

Expected response:
{
"totalMessages": "1"
}

Alternative Testing with curl:

# Write a message

curl -X POST http://localhost:3000/set-message -H "Content-Type: application/json" -d '{"message": "Hello, Polygon! This is a test message."}'

# Get a message by ID

curl http://localhost:3000/get-message/1

# Get total messages

curl http://localhost:3000/total-messages

Smart Contract Details
The MessageStorage contract (MessageStorage.sol) is a simple Solidity contract with the following functions:

setMessage(string memory \_message): Stores a message and returns its ID.
getMessage(uint256 \_id): Retrieves a message by its ID.
getTotalMessages(): Returns the total number of stored messages.

The contract is compiled with Solidity version >=0.8.2 <0.9.0 and deployed on the Polygon Mainnet.
Security Considerations

Private Key: Store the private key in the .env file and never commit it to version control. Consider using a key management system (e.g., AWS KMS) for production.
Gas Costs: The setMessage function requires MATIC for gas. Ensure the wallet has sufficient funds.
Access Control: The contract allows anyone to call setMessage. Add an onlyOwner modifier if you want to restrict access (requires redeploying the contract).
Environment Variables: The .env file is excluded via .gitignore to prevent exposing sensitive data.

Troubleshooting

Connection Errors: Verify the ALCHEMY_API_KEY and ensure the URL is https://polygon-mainnet.g.alchemy.com/v2/<API_KEY>.
Invalid Message ID: Ensure the ID used in GET /get-message/:id corresponds to a stored message.
Transaction Failures: Check that the wallet has enough MATIC and the CONTRACT_ADDRESS is correct.
Server Errors: Ensure message/message_storage.go exists and the .env file is properly configured.

Additional Notes

Testing on Testnet: Before using the Polygon Mainnet, test on the Polygon Amoy Testnet to avoid gas costs.
Scalability: Storing many messages on-chain can be expensive. Consider off-chain storage (e.g., IPFS) with on-chain hashes for large-scale use.
Extending the API: Add authentication or additional endpoints as needed (e.g., for batch message retrieval).

For further assistance, open an issue on the GitHub repository or contact the maintainer.
