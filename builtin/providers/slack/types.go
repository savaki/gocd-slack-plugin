package slack

type Attachment struct {
	Pretext string `json:"pretext"`
	Text    string `json:"text"`
}

type Properties map[string]interface{}

type User struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	Prefs          Properties `json:"prefs,omitempty"`
	Created        int        `json:"created"`
	ManualPresence string     `json:"manual_presence"`
}

type Team struct {
	Id                    string     `json:"id"`
	Name                  string     `json:"name"`
	EmailDomain           string     `json:"email_domain"`
	Domain                string     `json:"domain"`
	MessageEditWindowMins int        `json:"msg_edit_window_mins"`
	OverStorageLimit      bool       `json:"over_storage_limit"`
	Prefs                 Properties `json:"prefs,omitempty"`
	Icon                  Properties `json:"icon"`
}

type Edit struct {
	User string `json:"user"`
	Ts   string `json:"ts"`
}

type Message struct {
	Type      string `json:"type"`
	Inviter   string `json:"inviter"`
	Subtype   string `json:"subtype"`
	Hidden    bool   `json:"hidden"`
	Channel   string `json:"channel"`
	User      string `json:"user"`
	Text      string `json:"text"`
	Ts        string `json:"ts"`
	DeletedTs string `json:"deleted_ts,omitempty"`
	EventTs   string `json:"event_ts,omitempty"`
	Edited    *Edit  `json:"edited,omitempty"`
}

type Value struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

type Group struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	IsGroup            bool     `json:"is_group"`
	Created            int      `json:"created"`
	Creator            string   `json:"creator"`
	IsArchived         bool     `json:"is_archived"`
	IsOpen             bool     `json:"is_open"`
	LastRead           string   `json:"last_read"`
	Latest             *Message `json:"latest"`
	UnreadCount        int      `json:"unread_count"`
	UnreadCountDisplay int      `json:"unread_count_display"`
	Members            []string `json:"members"`
	Topic              Value    `json:"topic"`
	Purpose            Value    `json:"purpose"`
}

type Channel struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	IsChannel          bool     `json:"is_channel"`
	Created            int      `json:"created"`
	Creator            string   `json:"creator"`
	IsArchived         bool     `json:"is_archived"`
	IsGeneral          bool     `json:"is_general"`
	IsMember           bool     `json:"is_member"`
	LastRead           string   `json:"last_read"`
	Latest             *Message `json:"latest"`
	UnreadCount        int      `json:"unread_count"`
	UnreadCountDisplay int      `json:"unread_count_display"`
	Members            []string `json:"members"`
	Topic              Value    `json:"topic"`
	Purpose            Value    `json:"purpose"`
}

type Bot struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
	Deleted bool       `json:"deleted"`
	Icons   Properties `json:"icons"`
}
