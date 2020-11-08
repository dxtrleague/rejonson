package rejonson

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type redisProcessor struct {
	Process func(ctx context.Context, cmd redis.Cmder) error
}

/*
Client is an extended redis.Client, stores a pointer to the original redis.Client
*/
type Client struct {
	*redis.Client
	*redisProcessor
}

/*
Pipeline is an extended redis.Pipeline, stores a pointer to the original redis.Pipeliner
*/
type Pipeline struct {
	redis.Pipeliner
	*redisProcessor
}

func (cl *Client) Pipeline() *Pipeline {
	pip := cl.Client.Pipeline()
	return ExtendPipeline(pip)
}

func (cl *Client) TXPipeline() *Pipeline {
	pip := cl.Client.TxPipeline()
	return ExtendPipeline(pip)
}
func (pl *Pipeline) Pipeline() *Pipeline {
	pip := pl.Pipeliner.Pipeline()
	return ExtendPipeline(pip)
}

/*
JsonDel

returns intCmd -> deleted 1 or 0
read more: https://oss.redislabs.com/rejson/commands/#jsondel
*/
func (cl *redisProcessor) JsonDel(c context.Context, key, path string) *redis.IntCmd {
	return jsonDelExecute(c, cl, key, path)
}

/*
JsonGet

Possible args:

(Optional) INDENT + indent-string
(Optional) NEWLINE + line-break-string
(Optional) SPACE + space-string
(Optional) NOESCAPE
(Optional) path ...string

returns stringCmd -> the JSON string
read more: https://oss.redislabs.com/rejson/commands/#jsonget
*/
func (cl *redisProcessor) JsonGet(c context.Context, key string, args ...interface{}) *redis.StringCmd {
	return jsonGetExecute(c, cl, append([]interface{}{key}, args...)...)
}

/*
jsonSet

Possible args:
(Optional)
*/
func (cl *redisProcessor) JsonSet(c context.Context, key, path, json string, args ...interface{}) *redis.StatusCmd {
	return jsonSetExecute(c, cl, append([]interface{}{key, path, json}, args...)...)
}

func (cl *redisProcessor) JsonMGet(c context.Context, key string, args ...interface{}) *redis.StringSliceCmd {
	return jsonMGetExecute(c, cl, append([]interface{}{key}, args...)...)
}

func (cl *redisProcessor) JsonType(c context.Context, key, path string) *redis.StringCmd {
	return jsonTypeExecute(c, cl, key, path)
}

func (cl *redisProcessor) JsonNumIncrBy(c context.Context, key, path string, num int) *redis.StringCmd {
	return jsonNumIncrByExecute(c, cl, key, path, num)
}

func (cl *redisProcessor) JsonNumMultBy(c context.Context, key, path string, num int) *redis.StringCmd {
	return jsonNumMultByExecute(c, cl, key, path, num)
}

func (cl *redisProcessor) JsonStrAppend(c context.Context, key, path, appendString string) *redis.IntCmd {
	return jsonStrAppendExecute(c, cl, key, path, appendString)
}

func (cl *redisProcessor) JsonStrLen(c context.Context, key, path string) *redis.IntCmd {
	return jsonStrLenExecute(c, cl, key, path)
}

func (cl *redisProcessor) JsonArrAppend(c context.Context, key, path string, jsons ...interface{}) *redis.IntCmd {
	return jsonArrAppendExecute(c, cl, append([]interface{}{key, path}, jsons...)...)
}

func (cl *redisProcessor) JsonArrIndex(c context.Context, key, path string, jsonScalar interface{}, startAndStop ...interface{}) *redis.IntCmd {
	return jsoArrIndexExecute(c, cl, append([]interface{}{key, path, jsonScalar}, startAndStop...)...)
}

func (cl *redisProcessor) JsonArrInsert(c context.Context, key, path string, index int, jsons ...interface{}) *redis.IntCmd {
	return jsonArrInsertExecute(c, cl, append([]interface{}{key, path, index}, jsons...)...)
}

func (cl *redisProcessor) JsonArrLen(c context.Context, key, path string) *redis.IntCmd {
	return jsonArrLenExecute(c, cl, key, path)
}

func (cl *redisProcessor) JsonArrPop(c context.Context, key, path string, index int) *redis.StringCmd {
	return jsonArrPopExecute(c, cl, key, path, index)
}

func (cl *redisProcessor) JsonArrTrim(c context.Context, key, path string, start, stop int) *redis.IntCmd {
	return jsonArrTrimExecute(c, cl, key, path, start, stop)
}

func (cl *redisProcessor) JsonObjKeys(c context.Context, key, path string) *redis.StringSliceCmd {
	return jsonObjKeysExecute(c, cl, key, path)
}

func (cl *redisProcessor) JsonObjLen(c context.Context, key, path string) *redis.IntCmd {
	return jsonObjLen(c, cl, key, path)
}
