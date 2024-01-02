package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// เริ่มต้นฟังก์ชันหลัก

	// เชื่อมต่อกับแม่เซิร์ฟเวอร์ที่รอรับที่ localhost:5000
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	// สร้าง reader สำหรับอ่านข้อมูลจากผู้ใช้
	reader := bufio.NewReader(os.Stdin)

	for {
		// อ่านข้อมูลที่ผู้ใช้ป้อน
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		// ตัดช่องว่างและตัวอักษร newline ออกจากข้อความ
		message = strings.TrimSpace(message)

		// ส่งข้อความไปยังแม่เซิร์ฟเวอร์
		conn.Write([]byte(message))

		// แสดงจำนวน byte ที่ถูกส่ง
		fmt.Printf("Sent %d bytes\n", len(message))

		// ตรวจสอบว่าผู้ใช้ต้องการออกจากโปรแกรมหรือไม่
		if message == "quit" {
			fmt.Println("Quitting the program.")
			break
		}

		// รับและแสดงผลตอบกลับจากแม่เซิร์ฟเวอร์
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Printf("Server response: %s", buffer[:n])
	}
}
