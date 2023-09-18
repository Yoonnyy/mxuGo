package models

import "database/sql"

type File struct {
	Id               int
	OriginalFilename string
	Slug             string
	Size             int
	Expires          int
}

type FileStore struct {
	DB *sql.DB
}

func (s *FileStore) GetById(id int) (*File, error) {
	row := s.DB.QueryRow(`
	select 
	id,
	original_filename,
	slug,
	size,
	expires
	from files
	where id = $1;`, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	q := File{}
	err := row.Scan(&q.Id, &q.OriginalFilename, &q.Slug, &q.Size, &q.Expires)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (s *FileStore) GetBySlug(slug string) (*File, error) {
	row := s.DB.QueryRow(`
	select 
	id,
	original_filename,
	slug,
	size,
	expires
	from files
	where slug = $1;`, slug)
	if row.Err() != nil {
		return nil, row.Err()
	}

	q := File{}
	err := row.Scan(&q.Id, &q.OriginalFilename, &q.Slug, &q.Size, &q.Expires)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (s *FileStore) Insert(id int, originalFilename string, slug string, size, expires int) error {
	_, err := s.DB.Exec(`
	insert into 
	files(id, original_filename, slug, size, expires)
	values ($1, $2, $3, $4, $5);
	`, id, originalFilename, slug, size, expires)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStore) DeleteById(id int) error {
	// delete from table `slugs` so it cascades
	_, err := s.DB.Exec("delete from slugs where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStore) DeleteBySlug(slug string) error {
	// delete from table `slugs` so it cascades
	_, err := s.DB.Exec("delete from slugs where slug = $1", slug)
	if err != nil {
		return err
	}
	return nil
}
