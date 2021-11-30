package presenter

type ShortestPathURI struct {
	Name1 *string `uri:"name1" binding:"required"`
	Name2 *string `uri:"name2" binding:"required"`
}
