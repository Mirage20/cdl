/*
 * Copyright (c) 2018 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package v1alpha1

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Gateway struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewaySpec   `json:"spec"`
	Status GatewayStatus `json:"status"`
}

type GatewayTemplateSpec struct {
	ObjectMeta `json:"metadata,omitempty"`

	Spec GatewaySpec `json:"spec,omitempty"`
}

type GatewaySpec struct {
	APIRoutes []APIRoute `json:"apis"`
}

type APIRoute struct {
	Context     string          `json:"context"`
	Definitions []APIDefinition `json:"definitions"`
	Backend     string          `json:"backend"`
	Global      bool            `json:"global"`
}

type APIDefinition struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type GatewayStatus struct {
	OwnerCell string `json:"ownerCell"`
	HostName  string `json:"hostname"`
	Status    string `json:"status"`
}
