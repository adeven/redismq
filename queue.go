package redismq

import (
	"github.com/adeven/redis"
)

type Queue struct {
	redisClient *redis.Client
	Name        string
}

func NewQueue(redisUrl, redisPassword string, redisDB int64, name string) *Queue {
	q := &Queue{Name: name}
	q.redisClient = redis.NewTCPClient(redisUrl, redisPassword, redisDB)
	q.redisClient.SAdd(MasterQueueKey(), name)
	return q
}

func MasterQueueKey() string {
	return "redismq::queues"
}

func (self *Queue) WorkerKey() string {
	return self.InputName() + "::workers"
}

func (self *Queue) InputName() string {
	return "redismq::" + self.Name
}

func (self *Queue) FailedName() string {
	return "redismq::" + self.Name + "::failed"
}

func (self *Queue) InputCounterName() string {
	return self.InputName() + "::counter"
}

func (self *Queue) FailedCounterName() string {
	return self.FailedName() + "::counter"
}

func (self *Queue) ConsumerWorkingPrefix() string {
	return "redismq::" + self.Name + "::working"
}

func (self *Queue) ConsumerAckPrefix() string {
	return "redismq::" + self.Name + "::ack"
}
