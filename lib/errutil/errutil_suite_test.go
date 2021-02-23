package errutil

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/phsm/pandora/lib/ginkgoutil"
)

func TestErrutil(t *testing.T) {
	ginkgoutil.RunSuite(t, "Errutil Suite")
}

var _ = Describe("Iscoreutil.IsNotCtxErroror", func() {
	canceledContext, cancel := context.WithCancel(context.Background())
	cancel()

	It("nil error", func() {
		Expect(IsNotCtxError(context.Background(), nil)).To(BeFalse())
	})

	It("context error", func() {
		Expect(IsNotCtxError(canceledContext, context.Canceled)).To(BeFalse())
	})

	It("caused by context error", func() {
		Expect(IsNotCtxError(canceledContext, errors.Wrap(context.Canceled, "new err"))).To(BeFalse())
	})

	It("usual error", func() {
		err := errors.New("new err")
		Expect(IsNotCtxError(canceledContext, err)).To(BeTrue())
		Expect(IsNotCtxError(context.Background(), err)).To(BeTrue())
	})
})
