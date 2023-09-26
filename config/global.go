package config

func Debug(k *Koanf) bool {
	return k.Bool("debug")
}
