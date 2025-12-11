ALTER TABLE attendance_histories add CONSTRAINT fk_user FOREIGN KEY (user_id)
REFERENCES users(id) ON DELETE CASCADE;