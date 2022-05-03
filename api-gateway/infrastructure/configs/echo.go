package configs

import "github.com/spf13/viper"

type EchoConfig struct {
	port         string
	jwtSecretKey string
}

func getEchoConfig() EchoConfig {
	viper.BindEnv("ASL_ECHO_JWT_SECRET")
	viper.BindEnv("ASL_ECHO_PORT")
	viper.SetDefault("ASL_ECHO_JWT_SECRET","---")
	viper.SetDefault("ASL_ECHO_PORT", "1444")
	echoConfig := EchoConfig{
		port:         viper.Get("ASL_ECHO_PORT").(string),
		jwtSecretKey: viper.Get("ASL_ECHO_JWT_SECRET").(string),
	}
	return echoConfig
}

func (echoConfig *EchoConfig) GetPort() string {
	return echoConfig.port
}

func (echoConfig *EchoConfig) GetJwtSecretKey() string {
	return echoConfig.jwtSecretKey
}
