package presenter

type PersonNameURI struct {
	Name *string `uri:"name" binding:"required"`
}
