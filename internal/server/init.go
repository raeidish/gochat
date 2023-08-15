package server

import(
    "net/http"
)

func StartServer(p string) error {
    //init chatServer
    chatServer := NewChatServer()
    
    //start chatServer event loops
    go chatServer.serverEventLoop()
    http.HandleFunc("/connect",chatServer.connect)

    //start file server
    fs := http.FileServer(http.Dir("./web"))
    http.Handle("/",http.StripPrefix("/",fs))

    //serve
    err := http.ListenAndServe(p,nil)

    if err != nil{
        return err
    }

    return nil
}
