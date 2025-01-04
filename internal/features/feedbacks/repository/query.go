package repository

import (
	"chapter1/internal/features/feedbacks"
	"fmt"

	"gorm.io/gorm"
)

type FeedbackQuery struct {
	db *gorm.DB
}

func NewFeedbackQuery(connect *gorm.DB) feedbacks.FQuery {
	return &FeedbackQuery{
		db: connect,
	}
}

func (fq *FeedbackQuery) AddFeedback(newFeedback feedbacks.Feedback) error {
	cnvData := ToFeedbackQuery(newFeedback)
	qry := fq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	return nil

}

func (fq *FeedbackQuery) DeleteFeedback(userID uint, feedbackID uint) error {
	var feedback Feedback

	qry := fq.db.Where("id = ?", feedbackID).First(&feedback)
	if qry.Error != nil {
		return qry.Error
	}

	if feedback.MemberID != userID {
		return fmt.Errorf("unauthorized: you are not allowed to delete this feedback")
	}

	qry = fq.db.Where("id = ?", feedbackID).Delete(&Feedback{})
	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (fq *FeedbackQuery) GetAllFeedbacks(limit uint, page uint) ([]feedbacks.Feedback, uint, error) {
	var feedbacksList []Feedback
	var totalItem int64

	offset := (page - 1) * limit

	qry := fq.db.Model(&Feedback{}).Count(&totalItem)
	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	qry = fq.db.Preload("User").Preload("Book").Limit(int(limit)).Offset(int(offset)).Find(&feedbacksList)
	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	feedbacksEntities := make([]feedbacks.Feedback, len(feedbacksList))
	for i, feedback := range feedbacksList {
		feedbacksEntities[i] = feedback.ToFeedbackEntity()
	}

	return feedbacksEntities, uint(totalItem), nil
}
