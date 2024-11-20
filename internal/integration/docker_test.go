package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	permissionsGrpc "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serviceImageName = "permissions-service:test"
	networkName      = "integration-test-network"
	postgresAlias    = "postgres"
)

type testContainer struct {
	container testcontainers.Container
	uri       string
	postgres  testcontainers.Container
	pgDSN     string
	network   *testcontainers.DockerNetwork
}

func setupNetwork(ctx context.Context) (*testcontainers.DockerNetwork, error) {
	net, err := network.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create network: %v", err)
	}
	return net, nil
}

func setupPostgres(ctx context.Context, net *testcontainers.DockerNetwork) (testcontainers.Container, string, error) {
	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "postgres:17-alpine",
			Env: map[string]string{
				"POSTGRES_USER":     "myuser",
				"POSTGRES_PASSWORD": "mypassword",
				"POSTGRES_DB":       "postgres",
			},
			ExposedPorts: []string{"5432/tcp"},
			Networks:     []string{net.Name},
			NetworkAliases: map[string][]string{
				net.Name: {postgresAlias},
			},
			WaitingFor: wait.ForAll(
				wait.ForLog("database system is ready to accept connections"),
				wait.ForListeningPort("5432/tcp"),
			),
		},
		Started: true,
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to start postgres container: %v", err)
	}

	// Use the network alias instead of host:port
	pgDSN := fmt.Sprintf("host=%s user=myuser password=mypassword dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Rome",
		postgresAlias)

	return pgContainer, pgDSN, nil
}

func setupTestContainer(t *testing.T) (*testContainer, error) {
	ctx := context.Background()

	// Setup network first
	net, err := setupNetwork(ctx)
	if err != nil {
		return nil, err
	}

	// Setup PostgreSQL
	pgContainer, pgDSN, err := setupPostgres(ctx, net)
	if err != nil {
		net.Remove(ctx)
		return nil, err
	}

	req := testcontainers.ContainerRequest{
		Image:        serviceImageName,
		ExposedPorts: []string{"50051/tcp"},
		Networks:     []string{net.Name},
		WaitingFor:   wait.ForListeningPort("50051/tcp"),
		Env: map[string]string{
			"ENV":          "test",
			"POSTGRES_DSN": pgDSN,
			"GRPC_PORT":    "50051",
			"LOG_LEVEL":    "debug",
			"LOG_COLOR":    "true",
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		pgContainer.Terminate(ctx)
		net.Remove(ctx)
		return nil, fmt.Errorf("failed to start container: %v", err)
	}

	// Get the container's host and port
	host, err := container.Host(ctx)
	if err != nil {
		pgContainer.Terminate(ctx)
		container.Terminate(ctx)
		net.Remove(ctx)
		return nil, fmt.Errorf("failed to get container host: %v", err)
	}

	mappedPort, err := container.MappedPort(ctx, "50051/tcp")
	if err != nil {
		pgContainer.Terminate(ctx)
		container.Terminate(ctx)
		net.Remove(ctx)
		return nil, fmt.Errorf("failed to get container port: %v", err)
	}

	uri := fmt.Sprintf("%s:%s", host, mappedPort.Port())

	return &testContainer{
		container: container,
		uri:       uri,
		postgres:  pgContainer,
		pgDSN:     pgDSN,
		network:   net,
	}, nil
}

type IntegrationTestSuite struct {
	suite.Suite
	container *testContainer
	client    permissionsGrpc.PermissionsServiceClient
	conn      *grpc.ClientConn
}

func (s *IntegrationTestSuite) SetupSuite() {
	container, err := setupTestContainer(s.T())
	require.NoError(s.T(), err)
	s.container = container

	// Setup gRPC client
	conn, err := grpc.Dial(container.uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(s.T(), err)
	s.conn = conn
	s.client = permissionsGrpc.NewPermissionsServiceClient(conn)

	// Give the service some time to fully initialize
	time.Sleep(2 * time.Second)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	ctx := context.Background()
	if s.conn != nil {
		s.conn.Close()
	}
	if s.container != nil {
		if err := s.container.container.Terminate(ctx); err != nil {
			s.T().Logf("failed to terminate service container: %v", err)
		}
		if err := s.container.postgres.Terminate(ctx); err != nil {
			s.T().Logf("failed to terminate postgres container: %v", err)
		}
		if err := s.container.network.Remove(ctx); err != nil {
			s.T().Logf("failed to remove network: %v", err)
		}
	}
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
	suite.Run(t, new(IntegrationTestSuite))
}

// Test user role management
func (s *IntegrationTestSuite) TestUserRoles() {
	ctx := context.Background()
	testUser := "test-user"
	testRole := "test-role"

	// Test SetUserRole
	setResp, err := s.client.SetUserRole(ctx, &permissionsGrpc.SetUserRoleRequest{
		User: testUser,
		Role: testRole,
	})
	s.Require().NoError(err)
	s.True(setResp.Affected)
	s.Contains(setResp.Roles, testRole)

	// Test GetUserRoles
	getRolesResp, err := s.client.GetUserRoles(ctx, &permissionsGrpc.GetUserRolesRequest{
		User: testUser,
	})
	s.Require().NoError(err)
	s.Contains(getRolesResp.Roles, testRole)

	// Test UnsetUserRole
	unsetResp, err := s.client.UnsetUserRole(ctx, &permissionsGrpc.UnsetUserRoleRequest{
		User: testUser,
		Role: testRole,
	})
	s.Require().NoError(err)
	s.True(unsetResp.Affected)
	s.NotContains(unsetResp.Roles, testRole)
}

// Test scope management
func (s *IntegrationTestSuite) TestScopeManagement() {
	ctx := context.Background()
	testUser := "test-user"
	testRole := "test-role"
	testScope := "test-scope"

	// Test global scope
	_, err := s.client.SetGlobalScope(ctx, &permissionsGrpc.SetGlobalScopeRequest{
		Scope:   testScope,
		Allowed: true,
	})
	s.Require().NoError(err)

	globalAllowed, err := s.client.GlobalAllowed(ctx, &permissionsGrpc.GlobalAllowedRequest{
		Scope: testScope,
	})
	s.Require().NoError(err)
	s.True(globalAllowed.Allowed)

	// Test role scope
	_, err = s.client.SetRoleScope(ctx, &permissionsGrpc.SetRoleScopeRequest{
		Role:    testRole,
		Scope:   testScope,
		Allowed: true,
	})
	s.Require().NoError(err)

	roleAllowed, err := s.client.RoleAllowed(ctx, &permissionsGrpc.RoleAllowedRequest{
		Role:  testRole,
		Scope: testScope,
	})
	s.Require().NoError(err)
	s.True(roleAllowed.Allowed)

	// Test user scope
	_, err = s.client.SetUserScope(ctx, &permissionsGrpc.SetUserScopeRequest{
		User:    testUser,
		Scope:   testScope,
		Allowed: true,
	})
	s.Require().NoError(err)

	userAllowed, err := s.client.UserAllowed(ctx, &permissionsGrpc.UserAllowedRequest{
		User:  testUser,
		Scope: testScope,
	})
	s.Require().NoError(err)
	s.True(userAllowed.Allowed)

	// Test getting scopes
	globalScopes, err := s.client.GetGlobalScopes(ctx, &permissionsGrpc.GetGlobalScopesRequest{
		ScopePattern: testScope,
	})
	s.Require().NoError(err)
	s.Len(globalScopes.Scopes, 1)
	s.Equal(testScope, globalScopes.Scopes[0].Scope)
	s.True(globalScopes.Scopes[0].Allowed)

	roleScopes, err := s.client.GetRoleScopes(ctx, &permissionsGrpc.GetRoleScopesRequest{
		Role:         testRole,
		ScopePattern: testScope,
	})
	s.Require().NoError(err)
	s.Len(roleScopes.Scopes, 1)
	s.Equal(testScope, roleScopes.Scopes[0].Scope)
	s.True(roleScopes.Scopes[0].Allowed)

	userScopes, err := s.client.GetUserScopes(ctx, &permissionsGrpc.GetUserScopesRequest{
		User:         testUser,
		ScopePattern: testScope,
	})
	s.Require().NoError(err)
	s.Len(userScopes.Scopes, 1)
	s.Equal(testScope, userScopes.Scopes[0].Scope)
	s.True(userScopes.Scopes[0].Allowed)
}
