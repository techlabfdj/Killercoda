package stores

// Store interface is a storage implementation for Item
type Store interface {
	AddData(id string, itemToAdd interface{}) (err *StoreError)
	GetDatas(offset, limit int) (items []interface{}, totalCount int, err *StoreError)
	GetData(id string) (item interface{}, err *StoreError)
	UpdateData(id string, itemToUpdate interface{}) (err *StoreError)
	DeleteData(id string) (err *StoreError)
}

// StoreError is used toreport errors occuring during store operations
type StoreError struct {
	ErrorCode        string
	ErrorDescription string
}

func (e *StoreError) Error() string {
	return e.ErrorDescription
}
