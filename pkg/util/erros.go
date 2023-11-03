package util

type errMissingField struct {
	Field string
}

func ErrMissingField(field string) error {
	return errMissingField{field}
}

func (e errMissingField) Error() string {
	return "missing field: " + e.Field
}

type errInvalidField struct {
	Field string
}

func ErrInvalidField(field string) error {
	return errInvalidField{field}
}

func (e errInvalidField) Error() string {
	return "invalid field: " + e.Field
}
