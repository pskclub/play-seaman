package main

import (
	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/pam4/describe"
	"bitbucket.org/3dsinteractive/pam4/healthcheck"
	"bitbucket.org/3dsinteractive/seaman"
	"flag"
	"fmt"
	members2 "golang-guideline/members"
	products2 "golang-guideline/products"
	"strings"
)

type runner struct {
}

func newRunner() *runner {
	return &runner{}
}

func main() {
	flag.Parse()

	cfg := config.NewConfig()
	newRunner().start(cfg)
}

func (r *runner) environments(env string) []string {
	strs := strings.Split(env, ",")
	for i, str := range strs {
		strs[i] = strings.TrimSpace(str)
	}
	return strs
}

func (r *runner) start(cfg config.IConfig) {

	fmt.Printf("Start Seaman Microservices in %s environment\n", cfg.Env())

	s := seaman.NewSeaman(cfg)
	s.AddAllowDomains(cfg.AllowDomains())
	s.AddAllowHeaders(cfg.AllowHeaders())

	defer s.Close()

	s.RegisterService(products2.NewProductHTTP(cfg))
	s.RegisterService(members2.NewMemberHTTP(cfg))
	s.RegisterService(products2.NewProductMQ(cfg))

	// Register /healthz on every service
	s.RegisterService(healthcheck.NewHealthcheckHTTP(cfg))

	// Register /describe on every service
	s.RegisterService(describe.NewDescribeHTTP(cfg))

	// Start all services
	s.Start()
}
