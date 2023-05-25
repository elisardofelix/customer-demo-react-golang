package providers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGetAllCustomers(t *testing.T) {
	// create sqlmock database connection and mock
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.23"))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db})) //open gorm db

	customerProvider := NewCustomerProvider(gormDB)

	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "gender", "company", "city", "title", "create_at", "modified_at"}).
		AddRow(1, "Elisardo", "Felix", "elisardo123@gmail.com", "Male", "Company S.A.", "Santo Domingo", "Sr", "2023-01-01", "2023-01-01").
		AddRow(1, "Edina", "John", "ej1234@gmail.com", "Female", "Verizon", "MyCom", "Dr", "2023-01-01", "2023-01-01")

	mock.ExpectQuery("SELECT \\* FROM `customers`").WillReturnRows(rows)

	customers, err := customerProvider.GetAll()
	if err != nil {
		t.Errorf("this is the error getting the customers: %v\n", err)
	}

	// Check that we got the expected number of users
	if len(customers) != 2 {
		t.Errorf("expected 2 users, but got %d", len(customers))
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetCustomerById(t *testing.T) {
	// create sqlmock database connection and mock
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.23"))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db})) //open gorm db

	customerProvider := NewCustomerProvider(gormDB)

	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "gender", "company", "city", "title", "create_at", "modified_at"}).
		AddRow(1, "Elisardo", "Felix", "elisardo123@gmail.com", "Male", "Company S.A.", "Santo Domingo", "Sr", "2023-01-01", "2023-01-01")

	mock.ExpectQuery("SELECT \\* FROM `customers` WHERE .*`id` = .* ORDER BY `customers`.*`id` LIMIT 1").WillReturnRows(rows)

	customers, err := customerProvider.GetById(1)
	if err != nil {
		t.Errorf("this is the error getting the customer: %v\n", err)
	}

	require.Equal(t, customers.FirstName, "Elisardo")
	require.Equal(t, customers.LastName, "Felix")
	require.Equal(t, customers.Email, "elisardo123@gmail.com")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
