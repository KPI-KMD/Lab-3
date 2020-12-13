package telemetrics

import (
	"database/sql"
)
type SendData struct {
 Name         string `json: "name"`
 Battery      string `json:"battery"`
 DeviceTime   string `json:"deviceTime"`
 CurrentVideo string `json:"currentVideo"`
}
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
							ON t.tablet_id = tb.id WHERE tablet_id = ($1) LIMIT 50`, idOfTheTablet);
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
	
	res.Telemetry = arrOfTelemetries;

	return res, nil
}

func (s *Store) sendData(sdata *SendData) error {
	var id int64
	row := s.Db.QueryRow(`SELECT id FROM tablets WHERE name = $1`, sdata.Name)
	err := row.Scan(&id)
	
	if err != sql.ErrNoRows && err != nil {
		return err
	} else if err == sql.ErrNoRows {
	 	_, err := s.Db.Exec(`INSERT INTO tablets ("name") VALUES ($1)`, sdata.Name)
	  	if err != nil {
	   		return err
	  	}
	  	row = s.Db.QueryRow(`SELECT id FROM tablets WHERE name = $1`, sdata.Name)
	}
 
	_, err = s.Db.Exec(`INSERT INTO telemetry
	   					("battery", "devicetime", "servertime", "currentvideo", "tablet_id")
	  					VALUES ($1, CURRENT_TIMESTAMP, $2, $3, $4)`,
						sdata.Battery, sdata.DeviceTime, sdata.CurrentVideo, id)
	return err
}