package todo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	mongoDB *mongo.Database
}

func (r *TodoRepository) FindAll() ([]*Todo, error) {
	cursor, err := r.mongoDB.Collection("todo").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, errors.New("No documents.")
	}
	defer cursor.Close(context.TODO())

	todoList := make([]*Todo, 0)
	for cursor.Next(context.TODO()) {
		var todo Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, err
		}
		todoList = append(todoList, &todo)
	}
	return todoList, nil
}

func (r *TodoRepository) FindById(id string) (*Todo, error) {
	var todo Todo
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid provided id.")
	}
	err = r.mongoDB.Collection("todo").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&todo)
	if err != nil {
		return nil, errors.New("No documents.")
	}
	return &todo, nil
}

func (r *TodoRepository) Create(todo *Todo) (*Todo, error) {
	dataToInsert := bson.M{
		"name":        todo.Name,
		"description": todo.Description,
		"status":      todo.Status,
	}
	insertedResult, err := r.mongoDB.Collection("todo").InsertOne(context.TODO(), dataToInsert)
	if err != nil {
		return nil, errors.New("Error while creating todo.")
	}
	if oid, ok := insertedResult.InsertedID.(primitive.ObjectID); ok {
		todo.ID = oid
	} else {
		return nil, errors.New("Error while creating todo.")
	}
	return todo, nil
}

func (r *TodoRepository) Update(id string, todo *Todo) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid provided id.")
	}
	filter := bson.M{"_id": oid}
	dataToUpdate := make(bson.M)
	if todo.Name != "" {
		dataToUpdate["name"] = todo.Name
	}
	if todo.Description != "" {
		dataToUpdate["description"] = todo.Description
	}
	if todo.Status != "" {
		dataToUpdate["status"] = todo.Status
	}
	update := bson.M{
		"$set": dataToUpdate,
	}
	_, err = r.mongoDB.Collection("todo").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Error while updating todo.")
	}
	return nil
}

func (r *TodoRepository) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid provided id.")
	}
	filter := bson.M{"_id": oid}
	deletedResult, err := r.mongoDB.Collection("todo").DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("Error while deleting todo.")
	}
	if deletedResult.DeletedCount == 0 {
		return errors.New("No documents found.")
	}
	return nil
}

func NewTodoRepository(mongo *mongo.Database) *TodoRepository {
	return &TodoRepository{
		mongoDB: mongo,
	}
}
