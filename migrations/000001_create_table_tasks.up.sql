CREATE TABLE IF NOT EXISTS student_tasks (
    id TEXT NOT NULL,
    title TEXT NOT NULL,
    deadline TIMESTAMP NOT NULL,
    summary TEXT,
    student_id TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    type VARCHAR(30) CHECK (type IN('paper-based','presentation','speech')),
    status VARCHAR(30) NOT NULL CHECK(status IN('completed','incomplete','missed'))
); 