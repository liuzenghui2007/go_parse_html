package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"math/rand"
	"path/filepath"
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
	minDelay := time.Second * 10
	maxDelay := time.Second * 20

	// 创建 transport
	transport := &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     30 * time.Second,
		DisableCompression:  true,
		TLSHandshakeTimeout: 20 * time.Second,
	}

	// 可选的代理设置
	proxyStr := "" // 这里可以设置代理地址，例如 "http://your-proxy-here"
	if proxyStr != "" {
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			fmt.Printf("Warning: Invalid proxy URL: %v\n", err)
		} else {
			transport.Proxy = http.ProxyURL(proxyURL)
			fmt.Printf("Using proxy: %s\n", proxyStr)
		}
	} else {
		fmt.Println("No proxy configured, using direct connection")
	}

	client := &http.Client{
		Timeout: time.Second * 120,
		Transport: transport,
	}

	// 随机化 User-Agent
	rand.Seed(time.Now().UnixNano())
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:123.0) Gecko/20100101 Firefox/123.0",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	}

	fmt.Printf("Starting to process %d categories...\n", len(categories))

	for category, path := range categories {
		// 每次请求随机选择 User-Agent
		headers := map[string]string{
			"User-Agent": userAgents[rand.Intn(len(userAgents))],
			"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"Accept-Language": "en-US,en;q=0.9",
			"Accept-Encoding": "gzip, deflate, br",
			"Cache-Control": "max-age=0",
			"Connection": "keep-alive",
			"Upgrade-Insecure-Requests": "1",
			"Sec-Ch-Ua": "\"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Google Chrome\";v=\"122\"",
			"Sec-Ch-Ua-Mobile": "?0",
			"Sec-Ch-Ua-Platform": "\"Windows\"",
			"Sec-Fetch-Dest": "document",
			"Sec-Fetch-Mode": "navigate",
			"Sec-Fetch-Site": "none",
			"Sec-Fetch-User": "?1",
		}

		url := baseURL + path
		// Replace spaces and slashes with underscores in the file name
		fileName := filepath.Join("results", strings.ReplaceAll(strings.ReplaceAll(category, " & ", "_"), "/", "_") + ".csv")

		fmt.Printf("\n[%s] Starting to process category\n", category)
		fmt.Printf("Using User-Agent: %s\n", headers["User-Agent"])
		
		maxRetries := 5
		for retry := 0; retry < maxRetries; retry++ {
			if retry > 0 {
				waitTime := time.Second * time.Duration(retry*10)
				fmt.Printf("[%s] Retry attempt %d/%d, waiting %v\n", category, retry+1, maxRetries, waitTime)
				time.Sleep(waitTime)
			}

			err := FetchAndSave(url, fileName, client, headers)
			if err == nil {
				break
			}
			fmt.Printf("[%s] Error: %v\n", category, err)
			if retry == maxRetries-1 {
				fmt.Printf("[%s] Failed after %d attempts\n", category, maxRetries)
			}
		}

		delay := minDelay + time.Duration(rand.Int63n(int64(maxDelay-minDelay)))
		fmt.Printf("[%s] Waiting %v before next request\n", category, delay)
		time.Sleep(delay)
	}

	fmt.Println("\nAll categories processed!")
} 