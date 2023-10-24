package bot

import (
	"github.com/enescakir/emoji"
)

const (
	creatorChatID int64  = 5785150199
)

var (
	isBotRunning bool
)

var (
	LeafGreenEmoji  = emoji.Sprintf("%v", emoji.LeafyGreen)
	HelpEmoji       = emoji.Sprintf("%v", emoji.Information)
	StartOkEmoji    = emoji.Sprintf("%v", emoji.GreenCircle)
	InfoEmoji       = emoji.Sprintf("%v", emoji.Information)
	OkEmoji         = emoji.Sprintf("%v", emoji.GreenCircle)
	StopEmoji       = emoji.Sprintf("%v", emoji.RedCircle)
	GotchaEmoji     = emoji.Sprintf("%v", emoji.OkHand)
	GetdataEmoji    = emoji.Sprintf("%v", emoji.Cityscape)
	GreebHeartEmoji = emoji.Sprintf("%v", emoji.GreenHeart)
)

var (
	YesMsg                = GotchaEmoji + " one second please..."
	NoMsg                 = GotchaEmoji + " i gotcha!"
	StatusOkMsg           = OkEmoji + " everything is ok!"
	StatusNotOkMsg        = StopEmoji + " something is wrong!"
	AdviceMsg             = "would you like to receive advice on what is best to do on this day (based on data you got)?"
	FetchingMsg           = GetdataEmoji + " please select a station from the list:"
	HelpMsg               = HelpEmoji + " all available commands:\n\n/help - use if you need some help\n\n/getdata - use to get all ecoman data, specify the city name\n\n/status - use to see working status\n\n/info - use to see more info about creator and bot\n\n/support - use if you found a bug etc.\n\n/support_creator - It's an open-source free-to-use product, so I don't get any money from it\n\n/stop - use stop command to stop the bot"
	StartMsg              = LeafGreenEmoji + " ecoman is running!"
	StopMsg               = StopEmoji + " ecoman has been stopped"
	AlreadyRunningMsg     = LeafGreenEmoji + " ecoman is already running!"
	AlreadyStoppedMsg     = StopEmoji + " ecoman is already stopped!"
	DescribeTheProblemMsg = LeafGreenEmoji + " please describe the problem:"
	ThxMsg                = GreebHeartEmoji + " thank you for your feedback!"
	BugReportMsg          = "[ ! ] bug report from user @%s:\n[ ! ] %s"
	IdkMsg                = "[ ? ] i don't know what to do with this command\n[ ? ] use /help to see all available commands"
	InfoMsg               = InfoEmoji + " ecoman is a telegram bot that allows you to get the latest information about ecology in Ukraine\ndata updates every 15 minutes"
	NoDataMsg             = "no data found for city: %s"
)
