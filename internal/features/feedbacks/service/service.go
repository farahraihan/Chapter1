package service

import (
	"chapter1/internal/features/feedbacks"
	"chapter1/internal/features/users"
	"errors"
	"log"
)

type FeedbackServices struct {
	qry      feedbacks.FQuery
	uService users.UService
}

func NewFeedbackServices(q feedbacks.FQuery, u users.UService) feedbacks.FService {
	return &FeedbackServices{
		qry:      q,
		uService: u,
	}
}

func (fs *FeedbackServices) AddFeedback(newFeedback feedbacks.Feedback) error {
	err := fs.qry.AddFeedback(newFeedback)
	if err != nil {
		log.Println("add feedback query error: ", err)
		return errors.New("failed to add feedback, please try again later")
	}

	err = fs.uService.AddPoints(newFeedback.MemberID, 50)
	if err != nil {
		log.Println("add user point query error: ", err)
		return errors.New("failed to add user point, please try again later")
	}

	return nil
}

func (fs *FeedbackServices) DeleteFeedback(userID uint, feedbackID uint) error {
	err := fs.qry.DeleteFeedback(userID, feedbackID)
	if err != nil {
		log.Println("delete feedback query error")
		return errors.New("delete failed, please try again later")
	}

	return nil
}

func (fs *FeedbackServices) GetAllFeedbacks(limit uint, page uint) ([]feedbacks.Feedback, uint, error) {
	feedbacks, totalItems, err := fs.qry.GetAllFeedbacks(limit, page)

	if err != nil {
		log.Println("get all feedbacks query error: ", err)
		return nil, 0, errors.New("failed to retrieve feedbacks, please try again later")
	}

	return feedbacks, totalItems, nil
}
