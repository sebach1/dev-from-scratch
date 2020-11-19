package sample1

import (
	"fmt"
	"sync"
	"time"
)

// PriceService is a service that we can use to get prices for the items
// Calls to this service are expensive (they take time)
type PriceService interface {
	GetPriceFor(itemCode string) (float64, error)
}

// TransparentCache is a cache that wraps the actual service
// The cache will remember prices we ask for, so that we don't have to wait on every call
// Cache should only return a price if it is not older than "maxAge", so that we don't get stale prices
type TransparentCache struct {
	actualPriceService PriceService
	maxAge             time.Duration
	prices             map[string]*PriceResult
	sync.RWMutex
}

type PriceResult struct {
	Price     float64
	CreatedAt time.Time
}

func NewTransparentCache(actualPriceService PriceService, maxAge time.Duration) *TransparentCache {
	return &TransparentCache{
		actualPriceService: actualPriceService,
		maxAge:             maxAge,
		prices:             make(map[string]*PriceResult),
	}
}

func (c *TransparentCache) getCachedPriceFor(itemCode string) *PriceResult {
	c.RLock()
	defer c.RUnlock()
	result, ok := c.prices[itemCode]
	if !ok {
		return nil
	}
	if time.Since(result.CreatedAt) > c.maxAge {
		return nil
	}
	return result
}

func (c *TransparentCache) setCachedPriceFor(itemCode string, price float64) {
	c.Lock()
	defer c.Unlock()
	c.prices[itemCode] = &PriceResult{CreatedAt: time.Now(), Price: price}
}

// GetPriceFor gets the price for the item, either from the cache or the actual service if it was not cached or too old
func (c *TransparentCache) GetPriceFor(itemCode string) (float64, error) {
	result := c.getCachedPriceFor(itemCode)
	if result != nil {
		return result.Price, nil
	}
	price, err := c.actualPriceService.GetPriceFor(itemCode)
	if err != nil {
		return 0, fmt.Errorf("getting price from service : %v", err.Error())
	}

	c.setCachedPriceFor(itemCode, price)
	return price, nil
}

// GetPricesFor gets the prices for several items at once, some might be found in the cache, others might not
// If any of the operations returns an error, it should return an error as well
func (c *TransparentCache) GetPricesFor(itemCodes ...string) ([]float64, error) {
	var results []float64
	errCh := make(chan error, 1)
	resultsCh := make(chan float64)

	for _, itemCode := range itemCodes {
		itemCode := itemCode
		go func() {
			price, err := c.GetPriceFor(itemCode)
			if err != nil {
				errCh <- err
				return
			}
			resultsCh <- price
		}()
	}

	for {
		select {
		case err := <-errCh:
			if err != nil {
				return nil, err
			}
		case result := <-resultsCh:
			results = append(results, result)
			if len(results) == len(itemCodes) {
				return results, nil
			}
		}
	}
}
