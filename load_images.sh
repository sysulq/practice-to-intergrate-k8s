#/bin/bash

file="images.properties"

if [ -f "$file" ]
then
  echo "$file found."

  while IFS='=' read -r key value
  do
    #echo "${key}=${value}"
    value=$(echo "$value" | tr -d '\r')
    docker pull ${value}
    docker tag ${value} ${key}
    docker rmi ${value}
  done < "$file"

else
  echo "$file not found."
fi
