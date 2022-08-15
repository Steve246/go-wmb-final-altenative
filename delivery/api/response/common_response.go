package response

const (
	SuccesCode          = "00"
	SuccesMessage       = "Success"
	DefaultErrorCode    = "XX"
	DefaultErrorMessage = "Something Went Wrong"
)

type AppHttpResponse interface {
	Send()
}
