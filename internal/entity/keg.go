package entity

type Keg struct {
	Id       string  `json:"id"`
	DeviceId *string `json:"deviceId"`
	Sku      struct {
		Id          string `json:"id"`
		Size        string `json:"size"`
		Spear       string `json:"spear"`
		Description string `json:"description"`
	} `json:"sku"`
}
