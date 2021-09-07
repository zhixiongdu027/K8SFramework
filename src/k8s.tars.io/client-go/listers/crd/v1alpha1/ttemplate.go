/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.tars.io/api/crd/v1alpha1"
)

// TTemplateLister helps list TTemplates.
type TTemplateLister interface {
	// List lists all TTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TTemplate, err error)
	// TTemplates returns an object that can list and get TTemplates.
	TTemplates(namespace string) TTemplateNamespaceLister
	TTemplateListerExpansion
}

// tTemplateLister implements the TTemplateLister interface.
type tTemplateLister struct {
	indexer cache.Indexer
}

// NewTTemplateLister returns a new TTemplateLister.
func NewTTemplateLister(indexer cache.Indexer) TTemplateLister {
	return &tTemplateLister{indexer: indexer}
}

// List lists all TTemplates in the indexer.
func (s *tTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.TTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TTemplate))
	})
	return ret, err
}

// TTemplates returns an object that can list and get TTemplates.
func (s *tTemplateLister) TTemplates(namespace string) TTemplateNamespaceLister {
	return tTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TTemplateNamespaceLister helps list and get TTemplates.
type TTemplateNamespaceLister interface {
	// List lists all TTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TTemplate, err error)
	// Get retrieves the TTemplate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TTemplate, error)
	TTemplateNamespaceListerExpansion
}

// tTemplateNamespaceLister implements the TTemplateNamespaceLister
// interface.
type tTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TTemplates in the indexer for a given namespace.
func (s tTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TTemplate))
	})
	return ret, err
}

// Get retrieves the TTemplate from the indexer for a given namespace and name.
func (s tTemplateNamespaceLister) Get(name string) (*v1alpha1.TTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ttemplate"), name)
	}
	return obj.(*v1alpha1.TTemplate), nil
}