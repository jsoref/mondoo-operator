/*
Copyright 2022 Mondoo, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestMergeEnv_NoDuplicates(t *testing.T) {
	a := []corev1.EnvVar{
		{Name: "a", Value: "2"},
		{Name: "a1", Value: "3"},
	}

	b := []corev1.EnvVar{
		{Name: "b", Value: "6"},
		{Name: "b1", Value: "7"},
	}

	env := MergeEnv(a, b)
	expected := []corev1.EnvVar{
		{Name: "a", Value: "2"},
		{Name: "a1", Value: "3"},
		{Name: "b", Value: "6"},
		{Name: "b1", Value: "7"},
	}

	assert.ElementsMatch(t, expected, env)
}

func TestMergeEnv_Duplicates(t *testing.T) {
	a := []corev1.EnvVar{
		{Name: "a", Value: "2"},
		{Name: "a1", Value: "3"},
	}

	b := []corev1.EnvVar{
		{Name: "b", Value: "6"},
		{Name: "b1", Value: "7"},
		{Name: "a1", Value: "17"},
	}

	env := MergeEnv(a, b)
	expected := []corev1.EnvVar{
		{Name: "a", Value: "2"},
		{Name: "a1", Value: "17"}, // value is from b
		{Name: "b", Value: "6"},
		{Name: "b1", Value: "7"},
	}

	assert.ElementsMatch(t, expected, env)
}

func TestMergeEnv_AEmpty(t *testing.T) {
	a := []corev1.EnvVar{}

	b := []corev1.EnvVar{
		{Name: "b", Value: "6"},
		{Name: "b1", Value: "7"},
		{Name: "a1", Value: "17"},
	}

	env := MergeEnv(a, b)
	expected := []corev1.EnvVar{
		{Name: "b", Value: "6"},
		{Name: "b1", Value: "7"},
		{Name: "a1", Value: "17"},
	}

	assert.ElementsMatch(t, expected, env)
}

func TestMergeEnv_BEmpty(t *testing.T) {
	a := []corev1.EnvVar{
		{Name: "a", Value: "2"},
		{Name: "a1", Value: "3"},
	}

	b := []corev1.EnvVar{}

	env := MergeEnv(a, b)
	expected := []corev1.EnvVar{
		{Name: "a", Value: "2"},
		{Name: "a1", Value: "3"},
	}

	assert.ElementsMatch(t, expected, env)
}
