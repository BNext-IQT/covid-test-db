/*create users and database*/
CREATE USER IF NOT EXISTS covid_bug;
CREATE DATABASE IF NOT EXISTS covid_tests;
GRANT ALL ON DATABASE covid_tests TO covid_bug;

/*setup tables*/
CREATE TABLE IF NOT EXISTS covid_tests.pocs (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	phone STRING NULL,
	email STRING NULL,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_tests.companies (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	city STRING NOT NULL,
	state STRING NOT NULL,
	country STRING NOT NULL,
	poc_id UUID NULL REFERENCES covid_tests.pocs(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_tests.suppliers (
	company_id UUID NOT NULL REFERENCES covid_tests.companies(id) ON DELETE CASCADE,
	supplier_id UUID NOT NULL REFERENCES covid_tests.companies(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	PRIMARY KEY (company_id, supplier_id)
);

CREATE TABLE IF NOT EXISTS covid_tests.test_types(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_tests.test_target_types(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_tests.regulatory_approval_types(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	valid_country STRING NULL,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_tests.tests (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	description STRING NOT NULL,
	test_type_id UUID NULL REFERENCES covid_tests.test_types(id) ON DELETE CASCADE,
	poc_id UUID NULL REFERENCES covid_tests.pocs(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_tests.test_regulatory_approval (
	test_id UUID NULL REFERENCES covid_tests.tests(id) ON DELETE CASCADE,
	regulatory_approval_type_id UUID NULL REFERENCES covid_tests.regulatory_approval_types(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	PRIMARY KEY (test_id, regulatory_approval_type_id)
);

CREATE TABLE IF NOT EXISTS covid_tests.test_targets (
	test_id UUID NULL REFERENCES covid_tests.tests(id) ON DELETE CASCADE,
	test_target_type_id UUID NULL REFERENCES covid_tests.test_target_types(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	PRIMARY KEY (test_id, test_target_type_id)
);

/* fill in reference tables */
INSERT INTO covid_tests.test_types (name, created_by, updated_by) VALUES ('Immunoassays/serology', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.test_types (name, created_by, updated_by) VALUES ('molecular assays', 'initialize.sql', 'initialize.sql');

INSERT INTO covid_tests.test_types (name, created_by, updated_by) VALUES ('IgG', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.test_types (name, created_by, updated_by) VALUES ('IgM', 'initialize.sql', 'initialize.sql');

INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('CE-IVD', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('RUO', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('HSA', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('EUA', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('MFDS (by country)', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('TGA (by country)', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('FDA (by country)', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('WHO EUL', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_tests.regulatory_approval_types (name, created_by, updated_by) VALUES ('NRA', 'initialize.sql', 'initialize.sql');
