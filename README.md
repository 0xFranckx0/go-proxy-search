# go-proxy-search

- You can test online the 2 routes implemented:
```shell
curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/usage/top_search?size=10

Result:
{"searchCount":72,"lastSearchAt":"2017-08-20T19:00:00.000Z","topSearches":[{"query":"price","count":31,"avgHitCountWithoutTypos":1,"avgHitCount":254},{"query":"streaming media","count":19,"avgHitCountWithoutTypos":55,"avgHitCount":55},{"query":"10","count":6,"avgHitCountWithoutTypos":1000,"avgHitCount":3352},{"query":"amazon","count":2,"avgHitCountWithoutTypos":39,"avgHitCount":55},{"query":"ama","count":1,"avgHitCountWithoutTypos":80,"avgHitCount":80}]}
```
```shell
curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/search?query=price

Truncated result:
{"hits":[{"_highlightResult":{"brand":{"matchLevel":"none","matchedWords":[],"value":"Sennheiser"},"categories":[{"matchLevel":"none","matchedWords":[],"value":"Audio"},{"matchLevel":"none","matchedWords":[],"value":"Headphones"},{"matchLevel":"none","matchedWords":[],"value":"On-Ear Headphones"}],...

```
- To run locally you first need to set the ADMIN_API_KEY
```shell
export ADMIN_API_KEY=XXXYYYYYYYYYZZZZZZZZZAAAAAAAAA
```

- You also need a PORT variable:
```shell
export PORT=8181
```

- Build the proxy:
```
make install
```

- And then run the proxy:
```shell
./proxy
```

- And in another terminal run :
```shell
 curl -XGET -v  http://localhost:8181/1/search?query=price
 curl -XGET -v  http://localhost:8181/1/usage?size=10
 ```
 
