package slack

import (
	"encoding/json"
	"errors"
)

type UserGroup struct {
	ID          string `json:"id"`
	TeamID      string `json:"team_id"`
	IsUsergroup bool   `json:"is_usergroup"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Handle      string `json:"handle"`
	IsExternal  bool   `json:"is_external"`
	DateCreate  int    `json:"date_create"`
	DateUpdate  int    `json:"date_update"`
	DateDelete  int    `json:"date_delete"`
	AutoType    string `json:"auto_type"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
	Prefs       struct {
		Channels []string `json:"channels"`
		Groups   []string `json:"groups"`
	} `json:"prefs,omitempty"`
	Users     []string `json:"users,omitempty"`
	UserCount int      `json:"user_count,omitempty"`
}

// API usergroups.list: Lists all user groups in a Slack team.
func (sl *Slack) UserGroupsList() ([]*UserGroup, error) {
	uv := sl.urlValues()
	body, err := sl.GetRequest(userGroupsListApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(UserGroupsListAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.UserGroups()
}

// response type of `usergroups.list` api
type UserGroupsListAPIResponse struct {
	BaseAPIResponse
	RawGroups json.RawMessage `json:"usergroups"`
}

// matching func
func (res *UserGroupsListAPIResponse) UserGroups() ([]*UserGroup, error) {
	var usergroups []*UserGroup
	err := json.Unmarshal(res.RawGroups, &usergroups)
	if err != nil {
		return nil, err
	}
	return usergroups, nil
}
