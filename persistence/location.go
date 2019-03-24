package persistence

// AddLocation adds location
func (l *Location) AddLocation() (err error) {
	statement := "INSERT INTO location (name, address, country, open_time, close_time) VALUES($1, $2, $3, $4, $5) RETURNING id;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(l.Name, l.Address, l.Country, l.OpenTime, l.CloseTime).Scan(&l.ID)
	return
}
