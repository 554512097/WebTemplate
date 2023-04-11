package params

type UserParam struct {
	Nick     string `json:"nick"`
	Age      uint8  `json:"age"`
	Phone    string `binding:"required,phone" json:"phone"`
	Account  string `binding:"required,max=50" json:"account"`
	Password string `binding:"required,max=30" json:"password"`
}
