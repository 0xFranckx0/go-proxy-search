# go-proxy-search

- You can test online the 2 routes implemented:
```shell
curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/usage/top_search?size=10

curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/search?query=price
```
- To run locally you first need to set the ADMIN_API_KEY
```shell
export ADMIN_API_KEY=XXXYYYYYYYYYZZZZZZZZZAAAAAAAAA
```

- You also need a PORT variable:
```shell
export PORT=8181
```
