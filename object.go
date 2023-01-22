package sccrawler

type Reservation struct {
	Id    string
	Name  string
	Tel   string
	Place string
	Date  string
}

type AligoApiData struct {
	Key         string `json:"key"`         // **인증용 API Key
	User_id     string `json:"user_id"`     // **사용자 id
	Sender      string `json:"sender"`      // **발신자 전화번호 (최대 16bytes)
	Receiver    string `json:"receiver"`    // **수신자 전화번호 - 컴마(,)분기 입력으로 최대 1천명
	Msg         string `json:"msg"`         // **메세지 내용
	Msg_type    string `json:"msg_type"`    // SMS(단문) , LMS(장문), MMS(그림문자) 구분
	Title       string `json:"title"`       // 문자제목(LMS,MMS만 허용)
	Destination string `json:"destination"` // %고객명% 치환용 입력
	Rdate       string `json:"rdate"`       // 예약일 (현재일이상)
	Rtime       string `json:"rtime"`       // 예약시간 - 현재시간기준 10분이후
	Testmode_yn string `json:"testmode_yn"` // 연동테스트시 Y 적용
}

type AligoResponse struct {
	Result_code string `json:"result_code"`
	Message     string `json:"message"`
}

// Instantiation SC reservation
func NewObject() []Reservation {
	r := []Reservation{}
	return r
}

// Instantiation Aligo SMS data
func NewApiData() AligoApiData {
	return AligoApiData{}
}

// Instantiation Aligo SMS response
func NewResponseData() AligoResponse {
	return AligoResponse{}
}
