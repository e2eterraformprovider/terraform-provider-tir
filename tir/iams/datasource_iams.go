package iams

import (
	"context"
	
	"strconv"

	"github.com/e2eterraformprovider/terraform-provider-tir/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func DataSourceIAMS() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iams": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type : schema.TypeString,
				},
				Computed: true,
			},
		},
		ReadContext: dataSourceIAMS,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func dataSourceIAMS (ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	response, err := apiClient.GetIAMS()
	if err != nil {
		return diag.Errorf("Not able to find plans %s",err)
	}
	data := response["data"].([]interface{})
	var activeIams []string
	for _, item := range data {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			diag.Errorf("some problem occured please check apikey or auth_token")
		}
		if id, ok := itemMap["id"].(float64); ok {
			activeIams = append(activeIams, strconv.Itoa(int(id)))
		}
	}
	d.Set("iams",activeIams)
	d.SetId("iams")

	return diags
}