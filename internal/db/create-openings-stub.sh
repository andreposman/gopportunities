#!/bin/bash
export LC_ALL=C

lc="localhost:8080"
url="http://$lc/api/v1/opening"

echo $lc
echo $url

# Check if the number of users parameter is provided
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <number_of_users>"
  exit 1
fi

num_users=$1

# Function to generate a random string
generate_random_string() {
  local length=$1
  tr -dc A-Za-z0-9 </dev/urandom | head -c $length
}

# Function to generate a random salary
generate_random_salary() {
  echo $((RANDOM % 50000 + 50000))
}

# Loop to create and send the specified number of JSON payloads
for ((i=1; i<=num_users; i++)); do
  role="Software Engineer $(generate_random_string 5)"
  description="Responsible for developing and maintaining software applications. $(generate_random_string 10)"
  company="Tech Corp $(generate_random_string 3)"
  location="San Francisco, CA"
  isRemote=true
  link="https://techcorp.com/careers/software-engineer-$(generate_random_string 5)"
  salary=$(generate_random_salary)

  json_payload=$(cat <<EOF
{
  "role": "$role",
  "description": "$description",
  "company": "$company",
  "location": "$location",
  "isRemote": $isRemote,
  "link": "$link",
  "salary": $salary
}
EOF
  )

  # Make the POST request using curl
  response=$(curl -s -w "\nHTTP_STATUS:%{http_code}" -X POST -H "Content-Type: application/json" -d "$json_payload" "$url")

  # Extract the body and the status code
  body=$(echo "$response" | sed -e 's/HTTP_STATUS\:.*//g')
  status=$(echo "$response" | tr -d '\n' | sed -e 's/.*HTTP_STATUS://')

  # Check the status code
  if [ "$status" -eq 200 -o "$status" -eq 201 ]; then
    echo "Request $i successful"
  else
    echo "Request $i failed with status code: $status"
    echo "Response body: $body"
  fi
done