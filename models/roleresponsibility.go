package models

import "github.com/uadmin/uadmin"

type RoleResponsibility struct {
    uadmin.Model
    RoleID           uint
    Role             Role
    ResponsibilityID uint
    Responsibility   []Responsibility `gorm:"many2many:role_responsibility;"`
}

func (rr *RoleResponsibility) Save() {

    // Manually manage many-to-many relationships
    for _, responsibility := range rr.Responsibility {
        // Check if the RoleResponsibility already exists
        var existing RoleResponsibility
        err := uadmin.Get(&existing, "role_id = ? AND responsibility_id = ?", rr.RoleID, responsibility.ID)
        if err != nil || existing.ID == 0 {
            // If it does not exist, create it
            newRoleResponsibility := RoleResponsibility{
                RoleID:           rr.RoleID,
                ResponsibilityID: responsibility.ID,
            }
            uadmin.Save(&newRoleResponsibility)
        }
    }
}
