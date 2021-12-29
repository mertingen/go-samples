CREATE TABLE students_lectures
(
    id         int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    student_id int NOT NULL,
    lecture_id int NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (lecture_id) REFERENCES lectures(id)
);