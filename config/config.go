package config

import (
	"bitbucket.org/3dsinteractive/seaman"
)

type Config struct {
	seaman.Config
	idxCfg seaman.IIndexerConfig
}

type IConfig interface {
	seaman.IConfig

	IndexerConfig() seaman.IIndexerConfig
}

func (c *Config) IndexerConfig() seaman.IIndexerConfig {
	if c.idxCfg == nil {
		c.idxCfg = NewIndexerConfig()
	}
	return c.idxCfg
}
