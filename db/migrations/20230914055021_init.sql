-- migrate:up
CREATE TABLE "slugs" (
	"id" int PRIMARY KEY,
	"is_file" bool NOT NULL,
	"slug" varchar(255) UNIQUE NOT NULL
);

CREATE TABLE "files" (
	"id" int PRIMARY KEY,
	"original_filename" varchar(255) NOT NULL,
	"slug" varchar(255) UNIQUE NOT NULL,
	"size" int NOT NULL,
	"expires" bigint NOT NULL,
	CONSTRAINT fk_slug
		FOREIGN KEY("slug")
			REFERENCES "slugs"("slug")
			ON DELETE CASCADE
);

CREATE TABLE "urls" (
	"id" int PRIMARY KEY,
	"slug" varchar(255) UNIQUE NOT NULL,
	"destination" varchar(1024) NOT NULL,
	"expires" bigint NOT NULL,
	CONSTRAINT fk_slug
		FOREIGN KEY("slug")
			REFERENCES "slugs"("slug")
			ON DELETE CASCADE
);

CREATE INDEX ON "slugs" ("slug");

CREATE INDEX ON "files" ("slug");

CREATE INDEX ON "urls" ("slug");

-- migrate:down
DROP TABLE "slugs";

DROP TABLE "files";

DROP TABLE "urls";