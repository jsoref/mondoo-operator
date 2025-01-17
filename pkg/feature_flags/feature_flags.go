package feature_flags

import (
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

const FeatureFlagPrefix = "FEATURE_"

var (
	enablePodDiscovery      bool
	enableWorkloadDiscovery bool
	allFeatureFlags         = make(map[string]string)
)

func init() {
	envs := os.Environ()
	for _, e := range envs {
		// If it has the feature flag prefix, then parse the env var.
		if strings.HasPrefix(e, FeatureFlagPrefix) {
			val := strings.Split(e, "=")
			allFeatureFlags[val[0]] = val[1]
			setGlobalFlags(val[0], val[1])
		}
	}
}

func AllFeatureFlags() map[string]string {
	return allFeatureFlags
}

func AllFeatureFlagsAsEnv() []corev1.EnvVar {
	var env []corev1.EnvVar
	for k, v := range allFeatureFlags {
		env = append(env, corev1.EnvVar{Name: k, Value: v})
	}
	return env
}

func GetEnablePodDiscovery() bool {
	return enablePodDiscovery
}

func GetEnableWorkloadDiscovery() bool {
	return enableWorkloadDiscovery
}

func setGlobalFlags(k, v string) {
	switch k {
	case "FEATURE_DISCOVER_PODS":
		enablePodDiscovery = true
	case "FEATURE_DISCOVER_WORKLOADS":
		enableWorkloadDiscovery = true
	}
}
