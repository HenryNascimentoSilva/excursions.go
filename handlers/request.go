package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create Excursion
type CreateExcursionRequest struct {
	Image string `json:"img"`
  Title string `json:"title"`
  Description string `json:"desc"`
  Buy string `json:"buy"`
  FindMore string `json:"findMore"`	
}

func (c *CreateExcursionRequest) Validate() error {
	if c == nil {
		return fmt.Errorf("malformed request body")
	}
	if c.Image == "" {
		return errParamIsRequired("img", "string")
	}
	if c.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if c.Description == "" {
		return errParamIsRequired("desc", "string")
	}
	if c.Buy == "" {
		return errParamIsRequired("buy", "string")
	}
	if c.FindMore == "" {
		return errParamIsRequired("findMore", "string")
	}
	return nil
}