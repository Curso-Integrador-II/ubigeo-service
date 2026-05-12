package model

type Ubigeo struct {
	ID       string `gorm:"column:id;primaryKey;size:6" json:"idUbigeo"`
	State    string `gorm:"column:departamento;size:100" json:"state"`
	Province string `gorm:"column:provincia;size:100" json:"province"`
	County   string `gorm:"column:distrito;size:100" json:"county"`
}

func (Ubigeo) TableName() string {
	return "ubigeo"
}
