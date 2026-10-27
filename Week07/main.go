package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// โครงสร้างข้อมูลพัสดุ (ให้ตรงกับตารางใน MySQL)
type Parcel struct {
	ID        int    `json:"id"`
	Tracking  string `json:"tracking"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

var db *sql.DB

func main() {
	var err error
	// 1. เชื่อมต่อฐานข้อมูล MySQL (XAMPP User: root, Password: ว่างเปล่า)
	dsn := "root:@tcp(127.0.0.1:3306)/dorm_parcel_db"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("เชื่อมต่อ Database ไม่สำเร็จ:", err)
	}
	defer db.Close()

	app := fiber.New()

	// 2. อนุญาตให้ Vue.js (หน้าบ้าน) ยิง API เข้ามาหา Go ได้ (CORS)
	app.Use(cors.New())

	// 3. เส้น API (Routes)
	app.Get("/api/parcels", getParcels)          // ดึงพัสดุทั้งหมด
	app.Post("/api/parcels", addParcel)          // เพิ่มพัสดุใหม่
	app.Put("/api/parcels/:id", updateStatus)    // เปลี่ยนสถานะเป็นรับแล้ว
	app.Delete("/api/parcels/:id", deleteParcel) // ลบพัสดุ

	fmt.Println("🚀 Backend Go Fiber รันอยู่ที่ http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}

// Handler 1: ดึงรายการพัสดุทั้งหมดจาก MySQL
func getParcels(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, tracking, name, status, created_at FROM parcels ORDER BY id DESC")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	var parcels []Parcel
	for rows.Next() {
		var p Parcel
		if err := rows.Scan(&p.ID, &p.Tracking, &p.Name, &p.Status, &p.CreatedAt); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		parcels = append(parcels, p)
	}

	return c.JSON(parcels)
}

// Handler 2: บันทึกพัสดุใหม่ลง MySQL
func addParcel(c *fiber.Ctx) error {
	p := new(Parcel)
	if err := c.BodyParser(p); err != nil {
		return c.Status(400).SendString("ข้อมูลไม่ถูกต้อง")
	}

	query := "INSERT INTO parcels (tracking, name, status) VALUES (?, ?, 'รอรับ')"
	result, err := db.Exec(query, p.Tracking, p.Name)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)
	p.Status = "รอรับ"

	return c.JSON(p)
}

// Handler 3: อัปเดตสถานะพัสดุเป็น 'รับแล้ว'
func updateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "UPDATE parcels SET status = 'รับแล้ว' WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "อัปเดตสถานะเรียบร้อย"})
}

// Handler 4: ลบข้อมูลพัสดุ
func deleteParcel(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM parcels WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "ลบพัสดุเรียบร้อยแล้ว"})
}
