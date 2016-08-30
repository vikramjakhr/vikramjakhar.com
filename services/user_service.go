package services

import (
	"github.com/vikramjakhr/vikramjakhar.com/util"
)

func (regData *util.RegistrationInfo) Register() {
	regData.Validate()
}