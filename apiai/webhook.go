//    Copyright 2017 drillbits
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package apiai

import (
	"encoding/json"
	"fmt"
)

type jsonObject map[string]interface{}

// WebhookRequest represents an HTTP request to be send by api.ai.
// See https://docs.api.ai/docs/webhook#section-format-of-request-to-the-service
type WebhookRequest struct {
	QueryResponse
	OriginalRequest *OriginalRequest `json:"originalRequest"`
}

// Original returns a request data of the messaging platform.
func (r *WebhookRequest) Original() (interface{}, error) {
	b, err := json.Marshal(r.OriginalRequest.Data)
	if err != nil {
		return nil, err
	}

	switch r.OriginalRequest.Source {
	case "slack":
		v := &SlackRequest{}
		err := json.Unmarshal(b, v)
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, fmt.Errorf("unknown source: %s", r.OriginalRequest.Source)
	}
}

// OriginalRequest https://docs.api.ai/docs/query#query-parameters-and-json-fields.
type OriginalRequest struct {
	Source string     `json:"source"`
	Data   jsonObject `json:"data"`
}

// SlackRequest represents an HTTP request to be send by slack.
// See https://api.slack.com/events-api
type SlackRequest struct {
	Token       string      `json:"token"`
	TeamID      string      `json:"team_id"`
	APIAppID    string      `json:"api_app_id"`
	Event       *SlackEvent `json:"event"`
	Type        string      `json:"type"`
	AuthedUsers []string    `json:"authed_users"`
	EventID     string      `json:"event_id"`
	EventTime   int64       `json:"event_time"`
}

// SlackEvent represents an event of slack.
// See https://api.slack.com/events-api.
type SlackEvent struct {
	Type           string `json:"type"`
	EventTimestamp string `json:"event_ts"`
	User           string `json:"user"`
	Timestamp      string `json:"ts"`
	Channel        string `json:"channel,omitempty"`
	Text           string `json:"text,omitempty"`
}

// WebhookResponse represents an HTTP response to api.ai.
// See https://docs.api.ai/docs/webhook#section-format-of-response-from-the-service.
type WebhookResponse struct {
	Speech        string        `json:"speech"`
	DisplayText   string        `json:"displayText"`
	Data          jsonObject    `json:"data"`
	ContextOut    []*Context    `json:"contextOut"`
	Source        string        `json:"source"`
	FollowupEvent FollowupEvent `json:"followupEvent"`
}

// FollowupEvent is parameter to invoke events.
// See https://docs.api.ai/docs/concept-events#invoking-event-from-webhook
type FollowupEvent struct {
	Name string     `json:"name"`
	Data jsonObject `json:"data"`
}
