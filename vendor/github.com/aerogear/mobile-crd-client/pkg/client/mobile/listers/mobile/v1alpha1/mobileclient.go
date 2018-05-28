/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/aerogear/mobile-crd-client/pkg/apis/mobile/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MobileClientLister helps list MobileClients.
type MobileClientLister interface {
	// List lists all MobileClients in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MobileClient, err error)
	// MobileClients returns an object that can list and get MobileClients.
	MobileClients(namespace string) MobileClientNamespaceLister
	MobileClientListerExpansion
}

// mobileClientLister implements the MobileClientLister interface.
type mobileClientLister struct {
	indexer cache.Indexer
}

// NewMobileClientLister returns a new MobileClientLister.
func NewMobileClientLister(indexer cache.Indexer) MobileClientLister {
	return &mobileClientLister{indexer: indexer}
}

// List lists all MobileClients in the indexer.
func (s *mobileClientLister) List(selector labels.Selector) (ret []*v1alpha1.MobileClient, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MobileClient))
	})
	return ret, err
}

// MobileClients returns an object that can list and get MobileClients.
func (s *mobileClientLister) MobileClients(namespace string) MobileClientNamespaceLister {
	return mobileClientNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MobileClientNamespaceLister helps list and get MobileClients.
type MobileClientNamespaceLister interface {
	// List lists all MobileClients in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MobileClient, err error)
	// Get retrieves the MobileClient from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MobileClient, error)
	MobileClientNamespaceListerExpansion
}

// mobileClientNamespaceLister implements the MobileClientNamespaceLister
// interface.
type mobileClientNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MobileClients in the indexer for a given namespace.
func (s mobileClientNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MobileClient, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MobileClient))
	})
	return ret, err
}

// Get retrieves the MobileClient from the indexer for a given namespace and name.
func (s mobileClientNamespaceLister) Get(name string) (*v1alpha1.MobileClient, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mobileclient"), name)
	}
	return obj.(*v1alpha1.MobileClient), nil
}