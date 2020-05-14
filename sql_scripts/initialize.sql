/*create users and database*/
CREATE USER IF NOT EXISTS covid_bug;
CREATE DATABASE IF NOT EXISTS covid_diagnostics;
GRANT ALL ON DATABASE covid_diagnostics TO covid_bug;

/*setup tables*/
CREATE TABLE IF NOT EXISTS covid_diagnostics.pocs (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	phone STRING NULL,
	email STRING NULL,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_diagnostics.companies (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	street_address STRING NULL,
	city STRING NULL,
	state STRING NULL,
	country STRING NULL,
	postal_code STRING NULL,
	stage STRING NULL,
	valuation STRING NULL,
	poc_id UUID NULL REFERENCES covid_diagnostics.pocs(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_diagnostics.suppliers (
	company_id UUID NOT NULL REFERENCES covid_diagnostics.companies(id) ON DELETE CASCADE,
	supplier_id UUID NOT NULL REFERENCES covid_diagnostics.companies(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	PRIMARY KEY (company_id, supplier_id)
);

CREATE TABLE IF NOT EXISTS covid_diagnostics.diagnostic_types(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_diagnostics.diagnostic_target_types(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_diagnostics.regulatory_approval_types(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	valid_country STRING NULL,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_diagnostics.diagnostics (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	description STRING NOT NULL,
	company_id UUID NULL REFERENCES covid_diagnostics.companies(id) ON DELETE CASCADE,
	diagnostic_type_id UUID NULL REFERENCES covid_diagnostics.diagnostic_types(id) ON DELETE CASCADE,
	poc_id UUID NULL REFERENCES covid_diagnostics.pocs(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TABLE IF NOT EXISTS covid_diagnostics.diagnostic_regulatory_approval (
	diagnostic_id UUID NULL REFERENCES covid_diagnostics.diagnostics(id) ON DELETE CASCADE,
	regulatory_approval_type_id UUID NULL REFERENCES covid_diagnostics.regulatory_approval_types(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	PRIMARY KEY (diagnostic_id, regulatory_approval_type_id)
);

CREATE TABLE IF NOT EXISTS covid_diagnostics.diagnostic_targets (
	diagnostic_id UUID NULL REFERENCES covid_diagnostics.diagnostics(id) ON DELETE CASCADE,
	diagnostic_target_type_id UUID NULL REFERENCES covid_diagnostics.diagnostic_target_types(id) ON DELETE CASCADE,
	created_by STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by STRING NULL,
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	PRIMARY KEY (diagnostic_id, diagnostic_target_type_id)
);

/* fill in reference tables */
INSERT INTO covid_diagnostics.diagnostic_types (name, created_by, updated_by) VALUES ('Immunoassays/serology', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.diagnostic_types (name, created_by, updated_by) VALUES ('molecular assays', 'initialize.sql', 'initialize.sql');

INSERT INTO covid_diagnostics.diagnostic_target_types (name, created_by, updated_by) VALUES ('IgG', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.diagnostic_target_types (name, created_by, updated_by) VALUES ('IgM', 'initialize.sql', 'initialize.sql');

INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('CE-IVD', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('RUO', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('HSA', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('EUA', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('MFDS (by country)', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('TGA (by country)', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('FDA (by country)', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('WHO EUL', 'initialize.sql', 'initialize.sql');
INSERT INTO covid_diagnostics.regulatory_approval_types (name, created_by, updated_by) VALUES ('NRA', 'initialize.sql', 'initialize.sql');
