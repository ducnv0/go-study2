package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" // The FastAPI equivalent for Go!
)

// We define our incoming payload requirements (Pydantic BaseModel equivalent)
type CreateUserRequest struct {
	// The `binding:"required"` tag acts exactly like Pydantic's automatic validation!
	Username string `json:"username" binding:"required"`
	
	// We can even chain rules! (gte=18 means Greater Than or Equal to 18)
	Age int `json:"age" binding:"required,gte=18"` 
}

func main() {
	// 1. Initialize the router (Python: `app = FastAPI()`)
	r := gin.Default()

	// 2. A simple GET Endpoint (Python: `@app.get("/ping")`)
	r.GET("/ping", func(c *gin.Context) {
		// c.JSON returns an automatic JSON response payload. 
		// `gin.H` is just an incredibly quick shortcut for creating a Map (Dictionary).
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 3. A POST Endpoint with Payload Validation
	r.POST("/users", func(c *gin.Context) {
		// Create an empty struct in memory
		var payload CreateUserRequest

		// ShouldBindJSON looks at the incoming request body, matches the JSON tags, 
		// and runs our `binding` validation. We pass a Pointer `&` so it can mutate `payload`.
		if err := c.ShouldBindJSON(&payload); err != nil {
			// Early Return Pattern! If validation fails, immediately return 400 Bad Request.
			c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
			return 
		}

		// If we made it here, the payload is 100% valid type-safe data!
		responseMessage := fmt.Sprintf("Successfully processed %s!", payload.Username)
		
		c.JSON(http.StatusOK, gin.H{
			"message": responseMessage,
			"data":    payload,
		})
	})

	// 4. Start the server (Python: `uvicorn main:app --port 8080`)
	fmt.Println("🚀 Server starting on http://localhost:8080")
	r.Run(":8080") 
}
