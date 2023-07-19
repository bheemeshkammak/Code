package daos

import (
	"database/sql"
	"errors"
	"github.com/bheemeshkammak/Code/apitesting/pkg/rest/server/daos/clients/sqls"
	"github.com/bheemeshkammak/Code/apitesting/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type ApiDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateApis(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS apis(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Name TEXT NOT NULL,
		Age TEXT NOT NULL,
		Verified INTEGER NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewApiDao() (*ApiDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateApis(sqlClient)
	if err != nil {
		return nil, err
	}
	return &ApiDao{
		sqlClient,
	}, nil
}

func (apiDao *ApiDao) CreateApi(m *models.Api) (*models.Api, error) {
	insertQuery := "INSERT INTO apis(Name, Age, Verified)values(?, ?, ?)"
	res, err := apiDao.sqlClient.DB.Exec(insertQuery, m.Name, m.Age, m.Verified)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("api created")
	return m, nil
}

func (apiDao *ApiDao) UpdateApi(id int64, m *models.Api) (*models.Api, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	api, err := apiDao.GetApi(id)
	if err != nil {
		return nil, err
	}
	if api == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE apis SET Name = ?, Age = ?, Verified = ? WHERE Id = ?"
	res, err := apiDao.sqlClient.DB.Exec(updateQuery, m.Name, m.Age, m.Verified, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("api updated")
	return m, nil
}

func (apiDao *ApiDao) DeleteApi(id int64) error {
	deleteQuery := "DELETE FROM apis WHERE Id = ?"
	res, err := apiDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("api deleted")
	return nil
}

func (apiDao *ApiDao) ListApis() ([]*models.Api, error) {
	selectQuery := "SELECT * FROM apis"
	rows, err := apiDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var apis []*models.Api
	for rows.Next() {
		m := models.Api{}
		if err = rows.Scan(&m.Id, &m.Name, &m.Age, &m.Verified); err != nil {
			return nil, err
		}
		apis = append(apis, &m)
	}
	if apis == nil {
		apis = []*models.Api{}
	}

	log.Debugf("api listed")
	return apis, nil
}

func (apiDao *ApiDao) GetApi(id int64) (*models.Api, error) {
	selectQuery := "SELECT * FROM apis WHERE Id = ?"
	row := apiDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Api{}
	if err := row.Scan(&m.Id, &m.Name, &m.Age, &m.Verified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("api retrieved")
	return &m, nil
}
