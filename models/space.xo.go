package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Space represents a row from 'ultrachat.spaces'.
type Space struct {
	SpaceID string `json:"spaceId"` // spaceId
	Name    string `json:"name"`    // name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Space exists in the database.
func (s *Space) Exists() bool {
	return s._exists
}

// Deleted returns true when the Space has been marked for deletion from
// the database.
func (s *Space) Deleted() bool {
	return s._deleted
}

// Insert inserts the Space to the database.
func (s *Space) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO ultrachat.spaces (` +
		`spaceId, name` +
		`) VALUES (` +
		`?, ?` +
		`)`
	// run
	logf(sqlstr, s.SpaceID, s.Name)
	if _, err := db.ExecContext(ctx, sqlstr, s.SpaceID, s.Name); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Update updates a Space in the database.
func (s *Space) Update(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case s._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE ultrachat.spaces SET ` +
		`name = ? ` +
		`WHERE spaceId = ?`
	// run
	logf(sqlstr, s.Name, s.SpaceID)
	if _, err := db.ExecContext(ctx, sqlstr, s.Name, s.SpaceID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Space to the database.
func (s *Space) Save(ctx context.Context, db DB) error {
	if s.Exists() {
		return s.Update(ctx, db)
	}
	return s.Insert(ctx, db)
}

// Upsert performs an upsert for Space.
func (s *Space) Upsert(ctx context.Context, db DB) error {
	switch {
	case s._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO ultrachat.spaces (` +
		`spaceId, name` +
		`) VALUES (` +
		`?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`spaceId = VALUES(spaceId), name = VALUES(name)`
	// run
	logf(sqlstr, s.SpaceID, s.Name)
	if _, err := db.ExecContext(ctx, sqlstr, s.SpaceID, s.Name); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Delete deletes the Space from the database.
func (s *Space) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM ultrachat.spaces ` +
		`WHERE spaceId = ?`
	// run
	logf(sqlstr, s.SpaceID)
	if _, err := db.ExecContext(ctx, sqlstr, s.SpaceID); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// SpaceBySpaceID retrieves a row from 'ultrachat.spaces' as a Space.
//
// Generated from index 'spaces_spaceId_pkey'.
func SpaceBySpaceID(ctx context.Context, db DB, spaceID string) (*Space, error) {
	// query
	const sqlstr = `SELECT ` +
		`spaceId, name ` +
		`FROM ultrachat.spaces ` +
		`WHERE spaceId = ?`
	// run
	logf(sqlstr, spaceID)
	s := Space{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, spaceID).Scan(&s.SpaceID, &s.Name); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}