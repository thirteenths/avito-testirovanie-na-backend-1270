package storage

type EmployeeStorage interface {
	CheckUserIsExistByUsername(username string) (bool, error)
	CheckUserIsExistById(userId string) (bool, error)
	CheckUserOrganization(username string, organizationId string) (bool, error)
	CheckUserTender(username string, tenderId string) (bool, error)
}
