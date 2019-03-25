# Ip_location_finder

This README is a *quick start* document. 

What is Ip_location_finder?
--------------
Ip-Location-Finder provides a api to easily find the location information of a ip address, including country, province, the provincial capital city and the zip code. 
The api's route is /search/location, using query param ip. 
This project loads the ip-location file once just when the server starts. At every api call, ip-location information is got from the memory.

USAGE:
--------------
main [global options] command [command options] [arguments...]

 
Quick Start:
--------------
1. git clone 
2. go mod tidy
3. go run cmd/main.go start
4. curl '127.0.0.1:8081/search/location?ip=1.1.255.255'
