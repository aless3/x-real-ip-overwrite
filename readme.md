## What is this?

This plugin is a middleware that does only one thing: replace the X-Real-IP header with the value of another header.

## Usage

For a plugin to be active for a given Traefik instance, it must be declared in the static configuration.

The following declaration (given here in YAML) defines a plugin:

```yaml
# Static configuration

experimental:
  plugins:
    example:
      moduleName: github.com/aless3/x-real-ip-overwrite
      version: v0.0.1
```

Here is an example of a file provider dynamic configuration (given here in YAML), with only the `http.middlewares` section relevant to this plugin:

```yaml
# Dynamic configuration

http:
  middlewares:
    x-real-ip:
      plugin:
        x-real-ip-overwriter:
          header-name: "CF-Connecting-IP"
```

## Defining a Plugin

A plugin package must define the following exported Go objects:

- A type `type Config struct { ... }`. The struct fields are arbitrary.
- A function `func CreateConfig() *Config`.
- A function `func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error)`.

```go
// Package example a example plugin.
package example

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	// ...
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		// ...
	}
}

// Example a plugin.
type Example struct {
	next     http.Handler
	name     string
	// ...
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// ...
	return &Example{
		// ...
	}, nil
}

func (e *Example) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// ...
	e.next.ServeHTTP(rw, req)
}
```

## Logs

Currently, the only way to send logs to Traefik is to use `os.Stdout.WriteString("...")` or `os.Stderr.WriteString("...")`.

In the future, we will try to provide something better and based on levels.

## Plugins Catalog

Traefik plugins are stored and hosted as public GitHub repositories.

Every 30 minutes, the Plugins Catalog online service polls Github to find plugins and add them to its catalog.

### Prerequisites

To be recognized by Plugins Catalog, your repository must meet the following criteria:

- The `traefik-plugin` topic must be set.
- The `.traefik.yml` manifest must exist, and be filled with valid contents.

If your repository fails to meet either of these prerequisites, Plugins Catalog will not see it.

### Manifest

A manifest is also mandatory, and it should be named `.traefik.yml` and stored at the root of your project.

This YAML file provides Plugins Catalog with information about your plugin, such as a description, a full name, and so on.

Here is an example of a typical `.traefik.yml`file:

```yaml
# The name of your plugin as displayed in the Plugins Catalog web UI.
displayName: Name of your plugin

# For now, `middleware` is the only type available.
type: middleware

# The import path of your plugin.
import: github.com/username/my-plugin

# A brief description of what your plugin is doing.
summary: Description of what my plugin is doing

# Medias associated to the plugin (optional)
iconPath: foo/icon.png
bannerPath: foo/banner.png

# Configuration data for your plugin.
# This is mandatory,
# and Plugins Catalog will try to execute the plugin with the data you provide as part of its startup validity tests.
testData:
  Headers:
    Foo: Bar
```

Properties include:

- `displayName` (required): The name of your plugin as displayed in the Plugins Catalog web UI.
- `type` (required): For now, `middleware` is the only type available.
- `import` (required): The import path of your plugin.
- `summary` (required): A brief description of what your plugin is doing.
- `testData` (required): Configuration data for your plugin. This is mandatory, and Plugins Catalog will try to execute the plugin with the data you provide as part of its startup validity tests.
- `iconPath` (optional): A local path in the repository to the icon of the project.
- `bannerPath` (optional): A local path in the repository to the image that will be used when you will share your plugin page in social medias.

There should also be a `go.mod` file at the root of your project. Plugins Catalog will use this file to validate the name of the project.

### Tags and Dependencies

Plugins Catalog gets your sources from a Go module proxy, so your plugins need to be versioned with a git tag.

Last but not least, if your plugin middleware has Go package dependencies, you need to vendor them and add them to your GitHub repository.

If something goes wrong with the integration of your plugin, Plugins Catalog will create an issue inside your Github repository and will stop trying to add your repo until you close the issue.

## Troubleshooting

If Plugins Catalog fails to recognize your plugin, you will need to make one or more changes to your GitHub repository.

In order for your plugin to be successfully imported by Plugins Catalog, consult this checklist:

- The `traefik-plugin` topic must be set on your repository.
- There must be a `.traefik.yml` file at the root of your project describing your plugin, and it must have a valid `testData` property for testing purposes.
- There must be a valid `go.mod` file at the root of your project.
- Your plugin must be versioned with a git tag.
- If you have package dependencies, they must be vendored and added to your GitHub repository.
