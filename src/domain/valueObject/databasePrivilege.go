package valueObject

import (
	"errors"
	"slices"
	"strings"

	voHelper "github.com/speedianet/os/src/domain/valueObject/helper"
)

type DatabasePrivilege string

var ValidDatabasePrivileges = []string{
	"ALL PRIVILEGES",
	"ALL",
	"ALTER ROUTINE",
	"ALTER SYSTEM",
	"ALTER",
	"BYPASSRLS",
	"CONNECT",
	"CREATE DOMAIN",
	"CREATE FUNCTION",
	"CREATE GROUP",
	"CREATE INDEX",
	"CREATE LANGUAGE",
	"CREATE PROCEDURE",
	"CREATE ROLE",
	"CREATE ROUTINE",
	"CREATE SCHEMA",
	"CREATE TABLE",
	"CREATE TEMP",
	"CREATE TEMPORARY TABLES",
	"CREATE TRIGGER",
	"CREATE TYPE",
	"CREATE USER",
	"CREATE VIEW",
	"CREATE",
	"CREATEDB",
	"CREATEROLE",
	"DELETE HISTORY",
	"DELETE",
	"DROP",
	"EVENT",
	"EXECUTE",
	"FILE",
	"INDEX",
	"INSERT",
	"LOCK TABLES",
	"PASSWORDADMIN",
	"PROCESS",
	"PROXY",
	"REFERENCES",
	"RELOAD",
	"REPLICATION CLIENT",
	"REPLICATION SLAVE",
	"REPLICATION",
	"SELECT",
	"SET",
	"SHOW VIEW",
	"SHUTDOWN",
	"SUPER",
	"SUPERUSER",
	"TEMPORARY",
	"TRIGGER",
	"TRUNCATE",
	"UPDATE",
	"USAGE",
}

func NewDatabasePrivilege(value interface{}) (
	dbPrivilege DatabasePrivilege, err error,
) {
	stringValue, err := voHelper.InterfaceToString(value)
	if err != nil {
		return dbPrivilege, errors.New("DatabasePrivilegeMustBeString")
	}
	stringValue = strings.ReplaceAll(stringValue, "-", " ")
	stringValue = strings.ToUpper(stringValue)

	if !slices.Contains(ValidDatabasePrivileges, stringValue) {
		return dbPrivilege, errors.New("InvalidDatabasePrivilege")
	}
	return DatabasePrivilege(stringValue), nil
}

func (vo DatabasePrivilege) String() string {
	return string(vo)
}
