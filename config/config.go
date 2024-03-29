package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	JWT      JWTConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logger   LoggerConfig
	External ExternalConfig
}

type ServerConfig struct {
	Origin            string        `mapstructure:"ORIGIN"`
	Domain            string        `mapstructure:"DOMAIN"`
	AppVersion        string        `mapstructure:"APP_VERSION"`
	Port              string        `mapstructure:"PORT"`
	Mode              string        `mapstructure:"MODE"`
	ReadTimeout       time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout      time.Duration `mapstructure:"WRITE_TIMEOUT"`
	CtxDefaultTimeout time.Duration `mapstructure:"CTX_DEFAULT_TIMEOUT"`
	Debug             bool          `mapstructure:"DEBUG"`
}

type JWTConfig struct {
	JwtSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
	JwtIssuer     string `mapstructure:"JWT_ISSUER"`
	AccessExpMin  int    `mapstructure:"ACCESS_EXP_MIN"`
	RefreshExpMin int    `mapstructure:"REFRESH_EXP_MIN"`
}

type PostgresConfig struct {
	PostgresqlHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresqlUser     string `mapstructure:"POSTGRES_USER"`
	PostgresqlPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresqlDbname   string `mapstructure:"POSTGRES_DB"`
	PostgresqlSSLMode  bool   `mapstructure:"POSTGRES_SSL_MODE"`
	PgDriver           string `mapstructure:"POSTGRES_DRIVER"`
}

type RedisConfig struct {
	Address  string `mapstructure:"REDIS_ADDRESS"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int
}

type LoggerConfig struct {
	Development       bool   `mapstructure:"LOGGER_DEVELOPMENT"`
	DisableCaller     bool   `mapstructure:"LOGGER_DISABLE_CALLER"`
	DisableStacktrace bool   `mapstructure:"LOGGER_DISABLE_TRACE"`
	Encoding          string `mapstructure:"LOGGER_ENCODING"`
	Level             string `mapstructure:"LOGGER_LEVEL"`
}

type ExternalConfig struct {
	SlpURL             string `mapstructure:"SLP_URL"`
	SlpAPIKey          string `mapstructure:"SLP_API_KEY"`
	SlpMerchantCode    string `mapstructure:"SLP_MERCHANT_CODE"`
	SMTPHost           string `mapstructure:"SMTP_HOST"`
	SMTPPort           string `mapstructure:"SMTP_PORT"`
	SMTPPassword       string `mapstructure:"SMTP_PASSWORD"`
	SMTPFrom           string `mapstructure:"SMTP_FROM"`
	CloudinaryURL      string `mapstructure:"CLOUDINARY_URL"`
	OngkirAPIURL       string `mapstructure:"ONGKIR_API_URL"`
	OngkirAPIKey       string `mapstructure:"ONGKIR_API_KEY"`
	KodePosURL         string `mapstructure:"KODE_POS_URL"`
	GoogleClientID     string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleRedirectURL  string `mapstructure:"GOOGLE_OAUTH_REDIRECT_URL"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigFile(".env")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	if err := v.Unmarshal(&c.Server); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	if err := v.Unmarshal(&c.JWT); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	if err := v.Unmarshal(&c.Redis); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	if err := v.Unmarshal(&c.Postgres); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	if err := v.Unmarshal(&c.Logger); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	if err := v.Unmarshal(&c.External); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
