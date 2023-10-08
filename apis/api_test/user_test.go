package apitest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aniket-mdev/hr_managment/apis/dto"
	"github.com/aniket-mdev/hr_managment/utils"
	"github.com/stretchr/testify/require"
)

var base_url = "http://localhost:8080/api/"

func TestCreateUser(t *testing.T) {
	endpoint := base_url + "register-user"
	u_t, u_f := true, false
	args := []struct {
		RequestBody        dto.CreateUserRequestDTO
		ExpectedStatusCode int
	}{
		{
			RequestBody: dto.CreateUserRequestDTO{
				Name:            utils.RandomUserName(6),
				Email:           utils.RandomUserEmail(),
				Contact:         utils.RandomUserContact(),
				Password:        "user123",
				UserType:        "emp",
				IsAccountActive: &u_t,
			},
			ExpectedStatusCode: http.StatusOK,
		}, {
			RequestBody: dto.CreateUserRequestDTO{
				Name:            utils.RandomUserName(6),
				Email:           utils.RandomUserEmail(),
				Contact:         utils.RandomUserContact(),
				Password:        "user123",
				UserType:        "emp",
				IsAccountActive: &u_f,
			},
			ExpectedStatusCode: http.StatusOK,
		},
	}

	for _, req := range args {
		t.Run(req.RequestBody.Name, func(t *testing.T) {
			request_body, err := json.Marshal(&req.RequestBody)

			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(request_body))

			require.NoError(t, err)

			response, err := http.DefaultClient.Do(request)

			require.NoError(t, err)
			require.Equal(t, req.ExpectedStatusCode, response.StatusCode)
		})
	}
}
