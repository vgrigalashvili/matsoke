package types

type (
	Account struct {
		id    string
		owner *Person
	}
	AccountArguments struct {
		ID    string  `json:"id"`
		Owner *Person `json:"owner"`
	}
)

func (a *Account) ID() string {
	return a.id
}

func (a *Account) Owner() *Person {
	return &Person{
		firstName: a.owner.firstName,
	}
}
