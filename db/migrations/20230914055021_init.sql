-- migrate:up
CREATE TABLE "shortened" (
	"id" int PRIMARY KEY,
	"isFile" bool NOT NULL,
	"slug" varchar(255) UNIQUE NOT NULL
);

CREATE TABLE "files" (
	"id" int PRIMARY KEY,
	"originalFilename" varchar(255) NOT NULL,
	"slug" varchar(255) UNIQUE NOT NULL,
	"size" int NOT NULL,
	"expires" bigint NOT NULL
);

CREATE TABLE "urls" (
	"id" int PRIMARY KEY,
	"slug" varchar(255) UNIQUE NOT NULL,
	"destination" varchar(1024) NOT NULL,
	"expires" bigint NOT NULL
);

CREATE INDEX ON "shortened" ("slug");

CREATE INDEX ON "files" ("slug");

CREATE INDEX ON "urls" ("slug");

-- migrate:down
DROP TABLE "shortened";

DROP TABLE "files";

DROP TABLE "urls";