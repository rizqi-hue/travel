package repository

type paginate struct {
	Limit  int64
	Offset int64
}

func Pagination(page int64, size int64) *paginate {
	limit := size

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 5
	}

	offset := (page * limit) - limit

	return &paginate{
		Limit:  limit,
		Offset: offset,
	}
}
