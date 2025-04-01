package modelEndpoint

import (
	"context"
	"log"
	"reflect"

	"github.com/e2eterraformprovider/terraform-provider-tir/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func DataSourceSKUPlansModelEndpoint() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"plans": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"name": { //
						Type:     schema.TypeString,
						Computed: true,
					},
					"cpu": { ///
						Type:     schema.TypeString,
						Computed: true,
					},
					"gpu": { //
						Type:     schema.TypeString,
						Computed: true,
					},
					"sku_type": { //
						Type:     schema.TypeString,
						Computed: true,
					},
					"unit_price": { //
						Type:     schema.TypeFloat,
						Computed: true,
					},
					"committed_days": { //
						Type:     schema.TypeInt,
						Computed: true,
					},
					"currency": { //
						Type:     schema.TypeString,
						Computed: true,
					},
					"memory": { //
						Type:     schema.TypeString,
						Computed: true,
					},
				}},
				Computed: true,
			},
			"active_iam": {
				Type:     schema.TypeString,
				Required: true,
				Description: "This is IAM number for a particular user.",
			},
			"framework" : {
				Type : schema.TypeString,
				Required: true,
				Description: "There are different frameworks for which you want to run inference.",
			},
		},
		ReadContext: dataSourcePlansModelEndpointRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}


func dataSourcePlansModelEndpointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	
	apiClient := m.(*client.Client)
	var diags diag.Diagnostics
	active_iam := d.Get("active_iam").(string)
	response, err := apiClient.GetPlansModelEndpoint(active_iam, d.Get("framework").(string))
	if err != nil {
		return diag.Errorf("Not able to find plans %s",err)
	}
	var plans []interface{}
	// Process the CPU data
	data := response["data"].(map[string]interface{})
	if cpuItems, ok := data["CPU"].([]interface{}); ok {
		for _, cpuItem := range cpuItems {
			cpuMap := cpuItem.(map[string]interface{})
			if plansList, ok := cpuMap["plans"].([]interface{}); ok {
				for _, planItem := range plansList {
					planMap := planItem.(map[string]interface{})
					plans = append(plans, map[string]interface{}{
						"name":           cpuMap["name"],
						"cpu":            cpuMap["cpu"],
						"gpu":            cpuMap["gpu"],
						"memory":         cpuMap["memory"],
						"sku_type":       planMap["sku_type"],
						"committed_days": planMap["committed_days"],
						"unit_price":     planMap["unit_price"],
						"currency":       planMap["currency"],
					})
				}
			}
		}
	}
	// Process the GPU data
	if gpuItems, ok := data["GPU"].([]interface{}); ok {
		for _, gpuItem := range gpuItems {
			gpuMap := gpuItem.(map[string]interface{})
			if plansList, ok := gpuMap["plans"].([]interface{}); ok {
				for _, planItem := range plansList {
					planMap := planItem.(map[string]interface{})
					plans = append(plans, map[string]interface{}{
						"name":           gpuMap["name"],
						"cpu":            gpuMap["cpu"],
						"gpu":            gpuMap["gpu"],
						"memory":         gpuMap["memory"],
						"sku_type":       planMap["sku_type"],
						"committed_days": planMap["committed_days"],
						"unit_price":     planMap["unit_price"],
						"currency":       planMap["currency"],
					})
				}
			}
		}
	}
	log.Println("here i am", plans)
	log.Println("type", reflect.TypeOf(plans))
	d.SetId("plans")
	d.Set("plans", plans)
	log.Println("d.get", d.Get("plans"))
	return diags
}
