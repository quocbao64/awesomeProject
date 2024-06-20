package constant

type ResponseStatus int
type Headers int

const (
	Success ResponseStatus = iota + 1
	DataNotFound
	InvalidRequest
	InternalServerError
	Unauthorized
	NoContent
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "INVALID_REQUEST", "INTERNAL_SERVER_ERROR", "UNAUTHORIZED", "NO_CONTENT"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data not found", "Invalid request", "Internal server error", "Unauthorized", "No content"}[r-1]
}
