package products

import (
	"context"
	"encoding/json"
	"errors"

	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	ErrorNotFound = errors.New("Entity not found.")
)

type Repository interface {
	Close()
	Save(ctx context.Context, p Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	GetAllWithIDs(ctx context.Context, ids []string) ([]Product, error)
	Search(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type elasticRepository struct {
	client *elastic.Client
}

// This struct must have the same structure than document in Elasticsearch.
// This is an entity
type productDocument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

func NewElasticRepository(url string) (Repository, error) {

	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	return &elasticRepository(client), nil
}

func (r *elasticRepository) Close() {

}

func (r *elasticRepository) Save(ctx context.Context, p Product) error {

	_, err := r.client.Index().
		Index("catalog").
		Type("product").
		Id(p.ID).
		BodyJson(productDocument{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		}).
		Do(ctx)
	return err
}

func (r *elasticRepository) GetByID(ctx context.Context, id string) (*Product, error) {

	res, err := r.client.Get().
		Index("catalog").
		Type("product").
		Id(id).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	if !res.Found {
		return nil, ErrorNotFound
	}

	// Entity
	p := productDocument{}

	// We convert elasticsearch result to the entity
	if err = json.Unmarshal(*res.Source, &p); err != nil {
		return nil, err
	}

	return &Product{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, err

}

func (r *elasticRepository) GetAll(ctx context.Context, skip uint64, take uint64) ([]Product, error) {

}

func (r *elasticRepository) GetAllWithIDs(ctx context.Context, ids []string) ([]Product, error) {

}

func (r *elasticRepository) Search(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {

}
