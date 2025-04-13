package user

import "github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"

type GetParams struct {
	ID types.ID `json:"id"`
}

type GetResult struct {
	ID      types.ID `json:"id"`
	Message string   `json:"message"`
}
