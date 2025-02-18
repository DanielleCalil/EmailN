package endpoints

import "EmailN/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}