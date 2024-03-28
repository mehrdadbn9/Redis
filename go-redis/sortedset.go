/*
create a file and name it: food
meat kebab 5
meat stake 5
meat ice-cream 0
meat baklava 0
meat tea 0
meat coffee 0
meat burger 4
meat hot-dog 2
meat samosa 1
meat salad 0
meat falafel 0
meat ghormeh_sabzi 4
meat tahchin 3
meat fesenjan 4
meat grilled_chicken 5
meat sushi 3
meat pizza 2
meat pasta 1
meat tacos 3
meat mirza_ghasemi 0
meat kashk_e_bademjan 0
meat zereshk_polo 2
meat baghali_polo 2
meat dizi 4
meat koofteh_tabrizi 4
meat ash_reshteh 1
meat kabab_koobideh 5
meat kabab_barg 5
meat joojeh_kabab 5
meat sholezard 0
meat loobia_polo 2
meat kookoo_sabzi 0
meat dolmeh 1
iranian kebab 5
iranian stake 0
iranian ice-cream 0
iranian baklava 4
iranian tea 0
iranian coffee 0
iranian burger 0
iranian hot-dog 0
iranian samosa 3
iranian salad 0
iranian falafel 1
iranian ghormeh_sabzi 5
iranian tahchin 5
iranian fesenjan 5
iranian grilled_chicken 0
iranian sushi 0
iranian pizza 0
iranian pasta 0
iranian tacos 0
iranian fish 4
iranian mirza_ghasemi 5
iranian kashk_e_bademjan 5
iranian zereshk_polo 5
iranian baghali_polo 5
iranian dizi 5
iranian koofteh_tabrizi 5
iranian ash_reshteh 5
iranian kabab_koobideh 5
iranian kabab_barg 5
iranian joojeh_kabab 5
iranian sholezard 5
iranian loobia_polo 5
iranian kookoo_sabzi 5
iranian dolmeh 5
healthy kebab 2
healthy stake 2
healthy ice-cream 1
healthy baklava 1
healthy tea 4
healthy coffee 4
healthy burger 2
healthy hot-dog 1
healthy samosa 2
healthy salad 5
healthy falafel 3
healthy ghormeh_sabzi 3
healthy tahchin 2
healthy fesenjan 3
healthy grilled_chicken 4
healthy sushi 4
healthy pizza 2
healthy pasta 3
healthy tacos 3
healthy fish 4
healthy mirza_ghasemi 4
healthy kashk_e_bademjan 4
healthy zereshk_polo 3
healthy baghali_polo 3
healthy dizi 3
healthy koofteh_tabrizi 3
healthy ash_reshteh 3
healthy kabab_koobideh 2
healthy kabab_barg 2
healthy joojeh_kabab 4
healthy sholezard 2
healthy loobia_polo 3
healthy kookoo_sabzi 4
healthy dolmeh 4
vegetable kebab 0
vegetable stake 0
vegetable ice-cream 0
vegetable baklava 0
vegetable tea 0
vegetable coffee 0
vegetable burger 2
vegetable hot-dog 1
vegetable samosa 2
vegetable salad 5
vegetable falafel 4
vegetable ghormeh_sabzi 4
vegetable tahchin 1
vegetable fesenjan 1
vegetable grilled_chicken 0
vegetable sushi 2
vegetable pizza 2
vegetable pasta 2
vegetable tacos 3
vegetable fish 0
vegetable mirza_ghasemi 5
vegetable kashk_e_bademjan 5
vegetable zereshk_polo 1
vegetable baghali_polo 2
vegetable dizi 1
vegetable koofteh_tabrizi 2
vegetable ash_reshteh 5
vegetable kabab_koobideh 0
vegetable kabab_barg 0
vegetable joojeh_kabab 0
vegetable sholezard 0
vegetable loobia_polo 3
vegetable kookoo_sabzi 5
vegetable dolmeh 4
*/

package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:embed foods
var foodTags string

func main() {
	rdb := redis.NewClient(&redis.Options{}) //redis server must be running
	rdb.FlushAll(context.Background())       //database must be empty
	importTags(rdb)                          //import food tags
	listByTag(rdb, "healthy", "vegetable")   //list healthy and vegetable can change to iranian and ..
}

func importTags(rdb *redis.Client) {
	tags := strings.Split(strings.TrimSpace(foodTags), "\n")
	for _, row := range tags {
		items := strings.Split(row, " ")
		if len(items) != 3 {
			log.Fatalf("bad row: %s\n", row) //bad row
		}
		score, _ := strconv.Atoi(strings.TrimSpace(items[2]))                              //tag score
		if err := rdb.ZAdd(context.Background(), fmt.Sprintf("tag:%s", items[0]), redis.Z{ //add tag
			Score:  float64(score),
			Member: items[1],
		}).Err(); err != nil {
			panic(err)
		}
	}
	log.Println("added all rows")
}

func listByTag(rdb *redis.Client, tags ...string) {
	sort.Strings(tags)
	key := "tag:" + strings.Join(tags, ":") //tag:healthy:vegetable
	var keys []string
	for _, tag := range tags {
		keys = append(keys, "tag:"+tag)
	}
	// this tag does not exist
	if rdb.Exists(context.Background(), key).Val() == 0 {
		log.Printf("%s does not exist, calling ZINTERSTORE\n", key)
		if err := rdb.ZInterStore(context.Background(), key, &redis.ZStore{
			Keys:      keys,
			Aggregate: "SUM",
		}).Err(); err != nil {
			log.Printf("error while creating %s: %v\n", keys, err)
			return
		}
		log.Printf("%s was created\n", key)
		if err := rdb.Expire(context.Background(), key, time.Minute).Err(); err != nil { // ttl tag:healthy:vegetable
			log.Printf("error while expiring %s: %v\n", key, err)
		}
	}
	result, _ := rdb.ZRevRangeWithScores(context.Background(), key, 0, 10).Result()
	for _, z := range result {
		log.Printf("[%2.0f]  %s\n", z.Score, z.Member)
	}
}
