package bundler

import "github.com/sanglx-teko/opa-server/middleware/configurationmanager/model"

// ConfigurationManager ...
var ConfigurationManager model.IConfigurationManager

// InitCFManager ...
func InitCFManager(conf model.IConfigurationManager) {
	ConfigurationManager = conf
}
