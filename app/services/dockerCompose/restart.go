package dockerCompose

import (
    "fmt"
    l "log/slog"
    "os/exec"
)

func Restart() {
    
    l.Debug("Issueing Down")
    cmd := "/usr/local/bin/docker-compose down"
    
    out, err := exec.Command("bash", "-c", cmd).Output()
    
    if err != nil {
        l.With("Error", err).Error("Error in docker-compose")
        return
    }
    
    l.Debug("Issueing Up")
    cmd = "/usr/local/bin/docker-compose up -d"
    
    out, err = exec.Command("bash", "-c", cmd).Output()
    
    if err != nil {
        l.With("Error", err).Error("Error in docker-compose")
        return
    }
    fmt.Println(string(out))
}
