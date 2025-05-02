executor.go         # Spawns and manages docker containers
    │   │
    │   │   package docker
    │   │   import (
    │   │       "os/exec"
    │   │   )
    │   │   func RunInContainer(image string, command []string) error {
    │   │       cmd := exec.Command("docker", append([]string{"run", "--rm", image}, command...)...)
    │   │       return cmd.Run()
    │   │   }