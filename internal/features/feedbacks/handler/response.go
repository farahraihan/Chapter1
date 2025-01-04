package handler

import (
	"chapter1/internal/features/feedbacks"
	"time"
)

type FeedbackResponse struct {
	ID         uint      `json:"feedbackID"`
	Content    string    `json:"content"`
	Rating     uint      `json:"rating"`
	MemberID   uint      `json:"memberID"`
	BookID     uint      `json:"bookID"`
	MemberName string    `json:"memberName"`
	BookTitle  string    `json:"bookTitle"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToFeedbackResponse(input feedbacks.Feedback) FeedbackResponse {
	return FeedbackResponse{
		ID:         input.ID,
		Content:    input.Content,
		Rating:     input.Rating,
		MemberID:   input.MemberID,
		BookID:     input.BookID,
		MemberName: input.MemberName,
		BookTitle:  input.BookTitle,
	}
}

func ToFeedbackResponses(feedbacks []feedbacks.Feedback) []FeedbackResponse {
	responses := make([]FeedbackResponse, len(feedbacks))
	for i, feedback := range feedbacks {
		responses[i] = ToFeedbackResponse(feedback)
	}
	return responses
}
