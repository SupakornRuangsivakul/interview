# for i in {1..20}
# do
#   curl --location 'http://localhost:8080/card/fetch' \
# --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMSIsImV4cCI6MTY5OTE3ODE0M30.m6EkQfqVcaESu2zVevitigt1PVdrdRkWP9UGwMnEbDQ'
#   # Sleep 0.1 seconds to simulate 10 requests per second
#   sleep 0.1
# done

#!/bin/bash

# API endpoint URL
url="http://localhost:8080/card/fetch" 

# Authorization header
auth_header="Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMyIsImV4cCI6MTY5OTE4Nzk1Mn0.Zl5pBkdJVTKbGGCZp1_age35Wr43w1NDXcZLzLhRIwM"

# Function to send a request and print the status code
send_request () {
  status_code=$(curl --location --header "$auth_header" -o /dev/null -s -w "%{http_code}\n" $url)
  echo "Status code: $status_code"
}

echo "Sending request 1"
send_request

echo "Sending request 2"
send_request

echo "Sending request 3"
send_request

echo "Sending request 4"
send_request

echo "Sending request 5"
send_request

echo "Sending request 6"
send_request

echo "Sending request 7"
send_request

echo "Sending request 8"
send_request

sleep 1

echo "Sending request 8"
send_request
# Optional: Add a delay to send the third request after 1 second
# sleep 1
# echo "Sending request 3 (after delay)"
# send_request
