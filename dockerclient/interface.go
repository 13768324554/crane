package dockerclient

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	goclient "github.com/fsouza/go-dockerclient"
)

type DockerClientInterface interface {
	Ping() error

	NodeList(opts types.NodeListOptions) ([]swarm.Node, error)
	NodeInspect(nodeId string) (swarm.Node, error)
	NodeRemove(nodeId string) error

	InspectContainer(id string) (*goclient.Container, error)
	ListContainers(opts goclient.ListContainersOptions) ([]goclient.APIContainers, error)
	RemoveContainer(opts goclient.RemoveContainerOptions) error

	ConnectNetwork(id string, opts goclient.NetworkConnectionOptions) error
	CreateNetwork(opts goclient.CreateNetworkOptions) (*goclient.Network, error)
	DisconnectNetwork(id string, opts goclient.NetworkConnectionOptions) error
	InspectNetwork(id string) (*goclient.Network, error)
	ListNetworks(opts goclient.NetworkFilterOpts) ([]goclient.Network, error)
	RemoveNetwork(id string) error
}