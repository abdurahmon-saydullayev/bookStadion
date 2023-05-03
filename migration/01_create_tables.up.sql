CREATE TYPE stadion_type AS ENUM ('outside', 'inside');


create table IF NOT EXISTS "football"(
    "id" serial primary key not null,
    "name" varchar(50),
    "from_time"  timestamp not null,
    "to_time" timestamp not null,
    "price" integer not null,
    "description" varchar(250) not null,
    "image" varchar(50) not null,
    "stadion_type" stadion_type not null,
    "location" JSONB not null
);

