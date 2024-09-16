package postgres

import (
	"github.com/jackc/pgx"
	"github.com/pkg/errors"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/storage"
)

type bidStorage struct {
	conn *pgx.Conn
}

const checkBidIsExistQuery = `SELECT EXISTS (SELECT * FROM bid WHERE id = $1)`

func (b bidStorage) CheckBidIsExist(bidId string) (bool, error) {
	var exists bool
	err := b.conn.QueryRow(checkBidIsExistQuery, bidId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

const createBidVersionQuery = `INSERT INTO bid_version (name, description, number, bid_id) VALUES ($1, $2, $3, $4) RETURNING id, number`

const createBidQuery = `INSERT INTO bid (tender_id, author_type, author_id) VALUES ($1, $2, $3) RETURNING id, created_at`

func (b bidStorage) CreateBid(bid domain.Bid) (domain.Bid, error) {
	tx, err := b.conn.Begin()
	if err != nil {
		return domain.Bid{}, errors.Wrap(err, "could not begin transaction")
	}
	defer tx.Rollback()

	err = tx.QueryRow(createBidQuery, bid.TenderId, bid.AuthorType, bid.AuthorId).Scan(&bid.ID, &bid.CreatedAt)
	if err != nil {
		return domain.Bid{}, errors.Wrap(err, "could not insert bid version")
	}

	err = tx.QueryRow(createBidVersionQuery, bid.Name, bid.Description, 1, bid.ID).Scan(&bid.VersionId, &bid.Version)
	if err != nil {
		return domain.Bid{}, errors.Wrap(err, "could not insert bid version")
	}

	err = tx.Commit()
	if err != nil {
		return domain.Bid{}, errors.Wrap(err, "could not commit transaction")
	}

	return bid, nil
}

// const getBidsByFilterQuery = `SELECT bid.id, bid_version.id, status, tender_id, author_type, author_id, bid.created_at, name, description, number FROM bid
//    JOIN bid_version ON bid.id = bid_version.bid_id
//    JOIN employee ON bid.author_id = employee.id
//         WHERE username = $1
//         ORDER BY public.bid_version.name LIMIT $2 OFFSET $3`

const getBidsByFilterQuery = `WITH max_version AS (
    SELECT bid_id, MAX(number) AS max_number
    FROM bid_version
    GROUP BY bid_id
)
SELECT bid.id, bid_version.id, status, tender_id, author_type, author_id, bid.created_at, name, description, number
FROM bid
    JOIN bid_version ON bid.id = bid_version.bid_id
    JOIN employee ON bid.author_id = employee.id
    join MAX_VERSION on bid_version.bid_id = MAX_VERSION.bid_id
WHERE username = $1
ORDER BY public.bid_version.name LIMIT $2 OFFSET $3`

func (b bidStorage) GetBidsByFilter(limit, offset int, username string) ([]domain.Bid, error) {
	var bids []domain.Bid
	rows, err := b.conn.Query(getBidsByFilterQuery, username, limit, offset)
	if err != nil {
		return bids, errors.Wrap(err, "could not query get bids by filter")
	}

	defer rows.Close()
	for rows.Next() {
		var bid domain.Bid
		err = rows.Scan(
			&bid.ID,
			&bid.VersionId,
			&bid.Status,
			&bid.TenderId,
			&bid.AuthorType,
			&bid.AuthorId,
			&bid.CreatedAt,
			&bid.Name,
			&bid.Description,
			&bid.Version,
		)
		if err != nil {
			return bids, errors.Wrap(err, "could not scan bid row")
		}

		bids = append(bids, bid)
	}

	return bids, nil
}

// const getBidsByTenderIdByFilter = `SELECT bid.id, version_id, status, tender_id, author_type, author_id, bid.created_at, name, description, version.version FROM bid
//    JOIN version ON bid.version_id = version.id
//        WHERE tender_id = $1
//        ORDER BY public.version.name LIMIT $2 OFFSET $3`

const getBidsByTenderIdByFilter = `WITH max_version AS (
    SELECT bid_id, MAX(number) AS max_number
    FROM bid_version
    GROUP BY bid_id
)
SELECT bid.id, bid_version.id, status, tender_id, author_type, author_id, bid.created_at, name, description, number
FROM bid
    JOIN bid_version ON bid.id = bid_version.bid_id
    JOIN employee ON bid.author_id = employee.id
    join MAX_VERSION on bid_version.bid_id = MAX_VERSION.bid_id
WHERE tender_id = $1
ORDER BY public.bid_version.name LIMIT $2 OFFSET $3`

func (b bidStorage) GetBidsByTenderIdByFilter(limit, offset int, username string, tenderId string) ([]domain.Bid, error) {
	var bids []domain.Bid
	rows, err := b.conn.Query(getBidsByTenderIdByFilter, tenderId, limit, offset)
	if err != nil {
		return bids, errors.Wrap(err, "could not query get bids by filter")
	}

	defer rows.Close()
	for rows.Next() {
		var bid domain.Bid
		err = rows.Scan(
			&bid.ID,
			&bid.VersionId,
			&bid.Status,
			&bid.TenderId,
			&bid.AuthorType,
			&bid.AuthorId,
			&bid.CreatedAt,
			&bid.Name,
			&bid.Description,
			&bid.Version,
		)
		if err != nil {
			return bids, errors.Wrap(err, "could not scan bid row")
		}

		bids = append(bids, bid)
	}

	return bids, nil
}

const getBidByIdQuery = `SELECT bid.id, bid_version.id, status, tender_id, author_type, author_id, created_at, name, description, number
FROM bid JOIN bid_version ON bid.id = bid_version.bid_id WHERE bid.id = $1 AND number=(SELECT MAX(number) FROM bid_version WHERE bid_id = $1)`

func (b bidStorage) GetBidsById(bidId string) (domain.Bid, error) {
	var bid domain.Bid
	err := b.conn.QueryRow(getBidByIdQuery, bidId).Scan(
		&bid.ID,
		&bid.VersionId,
		&bid.Status,
		&bid.TenderId,
		&bid.AuthorType,
		&bid.AuthorId,
		&bid.CreatedAt,
		&bid.Name,
		&bid.Description,
		&bid.Version,
	)
	if err != nil {
		return bid, errors.Wrap(err, "could not query get bid row")
	}

	return bid, nil
}

const getBidStatusByIdQuery = `SELECT status FROM bid WHERE id = $1`

func (b bidStorage) GetBidStatusById(bidId string) (string, error) {
	var status string
	err := b.conn.QueryRow(getBidStatusByIdQuery, bidId).Scan(&status)
	if err != nil {
		return "", errors.Wrap(err, "could not query get bid status")
	}

	return status, nil
}

const updateBidStatusById = `UPDATE bid SET status = $1 WHERE id = $2`

func (b bidStorage) UpdateBidStatus(bidId string, status string) error {
	_, err := b.conn.Exec(updateBidStatusById, status, bidId)
	if err != nil {
		return errors.Wrap(err, "could not update bid status")
	}

	return nil
}

// const updateBidParamsById = `UPDATE bid SET version_id = $1 WHERE id = $2`

func (b bidStorage) UpdateBidById(bidId string, bid domain.Bid) error {
	tx, err := b.conn.Begin()
	if err != nil {
		return errors.Wrap(err, "could not begin transaction")
	}

	defer tx.Rollback()

	err = tx.QueryRow(createBidVersionQuery, bid.Name, bid.Description, bid.Version, bid.ID).Scan(&bid.VersionId, &bid.Version)
	if err != nil {
		return errors.Wrap(err, "could not insert bid version")
	}

	// _, err = tx.Exec(updateBidParamsById, bid.VersionId, bidId)
	// if err != nil {
	//	return errors.Wrap(err, "could not update bid params")
	// }

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "could not commit transaction")
	}

	return nil
}

const updateBidDecisionByIdQuery = `UPDATE bid SET status=$1 WHERE id = $2`

func (b bidStorage) UpdateBidDecisionById(bidId string, decision string) error {
	_, err := b.conn.Exec(updateBidDecisionByIdQuery, decision, bidId)
	if err != nil {
		return errors.Wrap(err, "could not update bid decision")
	}

	return nil
}

const createBidFeedbackQuery = `INSERT INTO review (description, bid_id) VALUES ($1, $2)`

func (b bidStorage) UpdateBidFeedbackById(bidId string, feedback string) error {
	_, err := b.conn.Exec(createBidFeedbackQuery, feedback, bidId)
	if err != nil {
		return errors.Wrap(err, "could not insert bid feedback")
	}

	return nil
}

const getBidVersionByIdQuery = `SELECT id, name, description, version.version FROM version WHERE id = $1`

func (b bidStorage) GetBidVersionById(bidId string, version int) (domain.Bid, error) {
	var bid domain.Bid
	err := b.conn.QueryRow(getBidVersionByIdQuery, bidId, version).Scan(
		&bid.VersionId,
		&bid.Name,
		&bid.Description,
		&bid.Version,
	)
	if err != nil {
		return bid, errors.Wrap(err, "could not query get bid version")
	}

	return bid, nil
}

const checkOrganizationIsExistQuery = `SELECT EXISTS (SELECT * FROM organization WHERE id = $1)`

func (b bidStorage) CheckOrganizationIsExist(organizationId string) (bool, error) {
	var exists bool
	err := b.conn.QueryRow(checkOrganizationIsExistQuery, organizationId).Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "could not check organization existence")
	}

	return exists, nil
}

func NewBidStorage(url string) (storage.BidStorage, error) {
	conf, err := pgx.ParseConnectionString(url)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse connection string")
	}
	conn, err := pgx.Connect(conf)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to connect to database")
	}

	return &bidStorage{
		conn: conn,
	}, nil
}
