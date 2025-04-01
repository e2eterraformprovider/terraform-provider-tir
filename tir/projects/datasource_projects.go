package projects

import (
	"context"
	
	"strconv"

	"github.com/e2eterraformprovider/terraform-provider-tir/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func DataSourceProjects() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"projects": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type : schema.TypeString,
				},
				Computed: true,
			},
			"team_id" : {
				Type : schema.TypeString,
				Required: true,
			},
			"active_iam" : {
				Type : schema.TypeString,
				Required: true,
			},
		},
		ReadContext: dataSourceProjects,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func dataSourceProjects (ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	activeIAM := d.Get("active_iam").(string)
	teamID := d.Get("team_id").(string)
	response, err := apiClient.GetProjects(activeIAM, teamID)
	if err != nil {
		return diag.Errorf("Not able to find projects %s",err)
	}
	data := response["data"].([]interface{})
	var projectList []string
	for _, item := range data {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			diag.Errorf("some problem occured please check the config file")
		}
		if id, ok := itemMap["project_id"].(float64); ok {
			projectList = append(projectList, strconv.Itoa(int(id)))
		}
	}
	d.Set("projects",projectList)
	d.SetId("projects" + teamID + activeIAM)

	return diags
}