package models

type Slug struct {
	Id     int
	IsFile bool
	Slug   string
}

type SlugStore struct {
}

func (s *SlugStore) GetById(id int) (*Slug, error) {
	row := db.QueryRow(`
	select 
	slugs.id, 
	slugs.is_file, 
	slugs.slug 
	from slugs
	where slugs.id = $1;`, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	q := Slug{}
	err := row.Scan(&q.Id, &q.IsFile, &q.Slug)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (s *SlugStore) GetBySlug(slug string) (*Slug, error) {
	row := db.QueryRow(`
	select 
	slugs.id, 
	slugs.is_file, 
	slugs.slug 
	from slugs
	where slugs.slug = $1;`, slug)
	if row.Err() != nil {
		return nil, row.Err()
	}

	q := Slug{}
	err := row.Scan(&q.Id, &q.IsFile, &q.Slug)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (s *SlugStore) Insert(isFile bool, slug string) error {
	_, err := db.Exec("insert into slugs(is_file, slug) values ($1, $2);", isFile, slug)
	if err != nil {
		return err
	}
	return nil
}

func (s *SlugStore) DeleteById(id int) error {
	_, err := db.Exec("delete from slugs where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SlugStore) DeleteBySlug(slug string) error {
	_, err := db.Exec("delete from slugs where slug = $1", slug)
	if err != nil {
		return err
	}
	return nil
}
