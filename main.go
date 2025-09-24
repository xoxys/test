package main

import (
	"fmt"
	"os"

	// github.com/container-storage-interface/spec
	csi "github.com/container-storage-interface/spec/lib/go/csi"
	// github.com/go-viper/mapstructure/v2
	"github.com/go-viper/mapstructure/v2"
	// github.com/google/uuid
	"github.com/google/uuid"
	// github.com/kubernetes-csi/csi-lib-utils (Corrected import path)
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	// github.com/onsi/ginkgo/v2
	"github.com/onsi/ginkgo/v2"
	// github.com/onsi/gomega
	"github.com/onsi/gomega"
	// github.com/prometheus/client_golang
	"github.com/prometheus/client_golang/prometheus"
	// github.com/spf13/cobra
	"github.com/spf13/cobra"
	// github.com/spf13/pflag
	"github.com/spf13/pflag"
	// github.com/stackitcloud/stackit-sdk-go/core
	"github.com/stackitcloud/stackit-sdk-go/core/config"
	// github.com/stackitcloud/stackit-sdk-go/services/iaas
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
	// github.com/stackitcloud/stackit-sdk-go/services/loadbalancer
	"github.com/stackitcloud/stackit-sdk-go/services/loadbalancer"
	// go.uber.org/mock
	"go.uber.org/mock/gomock"
	// golang.org/x/sync
	"golang.org/x/sync/errgroup"
	// golang.org/x/sys
	"golang.org/x/sys/unix"
	// google.golang.org/grpc
	"google.golang.org/grpc"
	// google.golang.org/protobuf
	"google.golang.org/protobuf/proto"
	// gopkg.in/gcfg.v1
	"gopkg.in/gcfg.v1"
	// gopkg.in/yaml.v3
	"gopkg.in/yaml.v3"
	// k8s.io/api
	v1 "k8s.io/api/core/v1"
	// k8s.io/apimachinery
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// k8s.io/client-go
	"k8s.io/client-go/kubernetes"
	// k8s.io/cloud-provider
	cloudprovider "k8s.io/cloud-provider"
	// k8s.io/component-base
	"k8s.io/component-base/version"
	// k8s.io/klog/v2
	"k8s.io/klog/v2"
	// k8s.io/mount-utils
	"k8s.io/mount-utils"
	// k8s.io/utils
	"k8s.io/utils/ptr"
)

func main() {
	fmt.Println("Using dependencies to prevent removal by 'go mod tidy'...")

	// Use a symbol from each direct dependency
	_ = csi.VolumeCapability{}
	_ = mapstructure.Metadata{}
	_ = uuid.New()
	_ = protosanitizer.StripSecrets("foo")
	_ = ginkgo.GinkgoWriter
	_ = gomega.Ω
	_ = prometheus.NewGauge(prometheus.GaugeOpts{})
	_ = cobra.Command{}
	_ = pflag.ContinueOnError
	_ = &config.Configuration{}
	_, _ = iaas.NewAPIClient()
	_, _ = loadbalancer.NewAPIClient()
	_ = gomock.Any()
	_ = errgroup.Group{}
	_ = unix.Stat_t{}
	_ = grpc.ErrClientConnClosing
	var _ proto.Message
	_ = gcfg.FatalOnly
	_ = yaml.Node{}
	_ = v1.Pod{}
	_ = metav1.ObjectMeta{}
	var _ kubernetes.Interface
	var _ cloudprovider.Interface
	_ = version.Get()
	klog.SetOutput(os.Stderr)
	klog.Info("Using klog")
	var _ mount.Interface
	_ = ptr.To(42)

	fmt.Println("All dependencies successfully referenced.")
}
