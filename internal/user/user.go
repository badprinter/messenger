package user

import "net"

type User struct {
	Username string
	tunnel   net.Conn
}

func NewUserWithParam(username string, tunnel net.Conn) *User {
	return &User{
		username,
		tunnel,
	}
}

func NewUser() *User {
	return &User{
		"",
		nil,
	}
}

func (u *User) Say(msg string) {
	u.tunnel.Write([]byte(msg))
}

func (u *User) GetIP() string {
	return u.tunnel.LocalAddr().String()
}

func (u *User) CloseTunnel() {
	u.tunnel.Close()
}
