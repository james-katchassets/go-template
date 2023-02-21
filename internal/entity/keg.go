package entity

import "time"

type Date struct {
	time.Time
}

type Keg struct {
	Id       string  `json:"id"`
	DeviceId *string `json:"deviceId"`
	Sku      struct {
		Id          string `json:"id"`
		Size        string `json:"size"`
		Spear       string `json:"spear"`
		Description string `json:"description"`
	} `json:"sku"`
	Manufacturer     string `json:"manufacturer"`
	ManufacturerDate Date   `json:"manufacturerDate"`
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Format("2006-01-02")), nil
}
