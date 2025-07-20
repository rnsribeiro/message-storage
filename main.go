package main

import (
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rnsribeiro/message-storage/message"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Configuration from .env
    alchemyURL := os.Getenv("ALCHEMY_URL")
    contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
    privateKeyHex := os.Getenv("PRIVATE_KEY")

    // Connect to Polygon Mainnet via Alchemy
    client, err := ethclient.Dial(alchemyURL)
    if err != nil {
        log.Fatalf("Failed to connect to Alchemy: %v", err)
    }
    defer client.Close()

    // Set up private key for signing transactions
    privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
    if err != nil {
        log.Fatalf("Failed to parse private key: %v", err)
    }

    // Instantiate the contract
    contract, err := message.NewMessageStorage(contractAddress, client)
    if err != nil {
        log.Fatalf("Failed to instantiate contract: %v", err)
    }

    // Set up Fiber
    app := fiber.New()

    // Route to write a message
    app.Post("/set-message", func(c *fiber.Ctx) error {
        type Request struct {
            Message string `json:"message"`
        }
        var req Request
        if err := c.BodyParser(&req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
        }

        if req.Message == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Message cannot be empty"})
        }

        // Set up transaction
        auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(137)) // Polygon Mainnet Chain ID
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transactor"})
        }

        // Send transaction to setMessage
        tx, err := contract.SetMessage(auth, req.Message)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"akers": err.Error()})
        }

        return c.JSON(fiber.Map{
            "message": "Message sent",
            "txHash":  tx.Hash().Hex(),
        })
    })

    // Route to read a message by ID
    app.Get("/get-message/:id", func(c *fiber.Ctx) error {
        id, err := c.ParamsInt("id")
        if err != nil || id <= 0 {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid message ID"})
        }

        message, err := contract.GetMessage(nil, big.NewInt(int64(id)))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }

        return c.JSON(fiber.Map{
            "id":      id,
            "message": message,
        })
    })

    // Route to get the total number of messages
    app.Get("/total-messages", func(c *fiber.Ctx) error {
        total, err := contract.GetTotalMessages(nil)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }

        return c.JSON(fiber.Map{
            "totalMessages": total.String(),
        })
    })

    // Start the server
    log.Fatal(app.Listen(":3000"))
}