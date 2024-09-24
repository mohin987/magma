/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#include "lte/gateway/c/core/oai/tasks/nas/ies/SecurityHeaderType.hpp"

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif
#include "lte/gateway/c/core/common/assertions.h"
#ifdef __cplusplus
}
#endif

//------------------------------------------------------------------------------
int decode_security_header_type(security_header_type_t *securityheadertype,
                                uint8_t iei, uint8_t *buffer, uint32_t len) {
  Fatal("TODO Implement decode_security_header_type");
  return -1;
}

//------------------------------------------------------------------------------
int encode_security_header_type(security_header_type_t *securityheadertype,
                                uint8_t iei, uint8_t *buffer, uint32_t len) {
  Fatal("TODO Implement encode_security_header_type");
  return -1;
}
