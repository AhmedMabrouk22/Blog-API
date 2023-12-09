package utils

import "main/models"

func GetBlogRes(blog models.Blog) models.BlogResponse {
	var blogRes models.BlogResponse
	blogRes.ID = blog.ID
	blogRes.Title = blog.Title
	blogRes.Content = blog.Content
	blogRes.ImageCover = blog.ImageCover
	blogRes.Author.ID = blog.Author.ID
	blogRes.Author.Name = blog.Author.Name
	blogRes.Author.Image = blog.Author.Image
	blogRes.Topics = blog.Topics

	var comments []models.CommentRes
	for _, val := range blog.Comments {
		comment := models.CommentRes{
			ID:      val.ID,
			Content: val.Content,
			Author: models.CommentAuthor{
				ID:    val.User.ID,
				Name:  val.User.Name,
				Image: val.User.Image,
			},
		}
		comments = append(comments, comment)
	}

	blogRes.Comments = comments

	var likesRes []models.LikeRes
	for _, val := range blog.Likes {
		like := models.LikeRes{
			ID: val.ID,
			User: models.LikeUser{
				ID:    val.User.ID,
				Name:  val.User.Name,
				Image: val.User.Image,
			},
		}
		likesRes = append(likesRes, like)
	}

	blogRes.Likes = likesRes

	return blogRes
}
