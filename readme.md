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
      version: v0.0.1 # Remember to use the last version
```

Here is an example of a file provider dynamic configuration (given here in YAML), with only the `http.middlewares` section relevant to this plugin:

```yaml
# Dynamic configuration

http:
  middlewares:
    x-real-ip:
      plugin:
        x-real-ip-overwriter:
          header-name: "CF-Connecting-IP" # or watever the header is called
```
