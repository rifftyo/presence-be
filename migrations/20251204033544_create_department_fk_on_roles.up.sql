ALTER TABLE roles add CONSTRAINT fk_department FOREIGN KEY (department_id)
REFERENCES departments(id) ON DELETE CASCADE;