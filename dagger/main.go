// This module uses a library call req to make a graphq request to a given url

package main

type GraphqlClient struct{}

func AlpineReq() *Container {
	return dag.Container().From("alpine:latest").
		WithExec([]string{"apk", "update"}).
		WithExec([]string{"apk", "add", "wget"}).
		WithExec([]string{"wget", "https://github.com/ernesto27/req/releases/download/v1.0.11/req_Linux_x86_64.tar.gz"}).
		WithExec([]string{"tar", "-xvf", "req_Linux_x86_64.tar.gz"})
}

// Make a request to a graphql endpoint and returns a json response
//
// Example usage: dagger call graphql-request --url "https://countries.trevorblades.com/" --message="query {countries {name}}" stdout
func (graphql *GraphqlClient) GraphqlRequest(
	// The url of the graphql server
	url string,
	// The graphql query definition
	message string,
	// The file name that have the graphql query definition
	// +optional
	file *File,
) *Container {

	if file == nil {
		return AlpineReq().
			WithExec([]string{"./req", "-t", "gq", "-u", url, "-p", message})
	} else {
		return AlpineReq().
			WithMountedFile("query.txt", file).
			WithExec([]string{"./req", "-t", "gq", "-u", url, "-p", message, "-p", "@query.txt"})
	}

}
