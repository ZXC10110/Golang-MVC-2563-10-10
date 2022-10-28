CREATE TABLE Test.student (
    student_id int PRIMARY KEY,
    first_name varchar(255),
    last_name varchar(255),
    year int
)

CREATE TABLE Test.sports (
    sport_no int,
    student_id int,
    first_name varchar(255),
    last_name varchar(255),
    year int
)