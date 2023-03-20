package user

type UserFormatter struct {
	ID           int    `json:"id"`
	LingkunganID int    `json:"lingkungan_id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	ImageURL     string `json:"image_url"`
	Role         string `json:"role"`
}

// type UserDetailFormatter struct {
// 	ID           int    `json:"id"`
// 	LingkunganID int    `json:"lingkungan_id"`
// 	Email        string `json:"email"`
// 	Token        string `json:"token"`
// 	ImageURL     string `json:"image_url"`
// 	Role         string `json:"role"`
// 	Lingkungan   UserLingkunganFormatter
// }

// type UserLingkunganFormatter struct {
// 	Nama string `json:"nama"`
// }

func FormatUser(user User, token string) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.LingkunganID = user.LingkunganID
	userFormatter.Email = user.Email
	userFormatter.Token = token
	userFormatter.ImageURL = user.Avatar
	userFormatter.Role = user.Role

	return userFormatter
}

// func FormatUserDetail(user User, token string) UserDetailFormatter {
// 	userDetailFormatter := UserDetailFormatter{}
// 	userDetailFormatter.ID = user.ID
// 	userDetailFormatter.LingkunganID = user.LingkunganID
// 	userDetailFormatter.Email = user.Email
// 	userDetailFormatter.Token = token
// 	userDetailFormatter.ImageURL = user.Avatar
// 	userDetailFormatter.Role = user.Role

// 	lingkungan := user.Lingkungan

// 	userLingkunganFormatter := UserLingkunganFormatter{}
// 	userLingkunganFormatter.Nama = lingkungan.Nama

// 	userDetailFormatter.Lingkungan = userLingkunganFormatter

// 	return userDetailFormatter
// }
