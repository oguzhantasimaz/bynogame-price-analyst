package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CoreRepository struct {
	mc *mongo.Database
	mo *options.CollectionOptions
	cn string
}

func NewCoreRepository(mc *mongo.Database, mo *options.CollectionOptions, cn string) *CoreRepository {
	return &CoreRepository{
		mc: mc,
		mo: mo,
		cn: cn,
	}
}

func (r *CoreRepository) GetCollection() *mongo.Collection {
	return r.mc.Collection(r.cn, r.mo)
}

func (r *CoreRepository) GetCollectionName() string {
	return r.mc.Collection(r.cn, r.mo).Name()
}

func (r *CoreRepository) GetDatabaseName() string {
	return r.mc.Name()
}

func (r *CoreRepository) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return r.mc.Collection(r.cn, r.mo).InsertOne(ctx, document, opts...)
}

func (r *CoreRepository) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return r.mc.Collection(r.cn, r.mo).InsertMany(ctx, documents, opts...)
}

func (r *CoreRepository) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return r.mc.Collection(r.cn, r.mo).FindOne(ctx, filter, opts...)
}

func (r *CoreRepository) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return r.mc.Collection(r.cn, r.mo).Find(ctx, filter, opts...)
}

func (r *CoreRepository) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return r.mc.Collection(r.cn, r.mo).UpdateOne(ctx, filter, update, opts...)
}

func (r *CoreRepository) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return r.mc.Collection(r.cn, r.mo).UpdateMany(ctx, filter, update, opts...)
}

func (r *CoreRepository) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return r.mc.Collection(r.cn, r.mo).DeleteOne(ctx, filter, opts...)
}

func (r *CoreRepository) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return r.mc.Collection(r.cn, r.mo).DeleteMany(ctx, filter, opts...)
}

func (r *CoreRepository) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return r.mc.Collection(r.cn, r.mo).Aggregate(ctx, pipeline, opts...)
}

func (r *CoreRepository) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return r.mc.Collection(r.cn, r.mo).CountDocuments(ctx, filter, opts...)
}
