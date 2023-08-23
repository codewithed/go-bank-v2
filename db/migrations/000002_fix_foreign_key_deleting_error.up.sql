
-- Add "ON DELETE CASCADE" to the foreign key constraints
ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
