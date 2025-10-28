package proto

type (
	DocumentURI string

	Position struct {
		Line      uint32 `json:"line"`
		Character uint32 `json:"character"`
	}

	Range struct {
		Start Position `json:"start"`
		End   Position `json:"end"`
	}

	URI string
)
