-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);

CREATE TABLE organization (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_responsible (
    id SERIAL PRIMARY KEY,
    organization_id uuid REFERENCES organization(id) ON DELETE CASCADE,
    user_id uuid REFERENCES employee(id) ON DELETE CASCADE
);

CREATE TYPE tender_service_type AS ENUM(
    'Construction',
    'Delivery',
    'Manufacture'
);

CREATE TYPE tender_status AS ENUM(
    'Created',
    'Published',
    'Closed'
);

CREATE TABLE tender (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    status tender_status DEFAULT  'Created',
    organization_id uuid REFERENCES organization(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE version (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL ,
    service_type tender_service_type,
    number INT DEFAULT 1 CHECK ( number >=1 ),
    tender_id uuid REFERENCES tender(id) ON DELETE CASCADE
);

CREATE TYPE bid_status AS ENUM(
    'Created',
    'Published',
    'Canceled',
    'Approved',
    'Rejected'
);

CREATE TYPE bid_author_type AS ENUM(
    'Organization',
    'User'
);

CREATE TABLE bid (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    status bid_status DEFAULT 'Created',
    tender_id uuid REFERENCES tender(id) ON DELETE CASCADE,
    author_type bid_author_type NOT NULL ,
    author_id uuid NOT NULL ,
    created_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE bid_version(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL ,
    number INT DEFAULT 1 CHECK ( number >=1 ),
    bid_id uuid REFERENCES bid(id) ON DELETE CASCADE
);

CREATE TABLE review (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    description TEXT NOT NULL,
    create_at TIMESTAMP DEFAULT current_timestamp,
    bid_id uuid REFERENCES bid(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE organization_responsible;
DROP TABLE organization;
DROP TYPE  organization_type;
DROP TABLE employee;

DROP TABLE version;
DROP TYPE tender_service_type;
DROP TYPE tender_status;
DROP TABLE tender;

DROP TYPE bid_status;
DROP TYPE bid_author_type;
DROP TABLE bid;
DROP TABLE review;
-- +goose StatementEnd
