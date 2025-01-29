package handler

import (
	"net/http"
	"time"

	authModel "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/models"
	service "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	// ...handle login logic...
	var login authModel.LoginRegister
	err := c.ShouldBindBodyWithJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.authService.Login(login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	createCookieWithJWT(c, tokens.AccessToken, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": " login successful"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// Remove Access Token Cookie
	c.SetCookie(
		"DAT",
		"",
		-1, // Set expiry to past to delete the cookie
		"/",
		"",
		true,
		true,
	)

	// Remove Refresh Token Cookie
	c.SetCookie(
		"RAT",
		"",
		-1, // Set expiry to past to delete the cookie
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	// ...handle register logic...
	var register authModel.AuthRegister
	err := c.ShouldBindBodyWithJSON(&register)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.authService.Register(register.Username, register.Email, register.Password)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("RAT")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	tokens, err := h.authService.RefreshJWTToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	createCookieWithJWT(c, tokens.AccessToken, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}

func (h *AuthHandler) ValidateJWT(c *gin.Context) {
	tokenString, err := c.Cookie("DAT")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not found"})
		return
	}

	valid, err := h.authService.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token is valid", "claims": valid})
}

func createCookieWithJWT(c *gin.Context, accessToken string, refreshToken string) {
	// Access Token Cookie
	c.SetCookie(
		"DAT",
		accessToken,
		int(15*time.Minute.Seconds()), // Set access token expiry (15 mins example)
		"/",                           // Cookie accessible throughout the site
		"",                            // Domain (empty means current domain)
		true,                          // Secure (use HTTPS for transmission)
		true,                          // HttpOnly (prevent access from JavaScript)
	)

	// Refresh Token Cookie
	c.SetCookie(
		"RAT",
		refreshToken,
		int(7*24*time.Hour.Seconds()), // Set refresh token expiry (7 days example)
		"/",                           // Cookie accessible throughout the site
		"",                            // Domain (empty means current domain)
		true,                          // Secure (use HTTPS for transmission)
		true,                          // HttpOnly (prevent access from JavaScript)
	)
}
