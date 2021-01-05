# platform-specs
Specifications for Polygon.io Platform APIs using Open API 3, as well as Polygon's WebSockets using Async API 2.0

# Spec tools

Currently, the spec is written in openapi 3. Openapi 3 is really close at doing multi files well, but no cigar.
We are using [swagger-cli](https://github.com/APIDevTools/swagger-cli), a community accepted solution, to breaking the spec
into multiple files (makes for better organization). Currently, the solution is to use this wrapper script `./_bin_/bundle_spec`
which runs the npm package `swagger-cli`.  
If you want to view the specs locally for a sanity check run `./_bin_/spec_ui`
and go to http://localhost:8080. Note: This is the Swagger UI view, not the custom Polygon.io view.
