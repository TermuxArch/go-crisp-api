// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// ConversationListData mapping
type ConversationListData struct {
  Data  *[]Conversation  `json:"data,omitempty"`
}

// ConversationData mapping
type ConversationData struct {
  Data  *Conversation  `json:"data,omitempty"`
}

// Conversation mapping
type Conversation struct {
  SessionID     *string                 `json:"session_id,omitempty"`
  WebsiteID     *string                 `json:"website_id,omitempty"`
  State         *string                 `json:"state,omitempty"`
  Status        *uint8                  `json:"status,omitempty"`
  IsBlocked     *bool                   `json:"is_blocked,omitempty"`
  Availability  *string                 `json:"availability,omitempty"`
  LastMessage   *string                 `json:"last_message,omitempty"`
  CreatedAt     *uint                   `json:"created_at,omitempty"`
  UpdatedAt     *uint                   `json:"updated_at,omitempty"`
  Unread        *ConversationUnread     `json:"unread,omitempty"`
  Meta          *ConversationMeta       `json:"meta,omitempty"`
  Messages      *[]ConversationMessage  `json:"message,omitempty"`
}

// ConversationUnread mapping
type ConversationUnread struct {
  Operator  *uint  `json:"operator,omitempty"`
  Visitor   *uint  `json:"visitor,omitempty"`
}

// ConversationMeta mapping
type ConversationMeta struct {
  Nickname              *string                                      `json:"nickname,omitempty"`
  Email                 *string                                      `json:"email,omitempty"`
  Avatar                *string                                      `json:"avatar,omitempty"`
  BrowsingInformations  *ConversationMetaBrowsingInformations        `json:"browsing_informations,omitempty"`
  Tags                  *[]string                                    `json:"tags,omitempty"`
  Pages                 *[]ConversationMetaBrowsingInformationsPage  `json:"pages,omitempty"`
}

// ConversationMessage mapping
type ConversationMessage struct {
  SessionID    *string                        `json:"session_id,omitempty"`
  WebsiteID    *string                        `json:"website_id,omitempty"`
  Type         *string                        `json:"type,omitempty"`
  From         *string                        `json:"from,omitempty"`
  Origin       *string                        `json:"origin,omitempty"`
  Content      *string                        `json:"content,omitempty"`
  Preview      *[]ConversationMessagePreview  `json:"preview,omitempty"`
  Read         *string                        `json:"read,omitempty"`
  Delivered    *string                        `json:"delivered,omitempty"`
  Fingerprint  *int                           `json:"fingerprint,omitempty"`
  Timestamp    *uint                          `json:"timestamp,omitempty"`
  User         *ConversationMessageUser       `json:"user,omitempty"`
}

// ConversationMessagePreview mapping
type ConversationMessagePreview struct {
  URL      *string                                 `json:"url,omitempty"`
  Website  *string                                 `json:"website,omitempty"`
  Title    *string                                 `json:"title,omitempty"`
  Preview  *ConversationMessagePreviewInformation  `json:"preview,omitempty"`
  Stamped  *bool                                   `json:"stamped,omitempty"`
}

// ConversationMessagePreviewInformation mapping
type ConversationMessagePreviewInformation struct {
  Excerpt  *string  `json:"excerpt,omitempty"`
  Image    *string  `json:"image,omitempty"`
}

// ConversationMessageUser mapping
type ConversationMessageUser struct {
  UserID    *string  `json:"user_id,omitempty"`
  Nickname  *string  `json:"nickname,omitempty"`
}

// ConversationMetaBrowsingInformations mapping
type ConversationMetaBrowsingInformations struct {
  Geolocation  *ConversationMetaBrowsingInformationsGeolocation  `json:"geolocation,omitempty"`
  System       *ConversationMetaBrowsingInformationsSystem       `json:"system,omitempty"`
}

// ConversationMetaBrowsingInformationsGeolocation mapping
type ConversationMetaBrowsingInformationsGeolocation struct {
  Metro    *int        `json:"metro,omitempty"`
  LL       *[]float32  `json:"ll,omitempty"`
  City     *string     `json:"city,omitempty"`
  Region   *string     `json:"region,omitempty"`
  Country  *string     `json:"country,omitempty"`
  Range    *[]uint     `json:"range,omitempty"`
}

// ConversationMetaBrowsingInformationsSystem mapping
type ConversationMetaBrowsingInformationsSystem struct {
  OS       *ConversationMetaBrowsingInformationsSystemOS       `json:"os,omitempty"`
  Engine   *ConversationMetaBrowsingInformationsSystemEngine   `json:"engine,omitempty"`
  Browser  *ConversationMetaBrowsingInformationsSystemBrowser  `json:"browser,omitempty"`
  UA       *string                                             `json:"ua,omitempty"`
}

// ConversationMetaBrowsingInformationsSystemOS mapping
type ConversationMetaBrowsingInformationsSystemOS struct {
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// ConversationMetaBrowsingInformationsSystemEngine mapping
type ConversationMetaBrowsingInformationsSystemEngine struct {
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// ConversationMetaBrowsingInformationsSystemBrowser mapping
type ConversationMetaBrowsingInformationsSystemBrowser struct {
  Major    *string  `json:"major,omitempty"`
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// ConversationMetaBrowsingInformationsPage mapping
type ConversationMetaBrowsingInformationsPage struct {
  PageTitle  *string  `json:"page_title,omitempty"`
  PageURL    *string  `json:"page_url,omitempty"`
  Timestamp  *uint    `json:"timestamp,omitempty"`
}

// ConversationNewData mapping
type ConversationNewData struct {
  Data  *ConversationNew  `json:"data,omitempty"`
}

// ConversationNew mapping
type ConversationNew struct {
  SessionID  *string  `json:"session_id,omitempty"`
}

// ConversationTextMessageNew mapping
type ConversationTextMessageNew struct {
  Type         string                          `json:"type,omitempty"`
  From         string                          `json:"from,omitempty"`
  Origin       string                          `json:"origin,omitempty"`
  Content      string                          `json:"content,omitempty"`
  Fingerprint  int                             `json:"fingerprint,omitempty"`
  User         ConversationTextMessageNewUser  `json:"user,omitempty"`
}

// ConversationTextMessageNewUser mapping
type ConversationTextMessageNewUser struct {
  Nickname  string  `json:"nickname,omitempty"`
  Avatar    string  `json:"avatar,omitempty"`
}

// ConversationFileMessageNew mapping
type ConversationFileMessageNew struct {
  Type         string                             `json:"type,omitempty"`
  From         string                             `json:"from,omitempty"`
  Origin       string                             `json:"origin,omitempty"`
  Content      ConversationFileMessageNewContent  `json:"content,omitempty"`
  Fingerprint  int                                `json:"fingerprint,omitempty"`
}

// ConversationFileMessageNewContent mapping
type ConversationFileMessageNewContent struct {
  Name  string  `json:"name,omitempty"`
  URL   string  `json:"url,omitempty"`
  Type  string  `json:"type,omitempty"`
}

// ConversationComposeMessageNew mapping
type ConversationComposeMessageNew struct {
  Type     string  `json:"type,omitempty"`
  From     string  `json:"from,omitempty"`
  Excerpt  string  `json:"excerpt,omitempty"`
}

// ConversationReadMessageMark mapping
type ConversationReadMessageMark struct {
  Type          string  `json:"type,omitempty"`
  Origin        string  `json:"origin,omitempty"`
  Fingerprints  []int   `json:"fingerprints,omitempty"`
}

// ConversationMetaUpdate mapping
type ConversationMetaUpdate struct {
  Nickname  string    `json:"nickname,omitempty"`
  Email     string    `json:"email,omitempty"`
  Tags      []string  `json:"tags,omitempty"`
}

// ConversationStateUpdate mapping
type ConversationStateUpdate struct {
  State  *string  `json:"state,omitempty"`
}

// ConversationBlockUpdate mapping
type ConversationBlockUpdate struct {
  Blocked  *bool  `json:"blocked,omitempty"`
}


// String returns the string representation of Conversation
func (instance Conversation) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationNew
func (instance ConversationNew) String() string {
  return Stringify(instance)
}


// SearchConversations searches conversations for website.
func (service *WebsiteService) SearchConversations(websiteID string, pageNumber uint, searchQuery string, searchType string) (*[]Conversation, *Response, error) {
  var resourceURL string

  if searchQuery != "" && searchType != "" {
    resourceURL = fmt.Sprintf("website/%s/conversations/%d?search_query=%s&search_type=%s", websiteID, pageNumber, url.QueryEscape(searchQuery), url.QueryEscape(searchType))
  } else {
    resourceURL = fmt.Sprintf("website/%s/conversations/%d", websiteID, pageNumber)
  }

  req, _ := service.client.NewRequest("GET", resourceURL, nil)

  conversations := new(ConversationListData)
  resp, err := service.client.Do(req, conversations)
  if err != nil {
    return nil, resp, err
  }

  return conversations.Data, resp, err
}


// ListConversations lists conversations for website.
func (service *WebsiteService) ListConversations(websiteID string, pageNumber uint) (*[]Conversation, *Response, error) {
  return service.SearchConversations(websiteID, pageNumber, "", "")
}


// CreateNewConversation creates a new conversation.
func (service *WebsiteService) CreateNewConversation(websiteID string) (*ConversationNew, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation", websiteID)
  req, _ := service.client.NewRequest("POST", url, nil)

  conversation := new(ConversationNewData)
  resp, err := service.client.Do(req, conversation)
  if err != nil {
    return nil, resp, err
  }

  return conversation.Data, resp, err
}


// CheckConversationExists checks if given conversation session identifier exists.
func (service *WebsiteService) CheckConversationExists(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// GetConversation resolves conversation information and messages.
func (service *WebsiteService) GetConversation(websiteID string, sessionID string) (*Conversation, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  conversation := new(ConversationData)
  resp, err := service.client.Do(req, conversation)
  if err != nil {
    return nil, resp, err
  }

  return conversation.Data, resp, err
}


// RemoveConversation removes a conversation in website.
func (service *WebsiteService) RemoveConversation(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// InitiateConversationWithExistingSession initiates a conversation from an existing session.
func (service *WebsiteService) InitiateConversationWithExistingSession(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/initiate", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// SendTextMessageInConversation sends a message in an existing conversation.
func (service *WebsiteService) SendTextMessageInConversation(websiteID string, sessionID string, message ConversationTextMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// SendFileMessageInConversation sends a message in an existing conversation.
func (service *WebsiteService) SendFileMessageInConversation(websiteID string, sessionID string, message ConversationFileMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// ComposeMessageInConversation starts or stop composing a message in an existing conversation.
func (service *WebsiteService) ComposeMessageInConversation(websiteID string, sessionID string, compose ConversationComposeMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/compose", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, compose)

  return service.client.Do(req, nil)
}


// MarkMessagesReadInConversation marks messages as read in conversation. Either using given message fingerprints, or all messages.
func (service *WebsiteService) MarkMessagesReadInConversation(websiteID string, sessionID string, compose ConversationReadMessageMark) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/read", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, compose)

  return service.client.Do(req, nil)
}


// UpdateConversationMetas updates conversation meta information.
func (service *WebsiteService) UpdateConversationMetas(websiteID string, sessionID string, metas ConversationMetaUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/meta", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, metas)

  return service.client.Do(req, nil)
}


// ChangeConversationState updates conversation state.
func (service *WebsiteService) ChangeConversationState(websiteID string, sessionID string, state string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/state", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationStateUpdate{&state})

  return service.client.Do(req, nil)
}


// BlockIncomingMessagesForConversation blocks further incoming messages from a conversation.
func (service *WebsiteService) BlockIncomingMessagesForConversation(websiteID string, sessionID string, blocked bool) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/block", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationBlockUpdate{&blocked})

  return service.client.Do(req, nil)
}
