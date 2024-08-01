package ui

import "net"

const DEFAULT_ADDR = "localhost:16400"

func (model *Model) NewConnection(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	model.CloseConnection()
	model.Connection = conn

	model.NewIncomingBubble(false, "Connected successfully")

	return nil
}

func (model *Model) CloseConnection() {
	if model.Connection == nil {
		return
	}

	model.Connection.Close()
	model.Connection = nil
	model.NewIncomingBubble(false, "Connection closed")
}
