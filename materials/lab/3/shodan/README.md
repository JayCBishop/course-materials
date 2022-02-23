# README

## Option

For the first part of lab 3b, I chose to go with Option 1.

## Modifications

To fulfill Option 1's requirements, I created the `utility.go` file that accesses the `tools/myip` method on shodan through the `Utility` function. This method returns the IP address of the client and does not take any additional route or query parameters outside the API key. 

Additionally, I added a new public function to the `host.go` file called `SearchByIp`. This function takes an IP address as a string and uses it as a route parameter for a `shodan/host/{ip}` request. It then returns the service information retrieved in the response from shodan.
