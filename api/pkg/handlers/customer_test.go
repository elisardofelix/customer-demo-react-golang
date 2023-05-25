package handlers

import (
	"context"
	"customer-demo/pkg/mocks"
	"customer-demo/pkg/models"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_customerHandler_GetAll(t *testing.T) {
	ctl := gomock.NewController(t)
	customerProviderMock := mocks.NewMockCustomerProvider(ctl)

	customerProviderMock.EXPECT().GetAll().Return([]*models.Customer{
		&models.Customer{
			BaseMySQLID: models.BaseMySQLID{
				ID: 1,
			},
			FirstName: "Juan",
			LastName:  "Perencejo",
			Email:     "jperencejo@gmail.com",
			Gender:    "Male",
			Company:   "Catojin S.A.",
			City:      "San Cristobal",
			Title:     "Sr.",
		},
		&models.Customer{
			BaseMySQLID: models.BaseMySQLID{
				ID: 2,
			},
			FirstName: "Maria Magdalena",
			LastName:  "Grano de Oro",
			Email:     "mamagrano@gmail.com",
			Gender:    "Female",
			Company:   "Catojin S.A.",
			City:      "Santo Domingo",
			Title:     "Sra.",
		},
	}, nil)

	costumerHandler := NewCustomerHandler(customerProviderMock)

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/api/customers", nil)
	if err != nil {
		t.Fatal(err)
	}

	costumerHandler.GetAll(rr, r)
	rs := rr.Result()
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	b, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	var responseStruct SuccessResponse

	json.Unmarshal(b, &responseStruct)

	require.Len(t, responseStruct.Data, 2)
}

func Test_customerHandler_GetById(t *testing.T) {
	ctl := gomock.NewController(t)
	customerProviderMock := mocks.NewMockCustomerProvider(ctl)

	customerProviderMock.EXPECT().GetById(1).
		Return(&models.Customer{
			BaseMySQLID: models.BaseMySQLID{
				ID: 1,
			},
			FirstName: "Juan",
			LastName:  "Perencejo",
			Email:     "jperencejo@gmail.com",
			Gender:    "Male",
			Company:   "Catojin S.A.",
			City:      "San Cristobal",
			Title:     "Sr.",
		}, nil)

	costumerHandler := NewCustomerHandler(customerProviderMock)

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/api/customers/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	costumerHandler.GetById(rr, r)
	rs := rr.Result()
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	b, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	var responseStruct SuccessResponse

	json.Unmarshal(b, &responseStruct)

	data := responseStruct.Data.(map[string]interface{})

	require.Equal(t, 1, int(data["id"].(float64)))
	require.Equal(t, "Juan", data["first_name"].(string))
	require.Equal(t, "Perencejo", data["last_name"].(string))

	require.NotNil(t, responseStruct.Data)

}
