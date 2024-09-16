package postgres

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/storage"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

type employeeStorage struct {
	conn *pgx.Conn
}

const checkUserIsExist = `SELECT EXISTS(SELECT 1 FROM employee WHERE username = $1)`

func (e *employeeStorage) CheckUserIsExistByUsername(username string) (bool, error) {
	var exists bool

	err := e.conn.QueryRow(checkUserIsExist, username).Scan(&exists)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

const checkUserOrganizationQuery = `SELECT EXISTS(SELECT 1 FROM organization_responsible 
    JOIN public.employee E ON E.id = organization_responsible.user_id WHERE username = $1 AND organization_id = $2)`

func (e *employeeStorage) CheckUserOrganization(username string, organizationId string) (bool, error) {
	var exists bool

	err := e.conn.QueryRow(checkUserOrganizationQuery, username, organizationId).Scan(&exists)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

const checkUserTenderQuery = `SELECT EXISTS(SELECT 1 FROM organization_responsible
                       JOIN public.employee E ON E.id = organization_responsible.user_id
                       JOIN public.tender ON organization_responsible.organization_id = tender.organization_id
                       WHERE username = $1 AND tender.id = $2)`

func (e *employeeStorage) CheckUserTender(username string, tenderId string) (bool, error) {
	var exists bool

	err := e.conn.QueryRow(checkUserTenderQuery, username, tenderId).Scan(&exists)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

const checkUserIsExistById = `SELECT EXISTS(SELECT 1 FROM employee WHERE id = $1)`

func (e *employeeStorage) CheckUserIsExistById(userId string) (bool, error) {
	var exists bool

	err := e.conn.QueryRow(checkUserIsExistById, userId).Scan(&exists)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func NewEmployeeStorage(url string) (storage.EmployeeStorage, error) {
	conf, err := pgx.ParseConnectionString(url)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse connection string")
	}

	conn, err := pgx.Connect(conf)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to connect to database")
	}

	return &employeeStorage{
		conn: conn,
	}, nil
}
