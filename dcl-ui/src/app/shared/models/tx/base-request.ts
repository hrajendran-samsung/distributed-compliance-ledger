/**
 * Copyright 2020 DSR Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { JsonObject, JsonProperty } from 'json2typescript';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('BaseReq')
export class BaseReq {

  @JsonProperty('chain_id', String)
  chainId: string = undefined;

  @JsonProperty('account_number', AminoNumberConverter)
  accountNumber: number = undefined;

  @JsonProperty('sequence', AminoNumberConverter)
  sequence: number = undefined;

  @JsonProperty('from', String)
  from = '';

  constructor(init?: Partial<BaseReq>) {
    Object.assign(this, init);
  }

  setFrom(from: string) {
    this.from = from;
  }
}