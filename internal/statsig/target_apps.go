package statsig

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type TargetAppAPIRequest struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Gates          []string `json:"gates"`
	DynamicConfigs []string `json:"dynamicConfigs"`
	Experiments    []string `json:"experiments"`
}

type CreateTargetAppAPIResponse struct {
	Message string      `json:"message"`
	Data    []TargetApp `json:"data"`
}

type TargetApp struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (c *Client) CreateTargetApp(ctx context.Context, targetApp TargetAppAPIRequest) (*TargetAppAPIRequest, error) {
	response, err := c.Post("target_app", targetApp)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error creating target app: %s", err))
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("Response Body: %s", map[string]interface{}{"response": string(response)}))
	createdTargetApp := APIResponse[TargetAppAPIRequest]{}
	if err := json.Unmarshal(response, &createdTargetApp); err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error unmarshalling target app response: %s", err))
		return nil, err
	}

	tflog.Trace(ctx, fmt.Sprintf("Target App created with Name: %s; and ID: %s", createdTargetApp.Data.Name, createdTargetApp.Data.Description))

	return &createdTargetApp.Data, nil
}
