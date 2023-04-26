# cli-server
A cli and server program written in Golang. Uses Echo, Cobra, and Viper.

To run the local server change into the 'server' directory and run main.go.

Then open another terminal and run main.go from the root directory to use the cli.

## CLI COMMANDS

### GET
#### FLAGS: 
###### REQUIRED:
--port (-p): enter a port number
--endpoint (-e): enter a string for the url endpoint to want to reach
--baseurl (-b): enter the base url you want to reach

###### Example:
get --port 8080 --endpoint / --baseurl http://localhost
get --port 8080 --endpoint /heartbeat --baseurl http://localhost

### POST

#### FLAGS: 
###### REQUIRED:
--port (-p): enter a port number
--endpoint (-e): enter a string for the url endpoint to want to reach
--baseurl (-b): enter the base url you want to reach

###### OPTIONAL:
-- frvrid (-f): enter a int id to post
--language (-l): enter a string to post

###### Example:
post --port 8080 --endpoint /HappyLang --baseurl http://localhost -f 123 -l asdfasdf
post --port 8080 --endpoint /HappyLang --baseurl http://localhost

### POSTJSON
#### FLAGS: 
###### REQUIRED:
--port (-p): enter a port number
--endpoint (-e): enter a string for the url endpoint to want to reach
--baseurl (-b): enter the base url you want to reach

###### OPTIONAL:
-- frvrid (-f): enter a int id to post
--language (-l): enter a string to post

###### Example:
post --port 8080 --endpoint /HappyLang --baseurl http://localhost -f 123 -l asdfasdf
post --port 8080 --endpoint /HappyLang --baseurl http://localhost
