package controllers

import (
	"context"
	"log"
	"strings"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/aggregates"
)

func ListAggregates(cloudName string) (map[int]string, map[string][]string, error) {
	provider, err := GetOpenstackProvider(cloudName)
	if err != nil {
		return nil, nil, err
	}

	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("Erro ao criar compute client %v", err)
	}

	allPages, err := aggregates.List(computeClient).AllPages(context.TODO())
	if err != nil {
		log.Fatalf("Error getting aggregates list: %v", err)
	}

	allAggregates, err := aggregates.ExtractAggregates(allPages)
	if err != nil {
		log.Fatalf("Error extracting aggregates: %v", err)
	}

	headers := map[int]string{
		0: "Aggregate",
		1: "Hosts",
		2: "Availability Zone",
	}

	columns := map[string][]string{
		"Aggregate":         []string{},
		"Hosts":             []string{},
		"Availability Zone": []string{},
	}

	for _, aggregate := range allAggregates {
		columns["Aggregate"] = append(columns["Aggregate"], aggregate.Name)
		columns["Availability Zone"] = append(columns["Availability Zone"], aggregate.AvailabilityZone)

		var formattedHosts []string
		for i, host := range aggregate.Hosts {
			formattedHosts = append(formattedHosts, host)
			if (i+3)%3 == 0 && i+1 < len(aggregate.Hosts) {
				formattedHosts = append(formattedHosts, "\n")
			}
		}
		columns["Hosts"] = append(columns["Hosts"], strings.Join(formattedHosts, " "))
	}

	return headers, columns, nil

	//aggregateMap := make(map[string][]string)

	//for _, aggregate := range allAggregates {
	//  hosts := fmt.Sprintf("%v", aggregate.Hosts)
	//	aggregateMap[aggregate.Name] = []string{hosts, aggregate.AvailabilityZone}
	//}

	//return aggregateMap, nil

	//hostAggregates := make(map[string][]string)

	//for _, aggregate := range allAggregates {
	//	for _, host := range aggregate.Hosts {
	//		hostAggregates[host] = append(hostAggregates[host], aggregate.Name)
	//	}
	//}

	//finalMap := make(map[string][]string)
	//finalMap["Host"] = []string{}
	//finalMap["Aggregates"] = []string{}

	//for host, aggregates := range hostAggregates {
	//	if len(aggregates) > 1 {
	//		finalMap["Host"] = append(finalMap["Host"], host)
	//		finalMap["Aggregates"] = append(finalMap["Aggregates"], strings.Join(aggregates, ", "))
	//	}
	//}

	//return finalMap, nil
}
