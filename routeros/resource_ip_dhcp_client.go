package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceDhcpClient https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-DHCPClient
func ResourceDhcpClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-client"),
		MetaId:           PropId(Id),

		"add_default_route": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Whether to install default route in routing table received from DHCP server.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no", "special-classless"}, false),
		},
		"address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IP address and netmask, which is assigned to DHCP Client from the Server.",
		},
		"allow_reconfigure": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"check_gateway": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Method on how to check gateway reachability.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"default_route_distance": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1,
			Description:  "Distance of default route. Applicable if add-default-route is set to yes.",
			ValidateFunc: validation.IntBetween(0, 255),
			// Default route distance returns as empty when the dhcp-client is searching.
			// This produces inconsistent results, for this case, we will suppress changes.
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new || new == "" {
					return true
				}
				return false
			},
		},
		"default_route_tables": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Default route tables.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dhcp_options": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "hostname,clientid",
			Description: "Options that are sent to the DHCP server.",
		},
		"dhcp_server": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the DHCP server.",
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"expires_after": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A time when the lease expires (specified by the DHCP server).",
		},
		"gateway": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the gateway which is assigned by DHCP server.",
		},
		KeyInterface: PropInterfaceRw,
		KeyInvalid:   PropInvalidRo,
		"primary_dns": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the first DNS resolver, that was assigned by the DHCP server.",
		},
		"primary_ntp": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the primary NTP server, assigned by the DHCP server.",
		},
		"script": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A script.",
		},
		"secondary_dns": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the second DNS resolver, assigned by the DHCP server.",
		},
		"secondary_ntp": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the secondary NTP server, assigned by the DHCP server.",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"use_peer_dns": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
			Description: "Whether to accept the DNS settings advertised by DHCP Server (will override the settings " +
				"put in the /ip dns submenu).",
		},
		"use_peer_ntp": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
			Description: "Whether to accept the NTP settings advertised by DHCP Server (will override the settings " +
				"put in the /system ntp client submenu).",
		},
		"use_reconfigure": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Allow the server to send Reconfigure messages to clients, prompting them to renew or " +
				"update their configuration without waiting for their lease to expire.",
		},
	}
	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
