package osenv

import (
	"os"
)

// notice of feedRichMsg push url like   https://xxx.domain
func GetNoticeOfFeedRichMsgPushUrl() string {
	return os.Getenv("NOTICE_OF_FEEDRICHMSG_PUSH_URL")
}

// notice of feedRichMsg push url path like   /ws/push/
func GetNoticeOfFeedRichMsgPushUrlPath() string {
	return os.Getenv("NOTICE_OF_FEEDRICHMSG_PUSH_URL_Path")
}
