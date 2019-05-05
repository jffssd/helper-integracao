package libraries

type Cliente struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	/*
		Chave  string `json:"chave"`
		ERP    string `json:"erp"`
		WsURL  string `json:"ws_url"`
		WsUser string `json:"ws_user"`
		WsPass string `json:"ws_pass"`
	*/
}

type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}
