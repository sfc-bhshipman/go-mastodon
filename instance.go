package mastodon

import (
	"context"
	"net/http"
)

// Instance holds information for a mastodon instance.
type Instance struct {
	URI            string            `json:"uri"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	EMail          string            `json:"email"`
	Version        string            `json:"version,omitempty"`
	Thumbnail      string            `json:"thumbnail,omitempty"`
	URLs           map[string]string `json:"urls,omitempty"`
	Stats          *InstanceStats    `json:"stats,omitempty"`
	Languages      []string          `json:"languages"`
	ContactAccount *Account          `json:"contact_account"`
	Configuration  *InstanceConfig   `json:"configuration"`
}

type InstanceConfigMap map[string]interface{}

// InstanceConfig holds configuration accessible for clients.
type InstanceConfig struct {
	Accounts         *InstanceConfigMap     `json:"accounts"`
	Statuses         *InstanceConfigMap     `json:"statuses"`
	MediaAttachments map[string]interface{} `json:"media_attachments"`
	Polls            *InstanceConfigMap     `json:"polls"`
}

// InstanceStats holds information for mastodon instance stats.
type InstanceStats struct {
	UserCount   int64 `json:"user_count"`
	StatusCount int64 `json:"status_count"`
	DomainCount int64 `json:"domain_count"`
}

// Instance holds information for a mastodon instance
type InstanceV2 struct {
	Domain      string `json:"domain"`
	Title       string `json:"title"`
	Version     string `json:"version"`
	SourceURL   string `json:"source_url"`
	Description string `json:"description"`
	Usage       struct {
		Users struct {
			ActiveMonth int `json:"active_month"`
		} `json:"users"`
	} `json:"usage"`
	Thumbnail struct {
		URL      string      `json:"url"`
		Blurhash interface{} `json:"blurhash"`
		Versions struct {
			One_X interface{} `json:"@1x"`
			Two_X interface{} `json:"@2x"`
		} `json:"versions"`
	} `json:"thumbnail"`
	Languages     []string `json:"languages"`
	Configuration struct {
		Urls struct {
			Streaming string `json:"streaming"`
		} `json:"urls"`
		Accounts struct {
			MaxFeaturedTags int `json:"max_featured_tags"`
		} `json:"accounts"`
		Statuses struct {
			MaxCharacters            int `json:"max_characters"`
			MaxMediaAttachments      int `json:"max_media_attachments"`
			CharactersReservedPerURL int `json:"characters_reserved_per_url"`
		} `json:"statuses"`
		MediaAttachments struct {
			SupportedMimeTypes  []string `json:"supported_mime_types"`
			ImageSizeLimit      int      `json:"image_size_limit"`
			ImageMatrixLimit    int      `json:"image_matrix_limit"`
			VideoSizeLimit      int      `json:"video_size_limit"`
			VideoFrameRateLimit int      `json:"video_frame_rate_limit"`
			VideoMatrixLimit    int      `json:"video_matrix_limit"`
		} `json:"media_attachments"`
		Polls struct {
			MaxOptions             int `json:"max_options"`
			MaxCharactersPerOption int `json:"max_characters_per_option"`
			MinExpiration          int `json:"min_expiration"`
			MaxExpiration          int `json:"max_expiration"`
		} `json:"polls"`
		Translation struct {
			Enabled bool `json:"enabled"`
		} `json:"translation"`
	} `json:"configuration"`
	Registrations struct {
		Enabled          bool        `json:"enabled"`
		ApprovalRequired bool        `json:"approval_required"`
		Message          interface{} `json:"message"`
	} `json:"registrations"`
	Contact struct {
		Email   string `json:"email"`
		Account *Account
	} `json:"contact"`
	Rules []Rule `json:"rules"`
}

type Rule struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// GetInstance returns Instance.
func (c *Client) GetInstance(ctx context.Context) (*Instance, error) {
	var instance Instance
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/instance", nil, &instance, nil)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func (c *Client) GetInstanceV2(ctx context.Context) (*InstanceV2, error) {
	var instance InstanceV2
	err := c.doAPI(ctx, http.MethodGet, "/api/v2/instance", nil, &instance, nil)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

// GetConfig returns InstanceConfig.
func (c *Instance) GetConfig() *InstanceConfig {
	return c.Configuration
}

// WeeklyActivity holds information for mastodon weekly activity.
type WeeklyActivity struct {
	Week          Unixtime `json:"week"`
	Statuses      int64    `json:"statuses,string"`
	Logins        int64    `json:"logins,string"`
	Registrations int64    `json:"registrations,string"`
}

// GetInstanceActivity returns instance activity.
func (c *Client) GetInstanceActivity(ctx context.Context) ([]*WeeklyActivity, error) {
	var activity []*WeeklyActivity
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/instance/activity", nil, &activity, nil)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

// GetInstancePeers returns instance peers.
func (c *Client) GetInstancePeers(ctx context.Context) ([]string, error) {
	var peers []string
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/instance/peers", nil, &peers, nil)
	if err != nil {
		return nil, err
	}
	return peers, nil
}
