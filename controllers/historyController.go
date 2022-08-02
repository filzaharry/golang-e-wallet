package controllers

import (
	"strconv"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/kibo/e-wallet/database"
	"github.com/kibo/e-wallet/models"
)

func History(c *fiber.Ctx) error {
	// cookie := c.Cookies("jwt")

	// token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
	// 	return []byte(SecretKey), nil

	// })

	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }
	
	// claims := token.Claims.(*jwt.StandardClaims)
	id, _ := strconv.Atoi(c.Params("id"))

	history := []models.History{}
	database.DB.Find(&history, "user_id = ?", id)
	// database.DB.Find(&history, "user_id = ?", claims.Issuer)

	return c.JSON(fiber.Map{
		"code" : "200",
		"message": "success",
		"data" : history,
	})

}

func Topup(c *fiber.Ctx) error {
	var data map[string]string
	var transferDefault string = "0"
	var toDefault string = "0"
	// cookie := c.Cookies("jwt")

	// token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
	// 	return []byte(SecretKey), nil

	// })

	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }

	if err := c.BodyParser((&data)); err != nil {
		return err
	}

	
	// claims := token.Claims.(*jwt.StandardClaims)

	topup := models.History{
		Topup: data["topup"],
		Transfer: transferDefault,
		To: toDefault,
		UserId: c.Params("id"),
		// UserId: claims.Issuer,
	}

	database.DB.Create(&topup)
	
	return c.JSON(fiber.Map{
		"code" : "200",
		"message": "success",
		"data" : topup,
	})
}
func Transfer(c *fiber.Ctx) error {

	var data map[string]string
	var topupDefault string = "0"
	var transferDefault string = "0"

	// cookie := c.Cookies("jwt")

	// token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
	// 	return []byte(SecretKey), nil

	// })

	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }

	if err := c.BodyParser((&data)); err != nil {
		return err
	}
	
	// claims := token.Claims.(*jwt.StandardClaims)


	checkSaldo := models.CheckSaldo{}
    database.DB.Table("histories").Select("SUM(topup) as TotalTopup, SUM(transfer) as TotalTransfer, SUM(received) as TotalReceived").Where("user_id = ?", c.Params("id")).Scan(&checkSaldo)
    // database.DB.Table("histories").Select("SUM(topup) as TotalTopup, SUM(transfer) as TotalTransfer, SUM(received) as TotalReceived").Where("user_id = ?", claims.Issuer).Scan(&checkSaldo)


	totalTopup := checkSaldo.TotalTopup
	totalTransfer := checkSaldo.TotalTransfer
	totalReceived := checkSaldo.TotalReceived
	totalSaldo := totalTopup + totalReceived - totalTransfer

	dataTf := data["transfer"]
	dataTfInt, _ := strconv.Atoi(dataTf)

	if totalSaldo < dataTfInt {
		return c.JSON(fiber.Map{
			"code" : "500",
			"message": "saldo tidak cukup",
		})
	}

	var user models.User
	database.DB.Where("phone = ?", data["to"]).First(&user)

	if user.Id == 0 {
		return c.JSON(fiber.Map{
			"code" : "500",
			"message": "nomor belum teregistrasi",
		})
	}

	transfer := models.History{
		Topup: topupDefault,
		Transfer: data["transfer"],
		To: strconv.FormatUint(uint64(user.Id),10),
		UserId: c.Params("id"),
		// UserId: claims.Issuer,
	}	
	database.DB.Create(&transfer).Where("user_id = ?", c.Params("id"))
	// database.DB.Create(&transfer).Where("user_id = ?", claims.Issuer)

	received := models.History{
		Topup: topupDefault,
		Transfer: transferDefault,
		From: c.Params("id"),
		// From: claims.Issuer,
		Received: data["transfer"],
		UserId: strconv.FormatUint(uint64(user.Id),10),
	}
	database.DB.Create(&received).Where("phone = ?", data["to"])

	return c.JSON(fiber.Map{
		"code" : "200",
		"message": "success",
		"data" : transfer,
	})
}