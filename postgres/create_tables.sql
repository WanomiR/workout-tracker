CREATE TABLE public.users
(
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    name          VARCHAR(255) NOT NULL,
    patronymic    VARCHAR(255),
    surname       VARCHAR(255),
    weight        REAL,
    height        REAL,
    dob           DATE,
    registered_at TIMESTAMP WITHOUT TIME ZONE
);

INSERT INTO public.users (email, password, name, patronymic, surname, weight, height, dob, registered_at)
VALUES ('admin@example.com', '$2a$10$aao6Yk/9XpdbpON377Xikui50JAeaoZQg.emqV54Ym9s56fF6KS7a', 'User', '', 'Admin', 104.5,
        196.7, '1995-01-21', now())
;