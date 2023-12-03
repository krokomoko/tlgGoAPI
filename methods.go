package tlgGoAPI

import "errors"

type SendMessage struct {
	ChatId                   int64                `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Text                     string               `json:"text,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity      `json:"entities,omitempty"`
	DisableWebPagePreview    bool                 `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type ForwardMessage struct {
	ChatId              string `json:"chat_id,omitempty"`
	MessageThreadId     int64  `json:"message_thread_id,omitempty"`
	FromChatId          string `json:"from_chat_id,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
	MessageId           int64  `json:"message_id,omitempty"`
}

type CopyMessage struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	FromChatId               string               `json:"from_chat_id,omitempty"`
	MessageId                int64                `json:"message_id,omitempty"`
	Caption                  string               `json:"caption,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity      `json:"caption_entities,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendPhoto struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Photo                    InputFile            `json:"photo,omitempty"`
	Caption                  string               `json:"caption,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity      `json:"caption_entities,omitempty"`
	HasSpoiler               bool                 `json:"has_spoiler,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendAudio struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Audio                    InputFile            `json:"audio,omitempty"`
	Caption                  string               `json:"caption,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity      `json:"caption_entities,omitempty"`
	Duration                 int64                `json:"duration,omitempty"`
	Performer                string               `json:"performer,omitempty"`
	Title                    string               `json:"title,omitempty"`
	Thumbnail                InputFile            `json:"thumbnail,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendDocument struct {
	ChatId                      string               `json:"chat_id,omitempty"`
	MessageThreadId             int64                `json:"message_thread_id,omitempty"`
	Document                    InputFile            `json:"document,omitempty"`
	Thumbnail                   InputFile            `json:"thumbnail,omitempty"`
	Caption                     string               `json:"caption,omitempty"`
	ParseMode                   string               `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity      `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                 `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                 `json:"disable_notification,omitempty"`
	ProtectContent              bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId            int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendVideo struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Video                    InputFile            `json:"video,omitempty"`
	Duration                 int64                `json:"duration,omitempty"`
	Width                    int64                `json:"width,omitempty"`
	Height                   int64                `json:"height,omitempty"`
	Thumbnail                InputFile            `json:"thumbnail,omitempty"`
	Caption                  string               `json:"caption,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity      `json:"caption_entities,omitempty"`
	HasSpoiler               bool                 `json:"has_spoiler,omitempty"`
	SupportsStreaming        bool                 `json:"supports_streaming,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendAnimation struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Animation                InputFile            `json:"animation,omitempty"`
	Duration                 int64                `json:"duration,omitempty"`
	Width                    int64                `json:"width,omitempty"`
	Height                   int64                `json:"height,omitempty"`
	Thumbnail                InputFile            `json:"thumbnail,omitempty"`
	Caption                  string               `json:"caption,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity      `json:"caption_entities,omitempty"`
	HasSpoiler               bool                 `json:"has_spoiler,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendVoice struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Voice                    InputFile            `json:"voice,omitempty"`
	Caption                  string               `json:"caption,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity      `json:"caption_entities,omitempty"`
	Duration                 int64                `json:"duration,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendVideoNote struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	VideoNote                InputFile            `json:"video_note,omitempty"`
	Duration                 int64                `json:"duration,omitempty"`
	Length                   int64                `json:"length,omitempty"`
	Thumbnail                InputFile            `json:"thumbnail,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendMediaGroup struct {
	ChatId                   string        `json:"chat_id,omitempty"`
	MessageThreadId          int64         `json:"message_thread_id,omitempty"`
	Media                    []interface{} `json:"media,omitempty"`
	DisableNotification      bool          `json:"disable_notification,omitempty"`
	ProtectContent           bool          `json:"protect_content,omitempty"`
	ReplyToMessageId         int64         `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool          `json:"allow_sending_without_reply,omitempty"`
}

type SendLocation struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Latitude                 float32              `json:"latitude,omitempty"`
	Longitude                float32              `json:"longitude,omitempty"`
	HorizontalAccuracy       float32              `json:"horizontal_accuracy,omitempty"`
	LivePeriod               int64                `json:"live_period,omitempty"`
	Heading                  int64                `json:"heading,omitempty"`
	ProximityAlertRadius     int64                `json:"proximity_alert_radius,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendVenue struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Latitude                 float32              `json:"latitude,omitempty"`
	Longitude                float32              `json:"longitude,omitempty"`
	Title                    string               `json:"title,omitempty"`
	Address                  string               `json:"address,omitempty"`
	FoursquareId             string               `json:"foursquare_id,omitempty"`
	FoursquareType           string               `json:"foursquare_type,omitempty"`
	GooglePlaceId            string               `json:"google_place_id,omitempty"`
	GooglePlaceType          string               `json:"google_place_type,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendContact struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	PhoneNumber              string               `json:"phone_number,omitempty"`
	FirstName                string               `json:"first_name,omitempty"`
	LastName                 string               `json:"last_name,omitempty"`
	Vcard                    string               `json:"vcard,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendPoll struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Question                 string               `json:"question,omitempty"`
	Options                  []string             `json:"options,omitempty"`
	IsAnonymous              bool                 `json:"is_anonymous,omitempty"`
	Type                     string               `json:"type,omitempty"`
	AllowsMultipleAnswers    bool                 `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId          int64                `json:"correct_option_id,omitempty"`
	Explanation              string               `json:"explanation,omitempty"`
	ExplanationParseMode     string               `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities      []MessageEntity      `json:"explanation_entities,omitempty"`
	OpenPeriod               int64                `json:"open_period,omitempty"`
	CloseDate                int64                `json:"close_date,omitempty"`
	IsClosed                 bool                 `json:"is_closed,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendDice struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Emoji                    string               `json:"emoji,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendChatAction struct {
	ChatId          string `json:"chat_id,omitempty"`
	MessageThreadId int64  `json:"message_thread_id,omitempty"`
	Action          string `json:"action,omitempty"`
}

type GetUserProfilePhotos struct {
	UserId int64 `json:"user_id,omitempty"`
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

type GetFile struct {
	FileId string `json:"file_id,omitempty"`
}

type BanChatMember struct {
	ChatId         string `json:"chat_id,omitempty"`
	UserId         int64  `json:"user_id,omitempty"`
	UntilDate      int64  `json:"until_date,omitempty"`
	RevokeMessages bool   `json:"revoke_messages,omitempty"`
}

type UnbanChatMember struct {
	ChatId       string `json:"chat_id,omitempty"`
	UserId       int64  `json:"user_id,omitempty"`
	OnlyIfBanned bool   `json:"only_if_banned,omitempty"`
}

type RestrictChatMember struct {
	ChatId                        string          `json:"chat_id,omitempty"`
	UserId                        int64           `json:"user_id,omitempty"`
	Permissions                   ChatPermissions `json:"permissions,omitempty"`
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int64           `json:"until_date,omitempty"`
}

type PromoteChatMember struct {
	ChatId              string `json:"chat_id,omitempty"`
	UserId              int64  `json:"user_id,omitempty"`
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
}

type SetChatAdministratorCustomTitle struct {
	ChatId      string `json:"chat_id,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	CustomTitle string `json:"custom_title,omitempty"`
}

type BanChatSenderChat struct {
	ChatId       string `json:"chat_id,omitempty"`
	SenderChatId int64  `json:"sender_chat_id,omitempty"`
}

type UnbanChatSenderChat struct {
	ChatId       string `json:"chat_id,omitempty"`
	SenderChatId int64  `json:"sender_chat_id,omitempty"`
}

type SetChatPermissions struct {
	ChatId                        string          `json:"chat_id,omitempty"`
	Permissions                   ChatPermissions `json:"permissions,omitempty"`
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"`
}

type ExportChatInviteLink struct {
	ChatId string `json:"chat_id,omitempty"`
}

type CreateChatInviteLink struct {
	ChatId             string `json:"chat_id,omitempty"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int64  `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type EditChatInviteLink struct {
	ChatId             string `json:"chat_id,omitempty"`
	InviteLink         string `json:"invite_link,omitempty"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int64  `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type RevokeChatInviteLink struct {
	ChatId     string `json:"chat_id,omitempty"`
	InviteLink string `json:"invite_link,omitempty"`
}

type ApproveChatJoinRequest struct {
	ChatId string `json:"chat_id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type DeclineChatJoinRequest struct {
	ChatId string `json:"chat_id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type SetChatPhoto struct {
	ChatId string    `json:"chat_id,omitempty"`
	Photo  InputFile `json:"photo,omitempty"`
}

type DeleteChatPhoto struct {
	ChatId string `json:"chat_id,omitempty"`
}

type SetChatTitle struct {
	ChatId string `json:"chat_id,omitempty"`
	Title  string `json:"title,omitempty"`
}

type SetChatDescription struct {
	ChatId      string `json:"chat_id,omitempty"`
	Description string `json:"description,omitempty"`
}

type PinChatMessage struct {
	ChatId              string `json:"chat_id,omitempty"`
	MessageId           int64  `json:"message_id,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
}

type UnpinChatMessage struct {
	ChatId    string `json:"chat_id,omitempty"`
	MessageId int64  `json:"message_id,omitempty"`
}

type UnpinAllChatMessages struct {
	ChatId string `json:"chat_id,omitempty"`
}

type LeaveChat struct {
	ChatId string `json:"chat_id,omitempty"`
}

type GetChat struct {
	ChatId string `json:"chat_id,omitempty"`
}

type GetChatAdministrators struct {
	ChatId string `json:"chat_id,omitempty"`
}

type GetChatMemberCount struct {
	ChatId string `json:"chat_id,omitempty"`
}

type GetChatMember struct {
	ChatId string `json:"chat_id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type SetChatStickerSet struct {
	ChatId         string `json:"chat_id,omitempty"`
	StickerSetName string `json:"sticker_set_name,omitempty"`
}

type DeleteChatStickerSet struct {
	ChatId string `json:"chat_id,omitempty"`
}

type CreateForumTopic struct {
	ChatId            string `json:"chat_id,omitempty"`
	Name              string `json:"name,omitempty"`
	IconColor         int64  `json:"icon_color,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type EditForumTopic struct {
	ChatId            string `json:"chat_id,omitempty"`
	MessageThreadId   int64  `json:"message_thread_id,omitempty"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type CloseForumTopic struct {
	ChatId          string `json:"chat_id,omitempty"`
	MessageThreadId int64  `json:"message_thread_id,omitempty"`
}

type ReopenForumTopic struct {
	ChatId          string `json:"chat_id,omitempty"`
	MessageThreadId int64  `json:"message_thread_id,omitempty"`
}

type DeleteForumTopic struct {
	ChatId          string `json:"chat_id,omitempty"`
	MessageThreadId int64  `json:"message_thread_id,omitempty"`
}

type UnpinAllForumTopicMessages struct {
	ChatId          string `json:"chat_id,omitempty"`
	MessageThreadId int64  `json:"message_thread_id,omitempty"`
}

type EditGeneralForumTopic struct {
	ChatId string `json:"chat_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

type CloseGeneralForumTopic struct {
	ChatId string `json:"chat_id,omitempty"`
}

type ReopenGeneralForumTopic struct {
	ChatId string `json:"chat_id,omitempty"`
}

type HideGeneralForumTopic struct {
	ChatId string `json:"chat_id,omitempty"`
}

type UnhideGeneralForumTopic struct {
	ChatId string `json:"chat_id,omitempty"`
}

type UnpinAllGeneralForumTopicMessages struct {
	ChatId string `json:"chat_id,omitempty"`
}

type AnswerCallbackQuery struct {
	CallbackQueryId string `json:"callback_query_id,omitempty"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       int64  `json:"cache_time,omitempty"`
}

type SetMyCommands struct {
	Commands     []BotCommand    `json:"commands,omitempty"`
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode string          `json:"language_code,omitempty"`
}

type DeleteMyCommands struct {
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode string          `json:"language_code,omitempty"`
}

type GetMyCommands struct {
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode string          `json:"language_code,omitempty"`
}

type SetMyName struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type GetMyName struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyDescription struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type GetMyDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyShortDescription struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

type GetMyShortDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetChatMenuButton struct {
	ChatId     int64      `json:"chat_id,omitempty"`
	MenuButton MenuButton `json:"menu_button,omitempty"`
}

type GetChatMenuButton struct {
	ChatId int64 `json:"chat_id,omitempty"`
}

type SetMyDefaultAdministratorRights struct {
	Rights      ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                    `json:"for_channels,omitempty"`
}

type GetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

type EditMessageText struct {
	ChatId                string               `json:"chat_id,omitempty"`
	MessageId             int64                `json:"message_id,omitempty"`
	InlineMessageId       string               `json:"inline_message_id,omitempty"`
	Text                  string               `json:"text,omitempty"`
	ParseMode             string               `json:"parse_mode,omitempty"`
	Entities              []MessageEntity      `json:"entities,omitempty"`
	DisableWebPagePreview bool                 `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageCaption struct {
	ChatId          string               `json:"chat_id,omitempty"`
	MessageId       int64                `json:"message_id,omitempty"`
	InlineMessageId string               `json:"inline_message_id,omitempty"`
	Caption         string               `json:"caption,omitempty"`
	ParseMode       string               `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageMedia struct {
	ChatId          string               `json:"chat_id,omitempty"`
	MessageId       int64                `json:"message_id,omitempty"`
	InlineMessageId string               `json:"inline_message_id,omitempty"`
	Media           InputMedia           `json:"media,omitempty"`
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageLiveLocation struct {
	ChatId               string               `json:"chat_id,omitempty"`
	MessageId            int64                `json:"message_id,omitempty"`
	InlineMessageId      string               `json:"inline_message_id,omitempty"`
	Latitude             float32              `json:"latitude,omitempty"`
	Longitude            float32              `json:"longitude,omitempty"`
	HorizontalAccuracy   float32              `json:"horizontal_accuracy,omitempty"`
	Heading              int64                `json:"heading,omitempty"`
	ProximityAlertRadius int64                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type StopMessageLiveLocation struct {
	ChatId          string               `json:"chat_id,omitempty"`
	MessageId       int64                `json:"message_id,omitempty"`
	InlineMessageId string               `json:"inline_message_id,omitempty"`
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageReplyMarkup struct {
	ChatId          string               `json:"chat_id,omitempty"`
	MessageId       int64                `json:"message_id,omitempty"`
	InlineMessageId string               `json:"inline_message_id,omitempty"`
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type StopPoll struct {
	ChatId      string               `json:"chat_id,omitempty"`
	MessageId   int64                `json:"message_id,omitempty"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type DeleteMessage struct {
	ChatId    string `json:"chat_id,omitempty"`
	MessageId int64  `json:"message_id,omitempty"`
}

type SendSticker struct {
	ChatId                   string               `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	Sticker                  InputFile            `json:"sticker,omitempty"`
	Emoji                    string               `json:"emoji,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type GetStickerSet struct {
	Name string `json:"name,omitempty"`
}

type GetCustomEmojiStickers struct {
	CustomEmojiIds []string `json:"custom_emoji_ids,omitempty"`
}

type UploadStickerFile struct {
	UserId        int64     `json:"user_id,omitempty"`
	Sticker       InputFile `json:"sticker,omitempty"`
	StickerFormat string    `json:"sticker_format,omitempty"`
}

type CreateNewStickerSet struct {
	UserId          int64          `json:"user_id,omitempty"`
	Name            string         `json:"name,omitempty"`
	Title           string         `json:"title,omitempty"`
	Stickers        []InputSticker `json:"stickers,omitempty"`
	StickerFormat   string         `json:"sticker_format,omitempty"`
	StickerType     string         `json:"sticker_type,omitempty"`
	NeedsRepainting bool           `json:"needs_repainting,omitempty"`
}

type AddStickerToSet struct {
	UserId  int64        `json:"user_id,omitempty"`
	Name    string       `json:"name,omitempty"`
	Sticker InputSticker `json:"sticker,omitempty"`
}

type SetStickerPositionInSet struct {
	Sticker  string `json:"sticker,omitempty"`
	Position int64  `json:"position,omitempty"`
}

type DeleteStickerFromSet struct {
	Sticker string `json:"sticker,omitempty"`
}

type SetStickerEmojiList struct {
	Sticker   string   `json:"sticker,omitempty"`
	EmojiList []string `json:"emoji_list,omitempty"`
}

type SetStickerKeywords struct {
	Sticker  string   `json:"sticker,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

type SetStickerMaskPosition struct {
	Sticker      string       `json:"sticker,omitempty"`
	MaskPosition MaskPosition `json:"mask_position,omitempty"`
}

type SetStickerSetTitle struct {
	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
}

type SetStickerSetThumbnail struct {
	Name      string    `json:"name,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	Thumbnail InputFile `json:"thumbnail,omitempty"`
}

type SetCustomEmojiStickerSetThumbnail struct {
	Name          string `json:"name,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

type DeleteStickerSet struct {
	Name string `json:"name,omitempty"`
}

type AnswerInlineQuery struct {
	InlineQueryId string                   `json:"inline_query_id,omitempty"`
	Results       []InlineQueryResult      `json:"results,omitempty"`
	CacheTime     int64                    `json:"cache_time,omitempty"`
	IsPersonal    bool                     `json:"is_personal,omitempty"`
	NextOffset    string                   `json:"next_offset,omitempty"`
	Button        InlineQueryResultsButton `json:"button,omitempty"`
}

type AnswerWebAppQuery struct {
	WebAppQueryId string            `json:"web_app_query_id,omitempty"`
	Result        InlineQueryResult `json:"result,omitempty"`
}

type SendInvoice struct {
	ChatId                    string               `json:"chat_id,omitempty"`
	MessageThreadId           int64                `json:"message_thread_id,omitempty"`
	Title                     string               `json:"title,omitempty"`
	Description               string               `json:"description,omitempty"`
	Payload                   string               `json:"payload,omitempty"`
	ProviderToken             string               `json:"provider_token,omitempty"`
	Currency                  string               `json:"currency,omitempty"`
	Prices                    []LabeledPrice       `json:"prices,omitempty"`
	MaxTipAmount              int64                `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64              `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string               `json:"start_parameter,omitempty"`
	ProviderData              string               `json:"provider_data,omitempty"`
	PhotoUrl                  string               `json:"photo_url,omitempty"`
	PhotoSize                 int64                `json:"photo_size,omitempty"`
	PhotoWidth                int64                `json:"photo_width,omitempty"`
	PhotoHeight               int64                `json:"photo_height,omitempty"`
	NeedName                  bool                 `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                 `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                 `json:"need_email,omitempty"`
	NeedShippingAddress       bool                 `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                 `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                 `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                 `json:"is_flexible,omitempty"`
	DisableNotification       bool                 `json:"disable_notification,omitempty"`
	ProtectContent            bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId          int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply  bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup               InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type CreateInvoiceLink struct {
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

type AnswerShippingQuery struct {
	ShippingQueryId string           `json:"shipping_query_id,omitempty"`
	Ok              bool             `json:"ok,omitempty"`
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string           `json:"error_message,omitempty"`
}

type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id,omitempty"`
	Ok                 bool   `json:"ok,omitempty"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

type SetPassportDataErrors struct {
	UserId int64                  `json:"user_id,omitempty"`
	Errors []PassportElementError `json:"errors,omitempty"`
}

type SendGame struct {
	ChatId                   int64                `json:"chat_id,omitempty"`
	MessageThreadId          int64                `json:"message_thread_id,omitempty"`
	GameShortName            string               `json:"game_short_name,omitempty"`
	DisableNotification      bool                 `json:"disable_notification,omitempty"`
	ProtectContent           bool                 `json:"protect_content,omitempty"`
	ReplyToMessageId         int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SetGameScore struct {
	UserId             int64  `json:"user_id,omitempty"`
	Score              int64  `json:"score,omitempty"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	ChatId             int64  `json:"chat_id,omitempty"`
	MessageId          int64  `json:"message_id,omitempty"`
	InlineMessageId    string `json:"inline_message_id,omitempty"`
}

type GetGameHighScores struct {
	UserId          int64  `json:"user_id,omitempty"`
	ChatId          int64  `json:"chat_id,omitempty"`
	MessageId       int64  `json:"message_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
}

type getUpdates struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int64    `json:"limit,omitempty"`
	Timeout        int64    `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type setWebhook struct {
	Url                string    `json:"url,omitempty"`
	Certificate        InputFile `json:"certificate,omitempty"`
	IpAddress          string    `json:"ip_address,omitempty"`
	MaxConnections     int64     `json:"max_connections,omitempty"`
	AllowedUpdates     []string  `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool      `json:"drop_pending_updates,omitempty"`
	SecretToken        string    `json:"secret_token,omitempty"`
}

type deleteWebhook struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

func getFuncName(fn_struct interface{}) (string, error) {
	switch fn_struct.(type) {
	case SetPassportDataErrors:
		return "setPassportDataErrors", nil
	case AnswerWebAppQuery:
		return "answerWebAppQuery", nil
	case AnswerInlineQuery:
		return "answerInlineQuery", nil
	case SendMessage:
		return "sendMessage", nil
	case ForwardMessage:
		return "forwardMessage", nil
	case CopyMessage:
		return "copyMessage", nil
	case SendPhoto:
		return "sendPhoto", nil
	case SendAudio:
		return "sendAudio", nil
	case SendDocument:
		return "sendDocument", nil
	case SendVideo:
		return "sendVideo", nil
	case SendAnimation:
		return "sendAnimation", nil
	case SendVoice:
		return "sendVoice", nil
	case SendVideoNote:
		return "sendVideoNote", nil
	case SendMediaGroup:
		return "sendMediaGroup", nil
	case SendLocation:
		return "sendLocation", nil
	case SendVenue:
		return "sendVenue", nil
	case SendContact:
		return "sendContact", nil
	case SendPoll:
		return "sendPoll", nil
	case SendDice:
		return "sendDice", nil
	case SendChatAction:
		return "sendChatAction", nil
	case GetUserProfilePhotos:
		return "getUserProfilePhotos", nil
	case GetFile:
		return "getFile", nil
	case BanChatMember:
		return "banChatMember", nil
	case UnbanChatMember:
		return "unbanChatMember", nil
	case RestrictChatMember:
		return "restrictChatMember", nil
	case PromoteChatMember:
		return "promoteChatMember", nil
	case SetChatAdministratorCustomTitle:
		return "setChatAdministratorCustomTitle", nil
	case BanChatSenderChat:
		return "banChatSenderChat", nil
	case UnbanChatSenderChat:
		return "unbanChatSenderChat", nil
	case SetChatPermissions:
		return "setChatPermissions", nil
	case ExportChatInviteLink:
		return "exportChatInviteLink", nil
	case CreateChatInviteLink:
		return "createChatInviteLink", nil
	case EditChatInviteLink:
		return "editChatInviteLink", nil
	case RevokeChatInviteLink:
		return "revokeChatInviteLink", nil
	case ApproveChatJoinRequest:
		return "approveChatJoinRequest", nil
	case DeclineChatJoinRequest:
		return "declineChatJoinRequest", nil
	case SetChatPhoto:
		return "setChatPhoto", nil
	case DeleteChatPhoto:
		return "deleteChatPhoto", nil
	case SetChatTitle:
		return "setChatTitle", nil
	case SetChatDescription:
		return "setChatDescription", nil
	case PinChatMessage:
		return "pinChatMessage", nil
	case UnpinChatMessage:
		return "unpinChatMessage", nil
	case UnpinAllChatMessages:
		return "unpinAllChatMessages", nil
	case LeaveChat:
		return "leaveChat", nil
	case GetChat:
		return "getChat", nil
	case GetChatAdministrators:
		return "getChatAdministrators", nil
	case GetChatMemberCount:
		return "getChatMemberCount", nil
	case GetChatMember:
		return "getChatMember", nil
	case SetChatStickerSet:
		return "setChatStickerSet", nil
	case DeleteChatStickerSet:
		return "deleteChatStickerSet", nil
	case CreateForumTopic:
		return "createForumTopic", nil
	case EditForumTopic:
		return "editForumTopic", nil
	case CloseForumTopic:
		return "closeForumTopic", nil
	case ReopenForumTopic:
		return "reopenForumTopic", nil
	case DeleteForumTopic:
		return "deleteForumTopic", nil
	case UnpinAllForumTopicMessages:
		return "unpinAllForumTopicMessages", nil
	case EditGeneralForumTopic:
		return "editGeneralForumTopic", nil
	case CloseGeneralForumTopic:
		return "closeGeneralForumTopic", nil
	case ReopenGeneralForumTopic:
		return "reopenGeneralForumTopic", nil
	case HideGeneralForumTopic:
		return "hideGeneralForumTopic", nil
	case UnhideGeneralForumTopic:
		return "unhideGeneralForumTopic", nil
	case UnpinAllGeneralForumTopicMessages:
		return "unpinAllGeneralForumTopicMessages", nil
	case AnswerCallbackQuery:
		return "answerCallbackQuery", nil
	case SetMyCommands:
		return "setMyCommands", nil
	case DeleteMyCommands:
		return "deleteMyCommands", nil
	case GetMyCommands:
		return "getMyCommands", nil
	case SetMyName:
		return "setMyName", nil
	case GetMyName:
		return "getMyName", nil
	case SetMyDescription:
		return "setMyDescription", nil
	case GetMyDescription:
		return "getMyDescription", nil
	case SetMyShortDescription:
		return "setMyShortDescription", nil
	case GetMyShortDescription:
		return "getMyShortDescription", nil
	case SetChatMenuButton:
		return "setChatMenuButton", nil
	case GetChatMenuButton:
		return "getChatMenuButton", nil
	case SetMyDefaultAdministratorRights:
		return "setMyDefaultAdministratorRights", nil
	case GetMyDefaultAdministratorRights:
		return "getMyDefaultAdministratorRights", nil
	case EditMessageText:
		return "editMessageText", nil
	case EditMessageCaption:
		return "editMessageCaption", nil
	case EditMessageMedia:
		return "editMessageMedia", nil
	case EditMessageLiveLocation:
		return "editMessageLiveLocation", nil
	case StopMessageLiveLocation:
		return "stopMessageLiveLocation", nil
	case EditMessageReplyMarkup:
		return "editMessageReplyMarkup", nil
	case StopPoll:
		return "stopPoll", nil
	case DeleteMessage:
		return "deleteMessage", nil
	case SendSticker:
		return "sendSticker", nil
	case GetStickerSet:
		return "getStickerSet", nil
	case GetCustomEmojiStickers:
		return "getCustomEmojiStickers", nil
	case UploadStickerFile:
		return "uploadStickerFile", nil
	case CreateNewStickerSet:
		return "createNewStickerSet", nil
	case AddStickerToSet:
		return "addStickerToSet", nil
	case SetStickerPositionInSet:
		return "setStickerPositionInSet", nil
	case DeleteStickerFromSet:
		return "deleteStickerFromSet", nil
	case SetStickerEmojiList:
		return "setStickerEmojiList", nil
	case SetStickerKeywords:
		return "setStickerKeywords", nil
	case SetStickerMaskPosition:
		return "setStickerMaskPosition", nil
	case SetStickerSetTitle:
		return "setStickerSetTitle", nil
	case SetStickerSetThumbnail:
		return "setStickerSetThumbnail", nil
	case SetCustomEmojiStickerSetThumbnail:
		return "setCustomEmojiStickerSetThumbnail", nil
	case DeleteStickerSet:
		return "deleteStickerSet", nil
	case SendInvoice:
		return "sendInvoice", nil
	case CreateInvoiceLink:
		return "createInvoiceLink", nil
	case AnswerShippingQuery:
		return "answerShippingQuery", nil
	case AnswerPreCheckoutQuery:
		return "answerPreCheckoutQuery", nil
	case SendGame:
		return "sendGame", nil
	case SetGameScore:
		return "setGameScore", nil
	case GetGameHighScores:
		return "getGameHighScores", nil
	case getUpdates:
		return "getUpdates", nil
	case setWebhook:
		return "setWebhook", nil
	case deleteWebhook:
		return "deleteWebhook", nil

	default:
		return "", errors.New("ERROR: unknown type of api methods (getFuncName)")
	}
}
