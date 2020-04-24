# Cache

Go in memory cache solution stores values in memory.

- Concurrent cache.
- Any type of data can be added to the same cache.
- Only one cache object available.
- Cache is removed automatically if the data has expired and it is time to clean.

```go
c := cache.New(timeToExpire, timeToClean) //To create the cache object.
c.Add("key", data) //Adds the data with the specified key.
c.Remove("key") //Removes the data under the specified key.
err, data := c.Get("key") //Gets the data from the cache, if everything is ok -> err == nil.
```

To add it to your code:
```shell
go get github.com/javierlgroba/cache
```
and:
```go
import "github.com/javierlgroba/cache"
```
