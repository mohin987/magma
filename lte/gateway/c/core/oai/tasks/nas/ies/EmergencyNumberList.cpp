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

#include "lte/gateway/c/core/oai/tasks/nas/ies/EmergencyNumberList.hpp"

#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif
#include "lte/gateway/c/core/oai/common/TLVDecoder.h"
#include "lte/gateway/c/core/oai/common/TLVEncoder.h"
#ifdef __cplusplus
}
#endif

int decode_emergency_number_list(EmergencyNumberList *emergencynumberlist,
                                 uint8_t iei, uint8_t *buffer, uint32_t len) {
  int decoded = 0;
  uint8_t ielen = 0;

  if (iei > 0) {
    CHECK_IEI_DECODER(iei, *buffer);
    decoded++;
  }

  ielen = *(buffer + decoded);
  decoded++;
  CHECK_LENGTH_DECODER(len - decoded, ielen);
  emergencynumberlist->lengthofemergency = *(buffer + decoded);
  decoded++;
  emergencynumberlist->emergencyservicecategoryvalue =
      *(buffer + decoded) & 0x1f;
  decoded++;
  return decoded;
}

int encode_emergency_number_list(EmergencyNumberList *emergencynumberlist,
                                 uint8_t iei, uint8_t *buffer, uint32_t len) {
  uint8_t *lenPtr;
  uint32_t encoded = 0;

  /*
   * Checking IEI and pointer
   */
  CHECK_PDU_POINTER_AND_LENGTH_ENCODER(
      buffer, EMERGENCY_NUMBER_LIST_MINIMUM_LENGTH, len);

  if (iei > 0) {
    *buffer = iei;
    encoded++;
  }

  lenPtr = (buffer + encoded);
  encoded++;
  *(buffer + encoded) = emergencynumberlist->lengthofemergency;
  encoded++;
  *(buffer + encoded) =
      0x00 | (emergencynumberlist->emergencyservicecategoryvalue & 0x1f);
  encoded++;
  *lenPtr = encoded - 1 - ((iei > 0) ? 1 : 0);
  return encoded;
}

void dump_emergency_number_list_xml(EmergencyNumberList *emergencynumberlist,
                                    uint8_t iei) {
  OAILOG_DEBUG(LOG_NAS, "<Emergency Number List>\n");

  if (iei > 0)
    /*
     * Don't display IEI if = 0
     */
    OAILOG_DEBUG(LOG_NAS, "    <IEI>0x%X</IEI>\n", iei);

  OAILOG_DEBUG(LOG_NAS, "    <Length of emergency>%u</Length of emergency>\n",
               emergencynumberlist->lengthofemergency);
  OAILOG_DEBUG(
      LOG_NAS,
      "    <Emergency service category value>%u</Emergency service category "
      "value>\n",
      emergencynumberlist->emergencyservicecategoryvalue);
  OAILOG_DEBUG(LOG_NAS, "</Emergency Number List>\n");
}
