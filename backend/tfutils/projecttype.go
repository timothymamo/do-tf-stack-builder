package tfutils

type Project struct {
	Size    string   `json:"size" validate:"required"`
	State   string   `json:"state_file" validate:"required"`
	Envs    []string `json:"envs" validate:"required"`
	Modules []string `json:"modules" validate:"required"`
}
