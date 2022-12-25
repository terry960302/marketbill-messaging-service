package models

// SENS는 참고로 네이버 SMS 전송 서비스 명칭입니다.
type SensRequest struct {
	Type            string        `json:"type"`        // "(SMS | LMS | MMS)",
	ContentType     *string       `json:"contentType"` //(COMM - 일반 | AD - 광고)"
	CountryCode     *string       `json:"countryCode"`
	From            string        `json:"from"`
	Subject         *string       `json:"subject"` // 제목
	Content         string        `json:"content"`
	Messages        []SensMessage `json:"messages"`
	Files           *[]SensFile   `json:"files"`
	ReserveTime     *string       `json:"reserveTime"`
	ReserveTimeZone *string       `json:"reserveTimeZone"`
	ScheduleCode    *string       `json:"scheduleCode"`
}

type SensMessage struct {
	To      string  `json:"to"`
	Subject *string `json:"subject"`
	Content *string `json:"content"`
}

type SensFile struct {
	Name *string `json:"name"`
	Body *string `json:"body"`
}
