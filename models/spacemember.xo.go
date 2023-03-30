package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// SpaceMember represents a row from 'ultrachat.space_members'.
type SpaceMember struct {
	UserID  string `json:"userId"`  // userId
	SpaceID string `json:"spaceId"` // spaceId
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the SpaceMember exists in the database.
func (sm *SpaceMember) Exists() bool {
	return sm._exists
}

// Deleted returns true when the SpaceMember has been marked for deletion from
// the database.
func (sm *SpaceMember) Deleted() bool {
	return sm._deleted
}

// Insert inserts the SpaceMember to the database.
func (sm *SpaceMember) Insert(ctx context.Context, db DB) error {
	switch {
	case sm._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case sm._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO ultrachat.space_members (` +
		`userId, spaceId` +
		`) VALUES (` +
		`?, ?` +
		`)`
	// run
	logf(sqlstr, sm.UserID, sm.SpaceID)
	if _, err := db.ExecContext(ctx, sqlstr, sm.UserID, sm.SpaceID); err != nil {
		return logerror(err)
	}
	// set exists
	sm._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the SpaceMember from the database.
func (sm *SpaceMember) Delete(ctx context.Context, db DB) error {
	switch {
	case !sm._exists: // doesn't exist
		return nil
	case sm._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM ultrachat.space_members ` +
		`WHERE userId = ? AND spaceId = ?`
	// run
	logf(sqlstr, sm.UserID, sm.SpaceID)
	if _, err := db.ExecContext(ctx, sqlstr, sm.UserID, sm.SpaceID); err != nil {
		return logerror(err)
	}
	// set deleted
	sm._deleted = true
	return nil
}

// SpaceMembersBySpaceID retrieves a row from 'ultrachat.space_members' as a SpaceMember.
//
// Generated from index 'spaceId'.
func SpaceMembersBySpaceID(ctx context.Context, db DB, spaceID string) ([]*SpaceMember, error) {
	// query
	const sqlstr = `SELECT ` +
		`userId, spaceId ` +
		`FROM ultrachat.space_members ` +
		`WHERE spaceId = ?`
	// run
	logf(sqlstr, spaceID)
	rows, err := db.QueryContext(ctx, sqlstr, spaceID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*SpaceMember
	for rows.Next() {
		sm := SpaceMember{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&sm.UserID, &sm.SpaceID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &sm)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SpaceMemberByUserIDSpaceID retrieves a row from 'ultrachat.space_members' as a SpaceMember.
//
// Generated from index 'space_members_userId_spaceId_pkey'.
func SpaceMemberByUserIDSpaceID(ctx context.Context, db DB, userID, spaceID string) (*SpaceMember, error) {
	// query
	const sqlstr = `SELECT ` +
		`userId, spaceId ` +
		`FROM ultrachat.space_members ` +
		`WHERE userId = ? AND spaceId = ?`
	// run
	logf(sqlstr, userID, spaceID)
	sm := SpaceMember{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, userID, spaceID).Scan(&sm.UserID, &sm.SpaceID); err != nil {
		return nil, logerror(err)
	}
	return &sm, nil
}

// User returns the User associated with the SpaceMember's (UserID).
//
// Generated from foreign key 'space_members_ibfk_1'.
func (sm *SpaceMember) User(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, sm.UserID)
}

// Space returns the Space associated with the SpaceMember's (SpaceID).
//
// Generated from foreign key 'space_members_ibfk_2'.
func (sm *SpaceMember) Space(ctx context.Context, db DB) (*Space, error) {
	return SpaceBySpaceID(ctx, db, sm.SpaceID)
}