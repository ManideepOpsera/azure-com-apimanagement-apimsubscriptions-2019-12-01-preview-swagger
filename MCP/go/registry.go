package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_subscription "github.com/apimanagementclient/mcp-server/tools/subscription"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_subscription.CreateSubscription_listTool(cfg),
		tools_subscription.CreateSubscription_deleteTool(cfg),
		tools_subscription.CreateSubscription_getTool(cfg),
		tools_subscription.CreateSubscription_getentitytagTool(cfg),
		tools_subscription.CreateSubscription_updateTool(cfg),
		tools_subscription.CreateSubscription_createorupdateTool(cfg),
		tools_subscription.CreateSubscription_listsecretsTool(cfg),
		tools_subscription.CreateSubscription_regenerateprimarykeyTool(cfg),
		tools_subscription.CreateSubscription_regeneratesecondarykeyTool(cfg),
	}
}
