package models

import (
    "fmt"
    "github.com/henvo/golang-gin-gorm-starter/helper"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// DB is the database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
    u := helper.GetEnv("POSTGRES_USER", "metanmai")
    // p := helper.GetEnv("POSTGRES_PASSWORD", "")
    h := helper.GetEnv("POSTGRES_HOST", "localhost")
    n := helper.GetEnv("POSTGRES_DB", "hydra_db")
    port := helper.GetEnv("POSTGRES_PORT", "5432")
    sslmode := helper.GetEnv("DATABASE_SSLMODE", "disable")
    timezone := helper.GetEnv("DATABASE_TIMEZONE", "Asia/Kolkata")

    dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        h, u, n, port, sslmode, timezone,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Could not open database connection")
    }

    // Migrate the schema (add any models you want to migrate here).
    db.AutoMigrate(&User{})

    // Set the global DB variable
    DB = db
}
