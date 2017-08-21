# go-proxy-search

- You can test online the 2 routes implemented:
```shell
curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/usage/top_search?size=10

Result:
$ curl -XGET https://calm-lowlands-40938.herokuapp.com/1/usage/top_search?size=4
{"topsearches":[{"avgHitCount":453,"avgHitCountWithoutTypos":405,"count":31,"query":"media"},{"avgHitCount":1404,"avgHitCountWithoutTypos":271,"count":28,"query":"compact"},{"avgHitCount":252,"avgHitCountWithoutTypos":1,"count":15,"query":"price"},{"avgHitCount":1250,"avgHitCountWithoutTypos":207,"count":14,"query":"file"}]}
```
```shell
curl -XGET -v https://calm-lowlands-40938.herokuapp.com/1/search?query=price

Truncated result:
{"hits":[{"_highlightResult":{"brand":{"matchLevel":"none","matchedWords":[],"value":"Sennheiser"},"categories":[{"matchLevel":"none","matchedWords":[],"value":"Audio"},{"matchLevel":"none","matchedWords":[],"value":"Headphones"},{"matchLevel":"none","matchedWords":[],"value":"On-Ear Headphones"}],...

```
- To run locally you first need to create a keyfile file containing your Algolia admin api key
  in the proxy binary directory. 
```shell
echo "12345234524542545145426676386" > keyfile
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
 
