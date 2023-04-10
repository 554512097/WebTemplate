package params

type RegisterParam struct {
	Nick     string `json:"nick"`
	Age      uint8  `json:"age"`
	Phone    string `validate:"required" json:"phone"`
	Account  string `validate:"required,max=50" json:"account"`
	Password string `validate:"required,max=30" json:"password"`
}

type UserParam struct {
	Nick string
	Age  int
}
