-- table for organization (main entity)
create table organization (
  org_id uuid primary key default gen_random_uuid(), -- uuidv4 primary key (always not null)
  org_ror_id text not null, -- original ror id (e.g., 'https://ror.org/04ttjf776')
  org_type_id integer not null,
  org_established integer not null,
  org_status text not null,
  org_created_date date not null,
  org_created_schema_version text not null,
  org_last_modified_date date not null,
  org_last_modified_schema_version text not null
  org_created_at timestamptz not null,
  org_modified_at timestamptz not null
);

create table schema_version (
  sve_id serial primary key,
  sve_name text not null unique,
  sve_created_at timestamptz not null,
  sve_modified_at timestamptz not null
);

create table continent (
  con_id smallserial primary key,
  con_code text not null unique,
  con_name text not null,
  con_created_at timestamptz not null,
  con_modified_at timestamptz not null
);

create table country (
  cou_id smallserial primary key,
  cou_continent_id smallint not null,
  cou_code text not null unique,
  cou_name text not null,
  cou_created_at timestamptz not null,
  cou_modified_at timestamptz not null
);

create table country_subdivision (
  csu_id uuid primary key default gen_random_uuid(),
  csu_country_id smallint not null,
  csu_code text not null,
  csu_name text not null,
  csu_lat decimal(9, 6) not null,
  csu_lng decimal(9, 6) not null,
  csu_created_at timestamptz not null,
  csu_modified_at timestamptz not null
);

create table location (
  loc_id uuid primary key default gen_random_uuid(),
  loc_country_subdivision_id uuid not null,
  loc_geonames_id integer not null unique,
  loc_name text not null,
  csu_created_at timestamptz not null,
  csu_modified_at timestamptz not null
);
create index on location (loc_country_subdivision_id);
create index on location (loc_geonames_id);






-- table for organization type (e.g., 'education', 'funder')
create table organization_type (
  oty_id serial primary key,
  oty_name text not null,
  oty_created_at timestamptz not null,
  oty_modified_at timestamptz not null
);

---

-- table for organization name (including different types like 'label', 'acronym', 'alias')
create table organization_name (
  organization_id uuid not null,
  value text not null,
  type_name text not null, -- e.g., 'ror_display', 'label', 'acronym', 'alias'
  lang text, -- lang might be null for some acronyms
  primary key (organization_id, value, type_name)
);

---

-- table for external id
create table external_id (
  id serial primary key, -- auto-incrementing, always not null
  organization_id uuid not null,
  type text not null, -- e.g., 'fundref', 'grid', 'isni', 'wikidata'
  preferred text, -- 'preferred' can be null in the ROR data
  unique (organization_id, type, preferred) -- ensures unique preferred ID for a type within an organization
);

---

-- table for all 'all' values within external_id
create table external_id_all_value (
  external_id_entry_id integer not null,
  value text not null,
  primary key (external_id_entry_id, value)
);

---

-- table for location


---

-- junction table for organization-location relationship
create table organization_location (
  organization_id uuid not null,
  geonames_id integer not null,
  primary key (organization_id, geonames_id)
);

---

-- table for link (website, wikipedia, etc.)
create table link (
  id serial primary key, -- auto-incrementing, always not null
  organization_id uuid not null,
  type text not null, -- e.g., 'website', 'wikipedia'
  value text not null, -- the url
  unique (organization_id, type, value)
);

---

-- new table for relationship types (e.g., 'parent', 'child')
create table relationship_type (
  id uuid primary key default gen_random_uuid(), -- always not null
  name text unique not null -- 'parent', 'child', 'related', 'predecessor', 'successor'
);

---

-- table for relationship (parent, child, related)
create table relationship (
  id serial primary key, -- auto-incrementing, always not null
  source_organization_id uuid not null,
  target_organization_id uuid not null,
  relationship_type_id uuid not null, -- references relationship_type table
  label text not null, -- the label as provided in the json (e.g., "arc centre of excellence...")
  unique (source_organization_id, target_organization_id, relationship_type_id)
);

---

-- table for domain (if any)
create table domain (
  organization_id uuid not null,
  domain_name text not null,
  primary key (organization_id, domain_name)
);