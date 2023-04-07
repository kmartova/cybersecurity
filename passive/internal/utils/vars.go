package utils

import (
	"os"

	"github.com/alpkeskin/mosint/cmd/models"
	"github.com/olekukonko/tablewriter"
)

var TitleMap = map[string]string{
	"Googling": "\nGoogling Results:",
	"Social":   "\nSocial Media Results:",
	"IPAPI":    "\nIPAPI Results:",
	"FN":       "\nFN Results:",
}

var (
	ConfigReturn map[string]interface{}
	LookupTable  *tablewriter.Table = tablewriter.NewWriter(os.Stdout)
)

var (
	Number_result      []string
	Address_result     []string
	Googling_result    []string
	Intelx_result      []string
	Social_result      []string
	Psbdmp_result      models.PsbdmpStruct
	Lookup_temp_result [][]string
	Ipapi_result       models.IPAPIStruct
	Fn_result          models.LinkedinProfile
	// FullName_result    models.Profile
)

const (
	IntelxDefaultMaxResults = 10
)

const (
	IPAPIURL              = "https://ipapi.co/"
	AdobeEndpoint         = "https://auth.services.adobe.com/signin/v2/users/accounts"
	DiscordEndpoint       = "https://discord.com/api/v9/auth/register"
	InstagramCSRFEndpoint = "https://www.instagram.com/accounts/emailsignup/"
	InstagramEndpoint     = "https://www.instagram.com/accounts/web_create_ajax/attempt/"
	SpotifyEndpoint       = "https://spclient.wg.spotify.com/signup/public/v1/account"
	TwitterEndpoint       = "https://api.twitter.com/i/users/email_available.json"
	EndpointProfile       = "https://api.linkedin.com/v2/me?projection=(id,firstName,lastName,vanityName,localizedHeadline,localizedFirstName,localizedLastName,localizedHeadline,headline,profilePicture(displayImage~:playableStreams))"
	fullRequestURL        = "https://api.linkedin.com/v1/people/~:(id,first-name,email-address,last-name,headline,picture-url,industry,summary,specialties,positions:(id,title,summary,start-date,end-date,is-current,company:(id,name,type,size,industry,ticker)),educations:(id,school-name,field-of-study,start-date,end-date,degree,activities,notes),associations,interests,num-recommenders,date-of-birth,publications:(id,title,publisher:(name),authors:(id,name),date,url,summary),patents:(id,title,summary,number,status:(id,name),office:(name),inventors:(id,name),date,url),languages:(id,language:(name),proficiency:(level,name)),skills:(id,skill:(name)),certifications:(id,name,authority:(name),number,start-date,end-date),courses:(id,name,number),recommendations-received:(id,recommendation-type,recommendation-text,recommender),honors-awards,three-current-positions,three-past-positions,volunteer)?format=json"
	basicRequestURL       = "https://api.linkedin.com/v1/people/~:(id,first-name,last-name,headline,picture-url,industry,summary,specialties,positions:(id,title,summary,start-date,end-date,is-current,company:(id,name,type,size,industry,ticker)),educations:(id,school-name,field-of-study,start-date,end-date,degree,activities,notes),associations,interests,num-recommenders,date-of-birth,publications:(id,title,publisher:(name),authors:(id,name),date,url,summary),patents:(id,title,summary,number,status:(id,name),office:(name),inventors:(id,name),date,url),languages:(id,language:(name),proficiency:(level,name)),skills:(id,skill:(name)),certifications:(id,name,authority:(name),number,start-date,end-date),courses:(id,name,number),recommendations-received:(id,recommendation-type,recommendation-text,recommender),honors-awards,three-current-positions,three-past-positions,volunteer)?format=json"
)
