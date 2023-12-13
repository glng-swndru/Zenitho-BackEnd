// Package campaign menyediakan fungsi-fungsi untuk memformat data campaign.
package campaign

// CampaignFormatter adalah struktur data yang digunakan untuk memformat data campaign sebelum dikirim sebagai respons JSON.
type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

// FormatCampaign mengonversi data campaign menjadi CampaignFormatter.
func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		ImageURL:         "",
	}

	if len(campaign.CampaignImage) > 0 {
		formatter.ImageURL = campaign.CampaignImage[0].FileName
	}

	return formatter
}

// FormatCampaigns mengonversi daftar campaign menjadi daftar CampaignFormatter.
func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	var campaignsFormatter []CampaignFormatter

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
