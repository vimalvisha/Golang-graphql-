package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/vishal/gqlgen-todos/graph/generated"
	"github.com/vishal/gqlgen-todos/graph/model"
	"github.com/vishal/gqlgen-todos/graph/postgres"
	"github.com/vishal/gqlgen-todos/graph/util"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Updatestudent(ctx context.Context, input model.Iperson) (*model.Operson, error) {
	postgres.Update(&input.ID, &input.FullName)
	var arr1 []*string
	arr := []string{"value updated"}
	for i := range arr {
		arr1 = append(arr1, &arr[i])
	}
	s := model.Operson{}
	return &s, nil
}

func (r *mutationResolver) Chart(ctx context.Context, input model.Iperson) (*model.Operson, error) {
	arr := postgres.Retrievedataforchart()
	fmt.Println(arr)
	output := model.Operson{"chart prepared in excel sheet taking data from psql"}
	return &output, nil
}

func (r *mutationResolver) Updatedatafromexcel(ctx context.Context) (string, error) {
	postgres.SetFromExcel()
	return "ok", nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, input model.Userregistration) (string, error) {
	var userdata model.Userregistration
	userdata.EmailID = input.EmailID
	userdata.Username = input.Username
	userdata.Pasword = input.Pasword
	userdata.PhoneNo = input.PhoneNo
	postgres.Registeration(input.EmailID, input.Username, util.Encode(input.Pasword), input.PhoneNo)
	return "Registered", nil
}

func (r *queryResolver) Login(ctx context.Context, input model.Loginip) ([]*model.Loginop, error) {
	userarr, err := postgres.UserResponse(input.EmailID, util.Encode(input.Pasword))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return userarr, nil
}

func (r *queryResolver) Fetch(ctx context.Context, input model.Getdata) ([]*model.Loginop, error) {
	getuser := postgres.FetchResponse(input.First, input.After)
	return getuser, nil
}

func (r *queryResolver) Insertvalue(ctx context.Context, input model.Iperson) (string, error) {
	var persondata model.Iperson
	persondata.ID = input.ID
	persondata.FullName = input.FullName
	postgres.AddStudent(input.ID, input.FullName)
	return "ok", nil
}

func (r *queryResolver) Makeexcel(ctx context.Context) (string, error) {
	postgres.CreateExcel()
	return "ok", nil
}

func (r *queryResolver) Insertperson(ctx context.Context, input model.Cperson) (string, error) {
	var pdata model.Cperson
	pdata.ID = input.ID
	pdata.FirstName = input.FirstName
	pdata.LastName = input.LastName
	pdata.Gender = input.Gender
	pdata.CarID = input.CarID
	postgres.Addperson(input.ID, input.FirstName, input.LastName, input.Gender, input.CarID)
	return "ok", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
