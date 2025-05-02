package handlers

import (
	"crypto/rand"
	"ctf/models"
	"encoding/hex"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GenerateRandomToken(n int) (string, error) {
	n = n / 2
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Index_GET(c *fiber.Ctx) error {
	// Check if the user is logged in
	cookie := c.Cookies("ctftoken")
	if cookie == "5is4y89j7mp71evu295f29ym24f8ydb7zbzmfuumjz3byxgx10y9yad2zq7q1fqe" {
		return c.SendFile("./views/imagepage.html")
	}
	return c.SendFile("./views/index.html")
}

func Register_GET(c *fiber.Ctx) error {
	return c.SendFile("./views/register.html")
}

func Login_GET(c *fiber.Ctx) error {
	return c.SendFile("./views/login.html")
}

func Dashboard_GET(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	// write cookie to console
	if len(cookie) != 64 {
		return c.Redirect("/adminlogin", fiber.StatusSeeOther)
	}
	return c.SendFile("./views/dashboard.html")
}

func Upload_GET(c *fiber.Ctx) error {
	cookie := c.Cookies("admintoken")
	if cookie != "o9uo73t2jj241gds7urpr1hv4eapvoxb50wkmb53bld8b04pqdlso3ok3ftoawec" {
		return c.Redirect("/admin/dashboard", fiber.StatusSeeOther)
	}
	return c.SendFile("./views/upload.html")
}

func Flag_GET(c *fiber.Ctx) error {
	cookie := c.Cookies("ctftoken")
	if cookie != "5is4y89j7mp71evu295f29ym24f8ydb7zbzmfuumjz3byxgx10y9yad2zq7q1fqe" {
		return c.Redirect("/admin/dashboard", fiber.StatusSeeOther)
	}
	return c.SendFile("./views/flagpage.html")
}

func Image_GET(c *fiber.Ctx) error {
	cookie := c.Cookies("ctftoken")
	if cookie != "5is4y89j7mp71evu295f29ym24f8ydb7zbzmfuumjz3byxgx10y9yad2zq7q1fqe" {
		return c.Redirect("/admin/dashboard", fiber.StatusSeeOther)
	}
	return c.SendFile("./views/imagepage.html")
}

func Login_POST(c *fiber.Ctx) error {
	var loginDTO models.UserLoginDTO

	if err := c.BodyParser(&loginDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if loginDTO.Username != "admin" || loginDTO.Password != "verysecurepassword" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong username or password!"})
	}

	token, _ := GenerateRandomToken(64)

	// Set cookie
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(48 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	}
	c.Cookie(&cookie)

	return c.Redirect("/admin/dashboard", fiber.StatusSeeOther)
}

func Upload_POST(c *fiber.Ctx) error {
	// Yüklenen dosyanın ismini al
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString("File upload failed")
	}

	// Dosyayı kaydet
	filePath := "./uploads/" + file.Filename
	if err := c.SaveFile(file, filePath); err != nil {
		return err
	}

	// Go dosyasını çalıştırmak için komut oluştur
	if strings.HasSuffix(filePath, ".go") {
		cmd := exec.Command("go", "run", filePath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("Error executing Go code:", err)
			return c.Status(500).SendString("Error running Go code")
		}
		return c.SendString(fmt.Sprintf("Go code output:\n%s", string(output)))
	}

	return c.SendString("File uploaded and ran successfully!")
}

func Flag_POST(c *fiber.Ctx) error {
	var flagDTO models.FlagDTO

	cookie := c.Cookies("ctftoken")

	if cookie != "5is4y89j7mp71evu295f29ym24f8ydb7zbzmfuumjz3byxgx10y9yad2zq7q1fqe" {
		return c.Redirect("/admin/dashboard", fiber.StatusSeeOther)
	}

	if err := c.BodyParser(&flagDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"response": "Invalid request!"})
	}

	if flagDTO.Flag != "flag{N3_Mutlu_Turkum_Diy3n3}" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "Wrong!"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "Correct!"})
}
