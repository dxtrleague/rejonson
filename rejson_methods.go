package rejonson

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func concatWithCmd(cmdName string, args []interface{}) []interface{} {
	res := make([]interface{}, 1)
	res[0] = cmdName
	for _, v := range args {
		if str, ok := v.(string); ok {
			if len(str) == 0 {
				continue
			}
		}
		res = append(res, v)
	}
	return res
}

func jsonDelExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.DEL", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonGetExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.StringCmd {
	cmd := redis.NewStringCmd(concatWithCmd("JSON.GET", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonSetExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.StatusCmd {
	cmd := redis.NewStatusCmd(concatWithCmd("JSON.SET", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonMGetExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.StringSliceCmd {
	cmd := redis.NewStringSliceCmd(concatWithCmd("JSON.MGET", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonTypeExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.StringCmd {
	cmd := redis.NewStringCmd(concatWithCmd("JSON.TYPE", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonNumIncrByExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.StringCmd {
	cmd := redis.NewStringCmd(concatWithCmd("JSON.NUMINCRBY", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonNumMultByExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.StringCmd {
	cmd := redis.NewStringCmd(concatWithCmd("JSON.NUMMULTBY", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonStrAppendExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.STRAPPEND", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonStrLenExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.STRLEN", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonArrAppendExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.ARRAPPEND", args)...)
	_ = c.Process(cmd)
	return cmd
}

func jsoArrIndexExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.ARRINDEX", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonArrInsertExecute(c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.ARRINSERT", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonArrLenExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.ARRLEN", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonArrPopExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.StringCmd {
	cmd := redis.NewStringCmd(concatWithCmd("JSON.ARRPOP", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonArrTrimExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.ARRTRIM", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonObjKeysExecute(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.StringSliceCmd {
	cmd := redis.NewStringSliceCmd(concatWithCmd("JSON.OBJKEYS", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}

func jsonObjLen(ctx context.Context, c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.OBJLEN", args)...)
	_ = c.Process(ctx, cmd)
	return cmd
}
