package config

import (
	"os"
	"strings"
)

// IndexerConfig config for Indexer
type IndexerConfig struct {
	endpoints []string
	sniff     *bool
	gzip      *bool
	debug     *bool
}

// NewIndexerConfig return new IndexerConfig
func NewIndexerConfig() *IndexerConfig {
	return &IndexerConfig{}
}

// Endpoints return endpoints
func (c *IndexerConfig) Endpoints() []string {
	if c.endpoints == nil {
		pies := os.Getenv("PAM4_INDEXER_ENDPOINTS")
		ies := strings.Split(pies, ",")
		endpoints := make([]string, 0)
		for _, e := range ies {
			e = strings.TrimSpace(e)
			if len(e) == 0 {
				continue
			}
			endpoints = append(endpoints, e)
		}
		c.endpoints = endpoints
	}
	return c.endpoints
}

// Sniff return sniff enable
func (c *IndexerConfig) Sniff() bool {
	if c.sniff == nil {
		c.sniff = new(bool)
		*c.sniff = strings.ToLower(os.Getenv("PAM4_INDEXER_SNIFF")) == "true"
	}
	return *c.sniff
}

// Gzip return gzip enable
func (c *IndexerConfig) Gzip() bool {
	if c.gzip == nil {
		c.gzip = new(bool)
		*c.gzip = strings.ToLower(os.Getenv("PAM4_INDEXER_GZIP")) == "true"
	}
	return *c.gzip
}

// Debug return debug
func (c *IndexerConfig) Debug() bool {
	if c.debug == nil {
		c.debug = new(bool)
		*c.debug = strings.ToLower(os.Getenv("DEBUG")) == "true"
	}
	return *c.debug
}
