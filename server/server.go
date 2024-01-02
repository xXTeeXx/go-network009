//แพคเกจหลักของโปรแกรม
package main

//นำแพคเกจเข้าใช้งาน
import (
	"fmt" //แพคเกจ fmt ใช้งานสำหรับการแสดงผลบนหน้าจอ
	"net" //แพคเกจ net ใช้งานสำหรับเว็บและเครือข่าย
)

func main() {
	// เริ่มต้นฟังก์ชันหลัก

	// สร้าง listener ที่รอรับการเชื่อมต่อที่ port 5000
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 5000")

	for {
		// รับการเชื่อมต่อจาก clients
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New connection established")

		// เริ่ม go routine เพื่อจัดการการเชื่อมต่อนี้
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// ฟังก์ชัน handleConnection จัดการกับการเชื่อมต่อที่เปิดขึ้น

	defer conn.Close()

	// สร้าง buffer เพื่อเก็บข้อมูลที่จะได้รับ
	buffer := make([]byte, 1024)

	for {
		// อ่านข้อมูลจาก client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// แสดงจำนวน byte ที่ได้รับ
		fmt.Printf("Received %d bytes\n", n)

		// แสดงข้อมูลที่ได้รับ
		fmt.Printf("Received message: %s", buffer[:n])

		// ส่งข้อความกลับไปยัง client
		response := "Message received successfully\n"
		conn.Write([]byte(response))
	}
}
