package database

import (
	"time"

	"github.com/lib/pq"
)

// DApp model
type DApp struct {
	ID                uint           `gorm:"primary_key" json:"-" yaml:"id"`
	CreatedAt         time.Time      `json:"-" yaml:"-"`
	UpdatedAt         time.Time      `json:"-" yaml:"-"`
	DeletedAt         *time.Time     `json:"-" yaml:"-"`
	Name              string         `json:"name" yaml:"name"`
	ShortDescription  string         `json:"short_description" yaml:"short_description"`
	FullDescription   string         `json:"full_description" yaml:"full_description"`
	Version           string         `json:"version" yaml:"version"`
	License           string         `json:"license" yaml:"license"`
	WebSite           string         `json:"website" yaml:"web_site"`
	Slug              string         `json:"slug,omitempty" yaml:"slug"`
	AgoraReviewPostID uint           `json:"agora_review_post_id,omitempty" yaml:"agora_review_post_id"`
	AgoraQAPostID     uint           `json:"agora_qa_post_id,omitempty" yaml:"agora_qa_post_id"`
	Authors           pq.StringArray `gorm:"type:varchar(128)[]" json:"authors" yaml:"authors"`
	SocialLinks       pq.StringArray `gorm:"type:varchar(1024)[]" json:"social_links" yaml:"social_links"`
	Interfaces        pq.StringArray `gorm:"type:varchar(64)[]" json:"interfaces" yaml:"interfaces"`
	Categories        pq.StringArray `gorm:"type:varchar(32)[]" json:"categories" yaml:"categories"`
	Contracts         pq.StringArray `gorm:"type:varchar(36)[]" json:"contracts" yaml:"contracts"`
	Order             uint           `json:"-" yaml:"order"`
	Soon              bool           `gorm:"default:false" json:"soon" yaml:"soon"`

	Pictures  []Picture  `json:"pictures,omitempty" yaml:"pictures"`
	DexTokens []DexToken `json:"dex_tokens,omitempty" yaml:"dex_tokens"`
}

// Picture model
type Picture struct {
	ID     uint   `gorm:"primary_key,AUTO_INCREMENT" json:"-" yaml:"id"`
	Link   string `json:"link" yaml:"link"`
	Type   string `json:"type" yaml:"type"`
	DAppID uint   `json:"-" yaml:"dapp_id"`
}

// DexToken -
type DexToken struct {
	ID       uint   `gorm:"primary_key,AUTO_INCREMENT" json:"-" yaml:"id"`
	DAppID   uint   `yaml:"dapp_id"`
	Contract string `yaml:"contract"`
	TokenID  uint   `yaml:"token_id"`
}

// GetDApps -
func (d *db) GetDApps() (dapps []DApp, err error) {
	err = d.Preload("Pictures").Order(`order`).Find(&dapps).Error
	return
}

// GetDApp -
func (d *db) GetDApp(id uint) (dapp DApp, err error) {
	err = d.Preload("Pictures").Preload("DexTokens").Scopes(idScope(id)).First(&dapp).Error
	return
}

// GetDAppBySlug -
func (d *db) GetDAppBySlug(slug string) (dapp DApp, err error) {
	err = d.Preload("Pictures").Preload("DexTokens").Where("slug = ?", slug).First(&dapp).Error
	return
}

// CreateDapp -
func (d *db) CreateDapp(dapp *DApp) error {
	return d.Create(dapp).Error
}

// DeleteDapps -
func (d *db) DeleteDapps() error {
	if err := d.DropTableIfExists(&DexToken{}, &Picture{}, &DApp{}).Error; err != nil {
		return err
	}

	return d.CreateTable(&DexToken{}, &Picture{}, &DApp{}).Error
}