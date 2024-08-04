package config

var DefaultConfig = map[string]interface{}{
	"name":         "STREAM-RPC-BOILERPLATE",
	"port":         "127.0.0.1:8090",
	"rpc_travel":   "127.0.0.1:2201",
	"postgres_dsn": "host=localhost user=postgres password=password dbname=travel port=5432 sslmode=disable TimeZone=Asia/Jakarta",
	"log_level":    "DEBUG",
	"log_format":   "json",
	"secret":       "travel-api-main",
	"exp_auth":     168,
	"redis": map[string]interface{}{
		"server":   "localhost:6379",
		"timeout":  10,
		"authPass": "",
		"poolSize": 10,
	},
}
