package draft

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/") //r.URL.Path returns the path of the request which we
	// can then use strings.TrimPrefix to trim away /players/ to get the requested player

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}

func TestGetPlayer(t *testing.T){
	t.Run("return pepper's score", func(t *testing.T) {
		request,_ := http.NewRequest("GET","/players/Pepper",nil)
		response := httptest.NewRecorder()

		PlayerServer(response,request)
		got := response.Body.String()
		wont := "20"

		if got!=wont{
			t.Errorf("got %q want %q",got,wont)
		}

	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/players/Floyd", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}