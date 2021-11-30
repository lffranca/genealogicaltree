package presenter

type Error struct {
	Code    *string `json:"code" yaml:"code" xml:"code"`
	Message *string `json:"message" yaml:"message" xml:"message"`
}
