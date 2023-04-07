package controllers

import (
	"encoding/json"
	"io"

	"github.com/markbates/goth"
)

// func FnRegex(arg string) {
// 	re := regexp.MustCompile("(^[A-Za-z]{3,16})([ ]{0,1})([A-Za-z]{3,16})?([ ]{0,1})?([A-Za-z]{3,16})?([ ]{0,1})?([A-Za-z]{3,16})*$")
// 	if !re.MatchString(arg) {
// 		color.Red("\nInvalid full name!")
// 		os.Exit(0)
// 	}
// }

// func FN(arg string) {
// 	manualStr := "https://api.linkedin.com/v1/people/~:(id,first-name,email-address,last-name,headline,picture-url,industry,summary,specialties,positions:(id,title,summary,start-date,end-date,is-current,company:(id,name,type,size,industry,ticker)),educations:(id,school-name,field-of-study,start-date,end-date,degree,activities,notes),associations,interests,num-recommenders,date-of-birth,publications:(id,title,publisher:(name),authors:(id,name),date,url,summary),patents:(id,title,summary,number,status:(id,name),office:(name),inventors:(id,name),date,url),languages:(id,language:(name),proficiency:(level,name)),skills:(id,skill:(name)),certifications:(id,name,authority:(name),number,start-date,end-date),courses:(id,name,number),recommendations-received:(id,recommendation-type,recommendation-text,recommender),honors-awards,three-current-positions,three-past-positions,volunteer)?format=json"
// 	fn := arg
// 	if fn != "" {
// 		client := &http.Client{}
// 		req, _ := http.NewRequest("GET", manualStr+fn+"/json/", nil)
// 		req.Header.Set("User-Agent", "mosint")
// 		resp, _ := client.Do(req)
// 		body, _ := io.ReadAll(resp.Body)
// 		json.Unmarshal(body, &utils.Fn_result)
// 	}
// }

// func Googling(fn string) {
// 	q := "intext:" + string('"') + fn + string('"')
// 	res, err := googlesearch.Search(nil, q)
// 	if err != nil {
// 		panic(err)
// 	}
// 	size := len(res)
// 	for i := 0; i < size; i++ {
// 		utils.Googling_result = append(utils.Googling_result, res[i].URL)
// 	}
// }

// func FN(arg string) {
// 	fn := arg
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", utils.EndpointProfile+fn, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	body, _ := io.ReadAll(resp.Body)
// 	json.Unmarshal(body, &utils.FullName_result)
// }

const userEndpoint string = "//api.linkedin.com/v2/me?projection=(id,firstName,lastName,profilePicture(displayImage~:playableStreams))"

func FN(reader io.Reader, user *goth.User) error {
	u := struct {
		ID        string `json:"id"`
		FirstName struct {
			PreferredLocale struct {
				Country  string `json:"country"`
				Language string `json:"language"`
			} `json:"preferredLocale"`
			Localized map[string]string `json:"localized"`
		} `json:"firstName"`
		LastName struct {
			Localized       map[string]string
			PreferredLocale struct {
				Country  string `json:"country"`
				Language string `json:"language"`
			} `json:"preferredLocale"`
		} `json:"lastName"`
		ProfilePicture struct {
			DisplayImage struct {
				Elements []struct {
					AuthorizationMethod string `json:"authorizationMethod"`
					Identifiers         []struct {
						Identifier     string `json:"identifier"`
						IdentifierType string `json:"identifierType"`
					} `json:"identifiers"`
				} `json:"elements"`
			} `json:"displayImage~"`
		} `json:"profilePicture"`
	}{}

	err := json.NewDecoder(reader).Decode(&u)
	if err != nil {
		return err
	}

	user.FirstName = u.FirstName.Localized[u.FirstName.PreferredLocale.Language+"_"+u.FirstName.PreferredLocale.Country]
	user.LastName = u.LastName.Localized[u.LastName.PreferredLocale.Language+"_"+u.LastName.PreferredLocale.Country]
	user.Name = user.FirstName + " " + user.LastName
	user.NickName = user.FirstName
	user.UserID = u.ID

	avatarURL := ""
	// loop all displayimage elements
	for _, element := range u.ProfilePicture.DisplayImage.Elements {
		// only retrieve data where the authorization method allows public (unauthorized) access
		if element.AuthorizationMethod == "PUBLIC" {
			for _, identifier := range element.Identifiers {
				// check to ensure the identifier type is a url linking to the image
				if identifier.IdentifierType == "EXTERNAL_URL" {
					avatarURL = identifier.Identifier
					// we only need the first image url
					break
				}
			}
		}
		// if we have a valid image, exit the loop as we only support a single avatar image
		if len(avatarURL) > 0 {
			break
		}
	}

	user.AvatarURL = avatarURL

	return err
}

type User struct {
	RawData           map[string]interface{}
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	IDToken           string
}
