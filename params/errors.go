package params

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shop-pkg/errors"
)

var (
	ErrServerInternal = errors.NewError(fiber.StatusInternalServerError, "server.internal_error")

	ErrServerInvalidQuery = errors.NewError(fiber.StatusBadRequest, "server.method.invalid_message_query")

	ErrServerInvalidBody = errors.NewError(fiber.StatusBadRequest, "server.method.invalid_message_body")

	ErrAuthzInvalidPermission = errors.NewError(fiber.StatusUnauthorized, "authz.invalid_permission")

	ErrAuthzCsrfTokenMismatch = errors.NewError(fiber.StatusUnauthorized, "authz.csrf_token_mismatch")

	ErrAuthzMissingCsrfToken = errors.NewError(fiber.StatusUnauthorized, "authz.missing_csrf_token")

	ErrAuthzClientSessionMismatch = errors.NewError(fiber.StatusUnauthorized, "authz.client_session_mismatch")

	ErrAuthzUserNotActive = errors.NewError(fiber.StatusUnauthorized, "authz.user_not_active")

	ErrAuthzUserNotExist = errors.NewError(fiber.StatusUnauthorized, "authz.user_not_exist")

	ErrAuthzInvalidSession = errors.NewError(fiber.StatusUnauthorized, "authz.invalid_session")

	ErrAuthzPermissionDenied = errors.NewError(fiber.StatusUnauthorized, "authz.permission_denied")

	ErrJWTDecodeAndVerify = errors.NewError(fiber.StatusUnauthorized, "jwt.decode_and_verify")

	//ErrMethodNotAllowed = errors.NewError(fiber.StatusMethodNotAllowed, "server.method.not_allowed")

	ErrRecordNotFound = errors.NewError(fiber.StatusNotFound, "record.not_found")

	ErrAuthzGuestOnly = errors.NewError(fiber.StatusBadRequest, "authz.guest_only")

	ErrAuthzInvalidNonce = errors.NewError(fiber.StatusUnauthorized, "authz.invalid_nonce")

	ErrAuthzNonceExpired = errors.NewError(fiber.StatusUnauthorized, "authz.nonce_expired")

	ErrAuthzInvalidSignature = errors.NewError(fiber.StatusUnauthorized, "authz.invalid_signature")

	//ErrAuthzInvalidApiKey = errors.NewError(fiber.StatusUnauthorized, "authz.invalid_api_key")

	ErrAuthzApiKeyNotActive = errors.NewError(fiber.StatusUnauthorized, "authz.apikey_not_active")

	ErrAuthzDisabled2FA = errors.NewError(fiber.StatusUnauthorized, "authz.disabled_2fa")
)
