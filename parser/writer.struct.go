package parser

type jHolderWriter struct {
	Hello   string `json:"hello"`
	Int     int64  `json:"integer"`
	String  string `json:"str"`
	Boolean bool   `json:"bool,omitempty"`
}
