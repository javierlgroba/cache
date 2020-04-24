//Package definition for the memory cache
package cache

import(
  "fmt"
  "sync"
  "time"
  "errors"
)

type (
  //Type definition for the values to store in the cache
  value struct {
    time time.Time
    value interface{}
  }
  //Type definition for the cache storage
  Cache struct {
	   cache map[string]value
	   cacheMutex sync.Mutex
     expire, maid time.Duration
     isValid bool
    })

const(
  //Default time for cache expire
  defaultExpiringDuration = 5
  //Default time to destroy cached objects
  defaultMaidDuration = 10)

//local variable for the cache
var cache *Cache = &Cache{
  isValid: false}

//Function that initializes a Cache object. Receives two integers indicating
//the time to expire and the time to destroy expired cached objects in minutes.
//When passing 0,0 default values are used.
func New(expire, maid int) (*Cache){
  if expire==0 {
    expire = defaultExpiringDuration
  }
  if maid==0 {
    maid = defaultMaidDuration
  }

  expireDuration, _ := time.ParseDuration(fmt.Sprintf("%dm", expire))
  maidDuration, _ := time.ParseDuration(fmt.Sprintf("%dm", maid))

  //Make sure that no one is calling New at the same time.
  //Lock and Unlock the same mutex and set the old cache as invalid.
  cache.cacheMutex.Lock()
  cache.isValid = false
  cache.cacheMutex.Unlock()

  //Create the new cache
  cache = &Cache{
    cache: map[string]value{},
    expire: expireDuration,
    maid: maidDuration,
    isValid: false}

  //TODO: Start background thread for maid

  //Set cache as valid before returning
  cache.isValid = true
  return cache
}

//Adds data to the cache
func (c *Cache) Add(key string, data interface{}) error{
  c.cacheMutex.Lock()
  defer c.cacheMutex.Unlock()

  if !cache.isValid {
    return errors.New("The cache is now invalid.")
  }

  c.cache[key] = value{
    time: time.Now(),
    value: data}

    return nil
}

//Removes data from the cache
func (c *Cache) Remove(key string) error {
  c.cacheMutex.Lock()
  defer c.cacheMutex.Unlock()

  if !cache.isValid {
    return errors.New("The cache is now invalid.")}

  _, ok := c.cache[key];
  if ok {
    delete(c.cache, key)
  }

  return nil
}

//Gets data from the cache
func (c *Cache) Get(key string) (error, interface{}) {
  c.cacheMutex.Lock()
  defer c.cacheMutex.Unlock()

  if !cache.isValid {
    return errors.New("The cache is now invalid."), nil
  }

  value, ok := c.cache[key];
  if ok {
    if value.expired(c){
      return errors.New("The data has expired."), value.value
    }
    return nil, value.value
  }

  return errors.New("There is no value for the given key."), nil
}

//Check if data has expired
func (v value) expired(c *Cache) bool{
  return time.Since(v.time)>c.expire
}
