package config

// Config ...
type Config struct {
	GDrive struct {
		Credential struct {
			ClientSecret string
			Token        string
		}
		List struct {
			PageSize int64
		}
	}
}
