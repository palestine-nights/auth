package api

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MYSQL.
	"github.com/jmoiron/sqlx"
	"github.com/palestine-nights/auth/src/db"
)

// GenericError error model.
//
// swagger:model
type GenericError struct {
	// Error massage.
	Error string `json:"error"`
}

// Server is composition of router and DB instances.
// swagger:ignore
type Server struct {
	Router *gin.Engine
	DB     *sqlx.DB
	DBConn *sql.Conn
}

// Token is model with JWT.
//
// swagger:model
type Token struct {
	// JWT Token.
	Token string `json:"token"`
}

/// swagger:route GET /auth menu listMenu
/// Authenticate user with valid.
/// Responses:
///   500: GenericError
///   200: Token
func (server *Server) authenticate(c *gin.Context) {
	requestedUser := db.User{}

	if err := c.BindJSON(&requestedUser); err != nil {
		c.JSON(http.StatusBadRequest, GenericError{Error: "Invalid request payload"})
		return
	}

	user, err := db.User.FindByPassword(db.User{}, server.DB, requestedUser.UserName, requestedUser.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, GenericError{Error: "Incorrect credentials"})
		return
	}

	// Generate token.
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = user.UserName
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// RSA_PRIVATE_KEY - environment variable responsible for JWT private key.
	EncPrivateKey := os.Getenv("RSA_PRIVATE_KEY")

	privateKey, err := base64.StdEncoding.DecodeString(EncPrivateKey)

	if err != nil {
		gin.DefaultErrorWriter.Write([]byte(err.Error()))
		c.JSON(http.StatusInternalServerError, GenericError{Error: "Something went wrong"})
		return
	}

	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericError{Error: err.Error()})
		return
	}

	// Get token string.
	tokenString, err := token.SignedString(signingKey)

	if err == nil {
		c.JSON(http.StatusOK, Token{Token: tokenString})
	} else {
		c.JSON(http.StatusInternalServerError, GenericError{Error: err.Error()})
	}
}

// GetServer returns newly created Server{} object.
func GetServer(user, password, database, host, port string) *Server {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)

	DB := db.Initialize(connectionString)

	router := gin.Default()
	router.Use(cors.Default())
	server := Server{Router: router, DB: DB}

	server.Router.StaticFile("/", "./html/home.html")
	server.Router.POST("/auth", server.authenticate)

	return &server
}
