package App

import "github.com/astaxie/beego/session"

var (
	SecretKey string
	Sessions  *session.Manager
)

func init() {
	SecretKey = "f3b462d93b24cb0538f5d864546bc3e0"
	SessionInit()
}

func SessionInit() {
	config := &session.ManagerConfig{
		CookieName:      "SESSION",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "",
	}
	Sessions, _ = session.NewManager("memory", config)
	go Sessions.GC()
}
