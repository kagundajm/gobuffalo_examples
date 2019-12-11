CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "categories" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL
);
CREATE UNIQUE INDEX "uq_category_name" ON "categories" (name);
CREATE TABLE IF NOT EXISTS "products" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"category_id" char(36) NOT NULL,
"price" INTEGER NOT NULL
);
CREATE UNIQUE INDEX "uq_product_name" ON "products" (name);
