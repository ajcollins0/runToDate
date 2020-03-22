package stravaApi

import "time"

// for strava api to ensure the accsess token is valid
type loginData struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresAt    int    `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// unmarhsalling the whole thing even though we really only need a few things.
// doing this cause I could build on it later if needed
type activity struct {
	ResourceState              int         `json:"resource_state"`
	Athlete                    athlete     `json:"athlete"`
	Name                       string      `json:"name"`
	Distance                   float64     `json:"distance"`
	MovingTime                 int         `json:"moving_time"`
	ElapsedTime                int         `json:"elapsed_time"`
	TotalElevationGain         float64     `json:"total_elevation_gain"`
	Type                       string      `json:"type"`
	WorkoutType                interface{} `json:"workout_type"`
	ID                         int64       `json:"id"`
	ExternalID                 string      `json:"external_id"`
	UploadID                   int64       `json:"upload_id"`
	StartDate                  time.Time   `json:"start_date"`
	StartDateLocal             time.Time   `json:"start_date_local"`
	Timezone                   string      `json:"timezone"`
	UtcOffset                  float64     `json:"utc_offset"`
	StartLatlng                []float64   `json:"start_latlng"`
	EndLatlng                  []float64   `json:"end_latlng"`
	LocationCity               interface{} `json:"location_city"`
	LocationState              interface{} `json:"location_state"`
	LocationCountry            string      `json:"location_country"`
	StartLatitude              float64     `json:"start_latitude"`
	StartLongitude             float64     `json:"start_longitude"`
	AchievementCount           int         `json:"achievement_count"`
	KudosCount                 int         `json:"kudos_count"`
	CommentCount               int         `json:"comment_count"`
	AthleteCount               int         `json:"athlete_count"`
	PhotoCount                 int         `json:"photo_count"`
	Map                        map2        `json:"map"`
	Trainer                    bool        `json:"trainer"`
	Commute                    bool        `json:"commute"`
	Manual                     bool        `json:"manual"`
	Private                    bool        `json:"private"`
	Visibility                 string      `json:"visibility"`
	Flagged                    bool        `json:"flagged"`
	GearID                     interface{} `json:"gear_id"`
	FromAcceptedTag            bool        `json:"from_accepted_tag"`
	UploadIDStr                string      `json:"upload_id_str"`
	AverageSpeed               float64     `json:"average_speed"`
	MaxSpeed                   float64     `json:"max_speed"`
	AverageCadence             float64     `json:"average_cadence"`
	AverageTemp                float64     `json:"average_temp"`
	HasHeartrate               bool        `json:"has_heartrate"`
	AverageHeartrate           float64     `json:"average_heartrate"`
	MaxHeartrate               float64     `json:"max_heartrate"`
	HeartrateOptOut            bool        `json:"heartrate_opt_out"`
	DisplayHideHeartrateOption bool        `json:"display_hide_heartrate_option"`
	ElevHigh                   float64     `json:"elev_high"`
	ElevLow                    float64     `json:"elev_low"`
	PrCount                    float64     `json:"pr_count"`
	TotalPhotoCount            float64     `json:"total_photo_count"`
	HasKudoed                  bool        `json:"has_kudoed"`
}

type athlete struct {
	ID            int `json:"id"`
	ResourceState int `json:"resource_state"`
}

type map2 struct {
	ID              string `json:"id"`
	SummaryPolyline string `json:"summary_polyline"`
	ResourceState   int    `json:"resource_state"`
}
