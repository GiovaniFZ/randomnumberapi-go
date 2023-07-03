package redditrandom

type redditcomment struct {
	Kind string `json:"kind"`
	Data struct {
		After     string `json:"after"`
		Dist      int    `json:"dist"`
		Modhash   string `json:"modhash"`
		GeoFilter string `json:"geo_filter"`
		Children  []struct {
			Kind string `json:"kind"`
			Data struct {
				SubredditId         string        `json:"subreddit_id"`
				ApprovedAtUtc       interface{}   `json:"approved_at_utc"`
				AuthorIsBlocked     bool          `json:"author_is_blocked"`
				CommentType         interface{}   `json:"comment_type"`
				LinkTitle           string        `json:"link_title"`
				ModReasonBy         interface{}   `json:"mod_reason_by"`
				BannedBy            interface{}   `json:"banned_by"`
				Ups                 int           `json:"ups"`
				NumReports          interface{}   `json:"num_reports"`
				AuthorFlairType     string        `json:"author_flair_type"`
				TotalAwardsReceived int           `json:"total_awards_received"`
				Subreddit           string        `json:"subreddit"`
				LinkAuthor          string        `json:"link_author"`
				Likes               interface{}   `json:"likes"`
				Replies             string        `json:"replies"`
				UserReports         []interface{} `json:"user_reports"`
				Saved               bool          `json:"saved"`
				Id                  string        `json:"id"`
				BannedAtUtc         interface{}   `json:"banned_at_utc"`
				ModReasonTitle      interface{}   `json:"mod_reason_title"`
				Gilded              int           `json:"gilded"`
				Archived            bool          `json:"archived"`
				CollapsedReasonCode interface{}   `json:"collapsed_reason_code"`
				NoFollow            bool          `json:"no_follow"`
				Author              string        `json:"author"`
				NumComments         int           `json:"num_comments"`
				CanModPost          bool          `json:"can_mod_post"`
				SendReplies         bool          `json:"send_replies"`
				ParentId            string        `json:"parent_id"`
				Score               int           `json:"score"`
				AuthorFullname      string        `json:"author_fullname"`
				Over18              bool          `json:"over_18"`
				ReportReasons       interface{}   `json:"report_reasons"`
				RemovalReason       interface{}   `json:"removal_reason"`
				ApprovedBy          interface{}   `json:"approved_by"`
				Controversiality    int           `json:"controversiality"`
				Body                string        `json:"body"`
				Edited              bool          `json:"edited"`
				TopAwardedType      interface{}   `json:"top_awarded_type"`
				Downs               int           `json:"downs"`
				AuthorFlairCssClass interface{}   `json:"author_flair_css_class"`
				IsSubmitter         bool          `json:"is_submitter"`
				Collapsed           bool          `json:"collapsed"`
				AuthorFlairRichtext []struct {
					E string `json:"e"`
					T string `json:"t,omitempty"`
					A string `json:"a,omitempty"`
					U string `json:"u,omitempty"`
				} `json:"author_flair_richtext"`
				AuthorPatreonFlair bool   `json:"author_patreon_flair"`
				BodyHtml           string `json:"body_html"`
				Gildings           struct {
				} `json:"gildings"`
				CollapsedReason              interface{}   `json:"collapsed_reason"`
				Distinguished                *string       `json:"distinguished"`
				AssociatedAward              interface{}   `json:"associated_award"`
				Stickied                     bool          `json:"stickied"`
				AuthorPremium                bool          `json:"author_premium"`
				CanGild                      bool          `json:"can_gild"`
				LinkId                       string        `json:"link_id"`
				UnrepliableReason            interface{}   `json:"unrepliable_reason"`
				AuthorFlairTextColor         *string       `json:"author_flair_text_color"`
				ScoreHidden                  bool          `json:"score_hidden"`
				Permalink                    string        `json:"permalink"`
				SubredditType                string        `json:"subreddit_type"`
				LinkPermalink                string        `json:"link_permalink"`
				Name                         string        `json:"name"`
				AuthorFlairTemplateId        *string       `json:"author_flair_template_id"`
				SubredditNamePrefixed        string        `json:"subreddit_name_prefixed"`
				AuthorFlairText              *string       `json:"author_flair_text"`
				TreatmentTags                []interface{} `json:"treatment_tags"`
				Created                      float64       `json:"created"`
				CreatedUtc                   float64       `json:"created_utc"`
				Awarders                     []interface{} `json:"awarders"`
				AllAwardings                 []interface{} `json:"all_awardings"`
				Locked                       bool          `json:"locked"`
				AuthorFlairBackgroundColor   *string       `json:"author_flair_background_color"`
				CollapsedBecauseCrowdControl interface{}   `json:"collapsed_because_crowd_control"`
				ModReports                   []interface{} `json:"mod_reports"`
				Quarantine                   bool          `json:"quarantine"`
				ModNote                      interface{}   `json:"mod_note"`
				LinkUrl                      string        `json:"link_url"`
			} `json:"data"`
		} `json:"children"`
		Before interface{} `json:"before"`
	} `json:"data"`
}