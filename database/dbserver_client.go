package database

type Client struct {
	id       int32   `json:"id"`
	clientId int32   `json:"clientId"`
	name     *string `json:"name"`
}

func (s *DBServer) GetAllClients() (*[]Client, error) {
	cs := []Client{}
	rows, err := s.connPool.Query(`SELECT id, clientId, name FROM public.client `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Client
		if err = rows.Scan(&c.id, &c.clientId, &c.name); err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return &cs, nil
}

func (s *DBServer) GetClient(id int32) (*Client, error) {
	c := Client{}
	row := s.connPool.QueryRow(`SELECT id, clientId, name FROM public.client where id = ? `, id)
	err := row.Scan(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
