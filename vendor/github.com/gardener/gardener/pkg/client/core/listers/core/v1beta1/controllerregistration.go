// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ControllerRegistrationLister helps list ControllerRegistrations.
type ControllerRegistrationLister interface {
	// List lists all ControllerRegistrations in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.ControllerRegistration, err error)
	// Get retrieves the ControllerRegistration from the index for a given name.
	Get(name string) (*v1beta1.ControllerRegistration, error)
	ControllerRegistrationListerExpansion
}

// controllerRegistrationLister implements the ControllerRegistrationLister interface.
type controllerRegistrationLister struct {
	indexer cache.Indexer
}

// NewControllerRegistrationLister returns a new ControllerRegistrationLister.
func NewControllerRegistrationLister(indexer cache.Indexer) ControllerRegistrationLister {
	return &controllerRegistrationLister{indexer: indexer}
}

// List lists all ControllerRegistrations in the indexer.
func (s *controllerRegistrationLister) List(selector labels.Selector) (ret []*v1beta1.ControllerRegistration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ControllerRegistration))
	})
	return ret, err
}

// Get retrieves the ControllerRegistration from the index for a given name.
func (s *controllerRegistrationLister) Get(name string) (*v1beta1.ControllerRegistration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("controllerregistration"), name)
	}
	return obj.(*v1beta1.ControllerRegistration), nil
}
