package v1

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesV1(v1 *fiber.Router) {

	addressNeeds := (*v1).Group("/address_needs")
	addressNeeds.Get("/", addressNeedsGetHandler)
	addressNeeds.Post("/", addressNeedsPostHandler)
	addressNeeds.Get("/:id", addressNeedsGetWithIdHandler)

	authentication := (*v1).Group("/auth")
	// account data managent
	authentication.Post("/change_email", authPostChangeEmailHandler)
	authentication.Post("/change_password", authPostChangePasswordHandler)
	authentication.Post("/change_phone", authPostChangePhoneHandler)
	authentication.Post("/confirm_email", authPostConfirmEmailHandler)
	authentication.Get("/donor_card", authGetDonorCardHandler)
	authentication.Patch("/donor_card", authPatchDonorCardHandler)
	authentication.Get("/me", authGetMeHandler)
	authentication.Patch("/me", authPatchMeHandler)

	// account recovery
	authentication.Post("/check_sms_code", redirect) // unused
	authentication.Post("/recover", redirect)
	authentication.Post("/resend_code", redirect)
	authentication.Post("/resend_email_code", redirect)
	authentication.Post("/set_password", authPostSetPasswordHandler)

	// account creation and confirmation
	authentication.Post("/confirm_email_reg", redirect)
	authentication.Post("/confirm_phone_reg", redirect)
	authentication.Post("/registration", redirect)

	// authentication
	authentication.Post("/login", authPostLoginHandler)
	authentication.Post("/logout", authPostLogoutHandler)

	// extra
	authentication.Post("/set_two_factor_auth", authPostSetTwoFactorAuthHandler)

	// blood station information
	bloodStations := (*v1).Group("/blood_stations")
	bloodStations.Get("/", bloodStationsGetHandler)
	bloodStations.Get("/:id", bloodStationsGetWithIdHandler)
	bloodStations.Get("/:id/planned", bloodStationsGetPlannedWithIdHandler)
	bloodStations.Get("/:id/top", bloodStationsGetTopWithIdHandler)
	bloodStations.Get("/:id/top_halfyear", bloodStationsGetTopHalfYearWithIdHandler)
	bloodStations.Get("/all_names", bloodStationsGetAllNamesHandler)
	bloodStations.Get("/map", bloodStationsGetMapHandler)
	bloodStations.Get("/search", bloodStationsGetSearchHandler)
	bloodStations.Get("/selected", bloodStationsGetSelectedHandler)
	needs := (*v1).Group("/needs")
	needs.Get("/", needsGetHandler)
	needs.Get("/:id", needsGetWithIdHandler)
	needs.Get("/available", needsGetAvailableHandler)

	// bonuses information
	bonuses := (*v1).Group("/bonuses")
	bonuses.Get("/", bonusesGetHandler)
	bonuses.Get("/:id", bonusesGetWithIdHandler)
	bonuses.Post("/:id/feedback", bonusesPostFeedbackHandler) // TODO переделать для публикации оценки и получения оценки

	// bonuses
	bonuses.Get("/:id/feedback", bonusesGetFeedbackWithIdHandler)
	bonuses.Patch("/:id/feedback", bonusesPatchFeedbackHandler) // TODO переделать для публикации оценки и получения оценки

	// regions
	cities := (*v1).Group("/cities")
	cities.Get("/", citiesGetHandler)
	cities.Get("/:id", citiesGetWithIdHandler)
	cities.Get("/by_location", citiesGetByLocationHandler)
	countries := (*v1).Group("/countries")
	countries.Get("/", countriesGetHandler)
	countries.Get("/:id", countriesGetWithIdHandler)
	regions := (*v1).Group("/regions")
	regions.Get("/", regionsGetHandler)
	regions.Get("/:id", regionsGetWithIdHandler)

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

	// Events
	events := (*v1).Group("/events")
	events.Get("/", redirect)
	events.Get("/:id", eventsGetWithIdHandler)

	// picture
	picture := (*v1).Group("/picture")
	picture.Post("/", picturePostHandler)

	// user information
	usersTop := (*v1).Group("/users_top")
	usersTop.Get("/", redirect)

	// (*v1).Use(redirect)
}
