package user

import "net"

type User struct {
	name  string
	tunel net.Conn
}
