package serializer

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
