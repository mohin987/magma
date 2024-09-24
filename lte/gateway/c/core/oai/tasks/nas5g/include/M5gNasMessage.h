/*
   Copyright 2020 The Magma Authors.
   This source code is licensed under the BSD-style license found in the
   LICENSE file in the root directory of this source tree.
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
 */

#pragma once
namespace magma5g {
#define M5G_SESSION_MANAGEMENT_MESSAGES 0x2e
#define M5G_MOBILITY_MANAGEMENT_MESSAGES 0x7e

// IEI for Mobility Management Message
#define M5GS_MOBILE_IDENTITY 0X77
#define PLMN_LIST 0x4a
#define TAI_LIST 0x54
#define ALLOWED_NSSAI 0x15
#define REJECTED_NSSAI 0x11
#define M5GS_NETWORK_FEATURE_SUPPORT 0X21
#define PDU_SESSION_STATUS 0x50
#define PDU_SESSION_REACTIVATION_RESULT 0x26
#define PDU_SESSION_REACTIVATION_RESULT_ERROR 0X72
#define LADN_INFORMATION 0x79
#define MICO_INDICATION 0xB0
#define NETWORK_SLICING_INDICATION 0x90
#define SERVICE_AREA_LIST 0X27
#define GPRS_TIMER3 0x5E
#define GPRS_TIMER2 0x5D
#define EMERGENCY_NUMBER_LIST 0x34
#define EXTENDED_NUMBER_EMERGENCY_LIST 0x7A
#define SOR_TRANSPARANT_CONTAINER 0x73
#define EAP_MESSAGE 0x78
#define NSSAI_INCLUSION_MODE 0xA0
#define OPERATOR_DEFINED_ACCESS_CATEGORY_DEF 0x76
#define M5GS_DRX_PARAMETERS 0x51
#define NON_3GPP_NW_PROVIDED_POLICIES 0xD0
#define EPS_BEARER_CONTEXT_STATUS 0x60
#define NAS_KEY_IDENTIFIER 0xC0
#define M5GMM_CAPABILITY 0x10
#define UE_SECURITY_CAPABILITY 0x2E
#define S1_UE_NETWORK_CAPABILITY 0X17
#define REQUESTED_NSSAI 0x2F
#define UP_LINK_DATA_STATUS 0x40
#define UE_STATUS 0x2B
#define ALLOWED_PDU_SESSION_STATUS 0x25
#define UE_USSAGE_SETTING 0x18
#define EPS_NAS_MESSAGE_CONTAINER 0x70
#define PAYLOAD_CONTAINER_TYPE 0X80
#define PAYLOAD_CONTAINER 0x7B
#define M5GS_UPDATE_TYPE 0x53
#define NAS_MESSAGE_CONTAINER 0x71
#define TAI 0x52
#define AUTH_PARAM_RAND 0x21
#define AUTH_PARAM_AUTN 0x20
#define AUTH_RESPONSE_PARAMETER 0x2D
#define NETWORK_FEATURE 0x21

// 5G Session Management IE Types
#define REQUEST_PDU_SESSION_TYPE_TYPE 0x90
#define REQUEST_SSC_MODE_TYPE 0xA0
#define M5GSM_CAUSE 0x59
#define REQUEST_5GSM_CAPABILITY_TYPE 0x28
#define MAXIMUM_NUMBER_OF_SUPPORTED_PACKET_FILTERS_TYPE 0x55
#define REQUEST_ALWAYS_ON_PDU_SESSION_REQUESTED_TYPE 0xB0
#define REQUEST_SM_PDU_DN_REQUEST_CONTAINER_TYPE 0x39
#define REQUEST_EXTENDED_PROTOCOL_CONFIGURATION_OPTIONS_TYPE 0x7B
#define REQUEST_HEADER_COMPRESSION_CONFIGURATION_TYPE 0x66
#define REQUEST_DS_TT_ETHERNET_PORT_MAC_ADDRESS_TYPE 0x6E
#define REQUEST_UE_DS_TT_RESIDENCE_TIME_TYPE 0x6F
#define REQUEST_PORT_MANAGEMENT_INFORMATION_CONTAINER_TYPE 0x7C
#define PDU_SESSION_QOS_RULES_IE_TYPE 0x7A
#define PDU_SESSION_QOS_FLOW_DESC_IE_TYPE 0x79
#define PDU_SESSION_AMBR_IE_TYPE 0x2a

// IE field values
#define M5GS_REGISTRATION_RESULT_MAXIMUM_LENGTH 1
#define M5GSMobileIdentityMsg_GUTI 0x02
#define M3GPP_ACCESS 1
#define NON_M3GPP_ACCESS 2
#define M3GPP_ACCESS_AND_NON_3GPP_ACCESS 3
#define NOT_ALLOWED 0
#define ALLOWED 1
#define EVEN_IDENTITY_DIGITS 0
#define ODD_IDENTITY_DIGITS 1
#define M5G_IDENTITY_SUCI 1
#define IMSI_LENGTH 15
#define NATIVE_SECURITY_CONTEXT 0
#define MAPPED_SECURITY_CONTEXT 1
#define SERVICE_REJ_T3346_LEN 1
#define SERVICE_REJ_T3346_VALUE 60

} // namespace magma5g
