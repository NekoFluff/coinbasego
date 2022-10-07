package coinbasego

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type WebSockerHandler interface {
	ReceiveHandler(conn *websocket.Conn)
}

type Channel struct {
	Name       string   `json:"name" binding:"required"`
	ProductIDs []string `json:"product_ids" binding:"required"`
}

type WebSocket struct {
	connection *websocket.Conn
	client     *Client
}

type WebsocketResponse struct {
	Type          string `json:"type"`
	Message       string `json:"message"`
	OrderID       string `json:"order_id"`
	OrderType     string `json:"order_type"`
	Size          string `json:"size"`
	RemainingSize string `json:"remaining_size"`
	Price         string `json:"price"`
	ClientOID     string `json:"client_oid"`
	Side          string `json:"side"`
	ProductID     string `json:"product_id"`
	Time          string `json:"time"`
	// Sequence      int  `json:"sequence"`
	ProfileID string `json:"profile_id"`
	UserID    string `json:"user_id"`
	Reason    string `json:"reason"`
}

type TickerMessage struct {
	WebsocketResponse
	// Sequence  int                `json:"sequence"`
	ProductID ProductID `json:"product_id"`
	Price     float64   `json:"price,string"`
	Open24H   float64   `json:"open_24h,string"`
	Volume24H float64   `json:"volume_24h,string"`
	Low24H    float64   `json:"low_24h,string"`
	High24H   float64   `json:"high_24h,string"`
	Volume30D float64   `json:"volume_30d,string"`
	BestBid   string    `json:"best_bid"`
	BestAsk   string    `json:"best_ask"`
	Side      string    `json:"side"`
	Time      time.Time `json:"time"`
	TradeID   int       `json:"trade_id"`
	LastSize  float64   `json:"last_size,string"`
}

func (client *Client) WebSocketCreate(channels []Channel, wsh WebSockerHandler) *WebSocket {
	socketUrl := "wss://ws-feed.exchange.coinbasego.com"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}

	ws := &WebSocket{
		connection: conn,
		client:     client,
	}
	ws.subscribe(channels)
	go wsh.ReceiveHandler(conn)
	return ws
}

type SubscribeRequest struct {
	Type       string    `json:"type" binding:"required"`
	Channels   []Channel `json:"channels" binding:"required"`
	Signature  string    `json:"signature"`
	Key        string    `json:"key"`
	Passphrase string    `json:"passphrase"`
	Timestamp  string    `json:"timestamp"`
}

func (ws *WebSocket) subscribe(channels []Channel) {
	timestamp := strconv.Itoa(int(time.Now().UTC().Unix()))
	signature, err := ws.client.accessSignValue(timestamp, Request{
		Method:  "GET",
		PathURL: "/users/self/verify",
		Body:    nil,
	})

	if err != nil {
		log.Fatal(err)
	}

	sr := SubscribeRequest{
		Type:       "subscribe",
		Channels:   channels,
		Signature:  signature,
		Key:        ws.client.credentials.ApiKey,
		Passphrase: ws.client.credentials.Passphrase,
		Timestamp:  timestamp,
	}

	jsonStr, err := json.Marshal(sr)
	if err != nil {
		log.Println("Error when decoding json subscription request:", err)
		return
	}

	if err = ws.connection.WriteMessage(websocket.TextMessage, jsonStr); err != nil {
		log.Println("Error during writing to websocket:", err)
		return
	}
}

// Close the websocket connection
func (ws *WebSocket) Close() error {
	// TODO: Do we need to check if the connection is already closed
	err := ws.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	return ws.connection.Close()
}
