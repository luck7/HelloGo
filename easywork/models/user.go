package models

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}
type Userslice struct {
	Users []User
}

func (e *User) GetString() string {
	var s Userslice
	s.Users = append(s.Users, User{ServerName: "上海_VPN", ServerIP: "127.0.0.1"})
	s.Users = append(s.Users, User{ServerName: "北京_VPN", ServerIP:  "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	return string(b)
}