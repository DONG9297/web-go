package model

type Result struct {
	Code int      `json:"status"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}
