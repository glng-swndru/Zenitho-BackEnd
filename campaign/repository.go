package campaign

import "gorm.io/gorm"

// Repository adalah interface untuk fungsi-fungsi database yang berkaitan dengan Campaign.
type Repository interface {
	FindAll() ([]Campaign, error)                // Fungsi untuk dapetin semua campaign.
	FindByUserID(userID int) ([]Campaign, error) // Fungsi untuk dapetin campaign berdasarkan ID user.
}

// repository adalah implementasi dari Repository, pakai GORM.
type repository struct {
	db *gorm.DB // db adalah pointer ke GORM DB, buat interaksi dengan database.
}

// NewRepository adalah fungsi pembuat repository baru.
func NewRepository(db *gorm.DB) *repository {
	return &repository{db} // Balikin struct repository baru dengan DB yang sudah di-set.
}

// FindAll adalah method dari repository untuk dapetin semua campaign.
func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign // Siapin slice untuk tampung data campaign.

	// Query ke database, preload CampaignImages dengan kondisi is_primary = 1.
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err // Kalo ada error, balikin errornya.
	}
	return campaigns, nil // Kalo sukses, balikin list campaign.
}

// FindByUserID adalah method dari repository untuk dapetin campaign berdasarkan ID user.
func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign // Siapin slice untuk tampung data campaign.

	// Query ke database, cari berdasarkan user_id dan preload CampaignImages.
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err // Kalo ada error, balikin errornya.
	}
	return campaigns, nil // Kalo sukses, balikin list campaign sesuai user ID.
}
