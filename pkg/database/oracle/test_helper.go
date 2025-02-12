package oracle

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type OracleContainer struct {
	Host      string
	Port      string
	container testcontainers.Container
}

func SetupOracleContainer(t testing.TB) *OracleContainer {
	ctx := context.Background()
	absPath, err := filepath.Abs("../testdata/init.sql")
	if err != nil {
		panic(err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "gvenzl/oracle-free:23.5-slim",
		ExposedPorts: []string{"1521/tcp"},
		Env: map[string]string{
			"ORACLE_PASSWORD": "SuperPassword",
		},
		Files: []testcontainers.ContainerFile{
			{
				FileMode:          0o755,
				HostFilePath:      absPath,
				ContainerFilePath: "/docker-entrypoint-initdb.d/init.sql",
			},
		},
		WaitingFor: wait.ForLog("Completed: ALTER DATABASE OPEN"),
	}

	oracleContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	host, err := oracleContainer.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := oracleContainer.MappedPort(ctx, "1521")
	if err != nil {
		t.Fatal(err)
	}

	return &OracleContainer{
		Host:      host,
		Port:      port.Port(),
		container: oracleContainer,
	}
}

func (r *OracleContainer) Teardown(t testing.TB) {
	if err := r.container.Terminate(context.Background()); err != nil {
		t.Fatal(err)
	}
}
