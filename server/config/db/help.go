package db

import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
import "go.mongodb.org/mongo-driver/bson"

import "time"
import "context"

import "github.com/amsokol/mongo-go-driver-protobuf/pmongo"

import "reflect"

import "errors"

// Collection ...
type Collection struct {
	collection *mongo.Collection
	database *mongo.Database
}

// NoContext ...
var NoContext = context.TODO()

// ErrorInvalidOutType ...
var ErrorInvalidOutType = errors.New("Output must be pointer type")

// ErrorInvalidPipeline ...
var ErrorInvalidPipeline = errors.New("Pipeline must be slice")


// Doc ...
type Doc interface {
	SetModified(time.Time)
	SetCreated(time.Time)
	SetID(interface{})
}

// SoftDoc ...
type SoftDoc interface {
	Doc
	SetDeleted(bool)
}

// Insert ...
func (c *Collection) Insert(doc Doc) error {
	con, err := c.collection.Clone()
	if err != nil {
		return err
	}
	now := time.Now()
	doc.SetCreated(now)
	doc.SetModified(now)

	res, err := con.InsertOne(NoContext, doc)
	if err != nil {
		return err
	}
	doc.SetID(res.InsertedID)
	
	return nil
}

// FindOne ...
func (c *Collection) FindOne(query interface{}, out Doc) error {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return ErrorInvalidOutType
	}
	con, err := c.collection.Clone()
	if err != nil {
		return err
	}
	return con.FindOne(NoContext, query).Decode(out)
}

// FindByID ...
func (c *Collection) FindByID(id *pmongo.ObjectId, out Doc) (error) {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return ErrorInvalidOutType
	}
	con, err := c.collection.Clone()
	if err != nil {
		return err
	}
	return con.FindOne(NoContext, bson.D{{Key: "_id", Value: id}}).Decode(out)
}

// Find ...
func (c *Collection) Find(query interface{}) (*mongo.Cursor, error) {
	con, err := c.collection.Clone()
	if err != nil {
		return nil, err
	}
	return con.Find(NoContext, query)
}

// UpdateOne ...
func (c *Collection) UpdateOne(query interface{}, doc Doc) error {
	con, err := c.collection.Clone()
	if err != nil {
		return err
	}
	
	doc.SetModified(time.Now())

	_, err = con.UpdateOne(NoContext, query, doc, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

// PartialUpdateOne ...
func (c *Collection) PartialUpdateOne(query interface{}, update interface{}, out Doc) error {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return ErrorInvalidOutType
	}
	con, err := c.collection.Clone()
	if err != nil {
		return err
	}
	res := con.FindOneAndUpdate(NoContext, query, update, options.FindOneAndUpdate().SetReturnDocument(options.After))

	return res.Decode(out)
}

// Update ...
func (c *Collection) Update(query interface{}, update interface{}) (int64, error) {

	con, err := c.collection.Clone()
	if err != nil {
		return 0, err
	}
	res, err := con.UpdateMany(NoContext, query, update, options.Update().SetUpsert(true))

	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

// DeleteOne ...
func (c *Collection) DeleteOne(query interface{}, out Doc) error {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return ErrorInvalidOutType
	}
	con, err := c.collection.Clone()
	if err != nil {
		return err
	}
	res := con.FindOneAndDelete(NoContext, query)

	return res.Decode(out)
}

// SoftDelete ...
func (c *Collection) SoftDelete(query interface{}, out SoftDoc) error {
	if out == nil {
		return nil
	}
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return ErrorInvalidOutType
	}

	if err := c.FindOne(query, out); err != nil {
		return nil
	}

	out.SetDeleted(true)
	return c.UpdateOne(query, out)
}

// SoftRecover ...
func (c *Collection) SoftRecover(query interface{}, out SoftDoc) error {
	if out == nil {
		return nil
	}
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return ErrorInvalidOutType
	}

	if err := c.FindOne(query, out); err != nil {
		return nil
	}

	out.SetDeleted(false)
	return c.UpdateOne(query, out)
}

// Aggregate ...
func (c *Collection) Aggregate(pipeline interface{}) (*mongo.Cursor, error) {
	switch reflect.TypeOf(pipeline).Kind() {
	case reflect.Ptr:
		switch reflect.PtrTo(reflect.TypeOf(pipeline)).Kind() {
		case reflect.Array, reflect.Slice:
			break
		default:
			return nil, ErrorInvalidPipeline
		}
	default:
		return nil, ErrorInvalidPipeline
	}
	con, err := c.collection.Clone()
	if err != nil {
		return nil, err
	}
	return con.Aggregate(NoContext, pipeline, options.Aggregate().SetAllowDiskUse(true))
}