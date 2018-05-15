package main

import (
	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func connectToLXDserver() error {
	// Connect to LXD over the Unix socket
	c, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		return err
	}

	// Container creation request
	req := api.ContainersPost{
		Name: "madewithapi",
		Source: api.ContainerSource{
			Type:  "image",
			Alias: "my-image",
		},
	}

	// Get LXD to create the container (background operation)
	op, err := c.CreateContainer(req)
	if err != nil {
		return err
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return err
	}

	// Get LXD to start the container (background operation)
	reqState := api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}

	op, err = c.UpdateContainerState("madewithapi", reqState, "")
	if err != nil {
		return err
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return err
	}
	return err
}

func main() {
	connectToLXDserver()
}
