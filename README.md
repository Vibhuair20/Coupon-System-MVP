# Coupon System MVP

A robust coupon management system built with Go, designed for a medicine ordering platform.

## Features

- Coupon creation and management
- Coupon validation with multiple constraints
- Support for different usage types (one-time, multi-use, time-based)
- Medicine and category-based applicability
- Minimum order value enforcement
- Usage limits per user
- Percentage and fixed amount discounts
- Concurrent validation support
- PostgreSQL database for persistence
- Dockerized deployment

## Architecture

The system is built using:
- Go (Golang) for the backend
- Fiber as the web framework
- GORM for database operations
- PostgreSQL for data persistence
- Docker for containerization

### Key Components

1. **Models**
   - Coupon: Core coupon data structure
   - CouponUsage: Tracks coupon usage per user

2. **Handlers**
   - ValidateCoupon: Validates and applies coupons
   - GetApplicableCoupons: Returns applicable coupons for a cart

3. **Database**
   - PostgreSQL with GORM ORM
   - Automatic schema migration
   - Connection pooling

## Setup Instructions

1. **Prerequisites**
   - Docker and Docker Compose
   - Go 1.21 or later (for local development)

2. **Environment Variables**
   Create a `.env` file in the backend directory:
   ```
   DB_HOST=db
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=coupon_db
   DB_PORT=5432
   APP_PORT=8080
   ```

3. **Running with Docker**
   ```bash
   docker-compose up --build
   ```

4. **Local Development**
   ```bash
   cd backend
   go mod download
   go run main.go
   ```

## API Endpoints

### GET /coupons/applicable
Get all applicable coupons for a cart.

**Request:**
```json
{
  "cart_items": [
    { "id": "med_123", "category": "painkiller" }
  ],
  "order_total": 700,
  "timestamp": "2025-05-05T15:00:00Z"
}
```

**Response:**
```json
{
  "applicable_coupons": [
    {
      "coupon_code": "SAVE20",
      "discount_value": 20
    }
  ]
}
```

### POST /coupons/validate
Validate and apply a coupon.

**Request:**
```json
{
  "coupon_code": "SAVE20",
  "cart_items": [...],
  "order_total": 700,
  "timestamp": "2025-05-05T15:00:00Z",
  "user_id": "user123"
}
```

**Success Response:**
```json
{
  "is_valid": true,
  "discount": {
    "items_discount": 50,
    "charges_discount": 20
  },
  "message": "coupon applied successfully"
}
```

**Failure Response:**
```json
{
  "is_valid": false,
  "reason": "coupon expired or not applicable"
}
```

## Concurrency and Caching

- The system uses database-level locking for concurrent coupon validations
- GORM's transaction support ensures data consistency
- Coupon usage tracking is atomic and thread-safe

## Development Notes

- The codebase follows Go best practices and standard project layout
- Error handling is comprehensive throughout the application
- Logging is implemented for debugging and monitoring
- Database migrations are automatic on startup

## Future Improvements

1. Add Redis caching for frequently accessed coupons
2. Implement rate limiting
3. Add admin API for coupon management
4. Implement coupon analytics
5. Add more sophisticated discount calculations
6. Implement coupon code generation
7. Add bulk coupon operations
