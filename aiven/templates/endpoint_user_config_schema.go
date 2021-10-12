// Code generated by go generate; DO NOT EDIT.
// This file was generated at Fri Oct  8 19:54:10 EEST 2021

package templates

import "encoding/json"

var (
	// inline JSON was taken from a file:
	// integration_endpoints_user_config_schema.json
	endpointJSON = `{
  "datadog": {
    "additionalProperties": false,
    "properties": {
      "datadog_api_key": {
        "example": "848f30907c15c55d601fe45487cce9b6",
        "maxLength": 32,
        "minLength": 32,
        "pattern": "^[A-Za-z0-9]{32}$",
        "title": "Datadog API key",
        "type": "string",
        "user_error": "Must consist of alpha-numeric characters and contain 32 characters"
      },
      "datadog_tags": {
        "example": [
          {
            "tag": "foo"
          },
          {
            "comment": "Useful tag",
            "tag": "bar:buzz"
          }
        ],
        "items": {
          "properties": {
            "comment": {
              "example": "Used to tag primary replica metrics",
              "maxLength": 1024,
              "title": "Optional tag explanation",
              "type": "string"
            },
            "tag": {
              "description": "Tag format and usage are described here: https://docs.datadoghq.com/getting_started/tagging. Tags with prefix 'aiven-' are reserved for Aiven.",
              "example": "replica:primary",
              "maxLength": 200,
              "minLength": 1,
              "pattern": "^(?!aiven-)[^\\W\\d_](?:[:\\w./-]*[\\w./-])?$",
              "title": "Tag value",
              "type": "string",
              "user_error": "Tags must start with a letter and after that may contain the characters listed below:\nalphanumerics, underscores, minuses, colons, periods, slashes.\nA tag cannot end with a colon.\nTags can be up to 200 characters long and support Unicode.\nTags with prefix 'aiven-' are reserved for Aiven.\nMore info: https://docs.datadoghq.com/getting_started/tagging.\n"
            }
          },
          "required": [
            "tag"
          ],
          "title": "Datadog tag defined by user",
          "type": "object"
        },
        "maxItems": 32,
        "title": "Custom tags provided by user",
        "type": "array"
      },
      "disable_consumer_stats": {
        "example": true,
        "title": "Disable consumer group metrics",
        "type": "boolean"
      },
      "kafka_consumer_check_instances": {
        "example": 8,
        "maximum": 100,
        "minimum": 1,
        "title": "Number of separate instances to fetch kafka consumer statistics with",
        "type": "integer"
      },
      "kafka_consumer_stats_timeout": {
        "example": 60,
        "maximum": 600,
        "minimum": 2,
        "title": "Number of seconds that datadog will wait to get consumer statistics from brokers",
        "type": "integer"
      },
      "max_partition_contexts": {
        "example": 32000,
        "maximum": 200000,
        "minimum": 200,
        "title": "Maximum number of partition contexts to send",
        "type": "integer"
      },
      "site": {
        "enum": [
          "datadoghq.com",
          "datadoghq.eu"
        ],
        "example": "datadoghq.com",
        "title": "Datadog intake site. Defaults to datadoghq.com",
        "type": "string"
      }
    },
    "required": [
      "datadog_api_key"
    ],
    "type": "object"
  },
  "external_aws_cloudwatch_logs": {
    "additionalProperties": false,
    "properties": {
      "access_key": {
        "example": "AAAAAAAAAAAAAAAAAAAA",
        "maxLength": 4096,
        "title": "AWS access key. Required permissions are logs:CreateLogGroup, logs:CreateLogStream, logs:PutLogEvents and logs:DescribeLogStreams",
        "type": "string"
      },
      "log_group_name": {
        "example": "my-log-group",
        "maxLength": 512,
        "minLength": 1,
        "pattern": "^[\\.\\-_/#A-Za-z0-9]+$",
        "title": "AWS CloudWatch log group name",
        "type": "string"
      },
      "region": {
        "example": "us-east-1",
        "maxLength": 32,
        "title": "AWS region",
        "type": "string"
      },
      "secret_key": {
        "example": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
        "maxLength": 4096,
        "title": "AWS secret key",
        "type": "string"
      }
    },
    "required": [
      "access_key",
      "secret_key",
      "region"
    ],
    "type": "object"
  },
  "external_aws_cloudwatch_metrics": {
    "additionalProperties": false,
    "properties": {
      "access_key": {
        "example": "AAAAAAAAAAAAAAAAAAAA",
        "maxLength": 4096,
        "title": "AWS access key. Required permissions are cloudwatch:PutMetricData",
        "type": "string"
      },
      "namespace": {
        "example": "my-metrics-namespace",
        "maxLength": 255,
        "minLength": 1,
        "pattern": "^(?!(AWS/|:))[:\\.\\-_/#A-Za-z0-9]+$",
        "title": "AWS CloudWatch Metrics Namespace",
        "type": "string"
      },
      "region": {
        "example": "us-east-1",
        "maxLength": 32,
        "title": "AWS region",
        "type": "string"
      },
      "secret_key": {
        "example": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
        "maxLength": 4096,
        "title": "AWS secret key",
        "type": "string"
      }
    },
    "required": [
      "access_key",
      "secret_key",
      "region",
      "namespace"
    ],
    "type": "object"
  },
  "external_elasticsearch_logs": {
    "additionalProperties": false,
    "properties": {
      "ca": {
        "example": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
        "maxLength": 16384,
        "title": "PEM encoded CA certificate",
        "type": [
          "string",
          "null"
        ]
      },
      "index_days_max": {
        "default": 3,
        "example": 3,
        "maximum": 10000,
        "minimum": 1,
        "title": "Maximum number of days of logs to keep",
        "type": "integer"
      },
      "index_prefix": {
        "default": "logs",
        "example": "logs",
        "maxLength": 1000,
        "minLength": 1,
        "pattern": "^[a-z0-9][a-z0-9-_.]+$",
        "title": "Elasticsearch index prefix",
        "type": "string",
        "user_error": "Must start with alpha-numeric character, and only contain alpha-numeric characters, dashes, underscores and dots"
      },
      "timeout": {
        "default": 10,
        "example": 10,
        "maximum": 120,
        "minimum": 10,
        "title": "Elasticsearch request timeout limit",
        "type": "number"
      },
      "url": {
        "example": "https://user:passwd@logs.example.com/",
        "maxLength": 2048,
        "minLength": 12,
        "title": "Elasticsearch connection URL",
        "type": "string"
      }
    },
    "required": [
      "url",
      "index_prefix"
    ],
    "type": "object"
  },
  "external_google_cloud_logging": {
    "additionalProperties": false,
    "properties": {
      "log_id": {
        "example": "syslog",
        "maxLength": 512,
        "title": "Google Cloud Logging log id",
        "type": "string"
      },
      "project_id": {
        "example": "snappy-photon-12345",
        "maxLength": 30,
        "minLength": 6,
        "title": "GCP project id.",
        "type": "string"
      },
      "service_account_credentials": {
        "description": "This is a JSON object with the fields documented in https://cloud.google.com/iam/docs/creating-managing-service-account-keys .",
        "example": "{\"type\": \"service_account\", ...",
        "maxLength": 4096,
        "title": "Google Service Account Credentials",
        "type": "string"
      }
    },
    "required": [
      "project_id",
      "log_id",
      "service_account_credentials"
    ],
    "title": "User configuration for Google Cloud Logging integration",
    "type": "object"
  },
  "external_kafka": {
    "additionalProperties": false,
    "properties": {
      "bootstrap_servers": {
        "example": "10.0.0.1:9092",
        "maxLength": 256,
        "minLength": 3,
        "title": "Bootstrap servers",
        "type": "string"
      },
      "sasl_mechanism": {
        "enum": [
          "PLAIN"
        ],
        "example": "PLAIN",
        "title": "The list of SASL mechanisms enabled in the Kafka server.",
        "type": [
          "string",
          "null"
        ]
      },
      "sasl_plain_password": {
        "example": "admin",
        "maxLength": 256,
        "minLength": 1,
        "title": "Password for SASL PLAIN mechanism in the Kafka server.",
        "type": [
          "string",
          "null"
        ]
      },
      "sasl_plain_username": {
        "example": "admin",
        "maxLength": 256,
        "minLength": 1,
        "title": "Username for SASL PLAIN mechanism in the Kafka server.",
        "type": [
          "string",
          "null"
        ]
      },
      "security_protocol": {
        "enum": [
          "PLAINTEXT",
          "SSL",
          "SASL_PLAINTEXT",
          "SASL_SSL"
        ],
        "example": "PLAINTEXT",
        "title": "Security protocol",
        "type": "string"
      },
      "ssl_ca_cert": {
        "example": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
        "maxLength": 16384,
        "title": "PEM-encoded CA certificate",
        "type": [
          "string",
          "null"
        ]
      },
      "ssl_client_cert": {
        "example": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
        "maxLength": 16384,
        "title": "PEM-encoded client certificate",
        "type": [
          "string",
          "null"
        ]
      },
      "ssl_client_key": {
        "example": "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n",
        "maxLength": 16384,
        "title": "PEM-encoded client key",
        "type": [
          "string",
          "null"
        ]
      },
      "ssl_endpoint_identification_algorithm": {
        "enum": [
          "https",
          ""
        ],
        "example": "https",
        "title": "The endpoint identification algorithm to validate server hostname using server certificate.",
        "type": [
          "string",
          "null"
        ]
      }
    },
    "required": [
      "bootstrap_servers",
      "security_protocol"
    ],
    "type": "object"
  },
  "external_schema_registry": {
    "additionalProperties": false,
    "properties": {
      "authentication": {
        "enum": [
          "none",
          "basic"
        ],
        "example": "basic",
        "title": "Authentication method",
        "type": "string"
      },
      "basic_auth_password": {
        "example": "Zm9vYg==",
        "maxLength": 256,
        "title": "Basic authentication password",
        "type": [
          "string",
          "null"
        ]
      },
      "basic_auth_username": {
        "example": "avnadmin",
        "maxLength": 256,
        "title": "Basic authentication user name",
        "type": [
          "string",
          "null"
        ]
      },
      "url": {
        "example": "https://schema-registry.kafka.company.com:28419",
        "maxLength": 2048,
        "title": "Schema Registry URL",
        "type": "string"
      }
    },
    "required": [
      "url",
      "authentication"
    ],
    "type": "object"
  },
  "jolokia": {
    "additionalProperties": false,
    "properties": {
      "basic_auth_password": {
        "example": "yhfBNFii4C",
        "maxLength": 64,
        "minLength": 8,
        "title": "Jolokia basic authentication password",
        "type": "string"
      },
      "basic_auth_username": {
        "example": "jol48k51",
        "maxLength": 32,
        "minLength": 5,
        "pattern": "^[a-z0-9\\-@_]{5,32}$",
        "title": "Jolokia basic authentication username",
        "type": "string",
        "user_error": "Must only contain a-z and 0-9"
      }
    },
    "type": "object"
  },
  "prometheus": {
    "additionalProperties": false,
    "properties": {
      "basic_auth_password": {
        "example": "fhyFNBjj3R",
        "maxLength": 64,
        "minLength": 8,
        "title": "Prometheus basic authentication password",
        "type": "string"
      },
      "basic_auth_username": {
        "example": "prom4851",
        "maxLength": 32,
        "minLength": 5,
        "pattern": "^[a-z0-9\\-@_]{5,32}$",
        "title": "Prometheus basic authentication username",
        "type": "string",
        "user_error": "Must only contain a-z and 0-9"
      }
    },
    "type": "object"
  },
  "rsyslog": {
    "additionalProperties": false,
    "properties": {
      "ca": {
        "example": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
        "maxLength": 16384,
        "title": "PEM encoded CA certificate",
        "type": [
          "string",
          "null"
        ]
      },
      "cert": {
        "example": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
        "maxLength": 16384,
        "title": "PEM encoded client certificate",
        "type": [
          "string",
          "null"
        ]
      },
      "format": {
        "default": "rfc5424",
        "enum": [
          "rfc5424",
          "rfc3164",
          "custom"
        ],
        "example": "rfc5424",
        "title": "message format",
        "type": "string"
      },
      "key": {
        "example": "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n",
        "maxLength": 16384,
        "title": "PEM encoded client key",
        "type": [
          "string",
          "null"
        ]
      },
      "logline": {
        "example": "<%pri%>%timestamp:::date-rfc3339% %HOSTNAME% %app-name% %msg%",
        "maxLength": 512,
        "minLength": 1,
        "title": "custom syslog message format",
        "type": "string"
      },
      "port": {
        "default": 514,
        "example": 514,
        "maximum": 65535,
        "minimum": 1,
        "title": "rsyslog server port",
        "type": "integer"
      },
      "sd": {
        "example": "TOKEN tag=\"LiteralValue\"",
        "maxLength": 1024,
        "title": "Structured data block for log message",
        "type": [
          "string",
          "null"
        ]
      },
      "server": {
        "example": "logs.example.com",
        "maxLength": 255,
        "minLength": 4,
        "title": "rsyslog server IP address or hostname",
        "type": "string"
      },
      "tls": {
        "default": true,
        "example": true,
        "title": "Require TLS",
        "type": "boolean"
      }
    },
    "required": [
      "server",
      "port",
      "format",
      "tls"
    ],
    "type": "object"
  },
  "signalfx": {
    "additionalProperties": false,
    "properties": {
      "enabled_metrics": {
        "default": [
          "*.*"
        ],
        "example": [
          "load.*",
          "vmpage.*"
        ],
        "items": {
          "maxLength": 1024,
          "title": "metric to send (glob pattern)",
          "type": "string"
        },
        "maxItems": 256,
        "title": "list of metrics to send",
        "type": "array"
      },
      "signalfx_api_key": {
        "example": "848f30907c15c55d601fe45487cce9b6",
        "maxLength": 255,
        "minLength": 8,
        "title": "SignalFX API key",
        "type": "string"
      },
      "signalfx_realm": {
        "default": "us0",
        "example": "us0",
        "maxLength": 8,
        "minLength": 2,
        "title": "SignalFX realm",
        "type": "string"
      }
    },
    "required": [
      "signalfx_api_key",
      "signalfx_realm"
    ],
    "type": "object"
  }
}`
)

func init() {
	var endpointSchema map[string]interface{}
	if err := json.Unmarshal([]byte(endpointJSON), &endpointSchema); err != nil {
		panic("cannot unmarshal user configuration options endpoint JSON', error :" + err.Error())
	}
	userConfigSchemas["endpoint"] = endpointSchema
}
