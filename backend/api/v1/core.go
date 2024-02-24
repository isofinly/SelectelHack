package v1

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/invalidteam/selectel_hack/api/auth"
	"github.com/invalidteam/selectel_hack/utils"
)

func SetupRoutesV1(v1 *fiber.Router) {

	authentication := (*v1).Group("/auth")
	// account data managent
	authentication.Get("/me", authGetMeHandler).Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    auth.Keys.PublicKey,
		},
	}))
	authentication.Patch("/me", authPatchMeHandler).Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    auth.Keys.PublicKey,
		},
	}))
	// authentication
	authentication.Post("/login", authPostLoginHandler)

	// bonuses information
	bonuses := (*v1).Group("/bonuses")
	bonuses.Get("/", bonusesGetHandler)
	bonuses.Get("/:id", bonusesGetWithIdHandler)
	bonuses.Post("/:id/feedback", bonusesPostFeedbackHandler) // TODO переделать для публикации оценки и получения оценки

	// bonuses
	bonuses.Patch("/:id/feedback", bonusesPatchFeedbackHandler) // TODO переделать для публикации оценки и получения оценки

	// account recovery
	authentication.Post("/set_password", authPostSetPasswordHandler)

	// extra
	authentication.Post("/set_two_factor_auth", authPostSetTwoFactorAuthHandler)

	// TODO ????
	// donation plan management
	donationPlan := (*v1).Group("/donation_plan")
	donationPlan.Get("/", donationPlanGetHandler)
	donationPlan.Post("/", donationPlanPostHandler)
	donationPlan.Get("/:id", donationPlanGetWithIdHandler)
	donationPlan.Put("/:id", donationPlanPutWithIdHandler)
	donationPlan.Patch("/:id", donationPlanPatchWithIdHandler)
	donationPlan.Delete("/:id", donationPlanDeleteWithIdHandler)
	donationPlan.Get("/latest", donationPlanGetLatestHandler)

	// Donation management
	donations := (*v1).Group("/donations")
	donations.Get("/", donationsGetHandler)
	donations.Post("/", donationsPostHandler)
	donations.Get("/:id", donationsGetWithIdHandler)
	donations.Put("/:id", donationsPutWithIdHandler)
	donations.Patch("/:id", donationsPatchWithIdHandler)
	donations.Delete("/:id", donationsDeleteWithIdHandler)
	donations.Get("/is-exists", donationsGetIsExistsHandler)

	// picture
	picture := (*v1).Group("/picture")
	picture.Post("/", picturePostHandler)

	(*v1).Use(utils.Redirect)
}
