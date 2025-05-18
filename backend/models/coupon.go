package models

import "time"

type UsageType string
type DiscountType string

const (
	OneTime    UsageType    = "one_time"
	MultiUse   UsageType    = "multi_use"
	TimeBased  UsageType    = "time_based"
	Percentage DiscountType = "percentage"
	Fixed      DiscountType = "fixed"
)

type Coupon struct {
	ID                    uint      `gorm:"primaryKey" json:"id"`
	CouponCode            string    `gorm:"uniqueIndex" json:"coupon_code"`
	ExpiryDate            time.Time `json:"expiry_date"`
	UsageType             UsageType `json:"usage_type"`
	ApplicableMedicineIDs []string  `gorm:"type:text[]" json:"applicable_medicine_ids"`
	ApplicableCategories  []string  `gorm:"type:text[]" json:"applicable_categories"`
	MinOrderValue         float64   `json:"min_order_value"`
	ValidTimeWindow       struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `gorm:"embedded" json:"valid_time_window"`
	TermsAndConditions string       `json:"terms_and_conditions"`
	DiscountType       DiscountType `json:"discount_type"`
	DiscountValue      float64      `json:"discount_value"`
	MaxUsagePerUser    int          `json:"max_usage_per_user"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

type CouponUsage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CouponID  uint      `json:"coupon_id"`
	UserID    string    `json:"user_id"`
	UsedAt    time.Time `json:"used_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartItem struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

type GetApplicableCouponsRequest struct {
	CartItems  []CartItem `json:"cart_item"`
	OrderTotal int        `json:"order_total"`
	Timestamp  time.Time  `json:"timestamp"`
}

type GetApplicableCouponsResponse struct {
	ApplicableCoupons []struct {
		CouponCode    string  `json:"coupon_code"`
		DsicountValue float64 `json:"discount_value"`
	} `json:"applicable_coupons"`
}

type ValidateCouponRequest struct {
	CouponCode string     `json:"coupon_code"`
	CartItems  []CartItem `json:"cart_items"`
	OrderTotal float64    `json:"order_total"`
	Timestamp  time.Time  `json:"timestamp"`
}

type ValidateCouponResponse struct {
	IsValid  bool `json:"is_valid"`
	Discount struct {
		ItemsDiscount   float64 `json:"items_discount"`
		ChargesDiscount float64 `json:"charges_discount"`
	} `json:"discount"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}
