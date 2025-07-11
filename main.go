package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templateFS embed.FS

const (
	dirPerm  = 0755
	filePerm = 0644
)

// TemplateData holds data for template processing
type TemplateData struct {
	ServiceName            string
	ServiceNameCapitalized string
	ModuleName             string
}

type FileTemplate struct {
	TemplatePath string
	OutputPath   string
}

func main() {
	if err := Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "golarsa",
	Short: "GoLarsa is a CLI tool for creating service folder structures",
	Long: `GoLarsa is a CLI application that helps you create standardized
service folder structures for your Go projects.`,
}

var createServiceCmd = &cobra.Command{
	Use:   "service [service-name]",
	Short: "Create a new service with standardized folder structure",
	Long: `Create a new service with the following folder structure:
pkg/services/<service-name>/di
pkg/services/<service-name>/handler
pkg/services/<service-name>/models
pkg/services/<service-name>/repo
pkg/services/<service-name>/<service-name>.go
pkg/services/<service-name>/repo/<service-name>.repo.go
pkg/services/<service-name>/models/<service-name>.models.go
pkg/services/<service-name>/handler/<service-name>.handler.go
pkg/services/<service-name>/di/di.go`,
	Args: cobra.ExactArgs(1),
	Run:  runCreateService,
}

func init() {
	rootCmd.AddCommand(createServiceCmd)
}

func runCreateService(cmd *cobra.Command, args []string) {
	if err := createService(args[0]); err != nil {
		log.Fatal(err)
	}
}

func createService(serviceName string) error {
	if err := checkProject(); err != nil {
		return err
	}

	serviceName = normalizeServiceName(serviceName)
	fmt.Printf("Creating service: %s\n", serviceName)

	moduleName, err := getModuleName()
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	servicePath := filepath.Join("pkg", "services", serviceName+"s")
	if err := createDirectoryStructure(servicePath); err != nil {
		return err
	}

	templateData := TemplateData{
		ServiceName:            serviceName,
		ServiceNameCapitalized: capitalize(serviceName),
		ModuleName:             moduleName,
	}

	fileTemplates := []FileTemplate{
		{"templates/services/service.go.tmpl", filepath.Join(servicePath, serviceName+"s.go")},
		{"templates/services/repo.go.tmpl", filepath.Join(servicePath, "repo", serviceName+"s.repo.go")},
		{"templates/services/models.go.tmpl", filepath.Join(servicePath, "models", serviceName+"s.models.go")},
		{"templates/services/handler.go.tmpl", filepath.Join(servicePath, "handler", serviceName+"s.handler.go")},
		{"templates/services/di.go.tmpl", filepath.Join(servicePath, "di", serviceName+"s.di.go")},
	}

	for _, ft := range fileTemplates {
		if err := createFileFromTemplate(ft.TemplatePath, ft.OutputPath, templateData); err != nil {
			return fmt.Errorf("failed to create %s: %w", ft.OutputPath, err)
		}
	}

	fmt.Printf("\n🎉 Service '%s' created successfully!\n", serviceName)
	fmt.Printf("📁 Service structure created at: %s\n", servicePath)
	return nil
}

func checkProject() error {
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return fmt.Errorf("cannot run this outside of a project\nInit functionality coming soon :)")
	}
	return nil
}

func getModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			moduleName := strings.TrimSpace(strings.TrimPrefix(line, "module"))
			if moduleName == "" {
				return "", fmt.Errorf("invalid module declaration in go.mod")
			}
			return moduleName, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading go.mod: %w", err)
	}

	return "", fmt.Errorf("module declaration not found in go.mod")
}

func createDirectoryStructure(basePath string) error {
	subdirs := []string{"di", "handler", "models", "repo"}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(basePath, subdir)
		if err := os.MkdirAll(dirPath, dirPerm); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dirPath, err)
		}
		fmt.Printf("✓ Created: %s\n", dirPath)
	}
	return nil
}

func createFileFromTemplate(templatePath, outputPath string, data TemplateData) error {
	templateContent, err := templateFS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading template file: %w", err)
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	var fileContent bytes.Buffer
	if err := tmpl.Execute(&fileContent, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	if err := os.WriteFile(outputPath, fileContent.Bytes(), filePerm); err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	fmt.Printf("✓ Created: %s\n", outputPath)
	return nil
}

func normalizeServiceName(name string) string {
	if len(name) > 0 && name[len(name)-1] == 's' {
		return name[:len(name)-1]
	}
	return name
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
