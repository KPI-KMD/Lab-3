package telemetrics

import (
	"database/sql"
)

type Telemetry struct {
	Battery string `json:"battery"`
	DeviceTime string `json:"devicetime"`
	ServerTime string `json: servertime`
	CurrentVideo *string `json: currentvideo`
}
type Tablet_Telemetry struct {
	Id int64 `json: "id"`
	Name string `json: name`
	Telemetry []*Telemetry `json: telemetry`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListTelemetries(idOfTheTablet int64) (Tablet_Telemetry, error) {
	res := Tablet_Telemetry{idOfTheTablet, "", nil};
	var arrOfTelemetries []*Telemetry
	
	rows, err := s.Db.Query(`SELECT tb.name, t.battery, t.devicetime, t.servertime, t.currentvideo 
							FROM telemetry AS t INNER JOIN tablets AS tb 
							ON t.tablet_id = tb.id WHERE tablet_id = ($1)`, idOfTheTablet);
	if err != nil {
		return res, err
	}
	
	defer rows.Close()
	
	for rows.Next() {
		var t Telemetry
		
		if err := rows.Scan(&res.Name, &t.Battery, &t.DeviceTime, &t.ServerTime, &t.CurrentVideo); err != nil {
			return res, err
		}
		
		arrOfTelemetries = append(arrOfTelemetries, &t)
	}
	if arrOfTelemetries == nil {
		arrOfTelemetries = make([]*Telemetry, 0)
	}
	if len(res.Name) == 0 {
		rows,_ := s.Db.Query("SELECT name FROM Tablets WHERE id = ($1)", idOfTheTablet);
		rows.Scan(&res.Name); 

	}
	
	res.Telemetry = arrOfTelemetries;

	return res, nil
}

/*func (s *Store) CreateChannel( Telemetry) error {
	if len(name) < 0 {
		return fmt.Errorf("channel name is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO channels (name) VALUES ($1)", name)
	return err
}*/