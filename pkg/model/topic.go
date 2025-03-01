package model

import (
	"github.com/TUM-Dev/meldeplattform/pkg/i18n"
)

type Admin struct {
	ID     uint `gorm:"primarykey"`
	UserID string
}

type Topic struct {
	ID       uint              `gorm:"primarykey"`
	Name     i18n.Translatable `yaml:"name" gorm:"embedded;embeddedPrefix:name_"`
	Summary  i18n.Translatable `yaml:"summary" gorm:"embedded;embeddedPrefix:summary_"`
	Fields   []Field           `yaml:"fields"`
	Contacts struct {
		Email   *EmailConfig   `yaml:"email"`
		Matrix  *MatrixConfig  `yaml:"matrix"`
		Webhook *WebhookConfig `yaml:"webhook"`
	} `yaml:"contacts" gorm:"-"`

	Admins []Admin `yaml:"admins" gorm:"many2many:topic_admins;"`

	Email string
}

type Field struct {
	ID          uint `gorm:"primarykey"`
	TopicID     uint
	Name        i18n.Translatable `yaml:"name" gorm:"embedded;embeddedPrefix:name_"`
	Type        string            `yaml:"type"` // e.g. file, text, email, textarea,
	Required    bool              `yaml:"required"`
	Description i18n.Translatable `yaml:"description" gorm:"embedded;embeddedPrefix:description_"`

	// For select inputs:
	Choices *[]string `yaml:"choices" gorm:"-"`
}

func (t *Topic) IsAdmin(userid string) bool {
	if t == nil || userid == "" {
		return false
	}
	for _, admin := range t.Admins {
		if admin.UserID == userid {
			return true
		}
	}
	return false
}
