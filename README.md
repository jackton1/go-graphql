Running Locally:

Open terminal run

```bash
$ make run
```


Open another terminal tab and run

```bash
$ curl -g 'http://localhost:12345/graphql?query={songs{title,duration}}'
```


Sample output:
```text
{"data":{"songs":[{"duration":"4:01","title":"Fearless"},{"duration":"4:54","title":"Fifteen"}]}}
```

