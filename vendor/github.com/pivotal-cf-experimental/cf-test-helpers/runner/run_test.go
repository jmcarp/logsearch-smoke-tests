package runner_test

import (
	"time"

	"github.com/pivotal-cf-experimental/cf-test-helpers/runner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

const cmdTimeout = 30 * time.Second

var _ = Describe("Run", func() {
	It("runs the given command in a cmdtest Session", func() {
		session := runner.Run("bash", "-c", "echo hi out; echo hi err 1>&2; exit 42").Wait(cmdTimeout)
		Expect(session).To(Exit(42))
		Expect(session.Out).To(Say("hi out"))
		Expect(session.Err).To(Say("hi err"))
	})
})

var _ = Describe("Curl", func() {
	It("outputs the body of the given URL", func() {
		session := runner.Curl("-I", "http://example.com").Wait(cmdTimeout)
		Expect(session).To(Exit(0))
		Expect(session.Out).To(Say("HTTP/1.1 200 OK"))
	})
})

var _ = Describe("cmdRunner", func() {

	Describe("Run with defaults", func() {
		It("does nothing when the command succeeds before the timeout", func() {
			failures := InterceptGomegaFailures(func() {
				session := runner.Run("bash", "-c", "echo hi out; echo hi err 1>&2; exit 0")
				session.Wait(cmdTimeout)
			})
			Expect(failures).To(BeEmpty())
		})
	})
})
