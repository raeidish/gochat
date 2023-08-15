package server

import (
	"encoding/json"
)

type htmxHeader struct {
    HX_Request string `json:"HX-Request"`
    HX_Trigger string `json:"HX-Trigger"`
    HX_Trigger_Name string `json:"HX-Trigger-Name"`
    HX_Target string `json:"HX-Target"`
    HX_Current_Url string `json:"HX-Current-URL"`
}

type msgContent struct {
    Text string `json:"chat_message"`
    Header htmxHeader `json:"HEADERS"`
}

type msg struct{
    sender *user
    content msgContent
}

func newMsg(t []byte, s *user) *msg{
    content := msgContent{}
    json.Unmarshal(t,&content)
    return &msg{sender:s,content:content}
}
