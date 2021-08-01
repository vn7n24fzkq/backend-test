package common

type CommonError struct {
	HttpCode      int    `json:"-"`
	Description   string `json:"description"`
	originalError error  `json:"-"`
}

func NewError(httpCode int, description string, err error) *CommonError {
	return &CommonError{
		HttpCode:      httpCode,
		Description:   description,
		originalError: err,
	}
}

func (b *CommonError) Error() string {
	return b.originalError.Error()
}
