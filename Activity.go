```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

func main() {
	serverAddr := "104.16.85.20:443"
	path := "/Telegram-MARAMBASHI/?ed=8"
	security := "tls"
	encryption := "none"
	host "jd.wxgqlfx.top"
	typ := "ws"
	sni := "jdgqlfx.top"
	uuid := "38a2070d-a-48f3-891b-c9b68b01c87"

	// สร้าง URL สำหับการเชื่อมต่อ
	u := url.URL		Scheme: "ws",
		Host:   serverAddr,
		Path:   path	}

	// เชื่อมต่อกับเิร์ฟเวอร์
	conn, err := net("tcp", serverAddr)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดใการเชื่อมต่อ:", err)
		return
}
	defer conn.Close()

	// ส่งข้อมูล HTTP สำหรับเชื่อมต่อ
	_, err = io.WriteString(conn, fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Key: %s\r\nSec-WebSocket-Version: 13\r\n\r\n", u.Path, u.Host, uuid))
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดนการส่งข้อมูล:", err)
		return
		// อ่านข้อมูลที่ไ้รับกลับจากเซิร์เวอร์
	resp, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
	

OK