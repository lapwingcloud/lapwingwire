package rest

import (
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"github.com/lapwingcloud/lapwingwire/controller/ent/oidcconfig"
	"golang.org/x/oauth2"
)

func (t *handler) GetAuthSignin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rw := NewResponder(w, r)
	providerKey := chi.URLParam(r, "provider_key")
	oidcConfigEnt, err := t.db.OIDCConfig.Query().Where(oidcconfig.ProviderKey(providerKey)).First(ctx)
	if err != nil {
		rw.InternalServerError().Err(err).Msg("failed to query oidc config")
		return
	}
	if oidcConfigEnt == nil {
		rw.NotFound().Err(err).Msg("oidc config not found")
		return
	}
	oidcProvider, err := oidc.NewProvider(ctx, oidcConfigEnt.DiscoveryURI)
	if err != nil {
		rw.InternalServerError().Err(err).Msg("failed to discover oidc config")
	}
	oauth2Config := oauth2.Config{
		Endpoint:     oidcProvider.Endpoint(),
		ClientID:     oidcConfigEnt.ClientID,
		ClientSecret: oidcConfigEnt.ClientSecret,
		RedirectURL:  oidcConfigEnt.RedirectURI,
		Scopes:       []string{"openid", "profile", "email"},
	}
	url := oauth2Config.AuthCodeURL("deadbeef")
	http.Redirect(w, r, url, http.StatusFound)
}

func (t *handler) GetAuthCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rw := NewResponder(w, r)

	code := r.FormValue("code")
	state := r.FormValue("state")

	providerKey := chi.URLParam(r, "provider_key")
	oidcConfigEnt, err := t.db.OIDCConfig.Query().Where(oidcconfig.ProviderKey(providerKey)).First(ctx)
	if err != nil {
		rw.InternalServerError().Err(err).Msg("failed to query oidc config")
		return
	}
	if oidcConfigEnt == nil {
		rw.NotFound().Err(err).Msg("oidc config not found")
		return
	}
	oidcProvider, err := oidc.NewProvider(ctx, oidcConfigEnt.DiscoveryURI)
	if err != nil {
		rw.InternalServerError().Err(err).Msg("failed to discover oidc config")
	}
	oauth2Config := oauth2.Config{
		Endpoint:     oidcProvider.Endpoint(),
		ClientID:     oidcConfigEnt.ClientID,
		ClientSecret: oidcConfigEnt.ClientSecret,
		RedirectURL:  oidcConfigEnt.RedirectURI,
		Scopes:       []string{"openid", "profile", "email"},
	}

	if state != "deadbeef" {
		rw.BadRequest().Msg("invalid oauth2 state")
		return
	}
	token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		rw.InternalServerError().Err(err).Msg("failed to exchange code with google")
		return
	}
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		rw.BadRequest().Msg("failed to extract id token from response")
		return
	}

	tokenVerifier := oidcProvider.Verifier(&oidc.Config{ClientID: oidcConfigEnt.ClientID})
	idToken, err := tokenVerifier.Verify(ctx, rawIDToken)
	if err != nil {
		rw.BadRequest().Msg("failed to verify id token")
		return
	}
	type Claims struct {
		Name          string `json:"name"`
		FamilyName    string `json:"family_name"`
		GivenName     string `json:"given_name"`
		MiddleName    string `json:"middle_name"`
		Picture       string `json:"picture"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
	}
	var claims Claims
	if err = idToken.Claims(&claims); err != nil {
		rw.BadRequest().Msg("failed to extract claims from id token")
		return
	}
	resp := struct {
		IDToken  string    `json:"id_token"`
		Issuer   string    `json:"iss"`
		Audience []string  `json:"aud"`
		Subject  string    `json:"sub"`
		Expiry   time.Time `json:"exp"`
		IssuedAt time.Time `json:"iat"`
		Claims   Claims    `json:"claims"`
	}{
		IDToken:  rawIDToken,
		Issuer:   idToken.Issuer,
		Audience: idToken.Audience,
		Subject:  idToken.Subject,
		Expiry:   idToken.Expiry,
		IssuedAt: idToken.IssuedAt,
		Claims:   claims,
	}
	rw.OK().Data(resp)
}
