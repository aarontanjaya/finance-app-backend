CREATE EXTENSION IF NOT EXISTS timescaledb;
CREATE TABLE IF NOT EXISTS "fund_snapshot"(
    "source" INTEGER,
    "id" INTEGER,
    "isin_code" VARCHAR(12),
    "name" TEXT,
    "im_name" TEXT,
    "im_fee" VARCHAR,
    "type" VARCHAR(32),
    "type_id" INTEGER,
    "index_flag" BOOLEAN,
    "etf_flag" BOOLEAN,
    "sharia_flag" BOOLEAN,
    "nav" NUMERIC,
    "daily_rr" NUMERIC,
    "monthly_rr" NUMERIC,
    "ytd_rr" NUMERIC,
    "yearly_rr" NUMERIC,
    "aum" NUMERIC,
    "aum_updated_at" NUMERIC,
    "units" NUMERIC,
    "last_update" NUMERIC,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

ALTER TABLE public.fund_snapshots ADD CONSTRAINT fund_snapshots_uk UNIQUE KEY ("last_update",id);


/* chunk_time_interval 1 week */
SELECT create_hypertable('fund_snapshot', 'last_update', chunk_time_interval => 604800);