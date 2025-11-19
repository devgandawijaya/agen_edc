-- agen_sahabat_full.sql
-- Complete PostgreSQL Schema for Aplikasi Agen Sahabat

CREATE SCHEMA IF NOT EXISTS agen_sahabat;
SET search_path = agen_sahabat, public;

-- =====================
-- TABLE: agents
-- =====================
CREATE TABLE IF NOT EXISTS agents (
    id BIGSERIAL PRIMARY KEY,
    agent_type VARCHAR(20) NOT NULL CHECK (agent_type IN ('perorangan','badan_usaha')),
    business_place_status VARCHAR(10) NOT NULL CHECK (business_place_status IN ('hak_milik','sewa')),
    edc_usage_activity JSONB,
    cooperation_type VARCHAR(20) NOT NULL CHECK (cooperation_type IN ('beli_putus','sewa_beli')),
    transaction_features JSONB,
    submission_status VARCHAR(30) DEFAULT 'pending',
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- =====================
-- TABLE: acquisition_info
-- =====================
CREATE TABLE IF NOT EXISTS acquisition_info (
    id BIGSERIAL PRIMARY KEY,
    agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
    acquisition_type VARCHAR(30) NOT NULL CHECK (acquisition_type IN ('independent_partner','master_dealer','teknisi','sales','lainnya')),
    acquisition_name VARCHAR(255),
    acquisition_nik VARCHAR(100),
    acquisition_city_or_email VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT now()
);

-- =====================
-- TABLE: owners
-- =====================
CREATE TABLE IF NOT EXISTS owners (
    id BIGSERIAL PRIMARY KEY,
    agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
    full_name VARCHAR(255) NOT NULL,
    birth_place VARCHAR(255),
    birth_date DATE,
    gender VARCHAR(20),
    religion VARCHAR(50),
    occupation VARCHAR(100),
    ktp_address TEXT,
    city VARCHAR(100),
    province VARCHAR(100),
    postal_code VARCHAR(20),
    phone VARCHAR(50),
    email VARCHAR(255),
    identity_number VARCHAR(100),
    tax_number VARCHAR(100),
    created_at TIMESTAMPTZ DEFAULT now()
);

-- =====================
-- TABLE: business_profiles
-- =====================
CREATE TABLE IF NOT EXISTS business_profiles (
    id BIGSERIAL PRIMARY KEY,
    agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
    business_name VARCHAR(255) NOT NULL,
    contact_person VARCHAR(255),
    contact_phone VARCHAR(50),
    phone VARCHAR(50),
    business_npwp VARCHAR(100),
    business_email VARCHAR(255),
    business_address TEXT,
    latitude NUMERIC(10,7),
    longitude NUMERIC(10,7),
    city VARCHAR(100),
    province VARCHAR(100),
    postal_code VARCHAR(20),
    product_description VARCHAR(255),
    business_duration VARCHAR(100),
    monthly_gross_profit VARCHAR(100),
    monthly_transaction_avg VARCHAR(100),
    operating_days VARCHAR(100),
    operating_hours VARCHAR(100),
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- =====================
-- TABLE: uploaded_documents
-- =====================
CREATE TABLE IF NOT EXISTS uploaded_documents (
    id BIGSERIAL PRIMARY KEY,
    agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
    document_type VARCHAR(50) NOT NULL CHECK (document_type IN ('ktp','selfie_ktp','foto_luar','foto_dalam','npwp','buku_rekening','skdu','lainnya')),
    file_path TEXT NOT NULL,
    file_name VARCHAR(255),
    file_size BIGINT,
    mime_type VARCHAR(100),
    uploaded_by VARCHAR(100),
    uploaded_at TIMESTAMPTZ DEFAULT now()
);

-- =====================
-- TABLE: bank_info
-- =====================
CREATE TABLE IF NOT EXISTS bank_info (
    id BIGSERIAL PRIMARY KEY,
    agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
    bank_account_number VARCHAR(100) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    bank_code VARCHAR(50),
    account_holder_name VARCHAR(255) NOT NULL,
    verification_status VARCHAR(30) DEFAULT 'unverified',
    created_at TIMESTAMPTZ DEFAULT now(),
    verified_at TIMESTAMPTZ
);

-- =====================
-- TABLE: signatures
-- =====================
CREATE TABLE IF NOT EXISTS signatures (
    id BIGSERIAL PRIMARY KEY,
    agent_id BIGINT NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
    owner_signature_path TEXT,
    company_signature_path TEXT,
    sign_date DATE,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- =====================
-- TABLE: audit_logs
-- =====================
CREATE TABLE IF NOT EXISTS audit_logs (
    id BIGSERIAL PRIMARY KEY,
    agent_id BIGINT,
    action VARCHAR(255) NOT NULL,
    actor VARCHAR(255),
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- =====================
-- INDEXES
-- =====================
CREATE INDEX IF NOT EXISTS idx_agents_agent_type ON agents(agent_type);
CREATE INDEX IF NOT EXISTS idx_agents_submission_status ON agents(submission_status);
CREATE INDEX IF NOT EXISTS idx_business_profiles_lat_long ON business_profiles(latitude, longitude);
CREATE INDEX IF NOT EXISTS idx_uploaded_documents_agent ON uploaded_documents(agent_id, document_type);

-- =====================
-- TRIGGERS
-- =====================
CREATE OR REPLACE FUNCTION agen_updated_at_trigger()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_agents_updated_at
BEFORE UPDATE ON agents
FOR EACH ROW
EXECUTE FUNCTION agen_updated_at_trigger();

CREATE TRIGGER trg_business_profiles_updated_at
BEFORE UPDATE ON business_profiles
FOR EACH ROW
EXECUTE FUNCTION agen_updated_at_trigger();

-- =====================
-- TABLE: users
-- =====================
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('user', 'admin')) DEFAULT 'user',
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- =====================
-- INDEXES for users
-- =====================
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- =====================
-- TRIGGER for users
-- =====================
CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION agen_updated_at_trigger();

-- =====================
-- VIEW: vw_agent_full
-- =====================
CREATE OR REPLACE VIEW vw_agent_full AS
SELECT
  a.id as agent_id,
  a.agent_type,
  a.business_place_status,
  a.cooperation_type,
  a.transaction_features,
  ai.acquisition_type,
  ai.acquisition_name,
  o.full_name as owner_name,
  bp.business_name,
  bp.business_address,
  bi.bank_name,
  bi.bank_account_number,
  (SELECT jsonb_agg(jsonb_build_object('type',ud.document_type,'path',ud.file_path))
     FROM uploaded_documents ud WHERE ud.agent_id = a.id) as documents
FROM agents a
LEFT JOIN acquisition_info ai ON ai.agent_id = a.id
LEFT JOIN owners o ON o.agent_id = a.id
LEFT JOIN business_profiles bp ON bp.agent_id = a.id
LEFT JOIN bank_info bi ON bi.agent_id = a.id;
