package config

import (
	"flag"
	"time"

	"github.com/MAXXXIMUS-tropical-milkshake/drop-audio-streaming/internal/lib/logger"
)

type (
	Config struct {
		HTTP
		Log
		DB
		TLS
		Audio
		Auth
	}

	HTTP struct {
		GRPCPort           string
		GRPCClientRetries  uint
		GRPCClientTimeout  time.Duration
		GRPCUserClientAddr string
		HTTPPort           string
		ReadTimeout        int
	}

	Log struct {
		Level string
	}

	DB struct {
		URL           string
		MinioPassword string
		MinioUser     string
		MinioEndpoint string
		MinioBucket   string
		MinioUseSSL   bool
		MinioLocation string
		RedisAddr     string
		RedisPassword string
		RedisDB       int
	}

	Audio struct {
		ChunkSize    int
		UploadURLTTL int
		UserHistory  int
	}

	Auth struct {
		JWTSecret string
	}

	TLS struct {
		Cert string
		Key  string
	}
)

func NewConfig() (*Config, error) {
	gRPCPort := flag.String("grpc_port", "localhost:50010", "GRPC Port")
	logLevel := flag.String("log_level", string(logger.InfoLevel), "logger level")
	dbURL := flag.String("db_url", "", "url for connection to database")
	httpPort := flag.String("http_port", "localhost:8080", "HTTP Port")
	readTimeout := flag.Int("read_timeout", 5, "read timeout")

	// TLS
	cert := flag.String("cert", "", "path to cert file")
	key := flag.String("key", "", "path to key file")

	// Minio S3 storage
	minioPassword := flag.String("minio_password", "minioadmin", "minio password")
	minioUser := flag.String("minio_user", "minioadmin", "minio user")
	minioEndpoint := flag.String("minio_endpoint", "192.168.0.170:9000", "minio endpoint")
	minioBucket := flag.String("minio_bucket", "drop-audio", "minio bucket")
	minioUseSSL := flag.Bool("minio_use_ssl", false, "minio use ssl")
	minioLocation := flag.String("minio_location", "us-east-1", "minio location")

	// audio
	chunkSize := flag.Int("chunk_size", 1024, "chunk size")
	urlTTL := flag.Int("upload_url_ttl", 4320, "upload url ttl in minutes")

	// Redis
	redisAddr := flag.String("redis_addr", "localhost:6379", "redis address")
	redisPassword := flag.String("redis_password", "", "redis password")
	redisDB := flag.Int("redis_db", 0, "redis db")
	userHistory := flag.Int("user_history", 50, "max user history size")

	// auth
	jwtSecret := flag.String("jwt_secret", "secret", "jwt secret")

	// grpc client
	grpcClientRetries := flag.Uint("grpc_client_retries", 1, "grpc client retries")
	grpcClientTimeout := flag.Duration("grpc_client_timeout", 2*time.Second, "grpc client timeout")
	grpcUserClientAddr := flag.String("grpc_user_client_addr", "", "grpc user client addr")

	flag.Parse()

	cfg := &Config{
		HTTP: HTTP{
			GRPCPort:           *gRPCPort,
			GRPCClientRetries:  *grpcClientRetries,
			GRPCClientTimeout:  *grpcClientTimeout,
			GRPCUserClientAddr: *grpcUserClientAddr,
			HTTPPort:           *httpPort,
			ReadTimeout:        *readTimeout,
		},
		Log: Log{
			Level: *logLevel,
		},
		DB: DB{
			URL:           *dbURL,
			MinioPassword: *minioPassword,
			MinioUser:     *minioUser,
			MinioEndpoint: *minioEndpoint,
			MinioBucket:   *minioBucket,
			MinioUseSSL:   *minioUseSSL,
			MinioLocation: *minioLocation,
			RedisAddr:     *redisAddr,
			RedisPassword: *redisPassword,
			RedisDB:       *redisDB,
		},
		TLS: TLS{
			Cert: *cert,
			Key:  *key,
		},
		Audio: Audio{
			ChunkSize:    *chunkSize,
			UploadURLTTL: *urlTTL,
			UserHistory:  *userHistory,
		},
		Auth: Auth{
			JWTSecret: *jwtSecret,
		},
	}

	return cfg, nil
}
