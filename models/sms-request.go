package models

type SmsRequest struct {
	Type            string       `json:"type"`        // "(SMS | LMS | MMS)",
	ContentType     *string      `json:"contentType"` //(COMM - 일반 | AD - 광고)"
	CountryCode     *string      `json:"countryCode"`
	From            string       `json:"from"`
	Subject         *string      `json:"subject"` // 제목
	Content         string       `json:"content"`
	Messages        []SmsMessage `json:"messages"`
	Files           *[]SmsFile   `json:"files"`
	ReserveTime     *string      `json:"reserveTime"`
	ReserveTimeZone *string      `json:"reserveTimeZone"`
	ScheduleCode    *string      `json:"scheduleCode"`
}

type SmsMessage struct {
	To      string  `json:"to"`
	Subject *string `json:"subject"`
	Content *string `json:"content"`
}

type SmsFile struct {
	Name *string `json:"name"`
	Body *string `json:"body"`
}
