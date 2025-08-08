package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/credondocr/github-workflow-showcase/models"
	"github.com/credondocr/github-workflow-showcase/routes"
)

var _ = Describe("User API", func() {
	var router http.Handler

	BeforeEach(func() {
		router = routes.SetupRouter()
	})

	Describe("GET /health", func() {
		It("should return health status", func() {
			req, _ := http.NewRequest("GET", "/health", http.NoBody)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			Expect(err).ToNot(HaveOccurred())
			Expect(response["success"]).To(BeTrue())
			Expect(response["message"]).To(Equal("API is working correctly"))
		})
	})

	Describe("GET /", func() {
		It("should return API information", func() {
			req, _ := http.NewRequest("GET", "/", http.NoBody)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			Expect(err).ToNot(HaveOccurred())
			Expect(response["message"]).To(Equal("Welcome to the example REST API!"))
			Expect(response["version"]).To(Equal("1.0.0"))
		})
	})

	Describe("User CRUD Operations", func() {
		Describe("GET /api/v1/users", func() {
			It("should return empty users list initially", func() {
				req, _ := http.NewRequest("GET", "/api/v1/users", http.NoBody)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response["success"]).To(BeTrue())
				Expect(response["total"]).To(Equal(float64(0)))
			})
		})

		Describe("POST /api/v1/users", func() {
			It("should create a new user with valid data", func() {
				user := models.User{
					Name:  "Juan Pérez",
					Email: "juan@example.com",
					Age:   30,
				}

				jsonData, _ := json.Marshal(user)
				req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonData))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusCreated))

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response["success"]).To(BeTrue())
				Expect(response["message"]).To(Equal("User created successfully"))

				data := response["data"].(map[string]interface{})
				Expect(data["name"]).To(Equal("Juan Pérez"))
				Expect(data["email"]).To(Equal("juan@example.com"))
				Expect(data["age"]).To(Equal(float64(30)))
				Expect(data["id"]).To(Equal(float64(1)))
			})

			It("should return error with invalid data", func() {
				user := models.User{
					Name:  "",
					Email: "invalid-email",
					Age:   -5,
				}

				jsonData, _ := json.Marshal(user)
				req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonData))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusBadRequest))

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response["error"]).ToNot(BeNil())
			})
		})

		Describe("GET /api/v1/users/:id", func() {
			It("should return user by ID after creating one", func() {
				// Primero crear un usuario
				user := models.User{
					Name:  "María García",
					Email: "maria@example.com",
					Age:   25,
				}

				jsonData, _ := json.Marshal(user)
				req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonData))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				// Luego obtener el usuario por ID
				req, _ = http.NewRequest("GET", "/api/v1/users/1", http.NoBody)
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response["success"]).To(BeTrue())

				data := response["data"].(map[string]interface{})
				Expect(data["name"]).To(Equal("María García"))
				Expect(data["email"]).To(Equal("maria@example.com"))
			})

			It("should return 404 for non-existent user", func() {
				req, _ := http.NewRequest("GET", "/api/v1/users/999", http.NoBody)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusNotFound))

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response["error"]).To(Equal("User not found"))
			})
		})

		Describe("PUT /api/v1/users/:id", func() {
			It("should update existing user", func() {
				// Primero crear un usuario
				user := models.User{
					Name:  "Carlos López",
					Email: "carlos@example.com",
					Age:   35,
				}

				jsonData, _ := json.Marshal(user)
				req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonData))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				// Luego actualizar el usuario
				updatedUser := models.User{
					Name:  "Carlos López Actualizado",
					Email: "carlos.updated@example.com",
					Age:   36,
				}

				jsonData, _ = json.Marshal(updatedUser)
				req, _ = http.NewRequest("PUT", "/api/v1/users/1", bytes.NewBuffer(jsonData))
				req.Header.Set("Content-Type", "application/json")
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response["success"]).To(BeTrue())
				Expect(response["message"]).To(Equal("User updated successfully"))
			})
		})

		Describe("DELETE /api/v1/users/:id", func() {
			It("should delete existing user", func() {
				// Primero crear un usuario
				user := models.User{
					Name:  "Ana Rodríguez",
					Email: "ana@example.com",
					Age:   28,
				}

				jsonData, _ := json.Marshal(user)
				req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonData))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				// Luego eliminar el usuario
				req, _ = http.NewRequest("DELETE", "/api/v1/users/1", http.NoBody)
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response["success"]).To(BeTrue())
				Expect(response["message"]).To(Equal("User deleted successfully"))
			})
		})

		Context("User Statistics API", func() {
			It("should return empty statistics when no users exist", func() {
				router := routes.SetupRouter()
				req, _ := http.NewRequest("GET", "/api/v1/users/stats", http.NoBody)
				recorder := httptest.NewRecorder()

				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))

				var response map[string]interface{}
				err := json.Unmarshal(recorder.Body.Bytes(), &response)
				Expect(err).To(BeNil())

				Expect(response["success"]).To(Equal(true))
				Expect(response["message"]).To(Equal("User statistics retrieved successfully"))

				data := response["data"].(map[string]interface{})
				Expect(data["total_users"]).To(Equal(float64(0)))
				Expect(data["average_age"]).To(Equal(float64(0)))
			})

			It("should return correct statistics when users exist", func() {
				router := routes.SetupRouter()

				// Create test users
				users := []map[string]interface{}{
					{"name": "Alice", "email": "alice@test.com", "age": 25},
					{"name": "Bob", "email": "bob@test.com", "age": 35},
					{"name": "Charlie", "email": "charlie@test.com", "age": 45},
				}

				for _, user := range users {
					userData, _ := json.Marshal(user)
					req, _ := http.NewRequest("POST", "/api/v1/users", strings.NewReader(string(userData)))
					req.Header.Set("Content-Type", "application/json")
					recorder := httptest.NewRecorder()
					router.ServeHTTP(recorder, req)
					Expect(recorder.Code).To(Equal(http.StatusCreated))
				}

				// Get statistics
				req, _ := http.NewRequest("GET", "/api/v1/users/stats", http.NoBody)
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))

				var response map[string]interface{}
				err := json.Unmarshal(recorder.Body.Bytes(), &response)
				Expect(err).To(BeNil())

				Expect(response["success"]).To(Equal(true))

				data := response["data"].(map[string]interface{})
				Expect(data["total_users"]).To(Equal(float64(3)))
				Expect(data["average_age"]).To(Equal(float64(35))) // (25+35+45)/3 = 35
				Expect(data["min_age"]).To(Equal(float64(25)))
				Expect(data["max_age"]).To(Equal(float64(45)))

				// Check age ranges
				ageRanges := data["age_ranges"].(map[string]interface{})
				Expect(ageRanges["18-25"]).To(Equal(float64(1))) // Alice
				Expect(ageRanges["26-35"]).To(Equal(float64(1))) // Bob
				Expect(ageRanges["36-50"]).To(Equal(float64(1))) // Charlie
				Expect(ageRanges["51+"]).To(Equal(float64(0)))
			})
		})
	})
})
