package main

import (
	"context"
	"errors"
	"log"

	"github.com/rs/xid"
)

var errInvalidNameOrSurname = errors.New("invalid name or surname given")
var errPersonNotFound = errors.New("person not found with the given ID")

type demoRPCServer struct {
	// This is for demo purposes, the RPC service should have a data provider
	// such a database or whatsoever. DON'T DO THIS IN PRODUCTION!
	personList map[string]Person
}

func newDemoRPCServer() *demoRPCServer {
	return &demoRPCServer{
		personList: make(map[string]Person),
	}
}

// CreatePerson instantiates a new Person with the given args, if they're correct, stores the Person
// in the local list and returns it
func (s *demoRPCServer) CreatePerson(ctx context.Context, args *CreatePersonArgs) (*Person, error) {
	log.Println("got CreatePerson request")

	newPerson, err := newPerson(args.GetName(), args.GetSurname(), args.GetAge())
	if err != nil {
		return nil, err
	}

	s.personList[newPerson.GetId()] = *newPerson

	return newPerson, nil
}

// ReadPerson returns the Person matching the given args.ID, if present, or an error otherwise
func (s *demoRPCServer) ReadPerson(ctx context.Context, args *ReadPersonArgs) (*Person, error) {
	log.Println("got ReadPerson request")

	person, personExists := s.personList[args.GetId()]
	if !personExists {
		return nil, errPersonNotFound
	}

	return &person, nil
}

func newPerson(name, surname string, age uint32) (*Person, error) {
	if name == "" || surname == "" {
		return nil, errInvalidNameOrSurname
	}

	return &Person{
		Id:      xid.New().String(),
		Name:    name,
		Surname: surname,
		Age:     age,
	}, nil
}
