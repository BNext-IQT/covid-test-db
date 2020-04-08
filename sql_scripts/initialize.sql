CREATE DATABASE IF NOT EXISTS covid_tests;
CREATE TABLE IF NOT EXISTS covid_tests.pocs (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	phone STRING NULL,
	email STRING NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);
CREATE TABLE IF NOT EXISTS covid_tests.companies (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	city STRING NOT NULL,
	state STRING NOT NULL,
	country STRING NOT NULL,
	poc_id UUID NULL REFERENCES covid_tests.pocs(id) ON DELETE CASCADE,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);
CREATE TABLE IF NOT EXISTS covid_tests.test_types(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);
INSERT INTO covid_tests.test_types (name) VALUES ('Immunoassays/serology');
INSERT INTO covid_tests.test_types (name) VALUES ('molecular assays');
CREATE TABLE IF NOT EXISTS covid_tests.tests (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name STRING NOT NULL,
	description STRING NOT NULL,
	test_type_id UUID NULL REFERENCES covid_tests.test_types(id) ON DELETE CASCADE,
	poc_id UUID NULL REFERENCES covid_tests.pocs(id) ON DELETE CASCADE,
	created TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);