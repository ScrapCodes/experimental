/*
Copyright 2020 The Knative Authors

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
	v1alpha1 "github.com/tektoncd/experimental/pipeline-in-pod/pkg/apis/colocatedpipelinerun/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ColocatedPipelineRunLister helps list ColocatedPipelineRuns.
// All objects returned here must be treated as read-only.
type ColocatedPipelineRunLister interface {
	// List lists all ColocatedPipelineRuns in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ColocatedPipelineRun, err error)
	// ColocatedPipelineRuns returns an object that can list and get ColocatedPipelineRuns.
	ColocatedPipelineRuns(namespace string) ColocatedPipelineRunNamespaceLister
	ColocatedPipelineRunListerExpansion
}

// colocatedPipelineRunLister implements the ColocatedPipelineRunLister interface.
type colocatedPipelineRunLister struct {
	indexer cache.Indexer
}

// NewColocatedPipelineRunLister returns a new ColocatedPipelineRunLister.
func NewColocatedPipelineRunLister(indexer cache.Indexer) ColocatedPipelineRunLister {
	return &colocatedPipelineRunLister{indexer: indexer}
}

// List lists all ColocatedPipelineRuns in the indexer.
func (s *colocatedPipelineRunLister) List(selector labels.Selector) (ret []*v1alpha1.ColocatedPipelineRun, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ColocatedPipelineRun))
	})
	return ret, err
}

// ColocatedPipelineRuns returns an object that can list and get ColocatedPipelineRuns.
func (s *colocatedPipelineRunLister) ColocatedPipelineRuns(namespace string) ColocatedPipelineRunNamespaceLister {
	return colocatedPipelineRunNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ColocatedPipelineRunNamespaceLister helps list and get ColocatedPipelineRuns.
// All objects returned here must be treated as read-only.
type ColocatedPipelineRunNamespaceLister interface {
	// List lists all ColocatedPipelineRuns in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ColocatedPipelineRun, err error)
	// Get retrieves the ColocatedPipelineRun from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ColocatedPipelineRun, error)
	ColocatedPipelineRunNamespaceListerExpansion
}

// colocatedPipelineRunNamespaceLister implements the ColocatedPipelineRunNamespaceLister
// interface.
type colocatedPipelineRunNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ColocatedPipelineRuns in the indexer for a given namespace.
func (s colocatedPipelineRunNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ColocatedPipelineRun, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ColocatedPipelineRun))
	})
	return ret, err
}

// Get retrieves the ColocatedPipelineRun from the indexer for a given namespace and name.
func (s colocatedPipelineRunNamespaceLister) Get(name string) (*v1alpha1.ColocatedPipelineRun, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("colocatedpipelinerun"), name)
	}
	return obj.(*v1alpha1.ColocatedPipelineRun), nil
}