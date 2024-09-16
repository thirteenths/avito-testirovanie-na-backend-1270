package consts

import "errors"

var UserIsNotExistError = errors.New("user is not exist")
var UserHasNoRights = errors.New("user has no rights")
var TenderIsNotExistError = errors.New("tender is not exist")
var BidIsNotExistError = errors.New("bid is not exist")
var VersionIsNotExistError = errors.New("version is not exist")
var OrganizationIsNotExistError = errors.New("organization is not exist")
