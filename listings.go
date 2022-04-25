package main

/*A listing is a house for sale
each listing has an id, street address, price, status
price is whole dollars between $0 and 100 billion
status can be for sale, sold or expired*/

type Listing struct {
	ID            int    `json:"id"`
	StreetAddress string `json:"streetAddress"`
	Price         uint64 `json:"price"`
	Status        string `json:"status"`
}

func (listing *Listing) IsStatusValid() bool {
	return listing.Status == "for sale" || listing.Status == "sold" || listing.Status == "expired"
}
