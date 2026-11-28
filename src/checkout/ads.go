// Demo change (portal Explore): checkout asks the ad service for an ad to show on the order confirmation.
package main

import (
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// adClient connects to the ad service, at AD_ADDR.
func adClient() (*grpc.ClientConn, error) {
	return grpc.NewClient(os.Getenv("AD_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
