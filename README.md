# golang-in-memory-cache

Sample in memory cache for "on demand" strategy in Go

#### In inMemCache domain first of all checks memory if it is not exists then checks db for data.
#### After find data in db ,set to memory with TTL.

### Patrickmn s go-cache package used for this

## Installation

go get github.com/patrickmn/go-cache

## Run
```bash
docker-compose up
```



## Usage

```go
func (c *cacheManager) GetUser(id int) *mockdata.UserModel {
	key := strconv.Itoa(id)
	result, hasItem := c.cache.Get(key)
	if hasItem {
		response, ok := result.(mockdata.UserModel)
		//check is casting ok
		if ok {
			return &response
		}
	}

	var response mockdata.UserModel
	//it is not ok go check db if it has record for id
	users := c.userRepo.GetUsersById(id)
	if users != nil && len(users.Records) > 0 {
		//db has record for id
		response = users.Records[0]
		//set to memory by setting ttl
		c.cache.Set(key, response, time.Second*TTL)
	}

	return &response
}
```

  
  
  
  
