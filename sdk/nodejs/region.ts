// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

export let APNortheast1Region: Region = "ap-northeast-1";
export let APNortheast2Region: Region = "ap-northeast-2";
export let APSouth1Region: Region     = "ap-south-1";
export let APSouthEast2Region: Region = "ap-southeast-2";
export let APSoutheast1Region: Region = "ap-southeast-1";
export let CACentralRegion: Region    = "ca-central-1";
export let EUCentral1Region: Region   = "eu-central-1";
export let EUWest1Region: Region      = "eu-west-1";
export let EUWest2Region: Region      = "eu-west-2";
export let EUWest3Region: Region      = "eu-west-3";
export let SAEast1Region: Region      = "sa-east-1";
export let USEast1Region: Region      = "us-east-1";
export let USEast2Region: Region      = "us-east-2";
export let USWest1Region: Region      = "us-west-1";
export let USWest2Region: Region      = "us-west-2";

/**
 * A Region represents any valid Amazon region that may be targeted with deployments.
 */
export type Region =
    "ap-northeast-1" |
    "ap-northeast-2" |
    "ap-south-1"     |
    "ap-southeast-2" |
    "ap-southeast-1" |
    "ca-central-1"   |
    "eu-central-1"   |
    "eu-west-1"      |
    "eu-west-2"      |
    "eu-west-3"      |
    "sa-east-1"      |
    "us-east-1"      |
    "us-east-2"      |
    "us-west-1"      |
    "us-west-2"      ;

