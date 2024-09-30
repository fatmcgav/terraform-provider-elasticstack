package agent_policy

import (
	"fmt"
	"slices"

	fleetapi "github.com/elastic/terraform-provider-elasticstack/generated/fleet"
	"github.com/elastic/terraform-provider-elasticstack/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type agentPolicyModel struct {
	ID                 types.String `tfsdk:"id"`
	PolicyID           types.String `tfsdk:"policy_id"`
	Name               types.String `tfsdk:"name"`
	Namespace          types.String `tfsdk:"namespace"`
	Description        types.String `tfsdk:"description"`
	DataOutputId       types.String `tfsdk:"data_output_id"`
	MonitoringOutputId types.String `tfsdk:"monitoring_output_id"`
	FleetServerHostId  types.String `tfsdk:"fleet_server_host_id"`
	DownloadSourceId   types.String `tfsdk:"download_source_id"`
	MonitorLogs        types.Bool   `tfsdk:"monitor_logs"`
	MonitorMetrics     types.Bool   `tfsdk:"monitor_metrics"`
	SysMonitoring      types.Bool   `tfsdk:"sys_monitoring"`
	SkipDestroy        types.Bool   `tfsdk:"skip_destroy"`
	GlobalDataTags     types.Map    `tfsdk:"global_data_tags"`
}

func (model *agentPolicyModel) populateFromAPI(data *fleetapi.AgentPolicy) diag.Diagnostics {
	if data == nil {
		return nil
	}

	model.ID = types.StringValue(data.Id)
	model.PolicyID = types.StringValue(data.Id)
	model.DataOutputId = types.StringPointerValue(data.DataOutputId)
	model.Description = types.StringPointerValue(data.Description)
	model.DownloadSourceId = types.StringPointerValue(data.DownloadSourceId)
	model.FleetServerHostId = types.StringPointerValue(data.FleetServerHostId)

	if data.MonitoringEnabled != nil {
		if slices.Contains(*data.MonitoringEnabled, fleetapi.AgentPolicyMonitoringEnabledLogs) {
			model.MonitorLogs = types.BoolValue(true)
		}
		if slices.Contains(*data.MonitoringEnabled, fleetapi.AgentPolicyMonitoringEnabledMetrics) {
			model.MonitorMetrics = types.BoolValue(true)
		}
	}
	if !utils.IsKnown(model.MonitorLogs) {
		model.MonitorLogs = types.BoolValue(false)
	}
	if !utils.IsKnown(model.MonitorLogs) {
		model.MonitorLogs = types.BoolValue(false)
	}

	model.MonitoringOutputId = types.StringPointerValue(data.MonitoringOutputId)
	model.Name = types.StringValue(data.Name)
	model.Namespace = types.StringValue(data.Namespace)

	if data.GlobalDataTags != nil && len(*data.GlobalDataTags) > 0 {
		gdt := map[string]attr.Value{}
		for _, tag := range *data.GlobalDataTags {
			name, err := tag["name"].AsAgentPolicyGlobalDataTags0()
			if err != nil {
				return diag.Diagnostics{
					diag.NewErrorDiagnostic(err.Error(), err.Error()),
				}
			}

			value, err := tag["value"].AsAgentPolicyGlobalDataTags0()
			if err != nil {
				return diag.Diagnostics{
					diag.NewErrorDiagnostic(err.Error(), err.Error()),
				}
			}

			gdt[name] = basetypes.NewStringValue(value)
		}
		tfGDT, diags := basetypes.NewMapValue(basetypes.StringType{}, gdt)
		if diags.HasError() {
			return diags
		}
		model.GlobalDataTags = tfGDT
	} else {
		model.GlobalDataTags = basetypes.NewMapNull(basetypes.StringType{})
	}

	return nil
}

func (model agentPolicyModel) toAPICreateModel() (fleetapi.AgentPolicyCreateRequest, error) {
	monitoring := make([]fleetapi.AgentPolicyCreateRequestMonitoringEnabled, 0, 2)
	if model.MonitorLogs.ValueBool() {
		monitoring = append(monitoring, fleetapi.AgentPolicyCreateRequestMonitoringEnabledLogs)
	}
	if model.MonitorMetrics.ValueBool() {
		monitoring = append(monitoring, fleetapi.AgentPolicyCreateRequestMonitoringEnabledMetrics)
	}

	globalDataTags, err := toAPIGlobalDataTags(
		model,
		func(key string) (fleetapi.AgentPolicyCreateRequest_GlobalDataTags_AdditionalProperties, error) {
			var name fleetapi.AgentPolicyCreateRequest_GlobalDataTags_AdditionalProperties
			err := name.FromAgentPolicyCreateRequestGlobalDataTags0(key)
			if err != nil {
				return fleetapi.AgentPolicyCreateRequest_GlobalDataTags_AdditionalProperties{}, fmt.Errorf("failed to parse global_data_tags key [%s]: %w", key, err)
			}

			return name, nil
		},
		func(stringVal basetypes.StringValue) (fleetapi.AgentPolicyCreateRequest_GlobalDataTags_AdditionalProperties, error) {
			val := fleetapi.AgentPolicyCreateRequest_GlobalDataTags_AdditionalProperties{}
			err := val.FromAgentPolicyCreateRequestGlobalDataTags0(stringVal.ValueString())
			if err != nil {
				return fleetapi.AgentPolicyCreateRequest_GlobalDataTags_AdditionalProperties{}, fmt.Errorf("failed to parse global_data_tags value [%s]: %w", stringVal.ValueString(), err)
			}
			return val, nil
		},
	)
	if err != nil {
		return fleetapi.AgentPolicyCreateRequest{}, err
	}

	body := fleetapi.AgentPolicyCreateRequest{
		DataOutputId:       model.DataOutputId.ValueStringPointer(),
		Description:        model.Description.ValueStringPointer(),
		DownloadSourceId:   model.DownloadSourceId.ValueStringPointer(),
		FleetServerHostId:  model.FleetServerHostId.ValueStringPointer(),
		Id:                 model.PolicyID.ValueStringPointer(),
		MonitoringEnabled:  &monitoring,
		MonitoringOutputId: model.MonitoringOutputId.ValueStringPointer(),
		Name:               model.Name.ValueString(),
		Namespace:          model.Namespace.ValueString(),
		GlobalDataTags:     globalDataTags,
	}

	return body, nil
}

func (model agentPolicyModel) toAPIUpdateModel() (fleetapi.AgentPolicyUpdateRequest, error) {
	monitoring := make([]fleetapi.AgentPolicyUpdateRequestMonitoringEnabled, 0, 2)
	if model.MonitorLogs.ValueBool() {
		monitoring = append(monitoring, fleetapi.Logs)
	}
	if model.MonitorMetrics.ValueBool() {
		monitoring = append(monitoring, fleetapi.Metrics)
	}

	globalDataTags, err := toAPIGlobalDataTags(
		model,
		func(key string) (fleetapi.AgentPolicyUpdateRequest_GlobalDataTags_AdditionalProperties, error) {
			var name fleetapi.AgentPolicyUpdateRequest_GlobalDataTags_AdditionalProperties
			err := name.FromAgentPolicyUpdateRequestGlobalDataTags0(key)
			if err != nil {
				return fleetapi.AgentPolicyUpdateRequest_GlobalDataTags_AdditionalProperties{}, fmt.Errorf("failed to parse global_data_tags key [%s]: %w", key, err)
			}

			return name, nil
		},
		func(stringVal basetypes.StringValue) (fleetapi.AgentPolicyUpdateRequest_GlobalDataTags_AdditionalProperties, error) {
			val := fleetapi.AgentPolicyUpdateRequest_GlobalDataTags_AdditionalProperties{}
			err := val.FromAgentPolicyUpdateRequestGlobalDataTags0(stringVal.ValueString())
			if err != nil {
				return fleetapi.AgentPolicyUpdateRequest_GlobalDataTags_AdditionalProperties{}, fmt.Errorf("failed to parse global_data_tags value [%s]: %w", stringVal.ValueString(), err)
			}
			return val, nil
		},
	)
	if err != nil {
		return fleetapi.AgentPolicyUpdateRequest{}, err
	}

	body := fleetapi.AgentPolicyUpdateRequest{
		DataOutputId:       model.DataOutputId.ValueStringPointer(),
		Description:        model.Description.ValueStringPointer(),
		DownloadSourceId:   model.DownloadSourceId.ValueStringPointer(),
		FleetServerHostId:  model.FleetServerHostId.ValueStringPointer(),
		MonitoringEnabled:  &monitoring,
		MonitoringOutputId: model.MonitoringOutputId.ValueStringPointer(),
		Name:               model.Name.ValueString(),
		Namespace:          model.Namespace.ValueString(),
		GlobalDataTags:     globalDataTags,
	}

	return body, nil
}

func toAPIGlobalDataTags[TValue any](
	model agentPolicyModel,
	makeName func(string) (TValue, error),
	makeValue func(types.String) (TValue, error),
) (*[]map[string]TValue, error) {
	if len(model.GlobalDataTags.Elements()) == 0 {
		return nil, nil
	}

	gdt := []map[string]TValue{}
	for key, value := range model.GlobalDataTags.Elements() {
		stringValue, ok := value.(types.String)
		if !ok {
			return nil, fmt.Errorf("failed to convert global data tag to string value. Expected String type but got %T", value)
		}
		name, err := makeName(key)
		if err != nil {
			return nil, fmt.Errorf("failed to parse global_data_tags key [%s]: %w", key, err)
		}

		val, err := makeValue(stringValue)
		if err != nil {
			return nil, fmt.Errorf("failed to parse global_data_tags value [%s]: %w", stringValue.ValueString(), err)
		}

		gdt = append(gdt, map[string]TValue{
			"name":  name,
			"value": val,
		})
	}

	return &gdt, nil
}
