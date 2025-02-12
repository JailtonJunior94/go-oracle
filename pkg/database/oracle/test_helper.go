package oracle

import (
	"context"
	"path"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestOracleDBWithContainer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "gvenzl/oracle-free:23.5-slim",
		ExposedPorts: []string{"1521/tcp"},
		Env: map[string]string{
			"ORACLE_PASSWORD": "SuperPassword",
		},
		Mounts: testcontainers.ContainerMounts{
			{
				Source: testcontainers.DockerVolumeMountSource{
					Name: path.Join("..", "scripts", "init.sql"),
				},
				Target: "/docker-entrypoint-initdb.d",
			},
		},
		WaitingFor: wait.ForLog("Database ready to use. Enjoy!").WithStartupTimeout(120),
	}

	oracleContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer oracleContainer.Terminate(ctx)

	if err != nil {
		t.Fatal(err)
	}
}
