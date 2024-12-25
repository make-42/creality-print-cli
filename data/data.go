package data

import (
	"creality-print-cli/config"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

type Err struct {
	Errcode int
	Key     int
}

type Data struct {
	TotalLayer           int
	AccelToDecelLimits   int
	AccelerationLimits   int
	AiDetection          int
	AiFirstFloor         int
	AiPausePrint         int
	AiSw                 int
	AutoLevelResult      string
	Autohome             string
	AuxiliaryFanPct      int
	BedTemp0             string
	BedTemp1             string
	BedTemp2             string
	BedTempAutoPid       int
	BoxTemp              int
	CaseFanPct           int
	Connect              int
	CornerVelocityLimits int
	CurFeedratePct       int
	CurFlowratePct       int
	CurPosition          string
	DProgress            int
	DeviceState          int
	EnableSelfTest       int
	Err                  Err
	Fan                  int
	FanAuxiliary         int
	FanCase              int
	Hostname             string
	Layer                int
	LightSw              int
	MaterialDetect       int
	MaterialStatus       int
	MaxBedTemp           int
	MaxNozzleTemp        int
	Model                string
	ModelFanPct          int
	ModelVersion         string
	NozzleMoveSnapshot   int
	NozzleTemp           string
	NozzleTempAutoPid    int
	PowerLoss            int
	PressureAdvance      string
	PrintFileName        string
	PrintId              string
	PrintJobTime         int // seconds
	PrintLeftTime        int // seconds
	PrintProgress        int
	PrintStartTime       int
	RealTimeFlow         string
	RealTimeSpeed        string
	RepoPlrStatus        int
	SmoothTime           string
	State                int
	TargetBedTemp0       int
	TargetBedTemp1       int
	TargetBedTemp2       int
	TargetNozzleTemp     int
	TfCard               int
	UpgradeStatus        int
	UsedMaterialLength   int
	VelocityLimits       int
	Video                int
	Video1               int
	VideoElapse          int
	VideoElapseFrame     int
	VideoElapseInterval  int
	WithSelfTest         int
}

var CurrentData Data

var C *websocket.Conn

func Init() {
	u := url.URL{Scheme: "ws", Host: config.Config.Address, Path: "/"}
	//log.Printf("connecting to %s", u.String())

	C, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		//log.Fatal("dial:", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := C.ReadMessage()
			if err != nil {
				//log.Println("read:", err)
				return
			}
			//log.Printf("recv: %s", message)
			err = json.Unmarshal([]byte(message), &CurrentData)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
}
