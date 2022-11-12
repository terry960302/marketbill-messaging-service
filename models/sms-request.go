package models

type SmsRequest struct {
	Type            string  // "(SMS | LMS | MMS)",
	ContentType     *string //(COMM - 일반 | AD - 광고)"
	CountryCode     *string
	From            string
	Subject         *string // 제목
	Content         string
	Messages        []SmsMessage
	Files           *[]SmsFile
	ReserveTime     *string
	ReserveTimeZone *string
	ScheduleCode    *string
}

type SmsMessage struct {
	To      string
	Subject *string
	Content *string
}

type SmsFile struct {
	Name *string
	Body *string
}
