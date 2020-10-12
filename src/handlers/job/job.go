package job

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/adryhappier/gin-kafka-mongo/config"
	model_job "github.com/adryhappier/gin-kafka-mongo/src/models/job"
	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Static Collection
// const JobCollection = "jobs"

// Get DB from Mongo Config
func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// Get All User Endpoint
func GetAllJobs(c *gin.Context) {
	log.Info(os.Getenv("MONGO_DB_NAME"))
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	jobs := model_job.Jobs{}
	err := db.C("jobs").Find(bson.M{}).All(&jobs)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get All Job",
		})
		return
	}

	c.JSON(200, gin.H{
		"jobs": &jobs,
	})
}

// Send data to Kafka
func SyncKafka(c *gin.Context) {
	var _job model_job.Job
	err := c.BindJSON(&_job)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	_job.CreatedAt = time.Now()
	_job.UpdatedAt = time.Now()

	saveJobToKafka(_job)

	c.JSON(200, gin.H{
		"message": "Succes POST Job",
		"user":    &_job,
	})
}

func saveJobToKafka(job model_job.Job) {
	jsonString, err := json.Marshal(job)

	jobString := string(jsonString)
	log.Info(jobString)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_HOST")})
	if err != nil {
		panic(err)
	}

	// Produce messages to topic (asynchronously)
	topic := os.Getenv("KAFKA_TOPIC")
	for _, word := range []string{string(jobString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
}
