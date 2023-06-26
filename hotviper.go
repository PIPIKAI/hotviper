package hotviper

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func marshal(v *viper.Viper) []byte {
	st, _ := json.Marshal(v.AllSettings())
	return st
}

var TemplateName = "template"

type HotViper struct {
	viper    *viper.Viper
	template *viper.Viper
	cache    any
	fileName string
	filePath string
	fileType string
}

func NewHotViper(filename, fileType, configPath string) (*HotViper, error) {
	// add default template
	template := viper.New()
	template.SetConfigName(TemplateName)
	template.SetConfigType(fileType)
	template.AddConfigPath(configPath)
	err := template.ReadInConfig() // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	tmpViper := viper.New()
	tmpViper.SetConfigName(filename)
	tmpViper.SetConfigType(fileType)
	tmpViper.AddConfigPath(configPath)

	err = tmpViper.ReadInConfig() // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		// Not have config file now
		// create it
		err = tmpViper.ReadConfig(bytes.NewBuffer(marshal(template)))
		if err != nil {
			return nil, err
		}
		err = tmpViper.WriteConfigAs(configPath + "/" + filename + "." + fileType)
		if err != nil {
			return nil, err
		}
	}
	tmpViper.WatchConfig()
	return &HotViper{
		viper:    tmpViper,
		template: template,
		cache:    template.AllSettings(),
		fileName: filename,
		filePath: configPath,
		fileType: fileType,
	}, nil
}

// set config to template
func (hp *HotViper) SetDefault() error {
	if hp.template == nil {
		return fmt.Errorf("template is not inited")
	}
	temp := hp.viper.AllSettings()
	err := hp.viper.ReadConfig(bytes.NewBuffer(marshal(hp.template)))
	if err != nil {
		return err
	}
	hp.cache = temp
	return hp.viper.WriteConfig()
}

// get viper
func (hp *HotViper) GetViper() *viper.Viper {
	return hp.viper
}

// get config
func (hp *HotViper) GetConfig() map[string]interface{} {
	return hp.viper.AllSettings()
}

// set config
func (hp *HotViper) SetConfig(s string) error {

	err := hp.viper.WriteConfigAs(hp.filePath + "/" + hp.fileName + ".old." + hp.fileType)
	if err != nil {
		return err
	}
	hp.cache = hp.viper.AllSettings()

	err = hp.viper.ReadConfig(bytes.NewBuffer([]byte(s)))
	if err != nil {
		return err
	}
	return hp.viper.WriteConfig()
}

// add watch func
func (hp *HotViper) AddWatchFunc(fc func(in fsnotify.Event)) {
	hp.viper.OnConfigChange(fc)
}

// roll back config
func (hp *HotViper) RollBack() error {
	if hp.cache == nil {
		return fmt.Errorf("cache is nil")
	}
	v, err := json.Marshal(hp.cache)
	if err != nil {
		return err
	}
	hp.cache = hp.viper.AllSettings()
	err = hp.viper.ReadConfig(bytes.NewBuffer(v))
	if err != nil {
		return err
	}
	return hp.viper.WriteConfig()
}
