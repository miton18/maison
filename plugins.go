package main

import (
	"path/filepath"
	"plugin"

	"github.com/miton18/maison/core"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const PLUGIN_STATE_OK = "plugin_state_ok"
const PLUGIN_STATE_KO = "plugin_state_ko"
const PLUGIN_STATE_INIT = "plugin_state_init"

type maisonPlugin struct {
	Name  string
	Init  func(*viper.Viper) error
	Run   func(maison *core.Maison)
	Close func()
	State string
}

var plugins = map[string]*maisonPlugin{}

func initPlugins() {
	for pName := range viper.GetStringMap("plugins") {
		log.Debugf("Loading " + pName)
		mp, err := fillPlugin(pName)
		if err != nil {
			log.Warn(err.Error())
			continue
		}
		plugins[mp.Name] = mp
	}

	for pluginName, p := range plugins {
		err := p.Init(viper.Sub("plugins." + pluginName))
		if err != nil {
			p.State = PLUGIN_STATE_KO
			log.Errorf("Plugin '%s' failed his initialization: %s", pluginName, err.Error())
			continue
		}
		p.State = PLUGIN_STATE_OK
	}
}

func closePlugins() {
	for pluginName := range plugins {
		delete(plugins, pluginName)
	}
}

func fillPlugin(name string) (*maisonPlugin, error) {
	prefix := viper.GetString("plugins-directory")
	mp := maisonPlugin{
		State: PLUGIN_STATE_INIT,
	}

	p, err := plugin.Open(filepath.Join(prefix, name+".so"))
	if err != nil {
		return nil, errors.Wrapf(err, "Cannot load plugin '%s'", name)
	}

	// Init
	nameS, err := p.Lookup("Name")
	if err != nil {
		return nil, errors.Wrapf(err, "Plugin '%s' don't export 'Name' constant", name)
	}
	pluginName, ok := nameS.(*string)
	if !ok {
		return nil, errors.Errorf("Plugin '%s' don't export 'Name' as a string", name)
	}
	if pluginName == nil {
		return nil, errors.Errorf("Plugin '%s''s 'Name' cannot be nil", name)
	}
	mp.Name = *pluginName

	// Init
	initS, err := p.Lookup("Init")
	if err != nil {
		return nil, errors.Wrapf(err, "Plugin '%s' don't export 'Init' method", name)
	}
	initFn, ok := initS.(func(*viper.Viper) error)
	if !ok {
		return nil, errors.Errorf("Plugin '%s' don't implement 'Init' as 'func(*viper.Viper) error'", name)
	}
	mp.Init = initFn

	// Run
	runS, err := p.Lookup("Run")
	if err != nil {
		return nil, errors.Wrapf(err, "Plugin '%s' don't export 'Run' method", name)
	}
	runFn, ok := runS.(func(maison *core.Maison))
	if !ok {
		return nil, errors.Errorf("Plugin '%s' don't implement 'Run' as 'func(maison *Maison)'", name)
	}
	mp.Run = runFn

	// Close
	closeS, err := p.Lookup("Close")
	if err != nil {
		return nil, errors.Wrapf(err, "Plugin '%s' don't export 'Close' method", name)
	}
	closeFn, ok := closeS.(func())
	if !ok {
		return nil, errors.Errorf("Plugin '%s' don't implement 'Run' as 'func(maison *Maison)'", name)
	}
	mp.Close = closeFn

	return &mp, nil
}
