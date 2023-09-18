package models

import "database/sql"

type Url struct {
	Id          int
	Slug        string
	Destination string
	Expires     int
}

type UrlStore struct {
	DB *sql.DB
}

func (s *UrlStore) GetById(id int) (*Url, error) {
	row := s.DB.QueryRow(`
	select 
	id,
	slug,
	destination,
	expires
	from urls
	where id = $1;`, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	q := Url{}
	err := row.Scan(&q.Id, &q.Slug, &q.Destination, &q.Expires)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (s *UrlStore) GetBySlug(slug string) (*Url, error) {
	row := s.DB.QueryRow(`
	select 
	id,
	slug,
	destination,
	expires
	from urls
	where slug = $1;`, slug)
	if row.Err() != nil {
		return nil, row.Err()
	}

	q := Url{}
	err := row.Scan(&q.Id, &q.Slug, &q.Destination, &q.Expires)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (s *UrlStore) Insert(id int, slug, destination string, expires int) error {
	_, err := s.DB.Exec(`
	insert into 
	urls(id, slug, destination, expires)
	values ($1, $2, $3, $4);
	`, id, slug, destination, expires)
	if err != nil {
		return err
	}
	return nil
}

func (s *UrlStore) DeleteById(id int) error {
	// delete from table `slugs` so it cascades
	_, err := s.DB.Exec("delete from slugs where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UrlStore) DeleteBySlug(slug string) error {
	// delete from table `slugs` so it cascades
	_, err := s.DB.Exec("delete from slugs where slug = $1", slug)
	if err != nil {
		return err
	}
	return nil
}
