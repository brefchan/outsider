package contract

import "net/http"

const KernelKey = "outsider:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}
