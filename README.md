# platform-specs
Specifications for Polygon.io REST APIs and WebSockets.

# Spec tools

Currently, the spec is written in openapi 3. Openapi 3 is really close at doing multi files well, but no cigar.
We are using [swagger-cli](https://github.com/APIDevTools/swagger-cli), a community accepted solution, to breaking the spec
into multiple files (makes for better organization). Currently, the solution is to use this wrapper script `./_bin_/bundle_spec`
which runs the npm package `swagger-cli`.  


If you want to view the specs locally for a sanity check, first run `./_bin_/bundle_spec` to build the spec, then run `./_bin_/spec_ui`
and go to http://localhost:8080. Note: This is the Swagger UI view, not the custom Polygon.io view.

To get a better view into our docs locally you could serve the docs on a local server using `./_bin_/serve_spec`, then you can 
run our ui-docs locally with the following env var set `NEXT_PUBLIC_API_SPECS="http://localhost:9000/rest.json"`
