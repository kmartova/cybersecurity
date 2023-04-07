package utils

import (
	"fmt"

	"github.com/alpkeskin/mosint/cmd/models"
	"github.com/fatih/color"
)

func PrintGoogle(googling_result []string) {
	fmt.Println(TitleMap["Googling"])
	if len(googling_result) == 0 {
		color.Red("|- No results found")
	} else {
		for _, v := range googling_result {
			fmt.Println("|- "+v, color.GreenString("\u2714"))
		}
	}
}

func PrintSocial(social_result []string) {
	fmt.Println(TitleMap["Social"])
	for _, v := range social_result {
		fmt.Println("|- " + v)
	}
}

func PrintIPAPI(apiapi_result models.IPAPIStruct) {
	fmt.Println(TitleMap["IPAPI"])
	fmt.Println("|- IP:", color.YellowString(apiapi_result.IP))
	fmt.Println("|- City:", color.WhiteString(apiapi_result.City))
	fmt.Println("|- Region:", color.WhiteString(apiapi_result.Region))
	fmt.Println("|- Country:", color.WhiteString(apiapi_result.CountryName))
	fmt.Println("|- Country Code:", color.WhiteString(apiapi_result.CountryCode))
	fmt.Println("|- Timezone:", color.WhiteString(apiapi_result.Timezone))
	fmt.Println("|- Organization:", color.WhiteString(apiapi_result.Org))
	fmt.Println("|- ASN:", color.WhiteString(apiapi_result.Asn))
}

// func PrintFN(Fn_result models.LinkedinProfile) {
// 	fmt.Println(TitleMap["FullName"])
// 	fmt.Println("|- CurrentShare:", color.YellowString(Fn_result.CurrentShare))
// 	fmt.Println("|- EmailAddress:", color.WhiteString(Fn_result.EmailAddress))
// 	fmt.Println("|- PictureURL:", color.WhiteString(Fn_result.PictureURL))
// 	fmt.Println("|- ProfileID:", color.WhiteString(Fn_result.ProfileID))
// 	fmt.Println("|- Specialties:", color.WhiteString(Fn_result.Specialties))
// 	fmt.Println("|- IsConnectionsCapped:", string(Fn_result.IsConnectionsCapped))
// 	fmt.Println("|- NumConnections:", int(Fn_result.NumConnections))
// 	fmt.Println("|- FirstName:", color.WhiteString(Fn_result.Headline))
// }

func PrintFN(FullName_result models.Profile) {
	fmt.Println(TitleMap["FullName"])
	fmt.Println("|- CurrentShare:", color.YellowString(FullName_result.ID))
	fmt.Println("|- EmailAddress:", color.WhiteString(FullName_result.LocalizedHeadline))
	fmt.Println("|- PictureURL:", color.WhiteString(FullName_result.ProfilePicture.DisplayImage))
	fmt.Println("|- ProfileID:", color.WhiteString(FullName_result.FirstName.PreferredLocale.Country))
	fmt.Println("|- Specialties:", color.WhiteString(FullName_result.LocalizedLastName))
	fmt.Println("|- IsConnectionsCapped:", string(FullName_result.LocalizedFirstName))
	fmt.Println("|- NumConnections:", string(FullName_result.ProfilePicture.DisplayImage))
	fmt.Println("|- FirstName:", color.WhiteString(FullName_result.ProfilePicture.DisplayImage))
}
