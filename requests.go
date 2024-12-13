package exaroton

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"pkg.icikowski.pl/exaroton/model"
)

const (
	contentTypeJSON           = "application/json"
	contentTypeOctetStream    = "application/octet-stream"
	contentTypeInodeDirectory = "inode/directory"
)

type requestsBuilder struct {
	apiURL string
	apiKey string
}

func newRequestBuilder(apiURL, apiKey string) *requestsBuilder {
	return &requestsBuilder{
		apiURL: apiURL,
		apiKey: apiKey,
	}
}

func (rb *requestsBuilder) setBearer(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+rb.apiKey)
}

func (rb *requestsBuilder) setContentType(req *http.Request, contentType string) {
	req.Header.Set("Content-Type", contentType)
}

func (rb *requestsBuilder) buildAccountGetRequest(ctx context.Context) *http.Request {
	target := fmt.Sprintf("%s/account", rb.apiURL)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildServersGetRequest(ctx context.Context) *http.Request {
	target := fmt.Sprintf("%s/servers", rb.apiURL)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildServerGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildLogsGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/logs", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildLogsShareGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/logs/share", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildRAMGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/ram", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildRAMPostRequest(ctx context.Context, serverID string, ram model.ServerRAM) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/ram", rb.apiURL, serverID)
	body, _ := json.Marshal(ram)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(body))
	rb.setBearer(req)
	rb.setContentType(req, contentTypeJSON)

	return req
}

func (rb *requestsBuilder) buildMOTDGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/motd", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildMOTDPostRequest(ctx context.Context, serverID string, params model.ServerMOTD) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/motd", rb.apiURL, serverID)
	body, _ := json.Marshal(params)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(body))
	rb.setBearer(req)
	rb.setContentType(req, contentTypeJSON)

	return req
}

func (rb *requestsBuilder) buildActionStartGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/start", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildActionStartPostRequest(ctx context.Context, serverID string, params model.ServerStartParams) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/start", rb.apiURL, serverID)
	body, _ := json.Marshal(params)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(body))
	rb.setBearer(req)
	rb.setContentType(req, contentTypeJSON)

	return req
}

func (rb *requestsBuilder) buildActionStopGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/stop", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildActionRestartGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/restart", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildActionCommandPostRequest(ctx context.Context, serverID string, params model.ServerCommand) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/command", rb.apiURL, serverID)
	body, _ := json.Marshal(params)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(body))
	rb.setBearer(req)
	rb.setContentType(req, contentTypeJSON)

	return req
}

func (rb *requestsBuilder) buildPlayerListsGetRequest(ctx context.Context, serverID string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/playerlists", rb.apiURL, serverID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildPlayerListGetRequest(ctx context.Context, serverID, listName string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/playerlists/%s", rb.apiURL, serverID, listName)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildPlayerListPutRequest(ctx context.Context, serverID, listName string, params model.PlayerListEntries) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/playerlists/%s", rb.apiURL, serverID, listName)
	body, _ := json.Marshal(params)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, target, bytes.NewReader(body))
	rb.setBearer(req)
	rb.setContentType(req, contentTypeJSON)

	return req
}

func (rb *requestsBuilder) buildPlayerListDeleteRequest(ctx context.Context, serverID, listName string, params model.PlayerListEntries) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/playerlists/%s", rb.apiURL, serverID, listName)
	body, _ := json.Marshal(params)

	req, _ := http.NewRequestWithContext(ctx, http.MethodDelete, target, bytes.NewReader(body))
	rb.setBearer(req)
	rb.setContentType(req, contentTypeJSON)

	return req
}

func (rb *requestsBuilder) buildFileInfoGetRequest(ctx context.Context, serverID, path string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/files/info/%s", rb.apiURL, serverID, path)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildFileDataGetRequest(ctx context.Context, serverID, path string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/files/data/%s", rb.apiURL, serverID, path)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildFileDataFilePutRequest(ctx context.Context, serverID, path string, src io.Reader) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/files/data/%s", rb.apiURL, serverID, path)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, target, src)
	rb.setBearer(req)
	rb.setContentType(req, contentTypeOctetStream)

	return req
}

func (rb *requestsBuilder) buildFileDataDirectoryPutRequest(ctx context.Context, serverID, path string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/files/data/%s", rb.apiURL, serverID, path)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, target, http.NoBody)
	rb.setBearer(req)
	rb.setContentType(req, contentTypeInodeDirectory)

	return req
}

func (rb *requestsBuilder) buildFileDataDeleteRequest(ctx context.Context, serverID, path string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/files/data/%s", rb.apiURL, serverID, path)

	req, _ := http.NewRequestWithContext(ctx, http.MethodDelete, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildConfigGetRequest(ctx context.Context, serverID, path string) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/files/config/%s", rb.apiURL, serverID, path)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildConfigPostRequest(ctx context.Context, serverID, path string, opts model.ConfigValues) *http.Request {
	target := fmt.Sprintf("%s/servers/%s/files/config/%s", rb.apiURL, serverID, path)
	body, _ := json.Marshal(opts)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(body))
	rb.setBearer(req)
	rb.setContentType(req, contentTypeJSON)

	return req
}

func (rb *requestsBuilder) buildCreditPoolsGetRequest(ctx context.Context) *http.Request {
	target := fmt.Sprintf("%s/billing/pools", rb.apiURL)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildCreditPoolGetRequest(ctx context.Context, poolID string) *http.Request {
	target := fmt.Sprintf("%s/billing/pools/%s", rb.apiURL, poolID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildCreditPoolMembersGetRequest(ctx context.Context, poolID string) *http.Request {
	target := fmt.Sprintf("%s/billing/pools/%s/members", rb.apiURL, poolID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}

func (rb *requestsBuilder) buildCreditPoolServersGetRequest(ctx context.Context, poolID string) *http.Request {
	target := fmt.Sprintf("%s/billing/pools/%s/servers", rb.apiURL, poolID)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, http.NoBody)
	rb.setBearer(req)

	return req
}
