package agent_policy

import (
	"context"
	"fmt"

	"github.com/elastic/terraform-provider-elasticstack/internal/clients/fleet"
	"github.com/elastic/terraform-provider-elasticstack/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func (r *agentPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var planModel agentPolicyModel

	diags := req.Plan.Get(ctx, &planModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := r.client.GetFleetClient()
	if err != nil {
		resp.Diagnostics.AddError(err.Error(), "")
		return
	}

	serverVersion, sdkDiags := r.client.ServerVersion(ctx)
	resp.Diagnostics.Append(utils.ConvertSDKDiagnosticsToFramework(sdkDiags)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(planModel.GlobalDataTags.Elements()) > 0 && serverVersion.LessThan(minVersionGlobalDataTags) {
		resp.Diagnostics.AddAttributeError(
			path.Root("global_data_tags"),
			fmt.Sprintf("'global_data_tags' is supported only for Elasticsearch v%s and above", minVersionGlobalDataTags.String()),
			fmt.Sprintf("'global_data_tags' is supported only for Elasticsearch v%s and above", minVersionGlobalDataTags.String()),
		)
	}

	body, err := planModel.toAPICreateModel()
	if err != nil {
		resp.Diagnostics.AddError(err.Error(), err.Error())
		return
	}

	sysMonitoring := planModel.SysMonitoring.ValueBool()
	policy, diags := fleet.CreateAgentPolicy(ctx, client, body, sysMonitoring)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(planModel.populateFromAPI(policy)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, planModel)...)
}
