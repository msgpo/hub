package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	ctx, write := ctx.CacheContext()
	defer write()

	node.BeginBlock(ctx, k.Node)
}

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case provider.MsgRegisterProvider:
			return provider.HandleRegisterProvider(ctx, k.Provider, msg)
		case provider.MsgUpdateProvider:
			return provider.HandleUpdateProvider(ctx, k.Provider, msg)
		case node.MsgRegisterNode:
			return node.HandleRegisterNode(ctx, k.Node, msg)
		case node.MsgUpdateNode:
			return node.HandleUpdateNode(ctx, k.Node, msg)
		case node.MsgSetNodeStatus:
			return node.HandleSetNodeStatus(ctx, k.Node, msg)
		case plan.MsgAddPlan:
			return plan.HandleAddPlan(ctx, k.Plan, msg)
		case plan.MsgSetPlanStatus:
			return plan.HandleSetPlanStatus(ctx, k.Plan, msg)
		case plan.MsgAddNodeForPlan:
			return plan.HandleAddNodeForPlan(ctx, k.Plan, msg)
		case plan.MsgRemoveNodeForPlan:
			return plan.HandleRemoveNodeForPlan(ctx, k.Plan, msg)
		case subscription.MsgStartSubscription:
			return subscription.HandleStartSubscription(ctx, k.Subscription, msg)
		case subscription.MsgAddQuotaForSubscription:
			return subscription.HandleAddQuotaForSubscription(ctx, k.Subscription, msg)
		case subscription.MsgUpdateQuotaForSubscription:
			return subscription.HandleUpdateQuotaForSubscription(ctx, k.Subscription, msg)
		case subscription.MsgEndSubscription:
			return subscription.HandleEndSubscription(ctx, k.Subscription, msg)
		case session.MsgUpdateSession:
			return session.HandleUpdateSession(ctx, k.Session, msg)
		default:
			return types.ErrorUnknownMsgType(msg.Type()).Result()
		}
	}
}
