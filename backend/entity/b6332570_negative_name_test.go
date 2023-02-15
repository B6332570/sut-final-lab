package entity

import (
	"testing"

	"github.com/onsi/gomega"

	"gorm.io/gorm"

	"github.com/asaskevich/govalidator"
)

type Employee2 struct {
	gorm.Model
	Name       string `valid:"required~Name not blank"`
	Email      string
	EmployeeID string
}

func TestNameNotBlank(t *testing.T) {
	t.Run("Name not blank", func(t *testing.T) {

		g := gomega.NewGomegaWithT(t)
		emp := Employee2{
			Name:       "",
			Email:      "Preechapat2544@gmail.com",
			EmployeeID: "J12345678",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Name not blank"))

	})

}
