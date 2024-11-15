package schemas

import ("gorm.io/gorm")

type Excursion struct {
  gorm.Model

  Image string
  Title string
  Description string
  Buy string
  FindMore string
}
