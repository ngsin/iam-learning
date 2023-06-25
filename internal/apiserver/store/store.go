package store

//go:generate mockgen -self_package=github.com/ngsin/iam-learning/internal/apiserver/store -destination mock_store.go -package store github.com/ngsin/iam-learning/internal/apiserver/store  Factory,UserStore,SecretStore,PolicyStore

var client Factory

// Factory defines the iam platform storage interface.
type Factory interface {
	Users() UserStore
	Secrets() SecretStore
	Policies() PolicyStore
	PolicyAudits() PolicyAuditStore
	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}