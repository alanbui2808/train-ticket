package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // user repo ur
// func (ur *UserRepo) GetUserInfo() string {
// 	return "Tipjs"
// }

// INTERFACE
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

// now private (everything works through UserRepo interface)
type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

// GetUserByEmail implements IUserRepository.
// turn userRepo into IUserRepo
func (u *userRepository) GetUserByEmail(email string) bool {
	return true
}
