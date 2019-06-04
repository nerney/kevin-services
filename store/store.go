package store

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"github.com/teris-io/shortid"
	"kevin_services/models"
)

var (
	c           = cache.New(-1, -1) // move to init function where we will do NewFrom with stored cache
	errNotFound = errors.New("not found")
)

// Add a quote
func Add(qt models.QuoteText) (models.Quote, error) {
	id, err := shortid.Generate()
	if err != nil {
		return models.Quote{}, err
	}
	quote := models.Quote{ID: id, Text: qt}
	c.SetDefault(id, quote)
	return quote, nil
}

// GetOne quote
func GetOne(id string) (models.Quote, error) {
	quote, found := c.Get(id)
	if found {
		return quote.(models.Quote), nil
	}
	return models.Quote{}, errNotFound
}

// GetAll quotes
func GetAll() ([]models.Quote, error) {
	quotes := []models.Quote{}
	for _, item := range c.Items() {
		if !item.Expired() {
			quotes = append(quotes, item.Object.(models.Quote))
		}
	}
	return quotes, nil
}

// Update a quote
func Update(q models.Quote) (models.Quote, error) {
	item, found := c.Get(q.ID)
	if !found {
		return models.Quote{}, errNotFound
	}
	quote := item.(models.Quote)
	quote = models.Quote{
		ID:   q.ID,
		Text: q.Text,
	}
	c.SetDefault(q.ID, quote)
	return quote, nil
}
