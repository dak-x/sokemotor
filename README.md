# sokemotor
A custom search engine over a highly configurable collection of sources.

Getting Started:

1.  Pack a browser extension for indexing (kilde)
2.  Add as a special flag/search engine on browsers. (other deploy options can be figured later) 
3.  HTML Source: Utility to process a given HTML Page (with the URI) and Update into Index.
4.  Sources: 
    1.  Some configurable sources
    2.  Plugin support for sources not available
5. Design the Search Frontend.
6. Include a polybar extension as well.

# Contrubting:
## docker:
* Export your GOPATH variable, alternatively create a `.env` file and add `GOPATH=<PATH TO GO DIR>`. To obtain the it use `go env GOPATH`.
* Run the command to launch the docker-container. Live reloading is already configured for the [gin](github.com/gin-gonic/gin) framework.
  
        docker-compose up
* [Reflex](github.com/cespare/reflex) is configured to auto-reload on changes to any `*.go` file in the workspace. This can be updates in the `reflex.conf` file in the directory.
  
* After making changes to the `Dockerfile` or `docker-compose.yml` always rebuild.

        docker-compose up --build
* To exit 

        docker-compose down