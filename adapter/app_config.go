package adapter

import (
	"github.com/codefluence-x/altair/core"
	"github.com/codefluence-x/altair/entity"
)

type (
	appConfig struct{ c entity.AppConfig }
	metric    struct{ m *entity.MetricConfig }
)

func AppConfig(c entity.AppConfig) core.AppConfig {
	return &appConfig{c: c}
}

func (a *appConfig) Port() int                           { return a.c.Port() }
func (a *appConfig) BasicAuthUsername() string           { return a.c.BasicAuthUsername() }
func (a *appConfig) BasicAuthPassword() string           { return a.c.BasicAuthPassword() }
func (a *appConfig) ProxyHost() string                   { return a.c.ProxyHost() }
func (a *appConfig) PluginExists(pluginName string) bool { return a.c.PluginExists(pluginName) }
func (a *appConfig) Plugins() []string                   { return a.c.Plugins() }
func (a *appConfig) Dump() string                        { return a.c.Dump() }
func (a *appConfig) Metric() core.MetricConfig           { return &metric{m: a.c.Metric()} }

func (m *metric) Interface() string { return m.m.Interface() }
