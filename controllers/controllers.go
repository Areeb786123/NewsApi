package controllers

import (
	"context"
	"fmt"
	dbconnection "newsApi/dbConnection"
	"newsApi/entitiy"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNews(entitiy.NewsModels) {

}

func DeleteNewsById(newsId string) {

}

func DeleteAllNews() {

}

func GetAllNews() entitiy.ResponseEntitity {
	var newsList []entitiy.NewsModels
	getAllNews, err := dbconnection.NewsCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
		return entitiy.ResponseEntitity{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	}
	// Iterate over the cursor and decode the data into the newsList
	for getAllNews.Next(context.Background()) {
		var newsData entitiy.NewsModels
		err := getAllNews.Decode(&newsData)
		if err != nil {
			fmt.Println(err)
			continue // Skip this document if there's an error decoding it
		}
		newsList = append(newsList, newsData)
	}

	// Check if there were any issues during the iteration
	if err := getAllNews.Err(); err != nil {
		fmt.Println(err)
		return entitiy.ResponseEntitity{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	}

	// Successfully return the newsList
	return entitiy.ResponseEntitity{
		Data:    newsList,
		Message: "success",
		Success: true,
	}
}

func GetNewsById(id string) entitiy.ResponseEntitity {
	var news entitiy.NewsModels
	newsId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entitiy.ResponseEntitity{
			Data:    nil,
			Success: false,
			Message: err.Error(),
		}
	}
	filter := bson.M{"_id": newsId}
	err = dbconnection.NewsCollection.FindOne(context.Background(), filter).Decode(&news)
	if err != nil {
		return entitiy.ResponseEntitity{
			Data:    nil,
			Success: false,
			Message: "no such news found",
		}
	}
	fmt.Println(news)
	return entitiy.ResponseEntitity{
		Data:    news,
		Success: true,
		Message: "success",
	}
}
