package errset

var ResponseOK = Response{Code: 0, Msg: "ok"}

// Errors
var (
	ErrInvalidParams    = NewResponseError(4001, "Invalid Parameters")
	ErrInvalidOperation = NewResponseError(4002, "Invalid Operation")
	ErrNoRow            = NewResponseError(4004, "Not found in DB")
	ErrInternalServer   = NewResponseError(4005, "Internal Server Error")
)
