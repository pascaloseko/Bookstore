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

// GetLocations gets all available locations
func GetLocations() (locations []Location, err error) {
	rows, err := Db.Query("SELECT * FROM location")
	if err != nil {
		return
	}

	for rows.Next() {
		loc := Location{}
		if err = rows.Scan(&loc.ID, &loc.Name, &loc.Address, &loc.Country, &loc.OpenTime, &loc.CloseTime); err != nil {
			return
		}

		locations = append(locations, loc)
	}

	rows.Close()
	return
}

//GetLocation gets a single location instance
func GetLocation(id int) (loc Location, err error) {
	loc = Location{}

	err = Db.QueryRow("select * from location where id = $1", id).Scan(&loc.ID, &loc.Name, &loc.Address, &loc.Country, &loc.OpenTime, &loc.CloseTime)
	return
}
