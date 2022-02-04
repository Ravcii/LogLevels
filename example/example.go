package main

import log "github.com/Ravcii/LogLevels"

func main() {
	log.SetLevel(log.ERROR)
	log.Debug("Debug test")
	log.Info("Info test")
	log.Warning("Warning test")
	log.Error("Error test")

	log.Errorf("Test %s, id: %d", "string", -50)
}
