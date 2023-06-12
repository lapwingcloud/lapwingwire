package rest

import (
	"encoding/json"
	"net/http"

	"github.com/lapwingcloud/lapwingwire/controller/ent"
	"github.com/lapwingcloud/lapwingwire/controller/ent/oidcconfig"
)

func (t *handler) PutOIDCConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rw := NewResponder(w, r)

	var req struct {
		ID           int    `json:"id"`
		ProviderKey  string `json:"provider_key"`
		ProviderName string `json:"provider_name"`
		DiscoveryURI string `json:"discovery_uri"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RedirectURI  string `json:"redirect_uri"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rw.BadRequest().Err(err).Msg("failed to decode request body to json")
		return
	}

	var err error
	var oidcConfigEnt *ent.OIDCConfig
	if req.ID == 0 {
		oidcConfigEnt, err = t.db.OIDCConfig.Create().
			SetProviderKey(req.ProviderKey).
			SetProviderName(req.ProviderName).
			SetDiscoveryURI(req.DiscoveryURI).
			SetClientID(req.ClientID).
			SetClientSecret(req.ClientSecret).
			SetRedirectURI(req.RedirectURI).
			Save(ctx)
		if err != nil {
			rw.InternalServerError().Err(err).Msg("failed to create oidc config")
			return
		}
	} else {
		oidcConfigEnt, err = t.db.OIDCConfig.Query().Where(oidcconfig.ID(req.ID)).First(ctx)
		if err != nil {
			rw.InternalServerError().Err(err).Msg("failed to query oidc config")
			return
		}
		if oidcConfigEnt == nil {
			rw.NotFound().Msg("oidc config not found")
			return
		}
		oidcConfigEnt, err = t.db.OIDCConfig.UpdateOne(oidcConfigEnt).
			SetProviderKey(req.ProviderKey).
			SetProviderName(req.ProviderName).
			SetDiscoveryURI(req.DiscoveryURI).
			SetClientID(req.ClientID).
			SetClientSecret(req.ClientSecret).
			SetRedirectURI(req.RedirectURI).
			Save(ctx)
		if err != nil {
			rw.InternalServerError().Err(err).Msg("failed to update oidc config")
			return
		}
	}
	resp := struct {
		ID           int    `json:"id"`
		ProviderKey  string `json:"provider_key"`
		ProviderName string `json:"provider_name"`
		DiscoveryURI string `json:"discovery_uri"`
		ClientID     string `json:"client_id"`
		RedirectURI  string `json:"redirect_uri"`
	}{
		ID:           oidcConfigEnt.ID,
		ProviderKey:  oidcConfigEnt.ProviderKey,
		ProviderName: oidcConfigEnt.ProviderName,
		DiscoveryURI: oidcConfigEnt.DiscoveryURI,
		ClientID:     oidcConfigEnt.ClientID,
		RedirectURI:  oidcConfigEnt.RedirectURI,
	}
	rw.OK().Data(resp)
}
