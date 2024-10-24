package controllers

import (
	"context"
	"log"
	"os"

	"github.com/AlissonBarbosa/alfred/models"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v3"
)

func LoadConfig(filePath string) (*models.Clouds, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var clouds models.Clouds
	err = decoder.Decode(&clouds)
	if err != nil {
		return nil, err
	}
	return &clouds, nil
}

func GetOpenstackProvider(cloudName string) (*gophercloud.ProviderClient, error) {
	ctx := context.Background()
	clouds, err := LoadConfig("/etc/openstack/clouds.yaml")
	if err != nil {
		log.Fatalf("Error to load yaml file: %v", err)
	}

	cloud, exists := clouds.Clouds[cloudName]
	if !exists {
		log.Fatalf("Cloud '%s' not found", cloudName)
	}

	authOptions := gophercloud.AuthOptions{
		IdentityEndpoint: cloud.Auth.AuthURL,
		Username:         cloud.Auth.Username,
		Password:         cloud.Auth.Password,
		TenantID:         cloud.Auth.ProjectID,
		DomainName:       cloud.Auth.UserDomainName,
	}

	provider, err := openstack.AuthenticatedClient(ctx, authOptions)
	if err != nil {
		log.Fatalf("Error to autenticate: %v", err)
	}

	return provider, nil
}

func RenderTable(headers map[int]string, columns map[string][]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	customStyle := table.Style{
		Name: "Custom",
		Box: table.BoxStyle{
			BottomLeft:       "+",
			BottomRight:      "+",
			BottomSeparator:  "+",
			Left:             "|",
			LeftSeparator:    "+",
			MiddleHorizontal: "-",
			MiddleSeparator:  "+",
			MiddleVertical:   "|",
			Right:            "|",
			RightSeparator:   "+",
			TopLeft:          "+",
			TopRight:         "+",
			TopSeparator:     "+",
		},
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: true,
			SeparateRows:    true,
		},
	}

	t.SetStyle(customStyle)

	headerRow := table.Row{}
	for i := 0; i < len(headers); i++ {
		headerRow = append(headerRow, headers[i])
	}
	t.AppendHeader(headerRow)

	numRows := len(columns[headers[0]])

	for rowIdx := 0; rowIdx < numRows; rowIdx++ {
		row := table.Row{}
		for i := 0; i < len(headers); i++ {
			header := headers[i]
			row = append(row, columns[header][rowIdx])
		}
		t.AppendRow(row)
	}

	t.Render()
}
