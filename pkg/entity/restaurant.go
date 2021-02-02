package entity

type Restaurant struct {
	Id          int     `db:"id"`
	OpenHourID  int     `db:"id_horario_abertura"`
	CloseHourID int     `db:"id_horario_fechamento"`
	CityID      int     `db:"id_cidade"`
	OpenDays    uint8   `db:"dias_funcionamento"`
	Name        string  `db:"nome"`
	Description *string `db:"descricao"`
	PhoneNumber string  `db:"telefone"`
	Address     string  `db:"endereco"`
}

func (r *Restaurant) IsEmpty() bool {
	return r == nil || r.Id == 0
}
