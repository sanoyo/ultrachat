package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// User represents a row from 'ultrachat.users'.
type User struct {
	UserID string `json:"userId"` // userId
	Name   string `json:"name"`   // name
	Email  string `json:"email"`  // email
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the User has been marked for deletion from
// the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO ultrachat.users (` +
		`userId, name, email` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`
	// run
	logf(sqlstr, u.UserID, u.Name, u.Email)
	if _, err := db.ExecContext(ctx, sqlstr, u.UserID, u.Name, u.Email); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE ultrachat.users SET ` +
		`name = ?, email = ? ` +
		`WHERE userId = ?`
	// run
	logf(sqlstr, u.Name, u.Email, u.UserID)
	if _, err := db.ExecContext(ctx, sqlstr, u.Name, u.Email, u.UserID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db DB) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for User.
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO ultrachat.users (` +
		`userId, name, email` +
		`) VALUES (` +
		`?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`userId = VALUES(userId), name = VALUES(name), email = VALUES(email)`
	// run
	logf(sqlstr, u.UserID, u.Name, u.Email)
	if _, err := db.ExecContext(ctx, sqlstr, u.UserID, u.Name, u.Email); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM ultrachat.users ` +
		`WHERE userId = ?`
	// run
	logf(sqlstr, u.UserID)
	if _, err := db.ExecContext(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UserByUserID retrieves a row from 'ultrachat.users' as a User.
//
// Generated from index 'users_userId_pkey'.
func UserByUserID(ctx context.Context, db DB, userID string) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`userId, name, email ` +
		`FROM ultrachat.users ` +
		`WHERE userId = ?`
	// run
	logf(sqlstr, userID)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, userID).Scan(&u.UserID, &u.Name, &u.Email); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}