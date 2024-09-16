package response

type Error struct {
	Reason string `json:"reason"`
}

func NewError(reason string) *Error {
	return &Error{
		Reason: reason,
	}
}
