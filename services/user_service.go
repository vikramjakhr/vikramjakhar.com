package services

import (
	"github.com/choudhary92/vikramjakhar.com/util"
)

func (regData *util.RegistrationInfo) Register() {
	regData.Validate()
}