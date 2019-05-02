/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// ClusterConsole represents the values of the 'cluster_console' type.
//
// Information about the console of a cluster.
type ClusterConsole struct {
	url *string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterConsole) Empty() bool {
	return o == nil || (o.url == nil &&
		true)
}

// URL returns the value of the 'URL' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The URL of the console of the cluster.
func (o *ClusterConsole) URL() string {
	if o != nil && o.url != nil {
		return *o.url
	}
	return ""
}

// GetURL returns the value of the 'URL' attribute and
// a flag indicating if the attribute has a value.
//
// The URL of the console of the cluster.
func (o *ClusterConsole) GetURL() (value string, ok bool) {
	ok = o != nil && o.url != nil
	if ok {
		value = *o.url
	}
	return
}

// ClusterConsoleList is a list of values of the 'cluster_console' type.
type ClusterConsoleList struct {
	items []*ClusterConsole
}

// Len returns the length of the list.
func (l *ClusterConsoleList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterConsoleList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterConsoleList) Slice() []*ClusterConsole {
	var slice []*ClusterConsole
	if l == nil {
		slice = make([]*ClusterConsole, 0)
	} else {
		slice = make([]*ClusterConsole, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterConsoleList) Each(f func(item *ClusterConsole) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterConsoleList) Range(f func(index int, item *ClusterConsole) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
