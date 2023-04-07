package models

import "net/http"

type EmailRepStruct struct {
	Email      string `json:"email"`
	Reputation string `json:"reputation"`
	Suspicious bool   `json:"suspicious"`
	References int    `json:"references"`
	Details    struct {
		Blacklisted             bool     `json:"blacklisted"`
		MaliciousActivity       bool     `json:"malicious_activity"`
		MaliciousActivityRecent bool     `json:"malicious_activity_recent"`
		CredentialsLeaked       bool     `json:"credentials_leaked"`
		CredentialsLeakedRecent bool     `json:"credentials_leaked_recent"`
		DataBreach              bool     `json:"data_breach"`
		FirstSeen               string   `json:"first_seen"`
		LastSeen                string   `json:"last_seen"`
		DomainExists            bool     `json:"domain_exists"`
		DomainReputation        string   `json:"domain_reputation"`
		NewDomain               bool     `json:"new_domain"`
		DaysSinceDomainCreation int      `json:"days_since_domain_creation"`
		SuspiciousTld           bool     `json:"suspicious_tld"`
		Spam                    bool     `json:"spam"`
		FreeProvider            bool     `json:"free_provider"`
		Disposable              bool     `json:"disposable"`
		Deliverable             bool     `json:"deliverable"`
		AcceptAll               bool     `json:"accept_all"`
		ValidMx                 bool     `json:"valid_mx"`
		PrimaryMx               string   `json:"primary_mx"`
		Spoofable               bool     `json:"spoofable"`
		SpfStrict               bool     `json:"spf_strict"`
		DmarcEnforced           bool     `json:"dmarc_enforced"`
		Profiles                []string `json:"profiles"`
	} `json:"details"`
}

type IPAPIStruct struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

type AdobeResponse []struct {
	HasT2ELinked bool `json:"hasT2ELinked"`
}

type DiscordResponse struct {
	Errors struct {
		Username struct {
			Errors []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"_errors"`
		} `json:"email"`
	} `json:"errors"`
}

type SpotifyResponse struct {
	Status int `json:"status"`
	Errors struct {
		Esername string `json:"email"`
	} `json:"errors"`
	Country                        string `json:"country"`
	CanAcceptLicensesInOneStep     bool   `json:"can_accept_licenses_in_one_step"`
	RequiresMarketingOptIn         bool   `json:"requires_marketing_opt_in"`
	RequiresMarketingOptInText     bool   `json:"requires_marketing_opt_in_text"`
	MinimumAge                     int    `json:"minimum_age"`
	CountryGroup                   string `json:"country_group"`
	SpecificLicenses               bool   `json:"specific_licenses"`
	TermsConditionsAcceptance      string `json:"terms_conditions_acceptance"`
	PrivacyPolicyAcceptance        string `json:"privacy_policy_acceptance"`
	SpotifyMarketingMessagesOption string `json:"spotify_marketing_messages_option"`
	PretickEula                    bool   `json:"pretick_eula"`
	ShowCollectPersonalInfo        bool   `json:"show_collect_personal_info"`
	UseAllGenders                  bool   `json:"use_all_genders"`
	UseOtherGender                 bool   `json:"use_other_gender"`
	DateEndianness                 int    `json:"date_endianness"`
	IsCountryLaunched              bool   `json:"is_country_launched"`
	AllowedCallingCodes            []struct {
		CountryCode string `json:"country_code"`
		CallingCode string `json:"calling_code"`
	} `json:"allowed_calling_codes"`
	PushNotifications bool `json:"push-notifications"`
}

type TwitterResponse struct {
	Valid bool   `json:"valid"`
	Msg   string `json:"msg"`
	Taken bool   `json:"taken"`
}

type LinkedinProfile struct {
	// ProfileID represents the Unique ID every Linkedin profile has.
	ProfileID string `json:"id"`
	// FirstName represents the user's first name.
	FirstName string `json:"first_name"`
	// LastName represents the user's last name.
	LastName string `json:"last-name"`
	// MaidenName represents the user's maiden name, if they have one.
	MaidenName string `json:"maiden-name"`
	// FormattedName represents the user's formatted name, based on locale.
	FormattedName string `json:"formatted-name"`
	// PhoneticFirstName represents the user's first name, spelled phonetically.
	PhoneticFirstName string `json:"phonetic-first-name"`
	// PhoneticFirstName represents the user's last name, spelled phonetically.
	PhoneticLastName string `json:"phonetic-last-name"`
	// Headline represents a short, attention grabbing description of the user.
	Headline string `json:"headline"`
	// Location represents where the user is located.
	Location Location `json:"location"`
	// Industry represents what industry the user is working in.
	Industry string `json:"industry"`
	// CurrentShare represents the user's last shared post.
	CurrentShare string `json:"current-share"`
	// NumConnections represents the user's number of connections, up to a maximum of 500.
	// The user's connections may be over this, however it will not be shown. (i.e. 500+ connections)
	NumConnections int `json:"num-connections"`
	// IsConnectionsCapped represents whether or not the user's connections are capped.
	IsConnectionsCapped bool `json:"num-connections-capped"`
	// Summary represents a long-form text description of the user's capabilities.
	Summary string `json:"summary"`
	// Specialties is a short-form text description of the user's specialties.
	Specialties string `json:"specialties"`
	// Positions is a Positions struct that describes the user's previously held positions.
	Positions Positions `json:"positions"`
	// PictureURL represents a URL pointing to the user's profile picture.
	PictureURL string `json:"picture-url"`
	// EmailAddress represents the user's e-mail address, however you must specify 'r_emailaddress'
	// to be able to retrieve this.
	EmailAddress string `json:"email-address"`
}

type Location struct {
	UserLocation string
	CountryCode  string
}

type Positions struct {
	total  int
	values []Position
}

// Position represents a job held by the authorized user.
type Position struct {
	// ID represents a unique ID representing the position
	ID string
	// Title represents a user's position's title, for example Jeff Bezos's title would be 'CEO'
	Title string
	// Summary represents a short description of the user's position.
	Summary string
	// StartDate represents when the user's position started.
	StartDate string
	// EndDate represents the user's position's end date, if any.
	EndDate string
	// IsCurrent represents if the position is currently held or not.
	// If this is false, EndDate will not be returned, and will therefore equal ""
	IsCurrent bool
	// Company represents the Company where the user is employed.
	Company PositionCompany
}

// PositionCompany represents a company that is described within a user's Profile.
// This is different from Company, which fully represents a company's data.
type PositionCompany struct {
	// ID represents a unique ID representing the company
	ID string
	// Name represents the name of the company
	Name string
	// Type represents the type of the company, either 'public' or 'private'
	Type string
	// Industry represents which industry the company is in.
	Industry string
	// Ticker represents the stock market ticker symbol of the company.
	// This will be blank if the company is privately held.
	Ticker string
}

type Profile struct {
	ErrorResponse
	ID                 string `json:"id"`
	LocalizedFirstName string `json:"localizedFirstName"`
	LocalizedLastName  string `json:"localizedLastName"`
	FirstName          struct {
		Localized struct {
			EnUS string `json:"en_US"`
		} `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"firstName"`
	LastName struct {
		Localized struct {
			EnUS string `json:"en_US"`
		} `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"lastName"`
	ProfilePicture struct {
		DisplayImage     string `json:"displayImage"`
		DisplayImageFull struct {
			Paging struct {
				Count int   `json:"count"`
				Start int   `json:"start"`
				Links []any `json:"links"`
			} `json:"paging"`
			Elements []struct {
				Artifact            string `json:"artifact"`
				AuthorizationMethod string `json:"authorizationMethod"`
				Data                struct {
					ComLinkedinDigitalmediaMediaartifactStillImage struct {
						MediaType    string `json:"mediaType"`
						RawCodecSpec struct {
							Name string `json:"name"`
							Type string `json:"type"`
						} `json:"rawCodecSpec"`
						DisplaySize struct {
							Width  float64 `json:"width"`
							Uom    string  `json:"uom"`
							Height float64 `json:"height"`
						} `json:"displaySize"`
						StorageSize struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"storageSize"`
						StorageAspectRatio struct {
							WidthAspect  float64 `json:"widthAspect"`
							HeightAspect float64 `json:"heightAspect"`
							Formatted    string  `json:"formatted"`
						} `json:"storageAspectRatio"`
						DisplayAspectRatio struct {
							WidthAspect  float64 `json:"widthAspect"`
							HeightAspect float64 `json:"heightAspect"`
							Formatted    string  `json:"formatted"`
						} `json:"displayAspectRatio"`
					} `json:"com.linkedin.digitalmedia.mediaartifact.StillImage"`
				} `json:"data"`
				Identifiers []struct {
					Identifier                 string `json:"identifier"`
					Index                      int    `json:"index"`
					MediaType                  string `json:"mediaType"`
					File                       string `json:"file"`
					IdentifierType             string `json:"identifierType"`
					IdentifierExpiresInSeconds int    `json:"identifierExpiresInSeconds"`
				} `json:"identifiers"`
			} `json:"elements"`
		} `json:"displayImage~"`
	} `json:"profilePicture"`
	Headline struct {
		Localized struct {
			EnUS string `json:"en_US"`
		} `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"headline"`
	LocalizedHeadline string `json:"localizedHeadline"`
	VanityName        string `json:"vanityName"`
}

type PrimaryContact struct {
	ErrorResponse
	Elements []struct {
		Handle      string `json:"handle"`
		HandleTilde struct {
			EmailAddress string `json:"emailAddress"`
			PhoneNumber  struct {
				Number string `json:"number"`
			} `json:"phoneNumber"`
		} `json:"handle~"`
		Primary bool   `json:"primary"`
		Type    string `json:"type"`
	} `json:"elements"`
}

type Client struct {
	http.Client
}

// Common struct for error response in linkedin API
type ErrorResponse struct {
	ServiceErrorCode int    `json:"serviceErrorCode"`
	Message          string `json:"message"`
	Status           int    `json:"status"`
}
