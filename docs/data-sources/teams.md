---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tir_teams Data Source for TeamIDs - tir"
subcategory: ""
description: |-
  
---

# tir_teams (Data Source)


## Example Usage

``` hcl 
data "tir_teams" "teams" {
    
    active_iam = <active_iam : string>
}

```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- active_iam (String) This will be string value which you can find in state file after runnign data sources for iams.



### Optional


### Read-Only

- `id` (String) The ID of this resource.
- `teams` (List of teams) 
