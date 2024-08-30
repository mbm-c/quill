package ws

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

type Client struct {
	ID   string
	Conn *websocket.Conn
}

var clients = make(map[*websocket.Conn]string)
var mutex = &sync.Mutex{}

func GetByID(c *websocket.Conn) {
	log.Println(c.Locals("allowed"))  // true
	log.Println(c.Params("id"))       // 123
	log.Println(c.Query("v"))         // 1.0
	log.Println(c.Cookies("session")) // ""

	clientID := c.Params("id")
	mutex.Lock()
	clients[c] = clientID
	mutex.Unlock()

	defer func() {
		// Unregister the client
		mutex.Lock()
		delete(clients, c)
		mutex.Unlock()
		c.Close()
	}()

	var (
		mt  int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv from %s: %s", clientID, msg)

		// Broadcast the message to all other clients
		mutex.Lock()
		for conn := range clients {
			if conn != c {
				if err = conn.WriteMessage(mt, msg); err != nil {
					log.Println("write:", err)
					break
				}
			}
		}
		mutex.Unlock()
	}
}
