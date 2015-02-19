package slack

type Event struct {
	Type string `json:"type"`
}

type HelloReceiver interface {
	OnHello(event Event) error
}

// https://api.slack.com/events/message
type MessageEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	Ts      string `json:"ts"`
}

type MessageReceiver interface {
	OnMessage(event MessageEvent) error
}

// https://api.slack.com/events/channel_marked
type ChannelMarkedEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

type ChannelMarkedReceiver interface {
	OnChannelMarked(event ChannelMarkedEvent) error
}

// https://api.slack.com/events/channel_created
type ChannelCreatedEvent struct {
	Type    string  `json:"type"`
	Channel Channel `json:"channel"`
}

// https://api.slack.com/events/channel_joined
type ChannelJoinedEvent struct {
	Type    string  `json:"type"`
	Channel Channel `json:"channel"`
}

// https://api.slack.com/events/channel_left
type ChannelLeftEvent struct {
	Type    string  `json:"type"`
	Channel Channel `json:"channel"`
}

// https://api.slack.com/events/channel_deleted
type ChannelDeletedEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

type ChannelRename struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Created int    `json:"created"`
}

// https://api.slack.com/events/channel_rename
type ChannelRenameEvent struct {
	Type    string        `json:"type"`
	Channel ChannelRename `json:"channel"`
}

// https://api.slack.com/events/channel_archive
type ChannelArchiveEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// https://api.slack.com/events/channel_unarchive
type ChannelUnarchiveEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// https://api.slack.com/events/channel_history_changed
type ChannelHistoryChanged struct {
	Type    string `json:"type"`
	Latest  string `json:"latest"`
	Ts      string `json:"ts"`
	EventTs string `json:"event_ts"`
}

// https://api.slack.com/events/im_created
type ImCreatedEvent struct {
	Type    string  `json:"type"`
	User    string  `json:"user"`
	Channel Channel `json:"channel"`
}

// https://api.slack.com/events/im_open
type ImOpenEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

// https://api.slack.com/events/im_close
type ImCloseEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

// https://api.slack.com/events/im_marked
type ImMarkedEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

// https://api.slack.com/events/im_history_changed
type ImHistoryChangedEvent struct {
	Type    string `json:"type"`
	Latest  string `json:"latest"`
	Ts      string `json:"ts"`
	EventTs string `json:"event_ts"`
}
