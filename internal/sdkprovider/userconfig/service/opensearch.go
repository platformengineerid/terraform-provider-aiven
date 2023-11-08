// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func opensearchUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Opensearch user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"additional_backup_regions": {
				Description: "Additional Cloud Regions for Backup Replication.",
				Elem: &schema.Schema{
					Description: "Target cloud.",
					Type:        schema.TypeString,
				},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"custom_domain": {
				Description: "Serve the web frontend using a custom CNAME pointing to the Aiven DNS name.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"disable_replication_factor_adjustment": {
				Description: "Disable automatic replication factor adjustment for multi-node services. By default, Aiven ensures all indexes are replicated at least to two nodes. Note: Due to potential data loss in case of losing a service node, this setting can no longer be activated.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"index_patterns": {
				Description: "Index patterns",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"max_index_count": {
						Description: "Maximum number of indexes to keep.",
						Required:    true,
						Type:        schema.TypeInt,
					},
					"pattern": {
						Description: "fnmatch pattern.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"sorting_algorithm": {
						Default:      "creation_date",
						Description:  "Deletion sorting algorithm. The default value is `creation_date`.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"alphabetical", "creation_date"}, false),
					},
				}},
				MaxItems: 512,
				Optional: true,
				Type:     schema.TypeList,
			},
			"index_template": {
				Description: "Template settings for all new indexes",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"mapping_nested_objects_limit": {
						Description: "The maximum number of nested JSON objects that a single document can contain across all nested types. This limit helps to prevent out of memory errors when a document contains too many nested objects. Default is 10000.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"number_of_replicas": {
						Description: "The number of replicas each primary shard has.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"number_of_shards": {
						Description: "The number of primary shards that an index should have.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"ip_filter": {
				Deprecated:  "Deprecated. Use `ip_filter_string` instead.",
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"ip_filter_object": {
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"description": {
						Description: "Description for IP filter list entry.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"network": {
						Description: "CIDR address block.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeList,
			},
			"ip_filter_string": {
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"keep_index_refresh_interval": {
				Description: "Aiven automation resets index.refresh_interval to default value for every index to be sure that indices are always visible to search. If it doesn't fit your case, you can disable this by setting up this flag to true.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"max_index_count": {
				Default:     0,
				Description: "use index_patterns instead. The default value is `0`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"openid": {
				Description: "OpenSearch OpenID Connect Configuration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"client_id": {
						Description: "The ID of the OpenID Connect client configured in your IdP. Required.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"client_secret": {
						Description: "The client secret of the OpenID Connect client configured in your IdP. Required.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"connect_url": {
						Description: "The URL of your IdP where the Security plugin can find the OpenID Connect metadata/configuration settings.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"enabled": {
						Default:     true,
						Description: "Enables or disables OpenID Connect authentication for OpenSearch. When enabled, users can authenticate using OpenID Connect with an Identity Provider. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"header": {
						Default:     "Authorization",
						Description: "HTTP header name of the JWT token. Optional. Default is Authorization. The default value is `Authorization`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"jwt_header": {
						Description: "The HTTP header that stores the token. Typically the Authorization header with the Bearer schema: Authorization: Bearer <token>. Optional. Default is Authorization.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"jwt_url_parameter": {
						Description: "If the token is not transmitted in the HTTP header, but as an URL parameter, define the name of the parameter here. Optional.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"refresh_rate_limit_count": {
						Default:     10,
						Description: "The maximum number of unknown key IDs in the time frame. Default is 10. Optional. The default value is `10`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"refresh_rate_limit_time_window_ms": {
						Default:     10000,
						Description: "The time frame to use when checking the maximum number of unknown key IDs, in milliseconds. Optional.Default is 10000 (10 seconds). The default value is `10000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"roles_key": {
						Description: "The key in the JSON payload that stores the user’s roles. The value of this key must be a comma-separated list of roles. Required only if you want to use roles in the JWT.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"scope": {
						Description: "The scope of the identity token issued by the IdP. Optional. Default is openid profile email address phone.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"subject_key": {
						Description: "The key in the JSON payload that stores the user’s name. If not defined, the subject registered claim is used. Most IdP providers use the preferred_username claim. Optional.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"opensearch": {
				Description: "OpenSearch settings",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"action_auto_create_index_enabled": {
						Description: "Explicitly allow or block automatic creation of indices. Defaults to true.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"action_destructive_requires_name": {
						Description: "Require explicit index names when deleting.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"auth_failure_listeners": {
						Description: "Opensearch Security Plugin Settings",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"internal_authentication_backend_limiting": {
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"allowed_tries": {
										Description: "The number of login attempts allowed before login is blocked.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"authentication_backend": {
										Description:  "internal_authentication_backend_limiting.authentication_backend.",
										Optional:     true,
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"internal"}, false),
									},
									"block_expiry_seconds": {
										Description: "The duration of time that login remains blocked after a failed login.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"max_blocked_clients": {
										Description: "internal_authentication_backend_limiting.max_blocked_clients.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"max_tracked_clients": {
										Description: "The maximum number of tracked IP addresses that have failed login.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"time_window_seconds": {
										Description: "The window of time in which the value for `allowed_tries` is enforced.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"type": {
										Description:  "internal_authentication_backend_limiting.type.",
										Optional:     true,
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"username"}, false),
									},
								}},
								MaxItems: 1,
								Optional: true,
								Type:     schema.TypeList,
							},
							"ip_rate_limiting": {
								Description: "IP address rate limiting settings",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"allowed_tries": {
										Description: "The number of login attempts allowed before login is blocked.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"block_expiry_seconds": {
										Description: "The duration of time that login remains blocked after a failed login.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"max_blocked_clients": {
										Description: "The maximum number of blocked IP addresses.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"max_tracked_clients": {
										Description: "The maximum number of tracked IP addresses that have failed login.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"time_window_seconds": {
										Description: "The window of time in which the value for `allowed_tries` is enforced.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"type": {
										Description:  "The type of rate limiting.",
										Optional:     true,
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"ip"}, false),
									},
								}},
								MaxItems: 1,
								Optional: true,
								Type:     schema.TypeList,
							},
						}},
						MaxItems: 1,
						Optional: true,
						Type:     schema.TypeList,
					},
					"cluster_max_shards_per_node": {
						Description: "Controls the number of shards allowed in the cluster per data node.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"cluster_routing_allocation_node_concurrent_recoveries": {
						Description: "How many concurrent incoming/outgoing shard recoveries (normally replicas) are allowed to happen on a node. Defaults to 2.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"email_sender_name": {
						Description: "Sender name placeholder to be used in Opensearch Dashboards and Opensearch keystore.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"email_sender_password": {
						Description: "Sender password for Opensearch alerts to authenticate with SMTP server.",
						Optional:    true,
						Sensitive:   true,
						Type:        schema.TypeString,
					},
					"email_sender_username": {
						Description: "Sender username for Opensearch alerts.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"http_max_content_length": {
						Description: "Maximum content length for HTTP requests to the OpenSearch HTTP API, in bytes.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"http_max_header_size": {
						Description: "The max size of allowed headers, in bytes.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"http_max_initial_line_length": {
						Description: "The max length of an HTTP URL, in bytes.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_fielddata_cache_size": {
						Description: "Relative amount. Maximum amount of heap memory used for field data cache. This is an expert setting; decreasing the value too much will increase overhead of loading field data; too much memory used for field data cache will decrease amount of heap available for other operations.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_memory_index_buffer_size": {
						Description: "Percentage value. Default is 10%. Total amount of heap used for indexing buffer, before writing segments to disk. This is an expert setting. Too low value will slow down indexing; too high value will increase indexing performance but causes performance issues for query performance.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_memory_max_index_buffer_size": {
						Description: "Absolute value. Default is unbound. Doesn't work without indices.memory.index_buffer_size. Maximum amount of heap used for query cache, an absolute indices.memory.index_buffer_size maximum hard limit.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_memory_min_index_buffer_size": {
						Description: "Absolute value. Default is 48mb. Doesn't work without indices.memory.index_buffer_size. Minimum amount of heap used for query cache, an absolute indices.memory.index_buffer_size minimal hard limit.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_queries_cache_size": {
						Description: "Percentage value. Default is 10%. Maximum amount of heap used for query cache. This is an expert setting. Too low value will decrease query performance and increase performance for other operations; too high value will cause issues with other OpenSearch functionality.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_query_bool_max_clause_count": {
						Description: "Maximum number of clauses Lucene BooleanQuery can have. The default value (1024) is relatively high, and increasing it may cause performance issues. Investigate other approaches first before increasing this value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_recovery_max_bytes_per_sec": {
						Description: "Limits total inbound and outbound recovery traffic for each node. Applies to both peer recoveries as well as snapshot recoveries (i.e., restores from a snapshot). Defaults to 40mb.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_recovery_max_concurrent_file_chunks": {
						Description: "Number of file chunks sent in parallel for each recovery. Defaults to 2.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ism_enabled": {
						Default:     true,
						Description: "Specifies whether ISM is enabled or not. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"ism_history_enabled": {
						Default:     true,
						Description: "Specifies whether audit history is enabled or not. The logs from ISM are automatically indexed to a logs document. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"ism_history_max_age": {
						Default:     24,
						Description: "The maximum age before rolling over the audit history index in hours. The default value is `24`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ism_history_max_docs": {
						Default:     2500000,
						Description: "The maximum number of documents before rolling over the audit history index. The default value is `2500000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ism_history_rollover_check_period": {
						Default:     8,
						Description: "The time between rollover checks for the audit history index in hours. The default value is `8`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ism_history_rollover_retention_period": {
						Default:     30,
						Description: "How long audit history indices are kept in days. The default value is `30`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"override_main_response_version": {
						Description: "Compatibility mode sets OpenSearch to report its version as 7.10 so clients continue to work. Default is false.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"reindex_remote_whitelist": {
						Description: "Whitelisted addresses for reindexing. Changing this value will cause all OpenSearch instances to restart.",
						Elem: &schema.Schema{
							Description: "Address (hostname:port or IP:port).",
							Type:        schema.TypeString,
						},
						MaxItems: 32,
						Optional: true,
						Type:     schema.TypeSet,
					},
					"script_max_compilations_rate": {
						Description: "Script compilation circuit breaker limits the number of inline script compilations within a period of time. Default is use-context.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"search_max_buckets": {
						Description: "Maximum number of aggregation buckets allowed in a single response. OpenSearch default value is used when this is not defined.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_analyze_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_analyze_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_force_merge_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_get_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_get_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_search_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_search_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_search_throttled_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_search_throttled_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_write_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_write_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"opensearch_dashboards": {
				Description: "OpenSearch Dashboards settings",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"enabled": {
						Default:     true,
						Description: "Enable or disable OpenSearch Dashboards. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"max_old_space_size": {
						Default:     128,
						Description: "Limits the maximum amount of memory (in MiB) the OpenSearch Dashboards process can use. This sets the max_old_space_size option of the nodejs running the OpenSearch Dashboards. Note: the memory reserved by OpenSearch Dashboards is not available for OpenSearch. The default value is `128`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"opensearch_request_timeout": {
						Default:     30000,
						Description: "Timeout in milliseconds for requests made by OpenSearch Dashboards towards OpenSearch. The default value is `30000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"opensearch_version": {
				Description:  "OpenSearch major version.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"opensearch": {
						Description: "Allow clients to connect to opensearch with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"opensearch_dashboards": {
						Description: "Allow clients to connect to opensearch_dashboards with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"privatelink_access": {
				Description: "Allow access to selected service components through Privatelink",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"opensearch": {
						Description: "Enable opensearch.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"opensearch_dashboards": {
						Description: "Enable opensearch_dashboards.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Enable prometheus.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"project_to_fork_from": {
				Description: "Name of another project to fork a service from. This has effect only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"public_access": {
				Description: "Allow access to selected service ports from the public Internet",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"opensearch": {
						Description: "Allow clients to connect to opensearch from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"opensearch_dashboards": {
						Description: "Allow clients to connect to opensearch_dashboards from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"recovery_basebackup_name": {
				Description: "Name of the basebackup to restore in forked service.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"saml": {
				Description: "OpenSearch SAML configuration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"enabled": {
						Description: "Enables or disables SAML-based authentication for OpenSearch. When enabled, users can authenticate using SAML with an Identity Provider. The default value is `true`.",
						Required:    true,
						Type:        schema.TypeBool,
					},
					"idp_entity_id": {
						Description: "The unique identifier for the Identity Provider (IdP) entity that is used for SAML authentication. This value is typically provided by the IdP.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"idp_metadata_url": {
						Description: "The URL of the SAML metadata for the Identity Provider (IdP). This is used to configure SAML-based authentication with the IdP.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"idp_pemtrustedcas_content": {
						Description: "This parameter specifies the PEM-encoded root certificate authority (CA) content for the SAML identity provider (IdP) server verification. The root CA content is used to verify the SSL/TLS certificate presented by the server.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"roles_key": {
						Description: "Optional. Specifies the attribute in the SAML response where role information is stored, if available. Role attributes are not required for SAML authentication, but can be included in SAML assertions by most Identity Providers (IdPs) to determine user access levels or permissions.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"sp_entity_id": {
						Description: "The unique identifier for the Service Provider (SP) entity that is used for SAML authentication. This value is typically provided by the SP.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"subject_key": {
						Description: "Optional. Specifies the attribute in the SAML response where the subject identifier is stored. If not configured, the NameID attribute is used by default.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"service_to_fork_from": {
				Description: "Name of another service to fork from. This has effect only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"static_ips": {
				Description: "Use static public IP addresses.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
