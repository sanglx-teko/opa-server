package dao

import "opatutorial/middleware/configurationmanager/model"


// ConfigurationManager ...
var ConfigurationManager model.IConfigurationManager

// InitCFManager ...
func InitCFManager(conf model.IConfigurationManager) {
	ConfigurationManager = conf
}