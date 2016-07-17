// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteBillingData mapping
type WebsiteBillingData struct {
  Data  *WebsiteBilling  `json:"data,omitempty"`
}

// WebsiteBilling mapping
type WebsiteBilling struct {
  CardID             *string                   `json:"card_id,omitempty"`
  NameOnCard         *string                   `json:"name_on_card,omitempty"`
  Address            *string                   `json:"address,omitempty"`
  CardNumberPreview  *string                   `json:"card_number_preview,omitempty"`
  CardProvider       *string                   `json:"card_provider,omitempty"`
  ExpirationDate     *string                   `json:"expiration_date,omitempty"`
  AddedDate          *string                   `json:"added_date,omitempty"`
  IsExpired          *bool                     `json:"is_expired,omitempty"`
  LinkedWebsites     *[]WebsiteBillingWebsite  `json:"linked_websites,omitempty"`
}

// WebsiteBillingWebsite mapping
type WebsiteBillingWebsite struct {
  WebsiteID  *string  `json:"website_id,omitempty"`
  Name       *string  `json:"name,omitempty"`
  Domain     *string  `json:"domain,omitempty"`
}

// WebsiteBillingLink mapping
type WebsiteBillingLink struct {
  CardID  *string  `json:"card_id,omitempty"`
}


// GetWebsiteBilling resolves website billing information (payment method linked to website).
// Reference: https://docs.crisp.im/api/v1/#website-website-billing-get
func (service *WebsiteService) GetWebsiteBilling(websiteID string) (*WebsiteBilling, *Response, error) {
  url := fmt.Sprintf("website/%s/billing", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  billing := new(WebsiteBillingData)
  resp, err := service.client.Do(req, billing)
  if err != nil {
    return nil, resp, err
  }

  return billing.Data, resp, err
}


// UpdateWebsiteBilling updates website billing information (payment method linked to website).
// Reference: https://docs.crisp.im/api/v1/#website-website-billing-patch
func (service *WebsiteService) UpdateWebsiteBilling(websiteID string, cardID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/billing", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteBillingLink{CardID: &cardID})

  return service.client.Do(req, nil)
}


// UnlinkWebsiteBilling unlink website billing information (payment method linked to website).
// Reference: https://docs.crisp.im/api/v1/#website-website-billing-delete
func (service *WebsiteService) UnlinkWebsiteBilling(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/billing", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}