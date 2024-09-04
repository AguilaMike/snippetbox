package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/AguilaMike/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	// Create a new instance of our application struct. For now, this just
	app := newTestApplication(t)

	// Create a new test server using the httptest.NewServer() function. We
	// pass the value returned by our app.routes() method as the parameter.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Make a GET request to the /ping endpoint on the test server. We use the
	code, _, body := ts.get(t, "/ping")

	// Check that the response code is 200.
	assert.Equal(t, code, http.StatusOK)
	// Check that the response body is "OK".
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	// Create a new instance of our application struct which uses the mocked
	// dependencies.
	app := newTestApplication(t)

	// Establish a new test server for running end-to-end tests.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Set up some table-driven tests to check the responses sent by our
	// application for different URLs.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond...",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			assert.Equal(t, code, tt.wantCode)

			if tt.wantBody != "" {
				assert.StringContains(t, body, tt.wantBody)
			}
		})
	}
}

func TestSnippetCreate(t *testing.T) {
	// Create a new instance of our application struct which uses the mocked
	// dependencies.
	app := newTestApplication(t)

	type ActionValidate struct {
		Method, Actual, Expectedstr string
		ExpectedInt                 int
	}

	tests := []struct {
		name, action string
		invokeLogin  bool
		validations  []ActionValidate
	}{
		{
			name:        "Authenticated",
			action:      "/snippet/create",
			invokeLogin: true,
			validations: []ActionValidate{
				{Method: "Equal", Actual: "code", ExpectedInt: http.StatusOK},
				{Method: "StringContains", Actual: "body", Expectedstr: "<form action='/snippet/create' method='POST'>"},
			},
		},
		{
			name:   "Unauthenticated",
			action: "/snippet/create",
			validations: []ActionValidate{
				{Method: "Equal", Actual: "code", ExpectedInt: http.StatusSeeOther},
				{Method: "Equal", Actual: "location", Expectedstr: "/user/login"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Establish a new test server for running end-to-end tests.
			ts := newTestServer(t, app.routes())
			defer ts.Close()

			if tt.invokeLogin {
				// Make a GET /user/login request and extract the CSRF token from the response
				_, _, body := ts.get(t, "/user/login")
				csrfToken := extractCSRFToken(t, body)

				// Make a POST /user/login request using the extracted CSRF token and
				// credentials from our the mock user model.
				form := url.Values{}
				form.Add("email", "alice@example.com")
				form.Add("password", "pa$$word")
				form.Add("csrf_token", csrfToken)
				ts.postForm(t, "/user/login", form)
			}

			code, headers, body := ts.get(t, tt.action)

			for _, v := range tt.validations {
				switch v.Method {
				case "Equal":
					switch v.Actual {
					case "code":
						assert.Equal(t, code, v.ExpectedInt)
					case "location":
						assert.Equal(t, headers.Get("Location"), v.Expectedstr)
					default:
						t.Fatalf("Unknown validation method: %s", v.Actual)
					}
				case "StringContains":
					assert.StringContains(t, body, v.Expectedstr)
				}
			}

			ts.Close()
		})
	}
}

func TestUserSignup(t *testing.T) {
	// Create the application struct containing our mocked dependencies and set
	// up the test server for running an end-to-end test.
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Make a GET /user/signup request and then extract the CSRF token from the
	// response body.
	_, _, body := ts.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	// Log the CSRF token value in our test output using the t.Logf() function.
	// The t.Logf() function works in the same way as fmt.Printf(), but writes
	// the provided message to the test output.
	// t.Logf("CSRF token is: %q", validCSRFToken )

	const (
		validName     = "Bob"
		validPassword = "validPassw0rd"
		validEmail    = "bob@example.com"
		formTag       = "<form action='/user/signup' method='POST' novalidate>"
	)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantFormTag  string
	}{
		{
			name:         "Valid submission",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusSeeOther,
		},
		{
			name:         "Invalid CSRF Token",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    "wrongToken",
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "Empty name",
			userName:     "",
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty email",
			userName:     validName,
			userEmail:    "",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty password",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: "",
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Invalid email",
			userName:     validName,
			userEmail:    "bob@example.",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Short password",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: "pa$$",
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Duplicate email",
			userName:     validName,
			userEmail:    "dupe@example.com",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)

			assert.Equal(t, code, tt.wantCode)

			if tt.wantFormTag != "" {
				assert.StringContains(t, body, tt.wantFormTag)
			}
		})
	}
}
