package config

const Version = "heplify-server 0.960"

var Setting HeplifyServer

type HeplifyServer struct {
	HEPAddr         string   `default:"0.0.0.0:9060"`
	HEPTCPAddr      string   `default:""`
	HEPTLSAddr      string   `default:"0.0.0.0:9060"`
	ESAddr          string   `default:""`
	ESDiscovery     bool     `default:"true"`
	MQDriver        string   `default:""`
	MQAddr          string   `default:""`
	MQTopic         string   `default:""`
	PromAddr        string   `default:""`
	PromTargetIP    string   `default:""`
	PromTargetName  string   `default:""`
	HoraclifixStats bool     `default:"false"`
	RTPAgentStats   bool     `default:"false"`
	DBShema         string   `default:"homer5"`
	DBDriver        string   `default:"mysql"`
	DBAddr          string   `default:"localhost:3306"`
	DBUser          string   `default:"root"`
	DBPass          string   `default:""`
	DBDataTable     string   `default:"homer_data"`
	DBConfTable     string   `default:"homer_configuration"`
	DBTableSpace    string   `default:""`
	DBBulk          int      `default:"200"`
	DBTimer         int      `default:"2"`
	DBRotate        bool     `default:"true"`
	DBPartLog       string   `default:"2h"`
	DBPartSip       string   `default:"1h"`
	DBPartQos       string   `default:"6h"`
	DBDropDays      int      `default:"0"`
	DBDropOnStart   bool     `default:"false"`
	Dedup           bool     `default:"false"`
	DiscardMethod   []string `default:""`
	DiscardProtoType []string `default:""`
	AlegIDs         []string `default:""`
	LogDbg          string   `default:""`
	LogLvl          string   `default:"info"`
	LogStd          bool     `default:"false"`
	Config          string   `default:"/root/heplify-server.toml"`
	Version         bool     `default:"false"`
	InsecurePort    int      `default:"80"`
	AdminPwd        string   `default:"test123"`
	SwaggerPath     string   `default:"/third_party/swagger-ui"`
	UIPath          string   `default:"/homer-ui/"`
}

func NewConfig() *HeplifyServer {
	return &HeplifyServer{
		HEPAddr:         "0.0.0.0:9060",
		HEPTCPAddr:      "0.0.0.0:9060",
		HEPTLSAddr:      "0.0.0.0:9060",
		ESAddr:          "",
		ESDiscovery:     true,
		MQDriver:        "",
		MQAddr:          "",
		MQTopic:         "",
		PromAddr:        "",
		PromTargetIP:    "",
		PromTargetName:  "",
		HoraclifixStats: false,
		RTPAgentStats:   false,
		DBShema:         "homer5",
		DBDriver:        "mysql",
		DBAddr:          "localhost:3306",
		DBUser:          "root",
		DBPass:          "",
		DBDataTable:     "homer_data",
		DBConfTable:     "homer_configuration",
		DBTableSpace:    "",
		DBBulk:          200,
		DBTimer:         2,
		DBRotate:        true,
		DBPartLog:       "6h",
		DBPartSip:       "2h",
		DBPartQos:       "12h",
		DBDropDays:      0,
		DBDropOnStart:   false,
		Dedup:           false,
		DiscardMethod:   nil,
		DiscardProtoType: nil,
		AlegIDs:         nil,
		LogDbg:          "",
		LogLvl:          "info",
		LogStd:          false,
		Config:          "./heplify-server.toml",
		Version:         false,
		InsecurePort:    80,
		AdminPwd:        "test123",
		SwaggerPath:     "/third_party/swagger-ui",
		UIPath:          "/homer-ui/",
		DropTableDays:   3,
	}
}

func Get() *HeplifyServer {
	return NewConfig()
}
