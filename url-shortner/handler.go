package main

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, url, http.StatusSeeOther)
		}

		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buildMap(parsedYaml)
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYAML(yml []byte) ([]PathToUrl, error) {
	parsedYAML := []PathToUrl{}
	err := yaml.Unmarshal(yml, &parsedYAML)

	return parsedYAML, err
}

func buildMap(parsedYAML []PathToUrl) map[string]string {
	pathsToUrls := make(map[string]string)

	for _, pathToUrl := range parsedYAML {
		pathsToUrls[pathToUrl.Path] = pathToUrl.Url
	}

	return pathsToUrls
}

type PathToUrl struct {
	Path string `yaml: path`
	Url  string `yaml: url`
}
