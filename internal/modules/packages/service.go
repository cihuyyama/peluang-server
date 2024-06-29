package packages

import (
	"peluang-server/domain"
	"peluang-server/dto"

	"github.com/google/uuid"
)

type packageService struct {
	packageRepo domain.PackageRepository
}

func NewService(repo domain.PackageRepository) domain.PackageService {
	return &packageService{
		packageRepo: repo,
	}
}

// Delete implements domain.PackageService.
func (p *packageService) Delete(id string) error {
	var packageData *domain.Packages
	packageData, err := p.packageRepo.FindByID(id)
	if err != nil {
		return domain.ErrPackageNotFound
	}

	if packageData.List != nil {
		err := p.packageRepo.DeleteList(packageData.List)
		if err != nil {
			return err
		}
	}

	if packageData.Aditional != nil {
		err := p.packageRepo.DeleteAditional(packageData.Aditional)
		if err != nil {
			return err
		}
	}

	err = p.packageRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAditional implements domain.PackageService.
func (p *packageService) DeleteAditional(id uint) error {
	err := p.packageRepo.DeleteAditional([]domain.AditionalList{
		{
			ID: id,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteList implements domain.PackageService.
func (p *packageService) DeleteList(id uint) error {
	err := p.packageRepo.DeleteList([]domain.PackageList{
		{
			ID: id,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements domain.PackageService.
func (p *packageService) FindAll() ([]domain.Packages, error) {
	var packages []domain.Packages
	packages, err := p.packageRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return packages, nil
}

// FindByID implements domain.PackageService.
func (p *packageService) FindByID(id string) (*domain.Packages, error) {
	var packageData *domain.Packages
	packageData, err := p.packageRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return packageData, nil
}

// Insert implements domain.PackageService.
func (p *packageService) Insert(data *dto.PackageRequest, merchantID string) error {
	var packageData domain.Packages
	packageData.MerchantID = merchantID
	packageData.Name = data.Name
	packageData.Price = data.Price
	packageData.ID = uuid.New().String()

	err := p.packageRepo.Insert(&packageData)
	if err != nil {
		return err
	}

	if data.List != nil {
		var list []domain.PackageList
		for i := range data.List {
			list = append(list, domain.PackageList{
				PackageID: packageData.ID,
				Name:      data.List[i].Name,
			})
		}
		err := p.packageRepo.InsertLists(list)
		if err != nil {
			return err
		}
	}

	if data.Aditional != nil {
		var aditional []domain.AditionalList
		for i := range data.Aditional {
			aditional = append(aditional, domain.AditionalList{
				PackageID: packageData.ID,
				Name:      data.Aditional[i].Name,
				Amount:    data.Aditional[i].Amount,
			})
		}
		err := p.packageRepo.InsertAditionals(aditional)
		if err != nil {
			return err
		}
	}

	return nil
}

// InsertAditionals implements domain.PackageService.
func (p *packageService) InsertAditionals(data []dto.Aditional, packageID string) error {
	var aditional []domain.AditionalList
	for i := range data {
		aditional = append(aditional, domain.AditionalList{
			PackageID: packageID,
			Name:      data[i].Name,
			Amount:    data[i].Amount,
		})
	}

	err := p.packageRepo.InsertAditionals(aditional)
	if err != nil {
		return err
	}
	return nil
}

// InsertLists implements domain.PackageService.
func (p *packageService) InsertLists(data []dto.List, packageID string) error {
	var list []domain.PackageList
	for i := range data {
		list = append(list, domain.PackageList{
			PackageID: packageID,
			Name:      data[i].Name,
		})
	}

	err := p.packageRepo.InsertLists(list)
	if err != nil {
		return err
	}
	return nil
}

// Update implements domain.PackageService.
func (p *packageService) Update(id string, data *dto.PackageRequest) error {
	packageData, err := p.packageRepo.FindByID(id)
	if err != nil {
		return domain.ErrPackageNotFound
	}

	packageData.Name = data.Name
	packageData.Price = data.Price

	if err := p.packageRepo.Update(packageData); err != nil {
		return err
	}

	return nil
}

// UpdateAditional implements domain.PackageService.
func (p *packageService) UpdateAditional(data domain.AditionalList) error {
	err := p.packageRepo.UpdateAditional(data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateList implements domain.PackageService.
func (p *packageService) UpdateList(data domain.PackageList) error {
	err := p.packageRepo.UpdateList(data)
	if err != nil {
		return err
	}
	return nil
}
