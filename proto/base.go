package proto

type (
	DocumentURI string

	Position struct {
		Line      uint32 `json:"line"`
		Character uint32 `json:"character"`
	}

	PositionEncoding string

	Range struct {
		Start Position `json:"start"`
		End   Position `json:"end"`
	}

	URI string
)

const (
	PositionEncodingNone  PositionEncoding = ""
	PositionEncodingUTF8  PositionEncoding = "utf-8"
	PositionEncodingUTF16 PositionEncoding = "utf-16"
	PositionEncodingUTF32 PositionEncoding = "utf-32"
)
