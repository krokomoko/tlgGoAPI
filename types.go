package tlgGoAPI

// Messages type

const (
	M_COMMAND = iota + 1
	M_TEXT
	M_IMAGE
	M_VIDEO
	M_VOICE
	M_AUDIO
	M_GIF
	M_DOCUMENT
	M_STICKER
	M_URL
)

type TelegramUpdate struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result,omitempty"`
}

type Update struct {
	UpdateId           int64              `json:"update_id,omitempty"`
	Message            Message            `json:"message,omitempty"`
	EditedMessage      Message            `json:"edited_message,omitempty"`
	ChannelPost        Message            `json:"channel_post,omitempty"`
	EditedChannelPost  Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               Poll               `json:"poll,omitempty"`
	PollAnswer         PollAnswer         `json:"poll_answer,omitempty"`
	MyChatMember       ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

type WebhookInfo struct {
	Url                          string   `json:"url,omitempty"`
	HasCustomCertificate         bool     `json:"has_custom_certificate,omitempty"`
	PendingUpdateCount           int64    `json:"pending_update_count,omitempty"`
	IpAddress                    string   `json:"ip_address,omitempty"`
	LastErrorDate                int64    `json:"last_error_date,omitempty"`
	LastErrorMessage             string   `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate int64    `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               int64    `json:"max_connections,omitempty"`
	AllowedUpdates               []string `json:"allowed_updates,omitempty"`
}

type User struct {
	Id                      int64  `json:"id,omitempty"`
	IsBot                   bool   `json:"is_bot,omitempty"`
	FirstName               string `json:"first_name,omitempty"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

type Chat struct {
	Id                                 int64     `json:"id,omitempty"`
	Type                               string    `json:"type,omitempty"`
	Title                              string    `json:"title,omitempty"`
	Username                           string    `json:"username,omitempty"`
	FirstName                          string    `json:"first_name,omitempty"`
	LastName                           string    `json:"last_name,omitempty"`
	IsForum                            bool      `json:"is_forum,omitempty"`
	Photo                              ChatPhoto `json:"photo,omitempty"`
	ActiveUsernames                    []string  `json:"active_usernames,omitempty"`
	EmojiStatusCustomEmojiId           string    `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int64     `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string    `json:"bio,omitempty"`
	HasPrivateForwards                 bool      `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool      `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool      `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool      `json:"join_by_request,omitempty"`
	Description                        string    `json:"description,omitempty"`
	InviteLink                         string    `json:"invite_link,omitempty"`

	Permissions                  ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay                int64           `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime        int64           `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled bool            `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers             bool            `json:"has_hidden_members,omitempty"`
	HasProtectedContent          bool            `json:"has_protected_content,omitempty"`
	StickerSetName               string          `json:"sticker_set_name,omitempty"`
	CanSetStickerSet             bool            `json:"can_set_sticker_set,omitempty"`
	LinkedChatId                 int64           `json:"linked_chat_id,omitempty"`
	Location                     ChatLocation    `json:"location,omitempty"`
	//PinnedMessage                      Message         `json:"pinned_message,omitempty"`
}

type Message struct {
	MessageId                     int64                         `json:"message_id,omitempty"`
	MessageThreadId               int64                         `json:"message_thread_id,omitempty"`
	From                          User                          `json:"from,omitempty"`
	SenderChat                    Chat                          `json:"sender_chat,omitempty"`
	Date                          int64                         `json:"date,omitempty"`
	Chat                          Chat                          `json:"chat,omitempty"`
	ForwardFrom                   User                          `json:"forward_from,omitempty"`
	ForwardFromChat               Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId          int64                         `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                        `json:"forward_signature,omitempty"`
	ForwardSenderName             string                        `json:"forward_sender_name,omitempty"`
	ForwardDate                   int64                         `json:"forward_date,omitempty"`
	IsTopicMessage                bool                          `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                          `json:"is_automatic_forward,omitempty"`
	ViaBot                        User                          `json:"via_bot,omitempty"`
	EditDate                      int64                         `json:"edit_date,omitempty"`
	HasProtectedContent           bool                          `json:"has_protected_content,omitempty"`
	MediaGroupId                  string                        `json:"media_group_id,omitempty"`
	AuthorSignature               string                        `json:"author_signature,omitempty"`
	Text                          string                        `json:"text,omitempty"`
	Entities                      []MessageEntity               `json:"entities,omitempty"`
	Animation                     Animation                     `json:"animation,omitempty"`
	Audio                         Audio                         `json:"audio,omitempty"`
	Document                      Document                      `json:"document,omitempty"`
	Photo                         []PhotoSize                   `json:"photo,omitempty"`
	Sticker                       Sticker                       `json:"sticker,omitempty"`
	Story                         Story                         `json:"story,omitempty"`
	Video                         Video                         `json:"video,omitempty"`
	VideoNote                     VideoNote                     `json:"video_note,omitempty"`
	Voice                         Voice                         `json:"voice,omitempty"`
	Caption                       string                        `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity               `json:"caption_entities,omitempty"`
	HasMediaSpoiler               bool                          `json:"has_media_spoiler,omitempty"`
	Contact                       Contact                       `json:"contact,omitempty"`
	Dice                          Dice                          `json:"dice,omitempty"`
	Game                          Game                          `json:"game,omitempty"`
	Poll                          Poll                          `json:"poll,omitempty"`
	Venue                         Venue                         `json:"venue,omitempty"`
	Location                      Location                      `json:"location,omitempty"`
	NewChatMembers                []User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                          `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                          `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                          `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int64                         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int64                         `json:"migrate_from_chat_id,omitempty"`

	Invoice                      Invoice                      `json:"invoice,omitempty"`
	SuccessfulPayment            SuccessfulPayment            `json:"successful_payment,omitempty"`
	UserShared                   UserShared                   `json:"user_shared,omitempty"`
	ChatShared                   ChatShared                   `json:"chat_shared,omitempty"`
	ConnectedWebsite             string                       `json:"connected_website,omitempty"`
	WriteAccessAllowed           WriteAccessAllowed           `json:"write_access_allowed,omitempty"`
	PassportData                 PassportData                 `json:"passport_data,omitempty"`
	ProximityAlertTriggered      ProximityAlertTriggered      `json:"proximity_alert_triggered,omitempty"`
	ForumTopicCreated            ForumTopicCreated            `json:"forum_topic_created,omitempty"`
	ForumTopicEdited             ForumTopicEdited             `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed             ForumTopicClosed             `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened           ForumTopicReopened           `json:"forum_topic_reopened,omitempty"`
	VideoChatScheduled           VideoChatScheduled           `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted             VideoChatStarted             `json:"video_chat_started,omitempty"`
	VideoChatEnded               VideoChatEnded               `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData                   WebAppData                   `json:"web_app_data,omitempty"`
	ReplyMarkup                  InlineKeyboardMarkup         `json:"reply_markup,omitempty"`
	//ReplyToMessage                Message                       `json:"reply_to_message,omitempty"`
	//PinnedMessage                 Message                       `json:"pinned_message,omitempty"`
	//GeneralForumTopicHidden      GeneralForumTopicHidden      `json:"general_forum_topic_hidden,omitempty"`
	//GeneralForumTopicUnhidden    GeneralForumTopicUnhidden    `json:"general_forum_topic_unhidden,omitempty"`
}

type MessageId struct {
	MessageId int64 `json:"message_id,omitempty"`
}

type MessageEntity struct {
	Type          string `json:"type,omitempty"`
	Offset        int64  `json:"offset,omitempty"`
	Length        int64  `json:"length,omitempty"`
	Url           string `json:"url,omitempty"`
	User          User   `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

type PhotoSize struct {
	FileId       string `json:"file_id,omitempty"`
	FileUniqueId string `json:"file_unique_id,omitempty"`
	Width        int64  `json:"width,omitempty"`
	Height       int64  `json:"height,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

type Animation struct {
	FileId       string    `json:"file_id,omitempty"`
	FileUniqueId string    `json:"file_unique_id,omitempty"`
	Width        int64     `json:"width,omitempty"`
	Height       int64     `json:"height,omitempty"`
	Duration     int64     `json:"duration,omitempty"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
}

type Audio struct {
	FileId       string    `json:"file_id,omitempty"`
	FileUniqueId string    `json:"file_unique_id,omitempty"`
	Duration     int64     `json:"duration,omitempty"`
	Performer    string    `json:"performer,omitempty"`
	Title        string    `json:"title,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
}

type Document struct {
	FileId       string    `json:"file_id,omitempty"`
	FileUniqueId string    `json:"file_unique_id,omitempty"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
}

type Story string

type Video struct {
	FileId       string    `json:"file_id,omitempty"`
	FileUniqueId string    `json:"file_unique_id,omitempty"`
	Width        int64     `json:"width,omitempty"`
	Height       int64     `json:"height,omitempty"`
	Duration     int64     `json:"duration,omitempty"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileId       string    `json:"file_id,omitempty"`
	FileUniqueId string    `json:"file_unique_id,omitempty"`
	Length       int64     `json:"length,omitempty"`
	Duration     int64     `json:"duration,omitempty"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
}

type Voice struct {
	FileId       string `json:"file_id,omitempty"`
	FileUniqueId string `json:"file_unique_id,omitempty"`
	Duration     int64  `json:"duration,omitempty"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji,omitempty"`
	Value int64  `json:"value,omitempty"`
}

type PollOption struct {
	Text       string `json:"text,omitempty"`
	VoterCount int64  `json:"voter_count,omitempty"`
}

type PollAnswer struct {
	PollId    string  `json:"poll_id,omitempty"`
	VoterChat Chat    `json:"voter_chat,omitempty"`
	User      User    `json:"user,omitempty"`
	OptionIds []int64 `json:"option_ids,omitempty"`
}

type Poll struct {
	Id                    string          `json:"id,omitempty"`
	Question              string          `json:"question,omitempty"`
	Options               []PollOption    `json:"options,omitempty"`
	TotalVoterCount       int64           `json:"total_voter_count,omitempty"`
	IsClosed              bool            `json:"is_closed,omitempty"`
	IsAnonymous           bool            `json:"is_anonymous,omitempty"`
	Type                  string          `json:"type,omitempty"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId       int64           `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int64           `json:"open_period,omitempty"`
	CloseDate             int64           `json:"close_date,omitempty"`
}

type Location struct {
	Longitude            float32 `json:"longitude,omitempty"`
	Latitude             float32 `json:"latitude,omitempty"`
	HorizontalAccuracy   float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int64   `json:"live_period,omitempty"`
	Heading              int64   `json:"heading,omitempty"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"`
}

type Venue struct {
	Location        Location `json:"location,omitempty"`
	Title           string   `json:"title,omitempty"`
	Address         string   `json:"address,omitempty"`
	FoursquareId    string   `json:"foursquare_id,omitempty"`
	FoursquareType  string   `json:"foursquare_type,omitempty"`
	GooglePlaceId   string   `json:"google_place_id,omitempty"`
	GooglePlaceType string   `json:"google_place_type,omitempty"`
}

type WebAppData struct {
	Data       string `json:"data,omitempty"`
	ButtonText string `json:"button_text,omitempty"`
}

type ProximityAlertTriggered struct {
	Traveler User  `json:"traveler,omitempty"`
	Watcher  User  `json:"watcher,omitempty"`
	Distance int64 `json:"distance,omitempty"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time,omitempty"`
}

type ForumTopicCreated struct {
	Name              string `json:"name,omitempty"`
	IconColor         int64  `json:"icon_color,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicClosed string

type ForumTopicEdited struct {
	Name              string
	IconCustomEmojiId string
}

type ForumTopicReopened string

type GeneralForumTopicHidden string

type GeneralForumTopicUnhidden string

type UserShared struct {
	RequestId int64 `json:"request_id,omitempty"`
	UserId    int64 `json:"user_id,omitempty"`
}

type ChatShared struct {
	RequestId int64 `json:"request_id,omitempty"`
	ChatId    int64 `json:"chat_id,omitempty"`
}

type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

type VideoChatScheduled struct {
	StartDate int64 `json:"start_date,omitempty"`
}

type VideoChatStarted string

type VideoChatEnded struct {
	Duration int64 `json:"duration,omitempty"`
}

type VideoChatParticipantsInvited struct {
	Users []User `json:"users,omitempty"`
}

type UserProfilePhotos struct {
	TotalCount int64         `json:"total_count,omitempty"`
	Photos     [][]PhotoSize `json:"photos,omitempty"`
}

type File struct {
	FileId       string `json:"file_id,omitempty"`
	FileUniqueId string `json:"file_unique_id,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type WebAppInfo struct {
	Url string `json:"url,omitempty"`
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard,omitempty"`
	IsPersistent          bool               `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
	Selective             bool               `json:"selective,omitempty"`
}

type KeyboardButton struct {
	Text            string                    `json:"text,omitempty"`
	RequestUser     KeyboardButtonRequestUser `json:"request_user,omitempty"`
	RequestChat     KeyboardButtonRequestChat `json:"request_chat,omitempty"`
	RequestContact  bool                      `json:"request_contact,omitempty"`
	RequestLocation bool                      `json:"request_location,omitempty"`
	RequestPoll     KeyboardButtonPollType    `json:"request_poll,omitempty"`
	WebApp          WebAppInfo                `json:"web_app,omitempty"`
}

type KeyboardButtonRequestUser struct {
	RequestId     int64 `json:"request_id,omitempty"`
	UserIsBot     bool  `json:"user_is_bot,omitempty"`
	UserIsPremium bool  `json:"user_is_premium,omitempty"`
}

type KeyboardButtonRequestChat struct {
	RequestId               int64                   `json:"request_id,omitempty"`
	ChatIsChannel           bool                    `json:"chat_is_channel,omitempty"`
	ChatIsForum             bool                    `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool                    `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool                    `json:"chat_is_created,omitempty"`
	UserAdministratorRights ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             bool                    `json:"bot_is_member,omitempty"`
}

type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"`
	Selective      bool `json:"selective,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

type InlineKeyboardButton struct {
	Text                         string                      `json:"text,omitempty"`
	Url                          string                      `json:"url,omitempty"`
	CallbackData                 string                      `json:"callback_data,omitempty"`
	WebApp                       WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                      `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                        `json:"pay,omitempty"`
}

type LoginUrl struct {
	Url                string `json:"url,omitempty"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

type CallbackQuery struct {
	Id              string  `json:"id,omitempty"`
	From            User    `json:"from,omitempty"`
	Message         Message `json:"message,omitempty"`
	InlineMessageId string  `json:"inline_message_id,omitempty"`
	ChatInstance    string  `json:"chat_instance,omitempty"`
	Data            string  `json:"data,omitempty"`
	GameShortName   string  `json:"game_short_name,omitempty"`
}

type ForceReply struct {
	ForceReply            bool   `json:"force_reply,omitempty"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id,omitempty"`
	SmallFileUniqueId string `json:"small_file_unique_id,omitempty"`
	BigFileId         string `json:"big_file_id,omitempty"`
	BigFileUniqueId   string `json:"big_file_unique_id,omitempty"`
}

type ChatInviteLink struct {
	InviteLink              string `json:"invite_link,omitempty"`
	Creator                 User   `json:"creator,omitempty"`
	CreatesJoinRequest      bool   `json:"creates_join_request,omitempty"`
	IsPrimary               bool   `json:"is_primary,omitempty"`
	IsRevoked               bool   `json:"is_revoked,omitempty"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int64  `json:"expire_date,omitempty"`
	MemberLimit             int64  `json:"member_limit,omitempty"`
	PendingJoinRequestCount int64  `json:"pending_join_request_count,omitempty"`
}

type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous,omitempty"`
	CanManageChat       bool `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   bool `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool `json:"can_change_info,omitempty"`
	CanInviteUsers      bool `json:"can_invite_users,omitempty"`
	CanPostMessages     bool `json:"can_post_messages,omitempty"`
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`
	CanPostStories      bool `json:"can_post_stories,omitempty"`
	CanEditStories      bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"`
	CanManageTopics     bool `json:"can_manage_topics,omitempty"`
}

type ChatMember interface{}

type ChatMemberOwner struct {
	Status      string `json:"status,omitempty"`
	User        User   `json:"user,omitempty"`
	IsAnonymous bool   `json:"is_anonymous,omitempty"`
	CustomTitle string `json:"custom_title,omitempty"`
}

type ChatMemberAdministrator struct {
	Status              string `json:"status,omitempty"`
	User                User   `json:"user,omitempty"`
	CanBeEdited         bool   `json:"can_be_edited,omitempty"`
	IsAnonymous         bool   `json:"is_anonymous,omitempty"`
	CanManageChat       bool   `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   bool   `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool   `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool   `json:"can_change_info,omitempty"`
	CanInviteUsers      bool   `json:"can_invite_users,omitempty"`
	CanPostMessages     bool   `json:"can_post_messages,omitempty"`
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`
	CanPostStories      bool   `json:"can_post_stories,omitempty"`
	CanEditStories      bool   `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool   `json:"can_delete_stories,omitempty"`
	CanManageTopics     bool   `json:"can_manage_topics,omitempty"`
	CustomTitle         string `json:"custom_title,omitempty"`
}

type ChatMemberMember struct {
	Status string `json:"status,omitempty"`
	User   User   `json:"user,omitempty"`
}

type ChatMemberRestricted struct {
	Status                string `json:"status,omitempty"`
	User                  User   `json:"user,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendAudios         bool   `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool   `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool   `json:"can_send_photos,omitempty"`
	CanSendVideos         bool   `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool   `json:"can_manage_topics,omitempty"`
	UntilDate             int64  `json:"until_date,omitempty"`
}

type ChatMemberLeft struct {
	Status string `json:"status,omitempty"`
	User   User   `json:"user,omitempty"`
}

type ChatMemberBanned struct {
	Status    string `json:"status,omitempty"`
	User      User   `json:"user,omitempty"`
	UntilDate int64  `json:"until_date,omitempty"`
}

type ChatMemberUpdated struct {
	Chat                    Chat           `json:"chat,omitempty"`
	From                    User           `json:"from,omitempty"`
	Date                    int64          `json:"date,omitempty"`
	OldChatMember           ChatMember     `json:"old_chat_member,omitempty"`
	NewChatMember           ChatMember     `json:"new_chat_member,omitempty"`
	InviteLink              ChatInviteLink `json:"invite_link,omitempty"`
	ViaChatFolderInviteLink bool           `json:"via_chat_folder_invite_link,omitempty"`
}

type ChatJoinRequest struct {
	Chat       Chat           `json:"chat,omitempty"`
	From       User           `json:"from,omitempty"`
	UserChatId int64          `json:"user_chat_id,omitempty"`
	Date       int64          `json:"date,omitempty"`
	Bio        string         `json:"bio,omitempty"`
	InviteLink ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

type ChatLocation struct {
	Location Location `json:"location,omitempty"`
	Address  string   `json:"address,omitempty"`
}

type ForumTopic struct {
	MessageThreadId   int64  `json:"message_thread_id,omitempty"`
	Name              string `json:"name,omitempty"`
	IconColor         int64  `json:"icon_color,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command,omitempty"`
	Description string `json:"description,omitempty"`
}

type BotCommandScope interface{}

type BotCommandScopeDefault struct {
	Type string `json:"type,omitempty"`
}

type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type,omitempty"`
}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type,omitempty"`
}

type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type,omitempty"`
}

type BotCommandScopeChat struct {
	Type   string `json:"type,omitempty"`
	ChatId int64 `json:"chat_id,omitempty"`
}

type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type,omitempty"`
	ChatId int64 `json:"chat_id,omitempty"`
}

type BotCommandScopeChatMember struct {
	Type   string `json:"type,omitempty"`
	ChatId int64 `json:"chat_id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type BotName struct {
	Name string `json:"name,omitempty"`
}

type BotDescription struct {
	Description string `json:"description,omitempty"`
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description,omitempty"`
}

type MenuButton interface{}

type MenuButtonCommands struct {
	Type string `json:"type,omitempty"`
}

type MenuButtonWebApp struct {
	Type   string     `json:"type,omitempty"`
	Text   string     `json:"text,omitempty"`
	WebApp WebAppInfo `json:"web_app,omitempty"`
}

type MenuButtonDefault struct {
	Type string `json:"type,omitempty"`
}

type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int64 `json:"retry_after,omitempty"`
}

type InputFile string

type InputMedia interface{}

type InputMediaPhoto struct {
	Type            string          `json:"type,omitempty"`
	Media           string          `json:"media,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler      bool            `json:"has_spoiler,omitempty"`
}

type InputMediaVideo struct {
	Type              string          `json:"type,omitempty"`
	Media             string          `json:"media,omitempty"`
	Thumbnail         InputFile       `json:"thumbnail,omitempty"`
	Caption           string          `json:"caption,omitempty"`
	ParseMode         string          `json:"parse_mode,omitempty"`
	CaptionEntities   []MessageEntity `json:"caption_entities,omitempty"`
	Width             int64           `json:"width,omitempty"`
	Height            int64           `json:"height,omitempty"`
	Duration          int64           `json:"duration,omitempty"`
	SupportsStreaming bool            `json:"supports_streaming,omitempty"`
	HasSpoiler        bool            `json:"has_spoiler,omitempty"`
}

type InputMediaAnimation struct {
	Type            string          `json:"type,omitempty"`
	Media           string          `json:"media,omitempty"`
	Thumbnail       InputFile       `json:"thumbnail,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Width           int64           `json:"width,omitempty"`
	Height          int64           `json:"height,omitempty"`
	Duration        int64           `json:"duration,omitempty"`
	HasSpoiler      bool            `json:"has_spoiler,omitempty"`
}

type InputMediaAudio struct {
	Type            string          `json:"type,omitempty"`
	Media           string          `json:"media,omitempty"`
	Thumbnail       InputFile       `json:"thumbnail,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        int64           `json:"duration,omitempty"`
	Performer       string          `json:"performer,omitempty"`
	Title           string          `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Type                        string          `json:"type,omitempty"`
	Media                       string          `json:"media,omitempty"`
	Thumbnail                   InputFile       `json:"thumbnail,omitempty"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   string          `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
}

type Sticker struct {
	FileId           string       `json:"file_id,omitempty"`
	FileUniqueId     string       `json:"file_unique_id,omitempty"`
	Type             string       `json:"type,omitempty"`
	Width            int64        `json:"width,omitempty"`
	Height           int64        `json:"height,omitempty"`
	IsAnimated       bool         `json:"is_animated,omitempty"`
	IsVideo          bool         `json:"is_video,omitempty"`
	Thumbnail        PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string       `json:"emoji,omitempty"`
	SetName          string       `json:"set_name,omitempty"`
	PremiumAnimation File         `json:"premium_animation,omitempty"`
	MaskPosition     MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiId    string       `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool         `json:"needs_repainting,omitempty"`
	FileSize         int64        `json:"file_size,omitempty"`
}

type StickerSet struct {
	Name        string    `json:"name,omitempty"`
	Title       string    `json:"title,omitempty"`
	StickerType string    `json:"sticker_type,omitempty"`
	IsAnimated  bool      `json:"is_animated,omitempty"`
	IsVideo     bool      `json:"is_video,omitempty"`
	Stickers    []Sticker `json:"stickers,omitempty"`
	Thumbnail   PhotoSize `json:"thumbnail,omitempty"`
}

type MaskPosition struct {
	Point  string  `json:"point,omitempty"`
	XShift float32 `json:"x_shift,omitempty"`
	YShift float32 `json:"y_shift,omitempty"`
	Scale  float32 `json:"scale,omitempty"`
}

type InputSticker struct {
	Sticker      InputFile    `json:"sticker,omitempty"`
	EmojiList    []string     `json:"emoji_list,omitempty"`
	MaskPosition MaskPosition `json:"mask_position,omitempty"`
	Keywords     []string     `json:"keywords,omitempty"`
}

type InlineQuery struct {
	Id       string   `json:"id,omitempty"`
	From     User     `json:"from,omitempty"`
	Query    string   `json:"query,omitempty"`
	Offset   string   `json:"offset,omitempty"`
	ChatType string   `json:"chat_type,omitempty"`
	Location Location `json:"location,omitempty"`
}

type InlineQueryResultsButton struct {
	Text           string     `json:"text,omitempty"`
	WebApp         WebAppInfo `json:"web_app,omitempty"`
	StartParameter string     `json:"start_parameter,omitempty"`
}

type InlineQueryResult interface{}

type InlineQueryResultArticle struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	Title               string               `json:"title,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Url                 string               `json:"url,omitempty"`
	HideUrl             bool                 `json:"hide_url,omitempty"`
	Description         string               `json:"description,omitempty"`
	ThumbnailUrl        string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int64                `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultPhoto struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	PhotoUrl            string               `json:"photo_url,omitempty"`
	ThumbnailUrl        string               `json:"thumbnail_url,omitempty"`
	PhotoWidth          int64                `json:"photo_width,omitempty"`
	PhotoHeight         int64                `json:"photo_height,omitempty"`
	Title               string               `json:"title,omitempty"`
	Description         string               `json:"description,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultGif struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	GifUrl              string               `json:"gif_url,omitempty"`
	GifWidth            int64                `json:"gif_width,omitempty"`
	GifHeight           int64                `json:"gif_height,omitempty"`
	GifDuration         int64                `json:"gif_duration,omitempty"`
	ThumbnailUrl        string               `json:"thumbnail_url,omitempty"`
	ThumbnailMimeType   string               `json:"thumbnail_mime_type,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	Mpeg4Url            string               `json:"mpeg4_url,omitempty"`
	Mpeg4Width          int64                `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int64                `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int64                `json:"mpeg4_duration,omitempty"`
	ThumbnailUrl        string               `json:"thumbnail_url,omitempty"`
	ThumbnailMimeType   string               `json:"thumbnail_mime_type,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVideo struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	VideoUrl            string               `json:"video_url,omitempty"`
	MimeType            string               `json:"mime_type,omitempty"`
	ThumbnailUrl        string               `json:"thumbnail_url,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	VideoWidth          int64                `json:"video_width,omitempty"`
	VideoHeight         int64                `json:"video_height,omitempty"`
	VideoDuration       int64                `json:"video_duration,omitempty"`
	Description         string               `json:"description,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultAudio struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	AudioUrl            string               `json:"audio_url,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	Performer           string               `json:"performer,omitempty"`
	AudioDuration       int64                `json:"audio_duration,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVoice struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	VoiceUrl            string               `json:"voice_url,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	VoiceDuration       int64                `json:"voice_duration,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultDocument struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	DocumentUrl         string               `json:"document_url,omitempty"`
	MimeType            string               `json:"mime_type,omitempty"`
	Description         string               `json:"description,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int64                `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultLocation struct {
	Type                 string               `json:"type,omitempty"`
	Id                   string               `json:"id,omitempty"`
	Latitude             float32              `json:"latitude,omitempty"`
	Longitude            float32              `json:"longitude,omitempty"`
	Title                string               `json:"title,omitempty"`
	HorizontalAccuracy   float32              `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int64                `json:"live_period,omitempty"`
	Heading              int64                `json:"heading,omitempty"`
	ProximityAlertRadius int64                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl         string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       int64                `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      int64                `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultVenue struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	Latitude            float32              `json:"latitude,omitempty"`
	Longitude           float32              `json:"longitude,omitempty"`
	Title               string               `json:"title,omitempty"`
	Address             string               `json:"address,omitempty"`
	FoursquareId        string               `json:"foursquare_id,omitempty"`
	FoursquareType      string               `json:"foursquare_type,omitempty"`
	GooglePlaceId       string               `json:"google_place_id,omitempty"`
	GooglePlaceType     string               `json:"google_place_type,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int64                `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultContact struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	PhoneNumber         string               `json:"phone_number,omitempty"`
	FirstName           string               `json:"first_name,omitempty"`
	LastName            string               `json:"last_name,omitempty"`
	Vcard               string               `json:"vcard,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int64                `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultGame struct {
	Type          string               `json:"type,omitempty"`
	Id            string               `json:"id,omitempty"`
	GameShortName string               `json:"game_short_name,omitempty"`
	ReplyMarkup   InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedPhoto struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	PhotoFileId         string               `json:"photo_file_id,omitempty"`
	Title               string               `json:"title,omitempty"`
	Description         string               `json:"description,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedGif struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	GifFileId           string               `json:"gif_file_id,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	Mpeg4FileId         string               `json:"mpeg4_file_id,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedSticker struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	StickerFileId       string               `json:"sticker_file_id,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedDocument struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	Title               string               `json:"title,omitempty"`
	DocumentFileId      string               `json:"document_file_id,omitempty"`
	Description         string               `json:"description,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVideo struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	VideoFileId         string               `json:"video_file_id,omitempty"`
	Title               string               `json:"title,omitempty"`
	Description         string               `json:"description,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVoice struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	VoiceFileId         string               `json:"voice_file_id,omitempty"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedAudio struct {
	Type                string               `json:"type,omitempty"`
	Id                  string               `json:"id,omitempty"`
	AudioFileId         string               `json:"audio_file_id,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	ParseMode           string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent  `json:"input_message_content,omitempty"`
}

type InputMessageContent interface{}

type InputTextMessageContent struct {
	MessageText           string          `json:"message_text,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float32 `json:"latitude,omitempty"`
	Longitude            float32 `json:"longitude,omitempty"`
	HorizontalAccuracy   float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int64   `json:"live_period,omitempty"`
	Heading              int64   `json:"heading,omitempty"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"`
}

type InputVenueMessageContent struct {
	Latitude        float32 `json:"latitude,omitempty"`
	Longitude       float32 `json:"longitude,omitempty"`
	Title           string  `json:"title,omitempty"`
	Address         string  `json:"address,omitempty"`
	FoursquareId    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceId   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title,omitempty"`
	Description               string         `json:"description,omitempty"`
	Payload                   string         `json:"payload,omitempty"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency,omitempty"`
	Prices                    []LabeledPrice `json:"prices,omitempty"`
	MaxTipAmount              int64          `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoUrl                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int64          `json:"photo_size,omitempty"`
	PhotoWidth                int64          `json:"photo_width,omitempty"`
	PhotoHeight               int64          `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

type ChosenInlineResult struct {
	ResultId        string   `json:"result_id,omitempty"`
	From            User     `json:"from,omitempty"`
	Location        Location `json:"location,omitempty"`
	InlineMessageId string   `json:"inline_message_id,omitempty"`
	Query           string   `json:"query,omitempty"`
}

type SentWebAppMessage struct {
	InlineMessageId string `json:"inline_message_id,omitempty"`
}

type LabeledPrice struct {
	Label  string `json:"label,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}

type Invoice struct {
	Title          string `json:"title,omitempty"`
	Description    string `json:"description,omitempty"`
	StartParameter string `json:"start_parameter,omitempty"`
	Currency       string `json:"currency,omitempty"`
	TotalAmount    int64  `json:"total_amount,omitempty"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code,omitempty"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	StreetLine1 string `json:"street_line1,omitempty"`
	StreetLine2 string `json:"street_line2,omitempty"`
	PostCode    string `json:"post_code,omitempty"`
}

type OrderInfo struct {
	Name            string          `json:"name,omitempty"`
	PhoneNumber     string          `json:"phone_number,omitempty"`
	Email           string          `json:"email,omitempty"`
	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingOption struct {
	Id     string         `json:"id,omitempty"`
	Title  string         `json:"title,omitempty"`
	Prices []LabeledPrice `json:"prices,omitempty"`
}

type SuccessfulPayment struct {
	Currency                string    `json:"currency,omitempty"`
	TotalAmount             int64     `json:"total_amount,omitempty"`
	InvoicePayload          string    `json:"invoice_payload,omitempty"`
	ShippingOptionId        string    `json:"shipping_option_id,omitempty"`
	OrderInfo               OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeId string    `json:"telegram_payment_charge_id,omitempty"`
	ProviderPaymentChargeId string    `json:"provider_payment_charge_id,omitempty"`
}

type ShippingQuery struct {
	Id              string          `json:"id,omitempty"`
	From            User            `json:"from,omitempty"`
	InvoicePayload  string          `json:"invoice_payload,omitempty"`
	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

type PreCheckoutQuery struct {
	Id               string    `json:"id,omitempty"`
	From             User      `json:"from,omitempty"`
	Currency         string    `json:"currency,omitempty"`
	TotalAmount      int64     `json:"total_amount,omitempty"`
	InvoicePayload   string    `json:"invoice_payload,omitempty"`
	ShippingOptionId string    `json:"shipping_option_id,omitempty"`
	OrderInfo        OrderInfo `json:"order_info,omitempty"`
}

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data,omitempty"`
	Credentials EncryptedCredentials       `json:"credentials,omitempty"`
}

type PassportFile struct {
	FileId       string `json:"file_id,omitempty"`
	FileUniqueId string `json:"file_unique_id,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
	FileDate     int64  `json:"file_date,omitempty"`
}

type EncryptedPassportElement struct {
	Type        string         `json:"type,omitempty"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   PassportFile   `json:"front_side,omitempty"`
	ReverseSide PassportFile   `json:"reverse_side,omitempty"`
	Selfie      PassportFile   `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash,omitempty"`
}

type EncryptedCredentials struct {
	Data   string `json:"data,omitempty"`
	Hash   string `json:"hash,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type PassportElementError interface{}

type PassportElementErrorDataField struct {
	Source    string `json:"source,omitempty"`
	Type      string `json:"type,omitempty"`
	FieldName string `json:"field_name,omitempty"`
	DataHash  string `json:"data_hash,omitempty"`
	Message   string `json:"message,omitempty"`
}

type PassportElementErrorFrontSide struct {
	Source   string `json:"source,omitempty"`
	Type     string `json:"type,omitempty"`
	FileHash string `json:"file_hash,omitempty"`
	Message  string `json:"message,omitempty"`
}

type PassportElementErrorReverseSide struct {
	Source   string `json:"source,omitempty"`
	Type     string `json:"type,omitempty"`
	FileHash string `json:"file_hash,omitempty"`
	Message  string `json:"message,omitempty"`
}

type PassportElementErrorSelfie struct {
	Source   string `json:"source,omitempty"`
	Type     string `json:"type,omitempty"`
	FileHash string `json:"file_hash,omitempty"`
	Message  string `json:"message,omitempty"`
}

type PassportElementErrorFile struct {
	Source   string `json:"source,omitempty"`
	Type     string `json:"type,omitempty"`
	FileHash string `json:"file_hash,omitempty"`
	Message  string `json:"message,omitempty"`
}

type PassportElementErrorFiles struct {
	Source     string   `json:"source,omitempty"`
	Type       string   `json:"type,omitempty"`
	FileHashes []string `json:"file_hashes,omitempty"`
	Message    string   `json:"message,omitempty"`
}

type PassportElementErrorTranslationFile struct {
	Source   string `json:"source,omitempty"`
	Type     string `json:"type,omitempty"`
	FileHash string `json:"file_hash,omitempty"`
	Message  string `json:"message,omitempty"`
}

type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source,omitempty"`
	Type       string   `json:"type,omitempty"`
	FileHashes []string `json:"file_hashes,omitempty"`
	Message    string   `json:"message,omitempty"`
}

type PassportElementErrorUnspecified struct {
	Source      string `json:"source,omitempty"`
	Type        string `json:"type,omitempty"`
	ElementHash string `json:"element_hash,omitempty"`
	Message     string `json:"message,omitempty"`
}

type Game struct {
	Title        string          `json:"title,omitempty"`
	Description  string          `json:"description,omitempty"`
	Photo        []PhotoSize     `json:"photo,omitempty"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    Animation       `json:"animation,omitempty"`
}

type CallbackGame string

type GameHighScore struct {
	Position int64 `json:"position,omitempty"`
	User     User  `json:"user,omitempty"`
	Score    int64 `json:"score,omitempty"`
}
