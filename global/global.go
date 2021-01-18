package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go-chains/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
)
