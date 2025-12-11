ALTER TABLE attendance_histories
ALTER COLUMN id TYPE VARCHAR(15);

ALTER TABLE attendance_histories
ADD COLUMN day_name VARCHAR(255);

ALTER TABLE attendance_histories
RENAME COLUMN check_out_photo TO check_out_phto;

ALTER TABLE attendance_histories
ALTER COLUMN user_id TYPE VARCHAR(10);

ALTER TABLE attendance_histories
ALTER COLUMN duration TYPE DECIMAL;