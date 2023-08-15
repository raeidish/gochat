package server

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type chatServer struct {
    upgr websocket.Upgrader
    users []*user
    msgQ []*msg
}

func (s *chatServer) addUsr(u *user){
    s.users = append(s.users, u)
}

func (s *chatServer) removeUsr(u *user){
    for i,usr := range s.users{
        if usr.id == u.id{
            s.users = append(s.users[:i], s.users[i+1:]...)
        }
    } 
}

func (s *chatServer) addMsg(m []byte,u *user){
    msg := newMsg(m,u)
    s.msgQ = append(s.msgQ, msg)
}

func (s *chatServer) serverEventLoop(){
    for {
        l := len(s.msgQ)
        
        if(l > 0){
            //take part of q
            pmsgQ := s.msgQ[0:l]

            //send messages to all users
            for _,msg := range pmsgQ{
                //escape text
                esMsg := html.EscapeString(msg.content.Text)

                time := time.Now();
                tMsg := []byte(fmt.Sprintf(`<div hx-swap-oob="afterbegin:#chat_room"><p>%s: %s</p></div>`,time.Format("2006/01/02, 15:04:05"),esMsg))
                


                log.Println(fmt.Sprintf("user: %d, msg: %s",msg.sender.id,msg.content.Text))
                //format text
                

                //send to all
                for _,s := range s.users{
                    s.con.WriteMessage(websocket.TextMessage,tMsg)
                }
            }

            //trim msgQ
            s.msgQ = s.msgQ[l:]
        }
        
        time.Sleep(100 * time.Millisecond) 
    }
}

func (s *chatServer) connect(w http.ResponseWriter, r *http.Request){
    c,err := s.upgr.Upgrade(w,r,nil)

    if err != nil {
        log.Print(err)
        return
    }

    user := NewUser(c,"",uint32(rand.Int()))
    s.addUsr(user)

    go func(){
        for {
            _,msg,err := c.ReadMessage()
            if err != nil {
                s.removeUsr(user)
                log.Println(err)
                return
            }
            s.addMsg(msg,user)
        }
    }()
}

func NewChatServer() *chatServer{
    upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true}}
    return &chatServer{upgr:upgrader}
}
