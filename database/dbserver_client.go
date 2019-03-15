package database

type Client struct {
	Id       int32   `json:"id"`
	ClientId *string `json:"clientId"`
	Name     *string `json:"name"`
}

type Clients struct {
	Clients *[]Client `json:"clients"`
}

func (s *DBServer) GetAllClients() (*Clients, error) {
	cs := []Client{}
	rows, err := s.connPool.Query(`SELECT id, clientId, name FROM client Where status = $1`, "ACTIVE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Client
		if err = rows.Scan(&c.Id, &c.ClientId, &c.Name); err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	s.log.Debugf("DB. Clients %v", cs)
	ca := Clients{}
	ca.Clients = &cs
	return &ca, nil
}

func (s *DBServer) GetClient(id int) (*Client, error) {
	c := Client{}
	row := s.connPool.QueryRow(`SELECT id, clientId, name FROM client where id=$1`, id)
	err := row.Scan(&c.Id, &c.ClientId, &c.Name)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *DBServer) DeleteClient(id int) (interface{}, error) {
	row, err := s.connPool.Exec("Update client set status = 'DEACTIVE' where id = $1", id)
	if err != nil {
		return nil, err
	}
	return row, nil

}
