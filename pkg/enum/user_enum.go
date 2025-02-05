// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package enum

import (
	"errors"
	"fmt"
)

const (
	// UserRoleSuperadmin is a UserRole of type superadmin.
	UserRoleSuperadmin UserRole = "superadmin"
	// UserRoleCustomer is a UserRole of type customer.
	UserRoleCustomer UserRole = "customer"
)

var ErrInvalidUserRole = errors.New("not a valid UserRole")

// String implements the Stringer interface.
func (x UserRole) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x UserRole) IsValid() bool {
	_, err := ParseUserRole(string(x))
	return err == nil
}

var _UserRoleValue = map[string]UserRole{
	"superadmin": UserRoleSuperadmin,
	"customer":   UserRoleCustomer,
}

// ParseUserRole attempts to convert a string to a UserRole.
func ParseUserRole(name string) (UserRole, error) {
	if x, ok := _UserRoleValue[name]; ok {
		return x, nil
	}
	return UserRole(""), fmt.Errorf("%s is %w", name, ErrInvalidUserRole)
}
