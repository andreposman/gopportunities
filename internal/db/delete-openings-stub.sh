#!/bin/bash
export LC_ALL=C

lc="localhost:8080"
url="http://$lc/api/v1/opening"

# Check if the number of arguments is correct
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <start_id> <end_id>"
  exit 1
fi

start_id=$1
end_id=$2

# Loop to delete entries by ID
for ((i=start_id; i<=end_id; i++)); do
  response=$(curl -s -w "\nHTTP_STATUS:%{http_code}" -X DELETE "$url?id=$i")

  # Extract the body and the status code
  body=$(echo "$response" | sed -e 's/HTTP_STATUS\:.*//g')
  status=$(echo "$response" | tr -d '\n' | sed -e 's/.*HTTP_STATUS://')

  # Check the status code
  if [ "$status" -eq 200 -o "$status" -eq 204 ]; then
    echo "Delete request for ID $i successful"
  else
    echo "Delete request for ID $i failed with status code: $status"
    echo "Response body: $body"
  fi
done