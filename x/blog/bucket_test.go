package blog

import (
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/weavetest/assert"
)

func TestBlogUserIDIndexer(t *testing.T) {
	now := weave.AsUnixTime(time.Now())

	userID := weavetest.SequenceID(1)

	blog := &Blog{
		Metadata:   &weave.Metadata{Schema: 1},
		PrimaryKey: weavetest.SequenceID(1),
		Owner:      userID,
		Title:      "Best hacker's blog",
		CreatedAt:  now,
	}

	cases := map[string]struct {
		obj      orm.Object
		expected []byte
		wantErr  *errors.Error
	}{
		"success": {
			obj:      orm.NewSimpleObj(nil, blog),
			expected: userID,
			wantErr:  nil,
		},
		"failure, obj is nil": {
			obj:      nil,
			expected: nil,
			wantErr:  nil,
		},
		"not blog": {
			obj:      orm.NewSimpleObj(nil, new(User)),
			expected: nil,
			wantErr:  errors.ErrState,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			index, err := blogUserIDIndexer(tc.obj)

			if !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
			assert.Equal(t, tc.expected, index)
		})
	}
}
func TestArticleBlogIDIndexer(t *testing.T) {
	now := weave.AsUnixTime(time.Now())
	future := now.Add(time.Hour)

	blogID := weavetest.SequenceID(1)

	article := &Article{
		Metadata:   &weave.Metadata{Schema: 1},
		PrimaryKey: weavetest.SequenceID(1),
		BlogKey:     blogID,
		Title:      "Best hacker's blog",
		Content:    "Best description ever",
		CreatedAt:  now,
		DeleteAt:   future,
	}

	cases := map[string]struct {
		obj      orm.Object
		expected []byte
		wantErr  *errors.Error
	}{
		"success": {
			obj:      orm.NewSimpleObj(nil, article),
			expected: blogID,
			wantErr:  nil,
		},
		"failure, obj is nil": {
			obj:      nil,
			expected: nil,
			wantErr:  nil,
		},
		"not article": {
			obj:      orm.NewSimpleObj(nil, new(Blog)),
			expected: nil,
			wantErr:  errors.ErrState,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			index, err := articleBlogIDIndexer(tc.obj)

			if !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
			assert.Equal(t, tc.expected, index)
		})
	}
}

func TestBlogTimedIndexer(t *testing.T) {
	now := weave.AsUnixTime(time.Unix(1, 0))
	invalidTime := weave.AsUnixTime(time.Unix(-1, 0))
	future := now.Add(time.Hour)

	blogID := weavetest.SequenceID(1)

	article := &Article{
		Metadata:   &weave.Metadata{Schema: 1},
		PrimaryKey: weavetest.SequenceID(1),
		BlogKey:     blogID,
		Title:      "Best hacker's blog",
		Content:    "Best description ever",
		CreatedAt:  now,
		DeleteAt:   future,
	}

	invalidArticle := &Article{
		Metadata:   &weave.Metadata{Schema: 1},
		PrimaryKey: weavetest.SequenceID(1),
		BlogKey:     blogID,
		Title:      "Best hacker's blog",
		Content:    "Best description ever",
		CreatedAt:  invalidTime,
		DeleteAt:   future,
	}

	// the index is by article and time, not by the blog
	successCaseExpectedValue := []byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}

	cases := map[string]struct {
		obj      orm.Object
		expected []byte
		wantErr  *errors.Error
	}{
		"success": {
			obj:      orm.NewSimpleObj(nil, article),
			expected: successCaseExpectedValue,
			wantErr:  nil,
		},
		"failure obj is nil": {
			obj:      nil,
			expected: nil,
			wantErr:  nil,
		},
		"not article": {
			obj:      orm.NewSimpleObj(nil, new(Blog)),
			expected: nil,
			wantErr:  errors.ErrState,
		},
		"empty obj has nil value": {
			obj:      orm.NewSimpleObj(nil, nil),
			expected: nil,
			wantErr:  nil,
		},
		"invalid creation time": {
			obj:      orm.NewSimpleObj(nil, invalidArticle),
			expected: nil,
			wantErr:  errors.ErrState,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			index, err := blogTimedIndexer(tc.obj)

			if !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
			assert.Equal(t, tc.expected, index)
		})
	}
}
