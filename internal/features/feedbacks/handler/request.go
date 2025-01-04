package handler

import "chapter1/internal/features/feedbacks"

type AddFeedbackRequest struct {
	Content  string `json:"content" form:"content"`
	Rating   uint   `json:"rating" form:"rating"`
	MemberID uint   `json:"memberID" form:"memberID"`
	BookID   uint   `json:"bookID" form:"bookID"`
}

func ToFeedbackModel(fr AddFeedbackRequest) feedbacks.Feedback {
	return feedbacks.Feedback{
		Content: fr.Content,
		Rating:  fr.Rating,
		BookID:  fr.BookID,
	}
}
