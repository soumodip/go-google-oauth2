package constants

func ParseConstants() map[string]string{
	//DECLARE ALL CONSTANTS
	constants := make(map[string]string) 
	constants["PORT"] = "4000"
	constants["GOOGLE_ID"] = "506824857156-n4u8kqfitis3psm4h8gt2nemv9nfqrhs.apps.googleusercontent.com"
	constants["GOOGLE_SECRET"] = "8DDvVTElAde3Lfxhs_wJ_6u9"
	constants["GOOGLE_REDIRECT_URI"] = "http://localhost:" + constants["PORT"] + "/google/authenticate/callback"
	constants["GOOGLE_SCOPE"] = "https://www.googleapis.com/auth/userinfo.email"
	//SEND THE MAP
	return constants
}

