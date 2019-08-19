package conf

import (
	"flag"
	"github.com/go-ini/ini"
	"go_gin_example/utils/common"
	"log"
	"path"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	PrefixUrl       string
	RuntimeRootPath string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExts  []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Server struct {
	RunMode      string
	HttpHost     string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var (
	AppConf      = new(App)
	ServiceConf  = new(Server)
	DatabaseConf = new(Database)
	RedisConf    = new(Redis)
)

var cfg *ini.File
var confPath string

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

func Init() {
	var err error
	if confPath == "" {
		confPath = path.Join(common.BaseDir, "conf/app.ini")

		log.Printf("no conf proivded, use default conf:%s\n", confPath)
	} else {
		log.Printf("conf path :%s", confPath)
	}

	if !common.IsFile(confPath) {
		log.Fatalf("invalid conf path: %s", confPath)
	}

	cfg, err = ini.Load(confPath)
	if err != nil {
		log.Fatalf("fail to parse conf file :%s ", confPath)
	}
	bindingTo("app", AppConf)
	bindingTo("server", ServiceConf)
	bindingTo("database", DatabaseConf)
	bindingTo("redis", RedisConf)

}

//binding .ini conf var to struct
func bindingTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("binding section %s to %s v err: %v", section, v, err)
	}
}

