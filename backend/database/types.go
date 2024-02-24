package database

import "time"

type User struct {
	Id           uint   `json:"id"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MiddleName   string `json:"middle_name"`
	MaidenName   string `json:"maiden_name"`
	BirthDate    string `json:"birth_date"`
	CityId       uint   `json:"city_id"`
	BloodGroup   string `json:"blood_group"`
	HashPassword string `json:"hash_password"`
}
type DonationCreate struct {
	ID               uint `json:"id"`
	BloodStationID   int  `json:"blood_station_id,omitempty"`
	Image            string
	ImageID          int       `json:"image_id,omitempty"`
	CityID           int       `json:"city_id"`
	LegacyImage      string    `json:"legacy_image,omitempty"`
	HasReply         string    `json:"has_reply,omitempty"`
	ReplyViewed      string    `json:"reply_viewed,omitempty"`
	AllowedModify    string    `json:"allowed_modify,omitempty"`
	FirstName        string    `json:"first_name,omitempty"`
	LastName         string    `json:"last_name,omitempty"`
	MiddleName       string    `json:"middle_name,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Status           string    `json:"status"`
	RejectReason     string    `json:"reject_reason,omitempty"`
	DonateAt         string    `json:"donate_at,omitempty"`
	BloodClass       string    `json:"blood_class,omitempty"`
	PaymentType      string    `json:"payment_type,omitempty"`
	IsOut            bool      `json:"is_out,omitempty"`
	Volume           float64   `json:"volume,omitempty"`
	PaymentCost      float64   `json:"payment_cost,omitempty"`
	OnModerationDate string    `json:"on_moderation_date,omitempty"`
	WithImage        bool      `json:"with_image,omitempty"`
	CreatedUsingOCR  bool      `json:"created_using_ocr,omitempty"`
	User             int       `json:"user"`
}

type DonationDetail struct {
	ID                     uint `json:"id"`
	BloodStationID         int  `json:"blood_station_id,omitempty"`
	Image                  string
	ImageID                int       `json:"image_id,omitempty"`
	CityID                 int       `json:"city_id"`
	LegacyImage            string    `json:"legacy_image,omitempty"`
	HasReply               string    `json:"has_reply,omitempty"`
	ReplyViewed            string    `json:"reply_viewed,omitempty"`
	AllowedModify          string    `json:"allowed_modify,omitempty"`
	Feedback               string    `json:"feedback,omitempty"`
	FavoriteBloodStationID string    `json:"favorite_blood_station_id,omitempty"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	Status                 string    `json:"status"`
	RejectReason           string    `json:"reject_reason,omitempty"`
	DonateAt               string    `json:"donate_at,omitempty"`
	BloodClass             string    `json:"blood_class,omitempty"`
	PaymentType            string    `json:"payment_type,omitempty"`
	IsOut                  bool      `json:"is_out,omitempty"`
	Volume                 string    `json:"volume,omitempty"`
	PaymentCost            string    `json:"payment_cost,omitempty"`
	OnModerationDate       time.Time `json:"on_moderation_date,omitempty"`
	WithImage              bool      `json:"with_image,omitempty"`
	CreatedUsingOCR        bool      `json:"created_using_ocr,omitempty"`
	User                   string    `json:"user"`
}

type DonationPlan struct {
	ID           uint   `json:"id"`
	Event        string `json:"event"`
	BloodStation string
	City         string
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ObjectID     int       `json:"object_id,omitempty"`
	BloodClass   string    `json:"blood_class,omitempty"`
	PlanDate     string    `json:"plan_date"`
	PaymentType  string    `json:"payment_type,omitempty"`
	Status       string    `json:"status,omitempty"`
	IsOut        bool      `json:"is_out"`
	User         int       `json:"user"`
	ContentType  int       `json:"content_type,omitempty"`
	Donation     int       `json:"donation,omitempty"`
}

type DonorCard struct {
	BloodGroup                 string `json:"blood_group,omitempty"`
	Kell                       string `json:"kell,omitempty"`
	Blood                      bool   `json:"blood,omitempty"`
	Plasma                     bool   `json:"plasma,omitempty"`
	Platelets                  bool   `json:"platelets,omitempty"`
	Erythrocytes               bool   `json:"erythrocytes,omitempty"`
	Leukocytes                 bool   `json:"leukocytes,omitempty"`
	CantDonate                 bool   `json:"cant_donate,omitempty"`
	CantDonateDisqualification bool   `json:"cant_donate_disqualification,omitempty"`
	CantDonateSurgery          bool   `json:"cant_donate_surgery,omitempty"`
	CantDonateLactation        bool   `json:"cant_donate_lactation,omitempty"`
	CantDonateOther            bool   `json:"cant_donate_other,omitempty"`
	CantDonateDonor            bool   `json:"cant_donate_donor,omitempty"`
	CantDonateNotice           string `json:"cant_donate_notice,omitempty"`
	BoneDonor                  bool   `json:"bone_donor,omitempty"`
	DonationsBeforeReg         int    `json:"donations_before_reg,omitempty"`
	DonorStatus                int    `json:"donor_status,omitempty"`
}