package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

func main() {
	serverAddr := "104.16.85.20:443"
	path := "/Telegram-MARAMBASHI/?ed=2048"
	security := "tls"
	encryption := "none"
	host := "jd.wxgqlfx.top"
	typ := "ws"
	sni := "jd.wxgqlfx.top"
	uuid := "38a2070d-a256-48f3-891b-c9b68b01c8f7"

	// สร้าง URL สำหรับการเชื่อมต่อ
	u := url.URL{
		Scheme: "ws",
		Host:   serverAddr,
		Path:   path,
	}

	// เชื่อมต่อกับเซิร์ฟเวอร์
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการเชื่อมต่อ:", err)
		return
	}
	defer conn.Close()

	// ส่งข้อมูล HTTP สำหรับเชื่อมต่อ
	_, err = io.WriteString(conn, fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Key: %s\r\nSec-WebSocket-Version: 13\r\n\r\n", u.Path, u.Host, uuid))
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการส่งข้อมูล:", err)
		return
	}

	// อ่านข้อมูลที่ได้รับกลับจากเซิร์ฟเวอร์
	resp, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการอ่านข้อมูล:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("เชื่อมต่อกับเซิร์ฟเวอร์เรียบร้อยแล้ว}
