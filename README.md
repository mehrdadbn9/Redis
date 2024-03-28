# Redis
usable command and feature of redis in use


### redis is single core
### redis is atomic

```bash
sudo apt install lsb-release curl gpg
Add the repository to the apt index, update it, and then install:

curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
```

```docker
docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
```
```bash
set myName mehrdad
get myName
KEYS *
```
## string, global lock 
wanna have display name for keys
```bash
set user:12:name ali
```
van use string for number because it would support and  do mathematics with it

## Hello World with Node.js and Redis
```bash
var redis = require("redis"); // 1
var client = redis.createClient(); // 2
client.set("my_key", "Hello World using Node.js and Redis"); // 3
client.get("my_key", redis.print); // 4
client.quit(); // 5

$ node hello.js
```

redis is using single thread so no need to consider mutex(prevent concurrent access) and prevent race condition can have automatic expiration(setex, ...) -->This is very useful when database queries take a long time to run and can be cached for a given period of time. Consequently, this avoids running those queries too frequently and can give a performance boost to applications. -->use for global lock and if has value it will reutn the error so two person would not have access to one value and only one can have access
```bash
127.0.0.1:6379> set a 12
OK
127.0.0.1:6379> INCRBY a -1
(integer) 11
```
## The MSET command sets the values of multiple keys at once. --> less syscall and overhead
```bash
127.0.0.1:6379> MSET first "First Key value" second "Second Key value"
OK

127.0.0.1:6379> MGET first second
1) "First Key value"
2) "Second Key value"
```
 ## The TTL (Time To Live) command returns one of the following:
•
 A positive integer: This is the amount of seconds a given key
has left to live
•
 -2: If the key is expired or does not exist
•
 -1: If the key exists but has no expiration time set
```bash
127.0.0.1:6379> EXPIRE current_chapter 10
(integer) 1

127.0.0.1:6379> TTL current_chapter
(integer) 3
```
[] getrange 
[] append
## consider redis is fast but RTT(round trip time) or network is not fast so we should consider it

hash can get a lot of key value but you can not use INCR  so we can use sting and hash with each other to use both feature
hash --> map fields to values within a single key

```bash
HSET page:20 title "hello world" content "<html>...</html>"
HGETALL page:20
HGET page:20
HMGET --> for multi one
```
[] HEXISTS
[] HINCRBYFLOAT mykey field 0.1
[]redis> HSET coin heads obverse tails reverse edge null
    (integer) 3
    redis> HRANDFIELD coin
    "heads"

## set is used for math work
```bash
SADD
SISMEMBER
```
when we use a list of thing and we wanna tag to it
```bash
SADD meat burger kebab
SADD vegt corn salad
SADD iranian kebab gheimeh
```
so we can use intersection and find which ones are in common
```bash
SINTER iranian vegt -->tell which is common in both
SUNION meat iranian  --> put all together
```
tell from root which goes to account as well
```bash
SADD page:/ 3 4 8
SADD page:/account 1 4 10
SINTER page:/ page:/account --> 4 has gone to account as well

SINTERSTORE temp:1  page:/account page:/
SMEMBERS temp:1
```
with sorted set and set can make secondary indexes and have awesome search
[] SSCAN --> iterate

## LIST
store an ordered collection 
use cases:
Implementing message queues or task queues where data needs to be processed in a particular order.
Managing timelines or activity streams, where you need to keep track of actions in a specific order.
Caching recent data like recent log entries, recent activities, etc.
Storing and managing a history of actions or events.
```bash
LPUSH  mlist 1 2 3 4
LRANGE mlist 0 -1
RPUSH  mlist 5
LPOP
RPOP mlist
```
we can use list for fan out, 

```bash
blpop -->block list pop
```

## sorted set
```bash
ZADD scores 1 ali 3.3 mehrdad 2 sara

zrange scores 0 2
1) "ali"
2) "mehrdad"

zrevrange scores 0 2 withscores
zrank scores  mehrdad
 ```
we can give weight to scores
```bash
ZINTERCARD --> give number of intersect
zinter --> consider complexity n^2 or nlogn
```
we wanna know healthy and meat how
```bash
zinterstore tag:meat:healthy 2 tag:healthy tag:meat 
EXPIRE  tag:meat:healthy 120
TTL tag:meat:healthy
BZPOPMAX --> can give higher priority to customer who wants to  better quality and they have paid so no need they wait in normal queue
```

geospatial and golang and POI(point of interest)
https://medium.com/@mhrlife/building-an-online-taxi-app-like-uber-with-golang-part-1-nearby-taxis-c509168ef59f

```bash
o(n) would answer about ten thousands request
o(logn) would answer about one hundered thousands req
```
for finding location of person after ten times of splitting with have only one car in one meter -->  at redis we have geospatial https://redis.io/commands/?group=geo
at google lat and longitude comes vice of versa in contrast with elastic and redis
```bash
```

```bash
```

```bash
```

```bash
```

```bash
```

```bash
```

```bash
```

```bash
```
