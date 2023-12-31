package usecase

import (
	"final-project-enigma-clean/exception"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/repository"
	"fmt"
)

type StaffUseCase interface {
	CreateNew(payload model.Staff) error
	FindByName(name string) ([]model.Staff, error)
	FindById(nik_staff string) (model.Staff, error)
	FindByAll() ([]model.Staff, error)
	Update(payload model.Staff) error
	Delete(nik_staff string) error
	Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error)
}

type staffUseCase struct {
	repo repository.StaffRepository
}

// FindById implements StaffUseCase.
func (s *staffUseCase) FindById(nik_staff string) (model.Staff, error) {
	staff, err := s.repo.FindById(nik_staff)
	if err != nil {
		return model.Staff{}, exception.BadRequestErr("staff not found")
	}
	return staff, nil

}

// CreateNew implements StaffUseCase.
func (s *staffUseCase) CreateNew(payload model.Staff) error {
	if payload.Nik_Staff == "" {
		return exception.BadRequestErr("nik staff cannot Empty")
	}
	if payload.Name == "" {
		return exception.BadRequestErr("name cannot Empty")
	}
	if len(payload.Phone_number) < 10 || len(payload.Phone_number) > 15 {
		return exception.BadRequestErr("phone number must be between 10 and 15 characters")
	}
	if payload.Address == "" {
		return exception.BadRequestErr("address cannot Empty")
	}
	if payload.Divisi == "" {
		return exception.BadRequestErr("divisi cannot Empty")
	}
	err := s.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new staff: %v", err)
	}
	return nil
}

// Delete implements StaffUseCase.
func (s *staffUseCase) Delete(nik_staff string) error {
	staff, err := s.FindById(nik_staff)
	if err != nil {
		return err
	}
	err = s.repo.Delete(staff.Nik_Staff)
	if err != nil {
		return fmt.Errorf("failed to delete staff: %v", err)
	}
	return nil
}

// FindAll implements StaffUseCase.
func (s *staffUseCase) FindByAll() ([]model.Staff, error) {
	staff, err := s.repo.FindByAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find by all staff: %v", err)
	}
	return staff, nil
}

// FindByName implements StaffUseCase.
func (s *staffUseCase) FindByName(name string) ([]model.Staff, error) {
	staff, err := s.repo.FindByName(name)
	if err != nil {
		return nil, exception.BadRequestErr("name staff not found")
	}
	return staff, nil

}

// Paging implements StaffUseCase.
func (s *staffUseCase) Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error) {
	return s.repo.Paging(payload)
}

// Update implements StaffUseCase.
func (s *staffUseCase) Update(payload model.Staff) error {
	if payload.Nik_Staff == "" {
		return exception.BadRequestErr("nik staff cannot Empty")
	}
	if payload.Name == "" {
		return exception.BadRequestErr("name cannot Empty")
	}
	if len(payload.Phone_number) < 10 || len(payload.Phone_number) > 15 {
		return exception.BadRequestErr("phone number must be between 10 and 15 characters")
	}
	if payload.Address == "" {
		return exception.BadRequestErr("address cannot Empty")
	}
	if payload.Divisi == "" {
		return exception.BadRequestErr("divisi cannot Empty")
	}
	_, err := s.FindById(payload.Nik_Staff)
	if err != nil {
		return err
	}
	err = s.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update staff: %v", err)
	}
	return nil
}

func NewStaffUseCase(repo repository.StaffRepository) StaffUseCase {
	return &staffUseCase{
		repo: repo,
	}
}
