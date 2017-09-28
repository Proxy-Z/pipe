// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package model

// Comment on types.
const (
	CommentOnTypeArticle = iota
	CommentOnTypePage
)

// Comment model.
type Comment struct {
	Model

	OnID                      uint   `json:"onID"`                 // ID of article or page
	OnType                    int    `gorm:"size:8" json:"onType"` // 0: article, 1: page
	AuthorName                string `gorm:"size:32" json:"authorName"`
	AuthorAvatarURL           string `gorm:"size:255" json:"authorAvatarURL"`
	Content                   string `gorm:"type:text" json:"content"`
	OriginalCommentID         uint   `json:"originalCommentID"` // ID of replied comment
	OriginalCommentAuthorName string `gorm:"size:32" json:"originalCommentAuthorName"`

	BlogID uint `json:"blogID"`
}