package request

// UserRequest represents the input data for creating a new user.
// @Summary User Input Data
// @Description Structure containing the required fields for creating a new user.
type UserRequest struct {
	// User's email (required and must be a valid email address).
	// Example: user@example.com
	// @json
	// @jsonTag email
	// @jsonExample user@example.com
	// @binding required,email
	Email string `json:"email" binding:"required,email"`

	// User's password (required, minimum of 6 characters, and must contain at least one of the characters: !@#$%*).
	// @json
	// @jsonTag password
	// @jsonExample P@ssw0rd!
	// @binding required,min=6,containsany=!@#$%*
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%&*"`

	// User's name (required, minimum of 4 characters, maximum of 100 characters).
	// Example: John Doe
	// @json
	// @jsonTag name
	// @jsonExample John Doe
	// @binding required,min=4,max=100
	Name string `json:"name" binding:"required,min=4,max=50"`

	// User's age (required, must be between 1 and 140).
	// @json
	// @jsonTag age
	// @jsonExample 30
	Age int `json:"age" binding:"required,min=2,max=120"`
}

type UserUpdateRequest struct {
	// User's name (required, minimum of 4 characters, maximum of 100 characters).
	// Example: John Doe
	// @json
	// @jsonTag name
	// @jsonExample John Doe
	// @binding required,min=4,max=100
	Name string `json:"name" binding:"omitempty,min=4,max=50"`

	// User's age (required, must be between 1 and 140).
	// @json
	// @jsonTag age
	// @jsonExample 30
	// @binding required,min=1,max=140
	Age int `json:"age" binding:"omitempty,min=2,max=120"`
}
