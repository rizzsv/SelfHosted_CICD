package executor

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func RunJob(repoUrl string) {
	timeStamp := time.Now().Format("20060102-150405")
	cloneDir := fmt.Sprintf("tmp/%s", timeStamp)
	logFile := fmt.Sprintf("logs/ci-log-%s.log", timeStamp)
	os.MkdirAll(cloneDir, 0755)

	f, _ := os.Create(logFile)
	defer f.Close()

	log := func(s string) {
		fmt.Fprint(f, s)
		fmt.Print(s)
	}
	log("🔄 Cloning repo...")
	cmd := exec.Command("git", "clone", repoUrl, cloneDir)
	cmd.Stdout = f
	cmd.Stderr = f
	_ = cmd.Run()

	log("⚙️ Building project...")
	buildCmd := exec.Command("go", "build", ".")
	buildCmd.Dir = cloneDir
	buildCmd.Stdout = f
	buildCmd.Stderr = f
	_ = buildCmd.Run()

	log("🧪 Running tests...")
	testCmd := exec.Command("go", "test", "./...")
	testCmd.Dir = cloneDir
	testCmd.Stdout = f
	testCmd.Stderr = f
	_ = testCmd.Run()

	log("✅ CI job complete.")
}
