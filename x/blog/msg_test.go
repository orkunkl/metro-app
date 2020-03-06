package blog

import (
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/weavetest/assert"
)

func TestValidateCreateUserMsg(t *testing.T) {
	cases := map[string]struct {
		msg      weave.Msg
		wantErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CreateUserMsg{
				Metadata: &weave.Metadata{Schema: 1},
				Username: "Crpto0X",
				Bio:      "Best hacker in the universe",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"Username": nil,
				"Bio":      nil,
			},
		},
		// add missing metadata test
		"failure missing username": {
			msg: &CreateUserMsg{
				Metadata: &weave.Metadata{Schema: 1},
				Bio:      "Best hacker in the universe",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"Username": errors.ErrModel,
				"Bio":      nil,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.msg.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}

func TestValidateCreateBlogMsg(t *testing.T) {
	cases := map[string]struct {
		msg      weave.Msg
		wantErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CreateBlogMsg{
				Metadata:    &weave.Metadata{Schema: 1},
				Title:       "insanely good title",
				Description: "best description in the existence",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":    nil,
				"Title":       nil,
				"Description": nil,
			},
		},
		// add missing metadata test
		"failure missing title": {
			msg: &CreateBlogMsg{
				Metadata:    &weave.Metadata{Schema: 1},
				Description: "best description in the existence",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":    nil,
				"Title":       errors.ErrModel,
				"Description": nil,
			},
		},
		"failure missing description": {
			msg: &CreateBlogMsg{
				Metadata: &weave.Metadata{Schema: 1},
				Title:    "insanely good title",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":    nil,
				"Title":       nil,
				"Description": errors.ErrModel,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.msg.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}

func TestChangeBlogOwnerMsg(t *testing.T) {
	cases := map[string]struct {
		msg      weave.Msg
		wantErrs map[string]*errors.Error
	}{
		"success": {
			msg: &ChangeBlogOwnerMsg{
				Metadata: &weave.Metadata{Schema: 1},
				BlogKey:  weavetest.SequenceID(1),
				NewOwner: weavetest.NewCondition().Address(),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  nil,
				"NewOwner": nil,
			},
		},
		// add missing metadata test
		"failure missing blog id": {
			msg: &ChangeBlogOwnerMsg{
				Metadata: &weave.Metadata{Schema: 1},
				NewOwner: weavetest.NewCondition().Address(),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  errors.ErrEmpty,
				"NewOwner": nil,
			},
		},
		"failure invalid blog id": {
			msg: &ChangeBlogOwnerMsg{
				Metadata: &weave.Metadata{Schema: 1},
				BlogKey:  []byte{0, 0},
				NewOwner: weavetest.NewCondition().Address(),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  errors.ErrInput,
				"NewOwner": nil,
			},
		},
		"failure missing owner": {
			msg: &ChangeBlogOwnerMsg{
				Metadata: &weave.Metadata{Schema: 1},
				BlogKey:  weavetest.SequenceID(1),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  nil,
				"NewOwner": errors.ErrInput,
			},
		},
		"failure invalid owner": {
			msg: &ChangeBlogOwnerMsg{
				Metadata: &weave.Metadata{Schema: 1},
				BlogKey:  weavetest.SequenceID(1),
				NewOwner: []byte{0, 0},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  nil,
				"NewOwner": errors.ErrInput,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.msg.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}

func TestValidateCreateArticleMsg(t *testing.T) {
	cases := map[string]struct {
		msg      weave.Msg
		wantErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CreateArticleMsg{
				Metadata: &weave.Metadata{Schema: 1},
				BlogKey:  weavetest.SequenceID(1),
				Title:    "insanely good title",
				Content:  "best content in the existence",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  nil,
				"Title":    nil,
				"Content":  nil,
			},
		},
		// add missing metadata test
		"failure missing blog id": {
			msg: &CreateArticleMsg{
				Metadata: &weave.Metadata{Schema: 1},
				Title:    "insanely good title",
				Content:  "best content in the existence",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  errors.ErrEmpty,
				"Title":    nil,
				"Content":  nil,
			},
		},
		"failure missing title": {
			msg: &CreateArticleMsg{
				Metadata: &weave.Metadata{Schema: 1},
				BlogKey:  weavetest.SequenceID(1),
				Content:  "best content in the existence",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  nil,
				"Title":    errors.ErrModel,
				"Content":  nil,
			},
		},
		"failure missing content": {
			msg: &CreateArticleMsg{
				Metadata: &weave.Metadata{Schema: 1},
				BlogKey:  weavetest.SequenceID(1),
				Title:    "insanely good title",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata": nil,
				"BlogKey":  nil,
				"Title":    nil,
				"Content":  errors.ErrModel,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.msg.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}

func TestValidateDeleteArticle(t *testing.T) {
	cases := map[string]struct {
		msg      weave.Msg
		wantErrs map[string]*errors.Error
	}{
		"success": {
			msg: &DeleteArticleMsg{
				Metadata:   &weave.Metadata{Schema: 1},
				ArticleKey: weavetest.SequenceID(1),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"ArticleKey": nil,
			},
		},
		// add missing metadata test
		"failure missing article id": {
			msg: &DeleteArticleMsg{
				Metadata: &weave.Metadata{Schema: 1},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"ArticleKey": errors.ErrEmpty,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.msg.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}

func TestCancelDeleteArticleTask(t *testing.T) {
	cases := map[string]struct {
		msg      weave.Msg
		wantErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CancelDeleteArticleTaskMsg{
				Metadata:   &weave.Metadata{Schema: 1},
				ArticleKey: weavetest.SequenceID(1),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"ArticleKey": nil,
			},
		},
		// add missing metadata test
		"failure missing task id": {
			msg: &CancelDeleteArticleTaskMsg{
				Metadata: &weave.Metadata{Schema: 1},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"ArticleKey": errors.ErrEmpty,
			},
		},
		"failure invalid task id": {
			msg: &CancelDeleteArticleTaskMsg{
				Metadata:   &weave.Metadata{Schema: 1},
				ArticleKey: []byte{0, 0},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"ArticleKey": errors.ErrInput,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.msg.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}
