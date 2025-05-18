package routes

import (
	"time"

	"github.com/gofiber/fiber"
)

type request struct {
}

type response struct {
	Coupon_Code string        `json:"coupon_code"`
	Expiry_Date time.Duration `json:"expiry_date"`
}

func GetApplicableCoupons(c *fiber.Ctx) error {

}
