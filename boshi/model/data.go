package model

type AutoGenerated struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
type PublishGroup struct {
	ID        string `json:"_id"`
	GroupName string `json:"groupName"`
	Inc       int    `json:"inc"`
}
type TopicOption struct {
	OptionContent string `json:"option_content"`
	OptionImg     string `json:"option_img"`
	CorrectOption int    `json:"correct_option"`
	OptionID      string `json:"option_id"`
}
type AnswerInfo struct {
	ID                string        `json:"_id"`
	Egroup            []int         `json:"egroup"`
	CTimestamp        int64         `json:"c_timestamp"`
	Thumbnail         interface{}   `json:"thumbnail"`
	TopicResourceID   interface{}   `json:"topic_resource_id"`
	GroupID           string        `json:"groupId"`
	CTime             string        `json:"c_time"`
	TopicType         int           `json:"topic_type"`
	TopicLabel        []int         `json:"topic_label"`
	TopicResourceName interface{}   `json:"topic_resource_name"`
	TopicScore        int           `json:"topic_score"`
	TopicOption       []TopicOption `json:"topic_option"`
	TopicSkill        []interface{} `json:"topic_skill"`
	TopicResource     interface{}   `json:"topic_resource"`
	UserID            string        `json:"user_id"`
	TopicContent      string        `json:"topic_content"`
	QuoteCount        int           `json:"quote_count"`
	TopicResolve      interface{}   `json:"topic_resolve"`
	TopicLevel        int           `json:"topic_level"`
	Status            int           `json:"status"`
	Username          string        `json:"username"`
	UTimestamp        int64         `json:"u_timestamp,omitempty"`
	UTime             string        `json:"u_time,omitempty"`
}
type Exam struct {
	ID            string         `json:"_id"`
	HaveTempTopic bool           `json:"haveTempTopic"`
	CTimestamp    int64          `json:"c_timestamp"`
	UserGroup     string         `json:"user_group"`
	TimeConsuming int            `json:"time_consuming"`
	GroupID       string         `json:"groupId"`
	CTime         string         `json:"c_time"`
	AnswerStatus  int            `json:"answer_status"`
	Ispass        int            `json:"ispass"`
	AnswerScore   int            `json:"answer_score"`
	UserID        string         `json:"user_id"`
	AnswerUserid  string         `json:"answer_userid"`
	ErrorSum      int            `json:"error_sum"`
	ExamID        string         `json:"exam_id"`
	Status        int            `json:"status"`
	Username      string         `json:"username"`
	PublishGroup  []PublishGroup `json:"publish_group"`
	AnswerInfo    []AnswerInfo   `json:"answer_info"`
	BeginTime     string         `json:"begin_time"`
	EndTime       string         `json:"end_time"`
	ExamName      string         `json:"exam_name"`
	CurrentStatus string         `json:"currentStatus"`
	Scorerate     string         `json:"scorerate"`
}
type User struct {
	Nickname string `json:"nickname"`
}
type Data struct {
	Exam Exam `json:"exam"`
	User User `json:"user"`
}