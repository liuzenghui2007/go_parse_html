package main

import (
	"fmt"
	"strings"
)

func main() {
	categories := map[string]string{
		"Accessories & Jewellery": "/top-instagram-accessories-jewellery-united-states/",
		"Adult content":           "/top-instagram-adult-content-united-states/",
		"Alcohol":                 "/top-instagram-alcohol-united-states/",
		"Animals":                 "/top-instagram-animals-united-states/",
		"Architecture & Urban Design": "/top-instagram-architecture-urban-design-united-states/",
		"Art/Artists":             "/top-instagram-art-artists-united-states/",
		"Beauty":                  "/top-instagram-beauty-united-states/",
		"Business & Careers":      "/top-instagram-business-careers-united-states/",
		"Cars & Motorbikes":       "/top-instagram-cars-motorbikes-united-states/",
		"Cinema & Actors/actresses": "/top-instagram-cinema-actors-actresses-united-states/",
		"Clothing & Outfits":      "/top-instagram-clothing-outfits-united-states/",
		"Comics & sketches":       "/top-instagram-comics-sketches-united-states/",
		"Computers & Gadgets":     "/top-instagram-computers-gadgets-united-states/",
		"Crypto":                  "/top-instagram-crypto-united-states/",
		"DIY & Design":            "/top-instagram-diy-design-united-states/",
		"Education":               "/top-instagram-education-united-states/",
		"Extreme Sports & Outdoor activity": "/top-instagram-extreme-sports-outdoor-united-states/",
		"Family":                  "/top-instagram-family-united-states/",
		"Fashion":                 "/top-instagram-fashion-united-states/",
		"Finance & Economics":     "/top-instagram-finance-economics-united-states/",
		"Fitness & Gym":           "/top-instagram-fitness-gym-united-states/",
		"Food & Cooking":          "/top-instagram-food-cooking-united-states/",
		"Gaming":                  "/top-instagram-gaming-united-states/",
		"Health & Medicine":       "/top-instagram-health-medicine-united-states/",
		"Humor & Fun & Happiness": "/top-instagram-humor-fun-happiness-united-states/",
		"Kids & Toys":             "/top-instagram-kids-toys-united-states/",
		"Lifestyle":               "/top-instagram-lifestyle-united-states/",
		"Literature & Journalism": "/top-instagram-literature-journalism-united-states/",
		"Luxury":                  "/top-instagram-luxury-united-states/",
		"Machinery & Technologies": "/top-instagram-machinery-technologies-united-states/",
		"Management & Marketing":  "/top-instagram-management-marketing-united-states/",
		"Mobile related":          "/top-instagram-mobile-related-united-states/",
		"Modeling":                "/top-instagram-modeling-united-states/",
		"Music":                   "/top-instagram-music-united-states/",
		"NFT":                     "/top-instagram-nft-united-states/",
		"Nature & landscapes":     "/top-instagram-nature-landscapes-united-states/",
		"Photography":             "/top-instagram-photography-united-states/",
		"Racing Sports":           "/top-instagram-racing-sports-united-states/",
		"Science":                 "/top-instagram-science-united-states/",
		"Shopping & Retail":       "/top-instagram-shopping-retail-united-states/",
		"Shows":                   "/top-instagram-shows-united-states/",
		"Sports with a ball":      "/top-instagram-sports-with-a-ball-united-states/",
		"Sweets & Bakery":         "/top-instagram-sweets-bakery-united-states/",
		"Tobacco & Smoking":       "/top-instagram-tobacco-smoking-united-states/",
		"Trainers & Coaches":      "/top-instagram-trainers-coaches-united-states/",
		"Travel":                  "/top-instagram-travel-united-states/",
		"Water sports":            "/top-instagram-water-sports-united-states/",
		"Winter sports":           "/top-instagram-winter-sports-united-states/",
	}

	baseURL := "https://hypeauditor.com"

	for category, path := range categories {
		url := baseURL + path
		fileName := strings.ReplaceAll(category, " & ", "_") + ".csv"
		fileName = strings.ReplaceAll(fileName, " ", "_")

		fmt.Printf("Fetching data for %s...\n", category)
		FetchAndSave(url, fileName)
	}
} 