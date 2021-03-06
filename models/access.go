// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"strings"
	"time"
)

// Access types.
const (
	AU_READABLE = iota + 1
	AU_WRITABLE
)

// Access represents the accessibility of user to repository.
type Access struct {
	Id       int64
	UserName string    `xorm:"unique(s)"`
	RepoName string    `xorm:"unique(s)"`
	Mode     int       `xorm:"unique(s)"`
	Created  time.Time `xorm:"created"`
}

// AddAccess adds new access record.
func AddAccess(access *Access) error {
	access.UserName = strings.ToLower(access.UserName)
	access.RepoName = strings.ToLower(access.RepoName)
	_, err := orm.Insert(access)
	return err
}

// UpdateAccess updates access information.
func UpdateAccess(access *Access) error {
	access.UserName = strings.ToLower(access.UserName)
	access.RepoName = strings.ToLower(access.RepoName)
	_, err := orm.Id(access.Id).Update(access)
	return err
}

// HasAccess returns true if someone can read or write to given repository.
func HasAccess(userName, repoName string, mode int) (bool, error) {
	return orm.Get(&Access{
		Id:       0,
		UserName: strings.ToLower(userName),
		RepoName: strings.ToLower(repoName),
		Mode:     mode,
	})
}
