package configs

type GlobalConfigs struct {
	echo       EchoConfig
}

func GetGlobalConfigs() GlobalConfigs {
	globalConfigs := GlobalConfigs{
		echo:       getEchoConfig(),
	}
	return globalConfigs
}

func (globalConfigs *GlobalConfigs) GetEchoConfig() EchoConfig {
	return globalConfigs.echo
}