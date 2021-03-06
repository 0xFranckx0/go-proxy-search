package algoliasearch

import "errors"

type indexIterator struct {
	cursor string
	index  Index
	opts   *RequestOptions
	page   BrowseRes
	params Map
	pos    int
}

var NoMoreHitsErr error = errors.New("No more hits")

// newIndexIterator instantiates a IndexIterator on the `index` and according
// to the given `params`. It is also trying to load the first page of results
// and return an error if something goes wrong.
func newIndexIterator(index Index, params Map, opts *RequestOptions) (it *indexIterator, err error) {
	it = &indexIterator{
		cursor: "",
		index:  index,
		opts:   opts,
		params: duplicateMap(params),
		pos:    0,
	}
	err = it.loadNextPage()
	return
}

func (it *indexIterator) Next() (res Map, err error) {
	// Abort if the user call `Next()` on a IndexIterator that has been
	// initialized without being able to load the first page.
	if len(it.page.Hits) == 0 {
		//err = errors.New("No more hits")
		err = NoMoreHitsErr
		return
	}

	// If the last element of the page has been reached, the next one is loaded
	// or returned an error if the last element of the last page has already
	// been returned.
	if it.pos == len(it.page.Hits) {
		if it.cursor == "" {
			//err = errors.New("No more hits")
			err = NoMoreHitsErr
		} else {
			err = it.loadNextPage()
		}

		if err != nil {
			return
		}
	}

	res = it.page.Hits[it.pos]
	it.pos++

	return
}

// loadNextPage is used internally to load the next page of results, using the
// underlying Browse cursor.
func (it *indexIterator) loadNextPage() (err error) {
	if it.page, err = it.index.BrowseWithRequestOptions(it.params, it.cursor, it.opts); err != nil {
		return
	}

	// Return an error if the newly loaded pages contains no results
	if len(it.page.Hits) == 0 {
		//err = errors.New("No more hits")
		err = NoMoreHitsErr
		return
	}

	it.cursor = it.page.Cursor
	it.pos = 0
	return
}
