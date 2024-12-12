package exaroton

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"pkg.icikowski.pl/exaroton/model"
)

type fileClient struct {
	*baseClient

	serverID string
	path     string
}

func newFileClient(base *baseClient, serverID, path string) FileAPI {
	return &fileClient{
		baseClient: base,
		serverID:   serverID,
		path:       path,
	}
}

// GetFile implements FileAPI.
func (c *fileClient) GetFile(ctx context.Context) (*model.File, *RawResponse[model.File], error) {
	resp, err := c.do(c.rb.buildFileInfoGetRequest(
		ctx,
		c.serverID,
		c.path,
	))
	if err != nil {
		return nil, nil, err
	}

	rawResp, err := decodeRawResponse[model.File](resp)
	if err != nil {
		return nil, rawResp, err
	}
	return rawResp.Data, rawResp, nil
}

// GetContent implements FileAPI.
func (c *fileClient) GetContent(ctx context.Context, dst io.Writer) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildFileDataGetRequest(
		ctx,
		c.serverID,
		c.path,
	))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		rawResp, err := decodeRawResponse[any](resp)
		if err != nil {
			return rawResp, err
		}
		return rawResp, nil
	}

	defer resp.Body.Close()
	if _, err := io.Copy(dst, resp.Body); err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return nil, nil
}

// SetContent implements FileAPI.
func (c *fileClient) SetContent(ctx context.Context, src io.Reader) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildFileDataFilePutRequest(
		ctx,
		c.serverID,
		c.path,
		src,
	))
	if err != nil {
		return nil, err
	}

	rawResp, err := decodeRawResponse[any](resp)
	if err != nil {
		return nil, err
	}

	return rawResp, nil
}

// CreateDirectory implements FileAPI.
func (c *fileClient) CreateDirectory(ctx context.Context) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildFileDataDirectoryPutRequest(
		ctx,
		c.serverID,
		c.path,
	))
	if err != nil {
		return nil, err
	}

	rawResp, err := decodeRawResponse[any](resp)
	if err != nil {
		return nil, err
	}

	return rawResp, nil
}

// Delete implements FileAPI.
func (c *fileClient) Delete(ctx context.Context) (*RawResponse[any], error) {
	resp, err := c.do(c.rb.buildFileDataDeleteRequest(
		ctx,
		c.serverID,
		c.path,
	))
	if err != nil {
		return nil, err
	}

	rawResp, err := decodeRawResponse[any](resp)
	if err != nil {
		return nil, err
	}

	return rawResp, nil
}
