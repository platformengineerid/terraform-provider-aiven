package clickhouse

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aiven/aiven-go-client/v2"
)

func isUnknownRole(err error) bool {
	aivenError, ok := err.(aiven.Error)
	if !ok {
		return false
	}
	return strings.Contains(aivenError.Message, "Code: 511")
}

func CreateRole(ctx context.Context, client *aiven.Client, projectName, serviceName, roleName string) error {
	query := createRoleStatement(roleName)

	log.Println("[DEBUG] Clickhouse: create role query: ", query)

	// TODO inspect result?
	_, err := client.ClickHouseQuery.Query(ctx, projectName, serviceName, defaultDatabase, query)
	return err
}

func RoleExists(ctx context.Context, client *aiven.Client, projectName, serviceName, roleName string) (bool, error) {
	query := showCreateRoleStatement(roleName)

	log.Println("[DEBUG] Clickhouse: role exists query: ", query)

	r, err := client.ClickHouseQuery.Query(ctx, projectName, serviceName, defaultDatabase, query)
	if err != nil {
		if isUnknownRole(err) {
			return false, nil
		}
		return false, err
	}
	return len(r.Data) > 0, nil
}

func DropRole(ctx context.Context, client *aiven.Client, projectName, serviceName, roleName string) error {
	query := dropRoleStatement(roleName)

	log.Println("[DEBUG] Clickhouse: drop role query: ", query)

	_, err := client.ClickHouseQuery.Query(ctx, projectName, serviceName, defaultDatabase, query)
	if err != nil && isUnknownRole(err) {
		return nil
	}
	return err
}

func createRoleStatement(roleName string) string {
	return fmt.Sprintf("CREATE ROLE IF NOT EXISTS %s", escape(roleName))
}

func dropRoleStatement(roleName string) string {
	return fmt.Sprintf("DROP ROLE IF EXISTS %s", escape(roleName))
}

func showCreateRoleStatement(roleName string) string {
	return fmt.Sprintf("SHOW CREATE ROLE %s", escape(roleName))
}
