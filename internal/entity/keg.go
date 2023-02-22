package entity

import (
	"github.com/google/uuid"
	"net/mail"
	"time"
)

type Date struct {
	time.Time
}

type EmailAddress struct {
	mail.Address
}

type Organisation struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	IsContractBrewer bool      `json:"isContractBrewer,omitempty"`
	Market           string    `json:"market,omitempty"`
}

type Location struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	State   string    `json:"state"`
	Notes   string    `json:"notes"`
}

type Movement struct {
	Timestamp   time.Time   `json:"timestamp"`
	Geolocation Geolocation `json:"geolocation"`
}

type Geolocation struct {
	Lat              float64  `json:"lat"`
	Lng              float64  `json:"lng"`
	Accuracy         int      `json:"accuracy"`
	Provider         string   `json:"provider"`
	OriginalLat      float64  `json:"originalLat"`
	OriginalLng      float64  `json:"originalLng"`
	OriginalAccuracy int      `json:"originalAccuracy"`
	Pluscode         string   `json:"pluscode"`
	OriginalPluscode string   `json:"originalPluscode"`
	LocalTimeZone    TimeZone `json:"LocalTimeZone"`
}

type TimeZone struct {
	time.Location
}

type Cycle struct {
	Id    uuid.UUID `json:"id"`
	Order []struct {
		Id uuid.UUID `json:"id"`
	} `json:"order"`
	Producer []struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"producer"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type Scheduler struct {
	Id                uuid.UUID `json:"id"`
	Product           string    `json:"product"`
	BatchId           uuid.UUID `json:"batchId"`
	Keg               string    `json:"keg"`
	StartCleaningDate time.Time `json:"startCleaningDate"`
	EndCleaningDate   time.Time `json:"endCleaningDate"`
	CreatedBy         string    `json:"createdBy"`
	UpdatedBy         string    `json:"updatedBy"`
	DoctrineCreatedAt time.Time `json:"doctrineCreatedAt"`
	DoctrineUpdatedAt time.Time `json:"doctrineUpdatedAt"`
}

type KegDto struct {
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
	Location         struct {
		Organisation Organisation `json:"organisation"`
		Location     Location     `json:"location"`
	} `json:"location,omitempty"`
	Product struct {
		Id   string `json:"id""`
		Name string `json:"name""`
	} `json:"product,omitempty"`
	Batch struct {
		Id             uuid.UUID `json:"id"`
		productId      uuid.UUID `json:"productId"`
		Name           string    `json:"name"`
		BestBeforeDate Date      `json:"bestBeforeDate"`
		Kegs           struct {
			Id        uuid.UUID `json:"id"`
			CreatedBy string    `json:"createdBy"`
			AddedBy   struct {
				Id    uuid.UUID    `json:"id"`
				Email EmailAddress `json:"email"`
			} `json:"notifications"`
			FirstName     string `json:"firstName"`
			LastName      string `json:"lastName"`
			Phone         string `json:"phone"`
			JobTitle      string `json:"jobTitle"`
			Notes         string `json:"notes"`
			Notifications []struct {
				Notifications string `json:"notifications"`
			} `json:"notifications"`
			Roles []struct {
				Notes string `json:"notes"`
			} `json:"roles,omitempty"`
			Organisation Organisation `json:"organisation,omitempty"`
		} `json:"kegs,omitempty"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"batch,omitempty"`
	Scheduler          Scheduler `json:"scheduler,omitempty"`
	KegBatchSchedulers []struct {
		Id                uuid.UUID `json:"id"`
		StartCleaningDate time.Time `json:"StartCleaningDate"`
		EndCleaningDate   time.Time `json:"endCleaningDate"`
		CreatedBy         string    `json:"createdBy"`
		CreatedAt         time.Time `json:"createdAt"`
	} `json:"kegBatchSchedulers,omitempty"`
	Movement       []Movement `json:"movement,omitempty"`
	Cycle          []Cycle    `json:"cycle"`
	IsVenueWarning bool       `json:"isVenueWarning"`
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Format("2006-01-02")), nil
}

func (e *EmailAddress) MarshalJSON() ([]byte, error) {
	return []byte(e.String()), nil
}

func (tz *TimeZone) MarshalJSON() ([]byte, error) {
	return []byte(tz.String()), nil
}

func (tz *TimeZone) UnmarshalJSON(data []byte) error {
	return tz.UnmarshalJSON(data)
}
