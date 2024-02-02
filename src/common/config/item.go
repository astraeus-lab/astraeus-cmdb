package config

// CoreConfig configure Astreaeus-CMDB core.
type CoreConfig struct {
	Web Web `json:"web" yaml:"web"`

	DB  DB  `json:"db" yaml:"db"`
	ES  ES  `json:"es" yaml:"es"`
	Log Log `json:"log" yaml:"log"`

	Redis Redis `json:"redis,omitempty" yaml:"redis,omitempty"`
}

type Web struct {
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}

type DB struct {
	Type   string `json:"type" yaml:"type"`
	Host   string `json:"host" yaml:"host"`
	User   string `json:"user" yaml:"user"`
	Passwd string `json:"passwd" yaml:"passwd"`
	DBName string `json:"dbName" yaml:"dbName"`

	Option DBOption `json:"option,omitempty" yaml:"option,omitempty"`
}

type DBOption struct {
	MaxOpenConns       int `json:"maxOpenConns,omitempty" yaml:"maxOpenConns,omitempty"`
	MaxIdleConns       int `json:"maxIdleConns,omitempty" yaml:"maxIdleConns,omitempty"`
	ConnMaxIdleTimeMin int `json:"connMaxIdleTimeMin,omitempty" yaml:"connMaxIdleTimeMin,omitempty"`
}

type ES struct {
	Endpoint []string `json:"endpoint" yaml:"endpoint"`
	User     string   `json:"user" yaml:"user"`
	Passwd   string   `json:"passwd" yaml:"passwd"`
}

type Log struct {
	Path   string `json:"path,omitempty" yaml:"path,omitempty"`
	Level  string `json:"level,omitempty" yaml:"level,omitempty"`
	Stdout bool   `json:"stdout,omitempty" yaml:"stdout,omitempty"`
}

type Redis struct {
	Enable   bool     `json:"enable" yaml:"enable"`
	Endpoint []string `json:"endpoint" yaml:"endpoint"`
	User     string   `json:"user" yaml:"user"`
	Passwd   string   `json:"passwd" yaml:"passwd"`

	Option RedisOption `json:"option,omitempty" yaml:"option,omitempty"`
}

type RedisOption struct {
	MaxOpenConns       int `json:"maxOpenConns,omitempty" yaml:"maxOpenConns,omitempty"`
	MaxIdleConns       int `json:"maxIdleConns,omitempty" yaml:"maxIdleConns,omitempty"`
	ConnMaxIdleTimeMin int `json:"connMaxIdleTimeMin,omitempty" yaml:"connMaxIdleTimeMin,omitempty"`
}

// completeConfig check the optional configuration is empty,
// fill in the default value if it is empty.
func (cc *CoreConfig) completeConfig() {
	if cc.DB.Option.MaxOpenConns == 0 {
		cc.DB.Option.MaxOpenConns = DefaultConfigDBMaxOpenConns
	}
	if cc.DB.Option.MaxIdleConns == 0 {
		cc.DB.Option.MaxIdleConns = DefaultConfigDBMaxIdleConns
	}
	if cc.DB.Option.ConnMaxIdleTimeMin == 0 {
		cc.DB.Option.ConnMaxIdleTimeMin = DefaultConfigDBConnMaxIdleTimeMin
	}

	if cc.Log.Path == "" {
		cc.Log.Path = DefaultConfigLogPath
	}
	if cc.Log.Level == "" {
		cc.Log.Level = DefaultConfigLogLevel
	}

	if cc.Redis.Enable {
		if cc.Redis.Option.MaxOpenConns == 0 {
			cc.Redis.Option.MaxOpenConns = DefaultConfigRedisMaxOpenConns
		}
		if cc.Redis.Option.MaxIdleConns == 0 {
			cc.Redis.Option.MaxIdleConns = DefaultConfigRedisMaxIdleConns
		}
		if cc.Redis.Option.ConnMaxIdleTimeMin == 0 {
			cc.Redis.Option.ConnMaxIdleTimeMin = DefaultConfigRedisConnMaxIdleTimeMin
		}
	}
}

// APIServerConfig configure Astreaeus-CMDB API Server.
type APIServerConfig struct {
}
