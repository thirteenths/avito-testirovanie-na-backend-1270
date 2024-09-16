-- +goose Up
-- +goose StatementBegin
-- Шаг 1: Заполнение таблицы employee
INSERT INTO employee (username, first_name, last_name) VALUES
                                                           ('john_doe', 'John', 'Doe'),
                                                           ('jane_smith', 'Jane', 'Smith'),
                                                           ('alice_johnson', 'Alice', 'Johnson');

-- Шаг 2: Заполнение таблицы organization
INSERT INTO organization (name, description, type) VALUES
                                                       ('ABC Construction', 'A construction company specializing in residential projects.', 'LLC'),
                                                       ('XYZ Delivery Services', 'Logistics and delivery services.', 'IE'),
                                                       ('123 Manufacturing', 'Manufacturing components and machinery.', 'JSC');

-- Шаг 3: Заполнение таблицы organization_responsible
INSERT INTO organization_responsible (organization_id, user_id) VALUES
                                                                    ((SELECT id FROM organization WHERE name = 'ABC Construction'), (SELECT id FROM employee WHERE username = 'john_doe')),
                                                                    ((SELECT id FROM organization WHERE name = 'XYZ Delivery Services'), (SELECT id FROM employee WHERE username = 'jane_smith')),
                                                                    ((SELECT id FROM organization WHERE name = '123 Manufacturing'), (SELECT id FROM employee WHERE username = 'alice_johnson'));

-- Шаг 4: Заполнение таблицы tender
INSERT INTO tender (status, organization_id) VALUES
                                                 ('Created', (SELECT id FROM organization WHERE name = 'ABC Construction')),
                                                 ('Published', (SELECT id FROM organization WHERE name = 'XYZ Delivery Services')),
                                                 ('Closed', (SELECT id FROM organization WHERE name = '123 Manufacturing'));

-- Шаг 5: Заполнение таблицы version
INSERT INTO version (name, description, service_type, number, tender_id) VALUES
                                                                             ('Version 1', 'Initial version of the tender.', 'Construction', 1, (SELECT id FROM tender WHERE organization_id = (SELECT id FROM organization WHERE name = 'ABC Construction'))),
                                                                             ('Version 1', 'Initial version for delivery services.', 'Delivery', 1, (SELECT id FROM tender WHERE organization_id = (SELECT id FROM organization WHERE name = 'XYZ Delivery Services'))),
                                                                             ('Version 2', 'Revised version for manufacturing.', 'Manufacture', 2, (SELECT id FROM tender WHERE organization_id = (SELECT id FROM organization WHERE name = '123 Manufacturing')));

-- Шаг 6: Заполнение таблицы bid
INSERT INTO bid (status, tender_id, author_type, author_id) VALUES
                                                                ('Created', (SELECT id FROM tender WHERE status = 'Created'), 'User', (SELECT id FROM employee WHERE username = 'john_doe')),
                                                                ('Published', (SELECT id FROM tender WHERE status = 'Published'), 'Organization', (SELECT id FROM organization WHERE name = 'XYZ Delivery Services')),
                                                                ('Canceled', (SELECT id FROM tender WHERE status = 'Closed'), 'User', (SELECT id FROM employee WHERE username = 'jane_smith'));

-- Шаг 7: Заполнение таблицы bid_version
INSERT INTO bid_version (name, description, number, bid_id) VALUES
                                                                ('Bid Version 1', 'First version of the bid for the tender.', 1, (SELECT id FROM bid WHERE status = 'Created')),
                                                                ('Bid Version 2', 'Second version of the bid.', 2, (SELECT id FROM bid WHERE status = 'Published')),
                                                                ('Bid Version 1', 'First version for the canceled bid.', 1, (SELECT id FROM bid WHERE status = 'Canceled'));

-- Шаг 8: Заполнение таблицы review
INSERT INTO review (description, bid_id) VALUES
                                             ('Review for Bid Version 1.', (SELECT id FROM bid WHERE status = 'Created')),
                                             ('Review for Bid Version 2.', (SELECT id FROM bid WHERE status = 'Published')),
                                             ('Review for canceled bid.', (SELECT id FROM bid WHERE status = 'Canceled'));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE bid;
TRUNCATE TABLE tender;
TRUNCATE TABLE organization_responsible;
TRUNCATE TABLE organization;
TRUNCATE TABLE employee;
-- +goose StatementEnd
