package orderbook

// Depth retruns the depth.
func (ob *OrderBook) Depth() *Depth {
	defer ob.RUnlock()
	ob.RLock()

	asks := make([]*PriceLevel, 0)
	level := ob.asks.MinPriceQueue()

	for level != nil {
		asks = append(asks, NewPriceLevel(level.price, level.amount))
		level = ob.asks.GreaterThan(level.price)
	}

	bids := make([]*PriceLevel, 0)
	level = ob.bids.MaxPriceQueue()

	for level != nil {
		bids = append(bids, NewPriceLevel(level.price, level.amount))
		level = ob.bids.LessThan(level.price)
	}

	return &Depth{bids, asks}
}

// NDepth retruns the depth from best bid/ask back nAsks and nBids.
func (ob *OrderBook) NDepth(nAsks int, nBids int) *Depth {
	defer ob.RUnlock()
	ob.RLock()

	asks := make([]*PriceLevel, 0)
	level := ob.asks.MinPriceQueue()
	i := 1
	for level != nil {
		asks = append(asks, NewPriceLevel(level.price, level.amount))
		level = ob.asks.GreaterThan(level.price)

		i++
		if i > nAsks {
			level = nil
		}
	}

	bids := make([]*PriceLevel, 0)
	level = ob.bids.MaxPriceQueue()
	i = 1
	for level != nil {
		bids = append(bids, NewPriceLevel(level.price, level.amount))
		level = ob.bids.LessThan(level.price)

		i++
		if i > nBids {
			level = nil
		}
	}

	return &Depth{bids, asks}
}