package api

// NpmTime holds creation and modification times for the entire package.
// We'll also parse the "last published" date from a specific version.
type NpmTime struct {
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

// DistInfo holds distribution-related fields for a specific version.
type DistInfo struct {
	Tarball      string `json:"tarball"`
	Shasum       string `json:"shasum"`
	Integrity    string `json:"integrity"`
	UnpackedSize int    `json:"unpackedSize"` // The registry gives a numeric "unpackedSize"
}

// VersionInfo holds metadata for a specific version (license, dist info, etc.).
type VersionInfo struct {
	Name    string   `json:"name"`
	Version string   `json:"version"`
	License any      `json:"license"` // License can be a string or object, so use `any`
	Dist    DistInfo `json:"dist"`
}

// RegistryResponse is the overall response from https://registry.npmjs.org/<packageName>.
type RegistryResponse struct {
	Name     string                   `json:"name"`
	Time     map[string]string        `json:"time"`     // Each version has a date, plus "created", "modified"
	Versions map[string]VersionInfo   `json:"versions"` // Keyed by version string
	DistTags map[string]string        `json:"dist-tags"`
}

// DownloadsResponse is the response from the downloads API at https://api.npmjs.org/downloads/point/last-week/<packageName>.
type DownloadsResponse struct {
	Downloads int    `json:"downloads"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Package   string `json:"package"`
}

// NpmPackageInfo is what we'll return with all fields we care about.
type NpmPackageInfo struct {
	Name            string
	Version         string
	License         string
	UnpackedSizeKB  string // We’ll convert from int bytes to a human-readable "KB"
	WeeklyDownloads string
	Created         string
	Modified        string
	LastPublished   string
}
