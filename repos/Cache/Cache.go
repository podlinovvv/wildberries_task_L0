package Cache

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"wb_l0/repos/database"
)

type Storage struct {
	Db *sql.DB
	C  map[string][]byte
}

func NewStorage() *Storage {
	db := database.ConnectToDb()
	cache := loadCacheFromDB(db)

	return &Storage{
		Db: db,
		C:  cache,
	}
}

func loadCacheFromDB(db *sql.DB) map[string][]byte {
	cache := make(map[string][]byte)
	rows, _ := db.Query("select * from orders")

	for rows.Next() {
		var id string
		var data []byte
		rows.Scan(&id, &data)
		cache[id] = data
	}
	return cache
}

//пишем заказ в кэш и бд
func (s *Storage) WriteOrder(id string, order []byte) error {
	_, err := s.Db.Exec(
		"insert into orders (id, data) "+
			"values ($1, $2) ",
		id, order)

	if err != nil {
		return err
	}

	s.C[id] = order
	return nil
}

//получаем заказ из кэша по id, если его нет в кэше, подгружаем из БД
func (s *Storage) ReadOrderById(id string) ([]byte, error) {
	order, ok := s.C[id]

	if ok {
		fmt.Println("из кэша")
		return order, nil
	}
	fmt.Println("не из кэша")
	err := s.Db.QueryRow(
		"select data from orders "+
			"where id = $1 ",
		id).Scan(&order)

	if err != nil {
		return nil, err
	}

	s.C[id] = order
	return order, nil
}

/*func (s *Storage) restoreCacheFromDB() {
	s.db
	var orders []structs.Order

	for _, order := range orders {
		s.WriteOrder(&order)
	}
}
*/
