package response

import (
	"github.com/jacky-htg/go-services/models"
)

//UserResponse : format json response for user
type AccessResponse struct {
	ID       uint32 `json:"id"`
	ParentID uint32 `json:"parent_id,omitempty"`
	Name     string `json:"name"`
}

//Transform from Access model to Access response
func (u *AccessResponse) Transform(access *models.Access) {
	u.ID = access.ID
	u.ParentID = uint32(access.ParentID.Int64)
	u.Name = access.Name
}
