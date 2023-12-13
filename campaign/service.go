package campaign

// Service adalah interface yang mendefinisikan fungsi yang harus ada di service campaign.
type Service interface {
	GetCampaigns(userID int) ([]Campaign, error) // Fungsi buat dapetin campaign berdasarkan userID.
}

// service adalah struct yang implementasi dari Service.
type service struct {
	repository Repository // Ini tempat nyimpen data, kaya database gitu.
}

// NewService adalah fungsi pembuat service baru.
func NewService(repository Repository) *service {
	return &service{repository} // Balikin instance service yang baru dengan repository.
}

// GetCampaigns adalah method dari service buat dapetin campaign.
func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	// Cek dulu, kalo userID nya nggak 0, berarti kita cari berdasarkan userID.
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err // Kalo ada error, langsung balikin errornya.
		}
		return campaigns, nil // Kalo nggak ada error, balikin campaignnya.
	}

	// Kalo userID nya 0, berarti kita cari semua campaign yang ada.
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err // Sama, kalo ada error, balikin errornya.
	}
	return campaigns, nil // Kalo nggak ada error, balikin semua campaign.
}
