package config

type ConfigS struct {
	Address                    string
	UpdateUIEveryXMilliseconds int
	UIPaddingIndentAmount      int
}

var Config = ConfigS{
	Address:                    "192.168.1.41:9999",
	UpdateUIEveryXMilliseconds: 1000,
	UIPaddingIndentAmount:      2,
}
