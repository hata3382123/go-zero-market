package mscoin_common

type BizCode int

const SuccessCode BizCode = 0

type Result struct {
	Code BizCode `json:"code"`
	Msg  string  `json:"msg"`
	Data any     `json:"data"`
}

func NewResult() *Result {
	return &Result{}
}
func (r *Result) Fail(code BizCode, msg string) {
	r.Code = code
	r.Msg = msg
}
func (r *Result) Success(data any) {
	r.Code = SuccessCode
	r.Msg = "success"
	r.Data = data
}
func (r *Result) Deal(data any, err error) *Result {
	if err != nil {
		r.Fail(-1, err.Error())
	} else {
		r.Success(data)
	}
	return r
}
