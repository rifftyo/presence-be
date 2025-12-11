ALTER TABLE attendance_histories
ALTER COLUMN id TYPE VARCHAR(255);

ALTER TABLE attendance_histories
DROP COLUMN day_name;

ALTER TABLE attendance_histories
RENAME COLUMN check_out_phto TO check_out_photo;

ALTER TABLE attendance_histories
ALTER COLUMN user_id TYPE VARCHAR(155);

ALTER TABLE attendance_histories
ALTER COLUMN duration TYPE VARCHAR(100);