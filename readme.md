## NOTICE: TESTS MISSING
This works for my use case, I will add tests probably early next years, busy with others things, if you want to add test you are more than welcome! <3
The file `x-real-ip-overwriter_test.go` should contain the logic of the tests, it is now absent.

## What is this?

This plugin is a middleware that does only one thing: replace the X-Real-IP header with the value of another header.

## But... Why only this small thing?
Well, I think middlewares should really try to avoid monolithic implementations, you should use a chain to define more complex behaviours.
There are quite a few plugins in the market that do this and other things, I tried a few of them... And none worked, I don't even know why, I am learning GO just for this plugin and surely cannot debug them.
They were doing more things - not an high number of things, but enough that I could not use them and also not fix them.
This plugins aim to do exactly one thing, but do it and do it right.
Well enough rambling, there the plugin documentation goes!

## Usage
For a plugin to be active for a given Traefik instance, it must be declared in the static configuration.
The following declaration in YAML installs it*:
```yaml
# Static configuration

experimental:
  plugins:
    example:
      moduleName: github.com/aless3/x-real-ip-overwrite
      version: v0.0.1 # Remember to use the last version
```
(* it might not work this way because I still have to publish it on the store I think, if that is the case install it as a localPlugin and omit the version that is not a supported field in localPlugins)

Here is an example of a file provider dynamic configuration (given here in YAML), with only the `http.middlewares` section relevant to this plugin:
```yaml
# Dynamic configuration

http:
  middlewares:
    real-ip: # the name to use in chains - you can change this
      plugin:
        x-real-ip-overwriter:
          header-name: "CF-Connecting-IP" # or watever the header is called
```

Note that the X-Real-IP is not overwritten if the "overwriting field" is empty, and in case is saved to the "X-Real-IP-overwritten" field, to avoid losing any info (I encountered a plugin that not only did not set X-Real-IP but also deleted the CF-Connecting-IP header...), so it will always be recoverable, like with a small middleware like this but working the opposite.


## Contributions
As stated above this plugin does only one thing, but if does't work let me know and I will do what I can to fix it, and if you have a fix PLEASE share it, honestly I don't even like the GO language.
If you think you have a small plugin that may work a charm with this and does one small thing (like setting X-Forwarded-For to the cloudflare link - I should probably do it and chain the two), let me know and I will add a section pointing to your plugin! This way users can find easily more small compatible plugins!
