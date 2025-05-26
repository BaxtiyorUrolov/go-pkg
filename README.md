# Go-Pkg - Foydali Go Paketi

`go-pkg` - bu Go dasturlash tilida yozilgan, turli xil foydali funksiyalarni o'z ichiga olgan paket. Ushbu paket yordamida logging (jurnal yozish), JWT tokenlarni boshqarish va xavfsizlik bilan bog'liq operatsiyalarni osonlik bilan amalga oshirishingiz mumkin. Loyiha @BaxtiyorUrolov tomonidan ishlab chiqilgan.

## Paketning Tarkibi

- **logger**: Moslashuvchan logging tizimi (Zap logger asosida).
- **jwt**: JWT tokenlarini yaratish va tekshirish.
- **security**: Xavfsizlik bilan bog'liq funksiyalar (hash va tasodifiy sonlar yaratish).

## O'rnatish

Paketni o'rnatish uchun quyidagi buyruqni ishlatishingiz mumkin:

```bash
go get github.com/BaxtiyorUrolov/go-pkg
```

## Foydalanish

### 1. Loggerdan Foydalanish

Logger paketidan foydalanish uchun quyidagi misolni ko'rib chiqing:

```go
package main

import (
	"github.com/BaxtiyorUrolov/go-pkg/logger"
)

func main() {
	// Logger yaratish
	log := logger.New("info", "main")
	defer logger.Cleanup(log)

	// Log yozish
	log.Info("Dastur ishga tushdi", logger.String("status", "success"))
	log.Error("Xatolik yuz berdi", logger.String("error", "database connection failed"))
}
```

### 2. JWT Token Yaratish

JWT tokenlarini yaratish va tekshirish uchun quyidagi misolni ko'ring:

```go
package main

import (
	"fmt"
	"time"
	"github.com/BaxtiyorUrolov/go-pkg/jwt"
)

func main() {
	// Ma'lumotlarni tayyorlash
	claims := map[string]interface{}{
		"user_id": "12345",
		"role":    "admin",
	}
	signKey := []byte("maxfiy-kalit")

	// Token yaratish
	accessToken, refreshToken, err := jwt.GenerateJWT(claims, signKey, 1*time.Hour, 24*time.Hour)
	if err != nil {
		fmt.Println("Xatolik:", err)
		return
	}
	fmt.Println("Access Token:", accessToken)
	fmt.Println("Refresh Token:", refreshToken)

	// Tokenni tekshirish
	extractedClaims, err := jwt.ExtractClaims(accessToken, signKey)
	if err != nil {
		fmt.Println("Xatolik:", err)
		return
	}
	fmt.Println("Claims:", extractedClaims)
}
```
