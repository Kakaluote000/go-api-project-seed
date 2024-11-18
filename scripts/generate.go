package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"unicode"
)

const (
	modelTemplatePath      = "templates/model.tmpl"
	repositoryTemplatePath = "templates/repository.tmpl"
	serviceTemplatePath    = "templates/service.tmpl"
	apiTemplatePath        = "templates/api.tmpl"
	outputDir              = "internal/"
	projectName            = "go-api-project-seed"
)

// Database configuration
var (
	dbUser     = "we_rummy_user_test"
	dbPassword = "ZxcVYuAbcDE2134566"
	dbHost     = "165.232.190.202"
	dbPort     = "63306"
	dbName     = "we_rummy"
)

func main() {
	// Connect to MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Fetch table list
	tables, err := fetchTables(db)
	if err != nil {
		log.Fatalf("Failed to fetch tables: %v", err)
	}

	// Display available tables
	fmt.Println("Available tables:")
	for i, table := range tables {
		fmt.Printf("%d. %s\n", i+1, table)
	}

	// Allow user to select tables
	fmt.Println("\nEnter table names to generate code for (comma-separated):")
	var input string
	fmt.Scanln(&input)
	selectedTables := strings.Split(input, ",")

	// Generate code for selected tables
	for _, table := range selectedTables {
		table = strings.TrimSpace(table)
		if err := generateCode(db, table); err != nil {
			log.Printf("Failed to generate code for table %s: %v", table, err)
		} else {
			fmt.Printf("Code generation completed for table %s.\n", table)
		}
	}
}

func fetchTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func generateCode(db *sql.DB, tableName string) error {
	// Fetch table structure
	columns, err := fetchTableColumns(db, tableName)
	if err != nil {
		return fmt.Errorf("failed to fetch columns for table %s: %v", tableName, err)
	}

	// Generate structure name
	structName := pascalCase(tableName)
	// Render templates
	if err := renderTemplate("model.tmpl", modelTemplatePath, outputDir+"model/"+tableName+".go", map[string]interface{}{
		"StructName":  structName,
		"TableName":   tableName,
		"Columns":     columns,
		"ProjectName": projectName,
	}); err != nil {
		return err
	}
	if err := renderTemplate("repository.tmpl", repositoryTemplatePath, outputDir+"repository/"+tableName+"_repository.go", map[string]interface{}{
		"StructName":  structName,
		"TableName":   tableName,
		"ProjectName": projectName,
	}); err != nil {
		return err
	}
	if err := renderTemplate("service.tmpl", serviceTemplatePath, outputDir+"service/"+tableName+"_service.go", map[string]interface{}{
		"StructName":  structName,
		"TableName":   tableName,
		"ProjectName": projectName,
	}); err != nil {
		return err
	}
	if err := renderTemplate("api.tmpl", apiTemplatePath, outputDir+"api/v1/"+tableName+"_api.go", map[string]interface{}{
		"StructName":  structName,
		"TableName":   tableName,
		"ProjectName": projectName,
	}); err != nil {
		return err
	}

	return nil
}

// renderTemplate renders a template and outputs the result to a file.
func renderTemplate(fileName, templatePath, outputPath string, data map[string]interface{}) error {
	// 打开模板文件并注册 goType 和 pascalCase 函数
	tmpl := template.Must(template.New(fileName).Funcs(template.FuncMap{
		"toSnakeCase": toSnakeCase,
		"pascalCase":  pascalCase,
		"goType":      goType,
	}).ParseFiles(templatePath))

	// 创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	//defer outputFile.Close()
	//
	//// 渲染模板
	//return tmpl.Execute(outputFile, data)

	// 渲染模板
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	// 写入未格式化的内容
	if _, err := outputFile.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	// 使用 goimports 修正导入
	if err := formatImports(outputPath); err != nil {
		return fmt.Errorf("goimports failed: %v", err)
	}

	// 格式化文件
	if err := formatGoFile(outputPath); err != nil {
		return fmt.Errorf("failed to format file %s: %v", outputPath, err)
	}
	return nil
}

func fetchTableColumns(db *sql.DB, tableName string) ([]map[string]string, error) {
	query := fmt.Sprintf("DESCRIBE %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []map[string]string
	for rows.Next() {
		//var field, dataType, null, key, defaultValue, extra string
		var field, dataType, null, key, defaultValue, extra sql.NullString
		if err := rows.Scan(&field, &dataType, &null, &key, &defaultValue, &extra); err != nil {
			return nil, err
		}
		columns = append(columns, map[string]string{
			"Field": field.String,
			"Type":  dataType.String,
			//"Key":   key.String,
		})
	}
	return columns, nil
}

// 渲染模板
// goType converts database field types to Go types.
func goType(dbType string) string {
	dbType = strings.ToLower(dbType)

	switch dbType {
	case "int", "bigint", "smallint":
		return "int"
	case "varchar", "text":
		return "string"
	case "boolean", "bool":
		return "bool"
	case "float", "double", "decimal":
		return "float64"
	case "date", "datetime", "timestamp":
		return "time.Time"
	default:
		return "string" // 默认返回 string 类型
	}
}

// pascalCase converts snake_case to PascalCase.
func pascalCase(input string) string {
	parts := strings.Split(input, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

// toSnakeCase converts a string to snake_case.
func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

// formatGoFile 使用 gofmt 格式化生成的文件
func formatGoFile(filePath string) error {
	cmd := exec.Command("gofmt", "-w", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gofmt failed: %v", err)
	}
	return nil
}

func formatImports(filePath string) error {
	cmd := exec.Command("goimports", "-w", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("goimports failed: %v", err)
	}
	return nil
}
