package enum

//go:generate go-enum -f=user.go

// ENUM(superadmin, customer)
type UserRole string
