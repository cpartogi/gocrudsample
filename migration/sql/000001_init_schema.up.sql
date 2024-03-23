CREATE TABLE tutorials (
	id uuid NOT NULL,
    tutorial_type_id uuid NOT NULL,
	keywords varchar(500) NULL,
    "sequence" int4 NULL,
    title varchar(255) NOT NULL,
    "description" text NULL,
    created_by varchar(255) NULL,
	created_at timestamptz NULL,
	updated_by varchar(255) NULL,
	updated_at timestamptz NULL,
    deleted_by varchar(255) NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT tutorials_pkey PRIMARY KEY (id)
);

CREATE TABLE tutorial_types (
	id uuid NOT NULL,
	type_name varchar(255) NOT NULL,
    created_by varchar(255) NULL,
	created_at timestamptz NULL,
	updated_by varchar(255) NULL,
	updated_at timestamptz NULL,
    deleted_by varchar(255) NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT tutorial_types_pkey PRIMARY KEY (id)
);

ALTER TABLE tutorials ADD CONSTRAINT tutorials_tutorial_type_id_fkey FOREIGN KEY (tutorial_type_id) REFERENCES tutorial_types(id);
