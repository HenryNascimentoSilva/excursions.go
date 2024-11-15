package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Excursion struct {
  gorm.Model

  Image string
  Title string
  Description string
  Buy string
  FindMore string
}

type ExcursionResponse struct {
  
  ID uint `json:"id"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
  DeletedAt time.Time `json:"deletedAt,omitempty"`
  Image string `json:"img"`
  Title string `json:"title"`
  Description string `json:"desc"`
  Buy string `json:"buy"`
  FindMore string `json:"findMore"`
}
