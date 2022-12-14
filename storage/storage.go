package storage

import (
	//"time"
	"errors"
)

var ErrFailedDeletion = errors.New("Failed deletion!")
var ErrFailedSearch = errors.New("Failed search!")
var ErrBadData = errors.New("Bad data!")

type Storage interface {
	GetCollectionReferences(int, string) (*[]NameReference, error)
	CreateCollectionByName(string) (*NameReference, error)
	GetReferenceCollectionById(string, int, string) (*ReferenceCollection, error)
	DeleteCollectionById(string) error

	AddDataPointsToCollectionById(string, []DataPoint) (*ReferenceGroups, error)
	GetDataInCollectionById(string, string, string) (*DataPoint, error)
	DeleteDataFromCollectionById(string, string, string) error

	GetNamedTypes() (*[]NamedType, error)
	GetNamedTypeById(string) (*NamedType, error)
	CreateNamedType(string, string) (*NamedType, error)
	DeleteNamedTypeById(string) error
}
