package main

// give mod vendor something to do
import (
	_ "github.com/getsentry/sentry-go"
	_ "github.com/gorilla/mux"
	_ "github.com/opentracing-contrib/go-stdlib/nethttp"
	_ "k8s.io/api/core/v1"
	_ "k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/apimachinery/pkg/types"
	_ "k8s.io/apimachinery/pkg/util/intstr"
	_ "k8s.io/utils/pointer"
	_ "sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	println("hello world")
}
