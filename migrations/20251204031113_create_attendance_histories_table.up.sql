CREATE TABLE attendance_histories(
    id VARCHAR(15) PRIMARY KEY,
    user_id VARCHAR(10) NOT NULL,
    attendance_date DATE NOT NULL,
    day_name VARCHAR(20),
    check_in_time TIMESTAMP WITH TIME ZONE NOT NULL,
    check_out_time TIMESTAMP WITH TIME ZONE,
    duration DECIMAL,
    status VARCHAR(20) NOT NULL,
    check_in_lat DECIMAL NOT NULL,
    check_in_lng DECIMAL NOT NULL,
    check_out_lat DECIMAL,
    check_out_lng DECIMAL,
    check_in_photo TEXT NOT NULL,
    check_out_phto TEXT
);