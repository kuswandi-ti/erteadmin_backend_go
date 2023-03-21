package user

type UserFormatter struct {
	ID           int    `json:"id"`
	LingkunganID int    `json:"lingkungan_id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	ImageURL     string `json:"image_url"`
	Role         string `json:"role"`
}

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
