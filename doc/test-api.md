# Test API

    curl --request POST \
      --url http://localhost:8081/wildfowl \
      --header 'cache-control: no-cache' \
      --header 'content-type: application/json' \
      --header 'postman-token: db523c0e-8c2b-5fc4-10b0-2b304eb20fb1' \
      --data '{\n	"s": "sssss"\n}'