package global

import (
	"github.com/Vzttfromxduszu/golang-1.git/common/config"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
)
