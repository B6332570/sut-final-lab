package entity

import (
	"testing"

	"github.com/onsi/gomega"

	"gorm.io/gorm"

	"github.com/asaskevich/govalidator"
)

type Employee3 struct {
	gorm.Model
	Name       string `valid:"required~Name not blank"`
	Email      string
	EmployeeID string `valid:"matches([JMS]\\d{8})~EmployeeID is not correct"`
}

func TestEmployeeID(t *testing.T) {
	t.Run("EmployeeID is not correct", func(t *testing.T) {

		g := gomega.NewGomegaWithT(t)
		emp := Employee3{
			Name:       "GG",
			Email:      "Preechapat2544@gmail.com",
			EmployeeID: "B6332570",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("EmployeeID is not correct"))

	})

}
