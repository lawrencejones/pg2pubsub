package imports_test

import (
	"context"
	"testing"

	"github.com/lawrencejones/pgsink/internal/dbschema/pgsink/model"

	kitlog "github.com/go-kit/kit/log"
	"github.com/jackc/pgx/v4"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var logger = kitlog.NewLogfmtLogger(GinkgoWriter)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "pkg/imports")
}

type importerFunc func(ctx context.Context, logger kitlog.Logger, tx pgx.Tx, job model.ImportJobs) error

func (f importerFunc) Do(ctx context.Context, logger kitlog.Logger, tx pgx.Tx, job model.ImportJobs) error {
	return f(ctx, logger, tx, job)
}

func noopImporter(err error) importerFunc {
	return importerFunc(func(_ context.Context, _ kitlog.Logger, _ pgx.Tx, _ model.ImportJobs) error {
		return err
	})
}

func waitImporter(ctx context.Context, done chan struct{}, err error) (chan model.ImportJobs, importerFunc) {
	acquired := make(chan model.ImportJobs, 1)
	return acquired, importerFunc(func(_ context.Context, _ kitlog.Logger, _ pgx.Tx, job model.ImportJobs) error {
		acquired <- job
		close(acquired)

		select {
		case <-ctx.Done():
		case <-done:
		}

		return err
	})
}
