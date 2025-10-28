package proto

type (
	WorkspaceFolder struct {
		URI  URI    `json:"uri"`
		Name string `json:"name"`
	}
)
