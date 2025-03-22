package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchNpmPackageInfo fetches metadata from the registry and weekly downloads from the npmjs.org API
func FetchNpmPackageInfo(packageName string) (*NpmPackageInfo, error) {
	// 1) Fetch the main registry data
	registryData, err := fetchRegistryData(packageName)
	if err != nil {
		return nil, err
	}

	// 2) Determine the "latest" version from dist-tags
	latestVersion, ok := registryData.DistTags["latest"]
	if !ok {
		return nil, fmt.Errorf("no 'latest' dist-tag found for package %s", packageName)
	}

	// 3) Pull the VersionInfo for that version
	versionInfo, ok := registryData.Versions[latestVersion]
	if !ok {
		return nil, fmt.Errorf("version '%s' not found in registry data for package %s", latestVersion, packageName)
	}

	// 4) Gather fields from the version data
	licenseStr := parseLicense(versionInfo.License)
	unpackedKB := fmt.Sprintf("%.2f KB", float64(versionInfo.Dist.UnpackedSize)/1024.0)

	// 5) Get last published date from the `time` map for that version
	lastPublished := registryData.Time[latestVersion]
	if lastPublished == "" {
		lastPublished = "N/A"
	}

	// 6) Created & Modified are top-level keys in `time`
	created := registryData.Time["created"]
	if created == "" {
		created = "N/A"
	}
	modified := registryData.Time["modified"]
	if modified == "" {
		modified = "N/A"
	}

	// 7) Fetch weekly downloads from the separate API
	downloadsData, err := fetchWeeklyDownloads(packageName)
	if err != nil {
		return nil, err
	}

	// 8) Combine everything into NpmPackageInfo
	npmInfo := &NpmPackageInfo{
		Name:            registryData.Name,
		Version:         latestVersion,
		License:         licenseStr,
		UnpackedSizeKB:  unpackedKB,
		WeeklyDownloads: fmt.Sprintf("%d", downloadsData.Downloads),
		Created:         created,
		Modified:        modified,
		LastPublished:   lastPublished,
	}

	return npmInfo, nil
}

// fetchRegistryData fetches JSON from https://registry.npmjs.org/<packageName>
func fetchRegistryData(packageName string) (*RegistryResponse, error) {
	url := fmt.Sprintf("https://registry.npmjs.org/%s", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch registry data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("registry responded with status %d", resp.StatusCode)
	}

	var data RegistryResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode registry response: %v", err)
	}
	return &data, nil
}

// fetchWeeklyDownloads fetches JSON from https://api.npmjs.org/downloads/point/last-week/<packageName>
func fetchWeeklyDownloads(packageName string) (*DownloadsResponse, error) {
	url := fmt.Sprintf("https://api.npmjs.org/downloads/point/last-week/%s", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch downloads data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("downloads API responded with status %d", resp.StatusCode)
	}

	var downloadsData DownloadsResponse
	if err := json.NewDecoder(resp.Body).Decode(&downloadsData); err != nil {
		return nil, fmt.Errorf("failed to decode downloads response: %v", err)
	}
	return &downloadsData, nil
}

// parseLicense extracts a string from the 'license' field, which might be a string or an object.
func parseLicense(license any) string {
	if license == nil {
		return "N/A"
	}
	switch val := license.(type) {
	case string:
		return val
	case map[string]any:
		// Some packages specify license like { "type": "MIT", "url": "..." }
		if t, ok := val["type"].(string); ok {
			return t
		}
		return "N/A"
	default:
		return "N/A"
	}
}
