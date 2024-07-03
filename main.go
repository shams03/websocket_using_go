package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader =websocket.Upgrader{
  ReadBufferSize:1024,
	WriteBufferSize:1024,
}

func wsHandler(w http.ResponseWriter,r *http.Request) {
  conn,err :=upgrader.Upgrade(w,r,nil)
 if err!=nil {
	log.Printf("connection error: %s",err);
	return
 }
 for{
	messageType,message,err:=conn.ReadMessage();
	if err!=nil {
		log.Printf("error reading message: %s",err);
		break;
 }
 message=[]byte("pong pong")
 conn.WriteMessage(messageType,message);
  
}
}

func main(){
	r:=gin.Default();
	r.LoadHTMLFiles("index.html")
	r.GET("/",func (c *gin.Context){
		c.HTML(http.StatusOK,"index.html",nil )
	})
	r.GET("/ws",func(c *gin.Context){
		wsHandler(c.Writer,c.Request)
	})
	r.Run("localhost:9000");


}
