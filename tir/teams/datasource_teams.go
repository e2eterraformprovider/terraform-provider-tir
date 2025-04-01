package teams

import (
	"context"
	
	"strconv"

	"github.com/e2eterraformprovider/terraform-provider-tir/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func DataSourceTeams() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"teams": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type : schema.TypeString,
				},
				Computed: true,
			},
			"active_iam" : {
				Type : schema.TypeString,
				Required: true,
			},
		},
		ReadContext: dataSourceTeams,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func dataSourceTeams (ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	activeIAM := d.Get("active_iam").(string)
	response, err := apiClient.GetTeams(activeIAM)
	if err != nil {
		return diag.Errorf("Not able to find teams %s",err)
	}
	data := response["data"].([]interface{})
	var teamsList []string
	for _, item := range data {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			diag.Errorf("some problem occured please check the config file")
		}
		if id, ok := itemMap["team_id"].(float64); ok {
			teamsList= append(teamsList, strconv.Itoa(int(id)))
		}
	}
	d.Set("teams",teamsList)
	d.SetId("teams" + activeIAM)

	return diags
}