package main
import (
	"net/http"
	"github.com/gorilla/websocket"
	"shopping/main/impl"
	"time"
)

var (

	upgrader =  websocket.Upgrader{
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)


func wsHandler(w http.ResponseWriter, r *http.Request){
	var (
		wsConn *websocket.Conn
		err error
		data []byte
		conn *impl.Connection
	)	

	if wsConn,err = upgrader.Upgrade(w,r,nil);err !=nil{
		return
	}

	if conn,err = impl.InitConnection(wsConn);err != nil {
		goto ERR
	}


	// for{
	// 	 if _,data,err =  conn.ReadMessage();err != nil {
	// 		 goto ERR
	// 	 }
	// 	 if err = conn.WriteMessage(websocket.TextMessage,data);err != nil{
	// 		 goto  ERR
	// 	}
	// }

	go func(){

		var (
			err error
		)

		for{
			if err = conn.WriteMessage([]byte("heartbeat"));err !=nil{
				return 
			}
			time.Sleep(1 * time.Second)
		}
	}()


	for{
		if data,err = conn.ReadMessage();err !=nil{
			goto ERR
		}
		if err = conn.WriteMessage(data);err != nil {
			goto ERR 
		}
	}
	ERR:
		// conn.Close()
}


func main(){
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe("0.0.0.0:7777",nil)
}

