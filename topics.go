//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
)

// TopicsService handles communication with the topics related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/topics.html
type TopicsService struct {
	client *Client
}

// Topic represents a GitLab project topic.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/topics.html
type Topic struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	TotalProjectsCount uint64 `json:"total_projects_count"`
	AvatarURL          string `json:"avatar_url"`
}

func (t Topic) String() string {
	return Stringify(t)
}

// ListTopicsOptions represents the available ListTopics() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/topics.html#list-topics
type ListTopicsOptions struct {
	ListOptions
	Search *string `url:"search,omitempty" json:"search,omitempty"`
}

// ListTopics returns a list of project topics in the GitLab instance ordered
// by number of associated projects.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/topics.html#list-topics
func (s *TopicsService) ListTopics(opt *ListTopicsOptions, options ...RequestOptionFunc) ([]*Topic, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "topics", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var t []*Topic
	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// GetTopic gets a project topic by ID.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/topics.html#get-a-topic
func (s *TopicsService) GetTopic(topic int, options ...RequestOptionFunc) (*Topic, *Response, error) {
	u := fmt.Sprintf("topics/%d", topic)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	t := new(Topic)
	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// CreateTopicOptions represents the available CreateTopic() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/topics.html#create-a-project-topic
type CreateTopicOptions struct {
	Name        *string `url:"name,omitempty" json:"name,omitempty"`
	Description *string `url:"description,omitempty" json:"description,omitempty"`
	//	Avatar      *string `url:"avatar,omitempty" json:"avatar,omitempty"`
}

// CreateTopic creates a new project topic.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/topics.html#create-a-project-topic
func (s *TopicsService) CreateTopic(opt *CreateTopicOptions, options ...RequestOptionFunc) (*Topic, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "topics", opt, options)
	if err != nil {
		return nil, nil, err
	}

	t := new(Topic)
	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// UpdateTopicOptions represents the available UpdateTopic() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/topics.html#update-a-project-topic
type UpdateTopicOptions struct {
	Name        *string `url:"name,omitempty" json:"name,omitempty"`
	Description *string `url:"description,omitempty" json:"description,omitempty"`
	//	Avatar      *string `url:"avatar,omitempty" json:"avatar,omitempty"`
}

// UpdateTopic updates a project topic. Only available to administrators.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/topics.html#update-a-project-topic
func (s *TopicsService) UpdateTopic(topic int, opt *UpdateTopicOptions, options ...RequestOptionFunc) (*Topic, *Response, error) {
	u := fmt.Sprintf("topics/%d", topic)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	t := new(Topic)
	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// DeleteTopic deletes a project topic. Only available to administrators.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/topics.html#delete-a-project-topic
func (s *TopicsService) DeleteTopic(topic int, ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("topics/%d", topic)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
