package model

// CreditPool represents the information about credit pool.
type CreditPool struct {
	// ID represents the unique credit pool ID.
	ID string `json:"id"`
	// Name represents the credit pool name.
	Name string `json:"name"`
	// Credits represents the amount of credits in this credit pool.
	Credits float64 `json:"credits"`
	// Servers represents the number of servers in this credit pool.
	Servers int `json:"servers"`
	// Owner represents the unique account ID of the pool owner.
	Owner string `json:"owner"`
	// IsOwner represents whether the current user is the owner of this credit pool.
	IsOwner bool `json:"isOwner"`
	// Members represents the number of members in this credit pool.
	Members int `json:"members"`
	// OwnShare represents the share of the credits int this pool that are owned by current user.
	OwnShare float64 `json:"ownShare"`
	// OwnCredits represents the number of credits in this pool that are owned by current user.
	OwnCredits float64 `json:"ownCredits"`
}

// CreditPools represents the information about credit pools.
type CreditPools []CreditPool

// CreditPoolMember represents the information about credit pool member.
type CreditPoolMember struct {
	// Account represents the unique account ID of the pool member.
	Account string `json:"account"`
	// Name represents the account name of the pool member.
	Name string `json:"name"`
	// Share represents the share of the credits in this pool that are owned by the account.
	Share float64 `json:"share"`
	// Credits represents the number of credits in this pool that are owned by the account.
	Credits float64 `json:"credits"`
	// IsOwner represents whether the account is the owner of this credit pool.
	IsOwner bool `json:"isOwner"`
}

// CreditPoolMembers represents the information about credit pool members.
type CreditPoolMembers []CreditPoolMember
