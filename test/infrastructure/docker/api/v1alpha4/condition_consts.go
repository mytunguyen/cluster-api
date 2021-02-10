/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha4

import clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"

// Conditions and condition Reasons for the DockerMachine object

const (
	// ContainerProvisionedCondition documents the status of the provisioning of the container
	// generated by a DockerMachine.
	//
	// NOTE: When the container provisioning starts the process completes almost immediately and within
	// the same reconciliation, so the user will always see a transition from Wait to Provisioned without
	// having evidence that the operation is started/is in progress.
	ContainerProvisionedCondition clusterv1.ConditionType = "ContainerProvisioned"

	// WaitingForClusterInfrastructureReason (Severity=Info) documents a DockerMachine waiting for the cluster
	// infrastructure to be ready before starting to create the container that provides the DockerMachine
	// infrastructure.
	WaitingForClusterInfrastructureReason = "WaitingForClusterInfrastructure"

	// WaitingForBootstrapDataReason (Severity=Info) documents a DockerMachine waiting for the bootstrap
	// script to be ready before starting to create the container that provides the DockerMachine infrastructure.
	WaitingForBootstrapDataReason = "WaitingForBootstrapData"

	// ContainerProvisioningFailedReason (Severity=Warning) documents a DockerMachine controller detecting
	// an error while provisioning the container that provides the DockerMachine infrastructure; those kind of
	// errors are usually transient and failed provisioning are automatically re-tried by the controller.
	ContainerProvisioningFailedReason = "ContainerProvisioningFailed"
)

const (
	// BootstrapExecSucceededCondition provides an observation of the DockerMachine bootstrap process.
	// 	It is set based on successful execution of bootstrap commands and on the existence of
	//	the /run/cluster-api/bootstrap-success.complete file.
	// The condition gets generated after ContainerProvisionedCondition is True.
	//
	// NOTE as a difference from other providers, container provisioning and bootstrap are directly managed
	// by the DockerMachine controller (not by cloud-init).
	BootstrapExecSucceededCondition clusterv1.ConditionType = "BootstrapExecSucceeded"

	// BootstrappingReason documents (Severity=Info) a DockerMachine currently executing the bootstrap
	// script that creates the Kubernetes node on the newly provisioned machine infrastructure.
	BootstrappingReason = "Bootstrapping"

	// BootstrapFailedReason documents (Severity=Warning) a DockerMachine controller detecting an error while
	// bootstrapping the Kubernetes node on the machine just provisioned; those kind of errors are usually
	// transient and failed bootstrap are automatically re-tried by the controller.
	BootstrapFailedReason = "BootstrapFailed"
)

// Conditions and condition Reasons for the DockerCluster object

const (
	// LoadBalancerAvailableCondition documents the availability of the container that implements the cluster load balancer.
	//
	// NOTE: When the load balancer provisioning starts the process completes almost immediately and within
	// the same reconciliation, so the user will always see a transition from no condition to available without
	// having evidence that the operation is started/is in progress.
	LoadBalancerAvailableCondition clusterv1.ConditionType = "LoadBalancerAvailable"

	// LoadBalancerProvisioningFailedReason (Severity=Warning) documents a DockerCluster controller detecting
	// an error while provisioning the container that provides the cluster load balancer.; those kind of
	// errors are usually transient and failed provisioning are automatically re-tried by the controller.
	LoadBalancerProvisioningFailedReason = "LoadBalancerProvisioningFailed"
)
