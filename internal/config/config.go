package config

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
}

func LoadConfig() *Config {
    return &Config{
        DBHost:     "localhost",
        DBPort:     "5432",
        DBUser:     "postgres",
        DBPassword: "postgres",
        DBName:     "notification_db",
        JWTSecret:  "your-secret-key",
    }
}