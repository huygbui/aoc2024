package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

const (
	AOC_SESSION_ENV = "AOC_SESSION"
	AOC_URL = "https://adventofcode.com/2024/day"
)

func getSessionToken() (string, error){
	_, currentFile, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(currentFile), "..", "")

	_ = godotenv.Load(filepath.Join(rootDir, ".env"))
	
	tok := os.Getenv(AOC_SESSION_ENV)
	if tok == "" {
		return "", fmt.Errorf("AOC_SESSION environment is not set")
	}
	return tok, nil
}

func GetInput(day int) (string, error) {
	if err := ensureInputDir(); err != nil {
		return "", err
	}

	inputPath := getInputPath(day)	
	if fileExists(inputPath) {
		content, err := os.ReadFile(inputPath)
		if err != nil {
			return "", err
		}
		return string(content), nil
	}

	content, err := downloadInput(day)
	if err != nil {
        return "", err
    }

	err = os.WriteFile(inputPath, content, 0644)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func getInputPath(day int) string {
	_, currentFile, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(currentFile), "..", "")
	
	return filepath.Join(rootDir, "inputs", fmt.Sprintf("day%02d.txt", day))
}

func ensureInputDir() error {
	return os.MkdirAll("inputs", 0755)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func downloadInput(day int) ([]byte, error) {
	url := fmt.Sprintf("%s/%d/input", AOC_URL, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return nil, err
	}
	
	tok, err := getSessionToken()
	if err != nil{
		return nil, err
	}
	
	cookie := &http.Cookie{
		Name: "session",
		Value: tok,
	}

	req.AddCookie(cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}