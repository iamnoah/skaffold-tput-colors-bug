package main

// give mod vendor something to do
import (
	_ "github.com/aws/aws-sdk-go"
	_ "github.com/docker/distribution"
	_ "github.com/evanphx/json-patch"
	_ "github.com/getsentry/sentry-go"
	_ "github.com/go-logr/logr"
	_ "github.com/google/uuid"
	_ "github.com/gorilla/handlers"
	_ "github.com/gorilla/mux"
	_ "github.com/mailgun/groupcache/v2"
	_ "github.com/mitchellh/mapstructure"
	_ "github.com/nxadm/tail"
	_ "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	_ "github.com/opentracing/opentracing-go"
	_ "github.com/pkg/errors"
	_ "github.com/spf13/viper"
	_ "github.com/stretchr/testify"
	_ "go.uber.org/atomic"
	_ "go.uber.org/zap"
	_ "google.golang.org/grpc"
	_ "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	_ "gopkg.in/launchdarkly/go-sdk-common.v2"
	_ "gopkg.in/launchdarkly/go-server-sdk.v5"
	_ "gotest.tools"
	_ "gotest.tools/v3"
	_ "k8s.io/api"
	_ "k8s.io/apimachinery/pkg/runtime/serializer"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/utils/exec"
	_ "sigs.k8s.io/controller-runtime"
	_ "sigs.k8s.io/kustomize/kyaml"
	_ "sigs.k8s.io/yaml"
)

func main() {
	println("hello world")
}
