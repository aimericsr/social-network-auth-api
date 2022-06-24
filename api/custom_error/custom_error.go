package custom_error

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

//generic responses
var (
	ErrInternalServer   = &ErrorResponse{StatusCode: 500, Message: "Internal server error", StatusText: "Please reache the developer team if this type of issue is happening often"}
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
)

// func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
// 	render.Status(r, e.StatusCode)
// 	return nil
// }

// func ErrorRenderer(err error) *ErrorResponse {
// 	return &ErrorResponse{
// 		Err:        err,
// 		StatusCode: 400,
// 		StatusText: "Bad request",
// 		Message:    err.Error(),
// 	}
// }

// func ServerErrorRenderer(err error) *ErrorResponse {
// 	return &ErrorResponse{
// 		Err:        err,
// 		StatusCode: 500,
// 		StatusText: "Internal server error",
// 		Message:    err.Error(),
// 	}
// }
