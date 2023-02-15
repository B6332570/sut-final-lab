package entity

import (
	"testing"

	"github.com/onsi/gomega"

	"gorm.io/gorm"

	"github.com/asaskevich/govalidator"
)

type Employee1 struct {
	gorm.Model
	Name       string `valid:"required~Name not blank"`
	Email      string
	EmployeeID string `valid:"matches([JMS]\\d{8})~EmployeeID is not correct"`
}

func TestSuccess(t *testing.T) {

	g := gomega.NewGomegaWithT(t)
	emp := Employee1{
		Name:       "GG",
		Email:      "Preechapat2544@gmail.com",
		EmployeeID: "J12345678",
	}
	ok, _ := govalidator.ValidateStruct(emp)
	g.Expect(ok).To(gomega.BeTrue())
	// g.Expect(err).To(gomega.BeNil())
	// g.Expect(err.Error()).To(gomega.Equal("GGG"))

}
