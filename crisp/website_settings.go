// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteSettingsData mapping
type WebsiteSettingsData struct {
  Data  *WebsiteSettings  `json:"data,omitempty"`
}

// WebsiteSettings mapping
type WebsiteSettings struct {
  WebsiteID  *string                  `json:"website_id,omitempty"`
  Name       *string                  `json:"name,omitempty"`
  Domain     *string                  `json:"domain,omitempty"`
  Logo       *string                  `json:"logo,omitempty"`
  Contact    *WebsiteSettingsContact  `json:"contact,omitempty"`
  Chatbox    *WebsiteSettingsChatbox  `json:"chatbox,omitempty"`
}

// WebsiteSettingsContact mapping
type WebsiteSettingsContact struct {
  Email  *string  `json:"email,omitempty"`
  Phone  *string  `json:"phone,omitempty"`
}

// WebsiteSettingsChatbox mapping
type WebsiteSettingsChatbox struct {
  AvailabilityTooltip  *bool    `json:"availability_tooltip,omitempty"`
  HideOnAway           *bool    `json:"hide_on_away,omitempty"`
  PositionReverse      *bool    `json:"position_reverse,omitempty"`
  EmailVisitors        *bool    `json:"email_visitors,omitempty"`
  ColorTheme           *string  `json:"color_theme,omitempty"`
  TextTheme            *string  `json:"text_theme,omitempty"`
  WelcomeMessage       *string  `json:"welcome_message,omitempty"`
  Locale               *string  `json:"locale,omitempty"`
}


// GetWebsiteSettings resolves the current settings for a website.
// Reference: https://docs.crisp.im/api/v1/#website-website-settings-get
func (service *WebsiteService) GetWebsiteSettings(websiteID string) (*WebsiteSettings, *Response, error) {
  url := fmt.Sprintf("website/%s/settings", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(WebsiteSettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// UpdateWebsiteSettings updates the current settings for a website.
// Reference: https://docs.crisp.im/api/v1/#website-website-settings-patch
func (service *WebsiteService) UpdateWebsiteSettings(websiteID string, name string, domain string, logo string, contactEmail string, contactPhone string, chatboxAvailabilityTooltip bool, chatboxHideOnAway bool, chatboxPositionReverse bool, chatboxEmailVisitors bool, chatboxColorTheme string, chatboxTextTheme string, chatboxWelcomeMessage string, chatboxLocale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/settings", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteSettings{Name: &name, Domain: &domain, Logo: &logo, Contact: &WebsiteSettingsContact{Email: &contactEmail, Phone: &contactPhone}, Chatbox: &WebsiteSettingsChatbox{AvailabilityTooltip: &chatboxAvailabilityTooltip, HideOnAway: &chatboxHideOnAway, PositionReverse: &chatboxPositionReverse, EmailVisitors: &chatboxEmailVisitors, ColorTheme: &chatboxColorTheme, TextTheme: &chatboxTextTheme, WelcomeMessage: &chatboxWelcomeMessage, Locale: &chatboxLocale}})

  return service.client.Do(req, nil)
}