package postgres

import (
	"github.com/jackc/pgx"
	"github.com/pkg/errors"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/storage"
)

type tenderStorage struct {
	conn *pgx.Conn
}

func NewTenderStorage(url string) (storage.TenderStorage, error) {
	conf, err := pgx.ParseConnectionString(url)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse connection string")
	}
	conn, err := pgx.Connect(conf)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to connect to database")
	}
	return &tenderStorage{conn: conn}, nil
}

const createTenderVersionQuery = `INSERT INTO version (name, description, service_type, number, tender_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

const createTenderQuery = `INSERT INTO tender (organization_id) VALUES ($1) RETURNING id, created_at, status`

func (s *tenderStorage) CreateTender(tender domain.Tender) (domain.Tender, error) {
	tx, err := s.conn.Begin()
	defer tx.Rollback()

	if err != nil {
		return domain.Tender{}, errors.Wrap(err, "failed to begin transaction")
	}

	tender.Version = 1
	err = tx.QueryRow(createTenderQuery, tender.OrganizationId).Scan(&tender.Id, &tender.CreatedAt, &tender.Status)
	if err != nil {
		return domain.Tender{}, errors.Wrap(err, "failed to insert versionId")
	}

	err = tx.QueryRow(createTenderVersionQuery, tender.Name, tender.Description, tender.ServiceType, 1, tender.Id).Scan(&tender.VersionId)
	if err != nil {
		return domain.Tender{}, errors.Wrap(err, "failed to insert version")
	}

	err = tx.Commit()
	if err != nil {
		return domain.Tender{}, errors.Wrap(err, "failed to commit transaction")
	}

	return tender, nil
}

const getAllTendersQuery = `SELECT tender.id, name, description, service_type, status, number, created_at 
							FROM tender JOIN version ON tender.id = version.tender_id
							WHERE status='Published'`

func (s *tenderStorage) GetAllTenders() ([]domain.Tender, error) {
	var tenders []domain.Tender
	rows, err := s.conn.Query(getAllTendersQuery)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to execute query")
	}
	for rows.Next() {
		var tender domain.Tender
		err = rows.Scan(
			&tender.Id,
			&tender.Name,
			&tender.Description,
			&tender.ServiceType,
			&tender.Status,
			&tender.Version,
			&tender.CreatedAt,
		)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to scan row")
		}

		tenders = append(tenders, tender)
	}

	return tenders, nil
}

// const getTenderByFilterQuery = `SELECT tender.id, name, description, status, service_type, number, created_at
//								FROM tender JOIN version ON tender.id = version.tender_id
//								WHERE status='Published'
//								ORDER BY name LIMIT $1 OFFSET $2`

const getTenderByFilterQuery = `WITH max_version AS (
    SELECT tender_id, MAX(number) AS max_number
    FROM version
    GROUP BY tender_id
)
SELECT
    tender.id,
    version.name,
    version.description,
    tender.status,
    version.service_type,
    version.number,
    tender.created_at
FROM tender
         JOIN organization_responsible ON tender.organization_id = organization_responsible.organization_id
         JOIN version ON tender.id = version.tender_id
         JOIN max_version ON version.tender_id = max_version.tender_id 
WHERE status='Published'
ORDER BY name LIMIT $2 OFFSET $3`

func (s *tenderStorage) GetTendersByFilter(limit, offset int) ([]domain.Tender, error) {
	var tenders []domain.Tender
	rows, err := s.conn.Query(getTenderByFilterQuery, limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to execute query")
	}
	for rows.Next() {
		var tender domain.Tender
		err = rows.Scan(
			&tender.Id,
			&tender.Name,
			&tender.Description,
			&tender.Status,
			&tender.ServiceType,
			&tender.Version,
			&tender.CreatedAt,
		)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to scan row")
		}

		tenders = append(tenders, tender)
	}

	return tenders, nil
}

// const getTenderByOneServiceTypeQuery = `SELECT tender.id, name, description, status, service_type, number, created_at
//								FROM tender  JOIN version ON tender.id = version.tender_id
//                            	WHERE service_type = $1 AND status='Published'
//								ORDER BY name LIMIT $2 OFFSET $3`

const getTenderByOneServiceTypeQuery = `WITH max_version AS (
    SELECT tender_id, MAX(number) AS max_number
    FROM version
    GROUP BY tender_id
)
SELECT
    tender.id,
    version.name,
    version.description,
    tender.status,
    version.service_type,
    version.number,
    tender.created_at
FROM tender
         JOIN organization_responsible ON tender.organization_id = organization_responsible.organization_id
         JOIN version ON tender.id = version.tender_id
         JOIN max_version ON version.tender_id = max_version.tender_id 
WHERE service_type = $1 AND status='Published'
ORDER BY name LIMIT $2 OFFSET $3`

func (s *tenderStorage) GetTendersByOneServiceType(limit, offset int, serviceType []string) ([]domain.Tender, error) {
	var tenders []domain.Tender
	rows, err := s.conn.Query(getTenderByOneServiceTypeQuery, serviceType[0], limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to execute query")
	}
	for rows.Next() {
		var tender domain.Tender
		err = rows.Scan(
			&tender.Id,
			&tender.Name,
			&tender.Description,
			&tender.Status,
			&tender.ServiceType,
			&tender.Version,
			&tender.CreatedAt,
		)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to scan row")
		}

		tenders = append(tenders, tender)
	}

	return tenders, nil
}

// const getTenderByTwoServiceTypeQuery = `SELECT tender.id, name, description, status, service_type, number, created_at FROM tender  JOIN version ON tender.id = version.tender_id
//                            	WHERE (service_type = $1 OR service_type = $2) AND status='Published'
//								ORDER BY name LIMIT $3 OFFSET $4`

const getTenderByTwoServiceTypeQuery = `WITH max_version AS (
    SELECT tender_id, MAX(number) AS max_number
    FROM version
    GROUP BY tender_id
)
SELECT
    tender.id,
    version.name,
    version.description,
    tender.status,
    version.service_type,
    version.number,
    tender.created_at
FROM tender
         JOIN organization_responsible ON tender.organization_id = organization_responsible.organization_id
         JOIN version ON tender.id = version.tender_id
         JOIN max_version ON version.tender_id = max_version.tender_id 
WHERE (service_type = $1 OR service_type = $2) AND status='Published'
ORDER BY name LIMIT $3 OFFSET $4`

func (s *tenderStorage) GetTendersByTwoServiceType(limit, offset int, serviceType []string) ([]domain.Tender, error) {
	var tenders []domain.Tender
	rows, err := s.conn.Query(getTenderByTwoServiceTypeQuery, serviceType[0], serviceType[1], limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to execute query")
	}
	for rows.Next() {
		var tender domain.Tender
		err = rows.Scan(
			&tender.Id,
			&tender.Name,
			&tender.Description,
			&tender.Status,
			&tender.ServiceType,
			&tender.Version,
			&tender.CreatedAt,
		)

		if err != nil {
			return nil, errors.WithMessage(err, "failed to scan row")
		}

		tenders = append(tenders, tender)
	}

	return tenders, nil
}

// const getTenderByUsernameQuery = `SELECT tender.id, name, description, status, service_type, number, public.tender.created_at
//									FROM tender JOIN organization_responsible ON tender.organization_id = organization_responsible.organization_id
//									    JOIN employee ON employee.id = organization_responsible.user_id
//									    JOIN version ON tender.id = version.tender_id
//                           			WHERE employee.username = $1 AND number=(SELECT MAX(number) FROM version GROUP BY tender_id)
//										ORDER BY name LIMIT $2 OFFSET $3`

const getTenderByUsernameQuery = `WITH max_version AS (
    SELECT tender_id, MAX(number) AS max_number
    FROM version
    GROUP BY tender_id
)
SELECT
    tender.id,
    version.name,
    version.description,
    tender.status,
    version.service_type,
    version.number,
    tender.created_at
FROM tender
         JOIN organization_responsible ON tender.organization_id = organization_responsible.organization_id
         JOIN employee ON employee.id = organization_responsible.user_id
         JOIN version ON tender.id = version.tender_id
         JOIN max_version ON version.tender_id = max_version.tender_id ---AND version.number = max_version.max_number
         JOIN organization ON tender.organization_id = organization.id
WHERE employee.username = $1
ORDER BY name LIMIT $2 OFFSET $3`

func (s *tenderStorage) GetTenderByUsername(username string, limit, offset int) ([]domain.Tender, error) {
	var tenders []domain.Tender
	rows, err := s.conn.Query(getTenderByUsernameQuery, username, limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to execute query")
	}

	for rows.Next() {
		var tender domain.Tender
		err = rows.Scan(
			&tender.Id,
			&tender.Name,
			&tender.Description,
			&tender.Status,
			&tender.ServiceType,
			&tender.Version,
			&tender.CreatedAt,
		)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to scan row")
		}

		tenders = append(tenders, tender)
	}

	return tenders, nil
}

const getTenderByIdQuery = `SELECT tender.id, name, description, service_type, status, organization_id, version.id, number, created_at FROM tender  JOIN version ON tender.id = version.tender_id WHERE tender.id = $1 AND number=(SELECT MAX(number) FROM version WHERE tender_id =$1)`

func (s *tenderStorage) GetTenderById(id string) (domain.Tender, error) {
	var tender domain.Tender
	err := s.conn.QueryRow(getTenderByIdQuery, id).Scan(
		&tender.Id,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.OrganizationId,
		&tender.VersionId,
		&tender.Version,
		&tender.CreatedAt,
	)
	if err != nil {
		return domain.Tender{}, errors.WithMessage(err, "failed to scan row")
	}

	return tender, nil
}

const getStatusTenderByIdQuery = `SELECT status FROM tender WHERE id = $1`

func (s *tenderStorage) GetStatusTenderById(id string) (string, error) {
	var status string
	err := s.conn.QueryRow(getStatusTenderByIdQuery, id).Scan(&status)
	if err != nil {
		return "", errors.WithMessage(err, "failed to scan row")
	}

	return status, nil
}

const updateStatusTenderByIdQuery = `UPDATE tender SET status = $1 WHERE id = $2`

func (s *tenderStorage) UpdateStatusTenderById(tenderId string, status string) error {
	_, err := s.conn.Exec(updateStatusTenderByIdQuery, status, tenderId)
	if err != nil {
		return errors.WithMessage(err, "failed to scan row")
	}

	return nil
}

const updateTenderVersionByIdQuery = `INSERT INTO version  (name, description, service_type, number, tender_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

// const updateTenderByIdQuery = `UPDATE tender SET version_id = $1 WHERE id = $2`

func (s *tenderStorage) UpdateTenderById(tender domain.Tender) error {
	tx, err := s.conn.Begin()
	if err != nil {
		return errors.WithMessage(err, "failed to begin transaction")
	}
	defer tx.Rollback()

	err = tx.QueryRow(updateTenderVersionByIdQuery, tender.Name, tender.Description, tender.ServiceType, tender.Version, tender.Id).Scan(&tender.VersionId)
	if err != nil {
		return errors.WithMessage(err, "failed to update version")
	}

	// _, err = tx.Exec(updateTenderByIdQuery, tender.VersionId, tender.Id)
	// if err != nil {
	//	return errors.WithMessage(err, "failed to update version")
	// }

	err = tx.Commit()
	if err != nil {
		return errors.WithMessage(err, "failed to commit transaction")
	}

	return nil
}

const getTenderVersionByIdQuery = `SELECT id, name, description, service_type FROM version  WHERE tender_id=$1 AND number = $2`

func (s *tenderStorage) GetTenderVersion(tenderId string, version int) (domain.Tender, error) {
	var tender domain.Tender

	err := s.conn.QueryRow(getTenderVersionByIdQuery,
		tenderId,
		version,
	).Scan(
		&tender.VersionId,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
	)
	if err != nil {
		return domain.Tender{}, errors.WithMessage(err, "failed to scan row")
	}

	return tender, nil
}

const checkTenderIsExistQuery = `SELECT EXISTS (SELECT id FROM tender WHERE id = $1)`

func (s *tenderStorage) CheckTenderIsExist(tenderId string) (bool, error) {
	var exists bool
	err := s.conn.QueryRow(checkTenderIsExistQuery, tenderId).Scan(&exists)
	if err != nil {
		return false, errors.WithMessage(err, "failed to check exists")
	}

	return exists, nil
}

const checkVersionTenderIsExistQuery = `SELECT EXISTS (SELECT version FROM version  WHERE number = $1 AND tender_id = $2)`

func (s *tenderStorage) CheckVersionTenderIsExist(version int, tenderId string) (bool, error) {
	var exist bool
	err := s.conn.QueryRow(checkVersionTenderIsExistQuery, version, tenderId).Scan(&exist)
	if err != nil {
		return false, errors.WithMessage(err, "failed to check exists")
	}

	return exist, nil
}

const getTenderLastVersion = `SELECT MAX(number) FROM version WHERE tender_id=$1`

func (s *tenderStorage) GetTenderLastVersion(tenderId string) (int, error) {
	var version int
	err := s.conn.QueryRow(getTenderLastVersion, tenderId).Scan(&version)
	if err != nil {
		return 0, errors.WithMessage(err, "failed to scan row")
	}

	return version, nil
}
