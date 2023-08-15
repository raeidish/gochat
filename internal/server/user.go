package server

import(
    "github.com/gorilla/websocket"
)

type user struct{
    id uint32
    con *websocket.Conn
    name string
}

func NewUser(con *websocket.Conn,name string,id uint32) *user{
    return &user{id:id,name:name,con: con}
}
