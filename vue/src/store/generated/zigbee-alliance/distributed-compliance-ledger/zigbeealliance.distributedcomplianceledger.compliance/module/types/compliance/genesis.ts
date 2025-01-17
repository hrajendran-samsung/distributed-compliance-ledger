/* eslint-disable */
import { ComplianceInfo } from '../compliance/compliance_info'
import { CertifiedModel } from '../compliance/certified_model'
import { RevokedModel } from '../compliance/revoked_model'
import { ProvisionalModel } from '../compliance/provisional_model'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance'

/** GenesisState defines the compliance module's genesis state. */
export interface GenesisState {
  complianceInfoList: ComplianceInfo[]
  certifiedModelList: CertifiedModel[]
  revokedModelList: RevokedModel[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  provisionalModelList: ProvisionalModel[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.complianceInfoList) {
      ComplianceInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.certifiedModelList) {
      CertifiedModel.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    for (const v of message.revokedModelList) {
      RevokedModel.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    for (const v of message.provisionalModelList) {
      ProvisionalModel.encode(v!, writer.uint32(34).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.complianceInfoList = []
    message.certifiedModelList = []
    message.revokedModelList = []
    message.provisionalModelList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.complianceInfoList.push(ComplianceInfo.decode(reader, reader.uint32()))
          break
        case 2:
          message.certifiedModelList.push(CertifiedModel.decode(reader, reader.uint32()))
          break
        case 3:
          message.revokedModelList.push(RevokedModel.decode(reader, reader.uint32()))
          break
        case 4:
          message.provisionalModelList.push(ProvisionalModel.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.complianceInfoList = []
    message.certifiedModelList = []
    message.revokedModelList = []
    message.provisionalModelList = []
    if (object.complianceInfoList !== undefined && object.complianceInfoList !== null) {
      for (const e of object.complianceInfoList) {
        message.complianceInfoList.push(ComplianceInfo.fromJSON(e))
      }
    }
    if (object.certifiedModelList !== undefined && object.certifiedModelList !== null) {
      for (const e of object.certifiedModelList) {
        message.certifiedModelList.push(CertifiedModel.fromJSON(e))
      }
    }
    if (object.revokedModelList !== undefined && object.revokedModelList !== null) {
      for (const e of object.revokedModelList) {
        message.revokedModelList.push(RevokedModel.fromJSON(e))
      }
    }
    if (object.provisionalModelList !== undefined && object.provisionalModelList !== null) {
      for (const e of object.provisionalModelList) {
        message.provisionalModelList.push(ProvisionalModel.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.complianceInfoList) {
      obj.complianceInfoList = message.complianceInfoList.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined))
    } else {
      obj.complianceInfoList = []
    }
    if (message.certifiedModelList) {
      obj.certifiedModelList = message.certifiedModelList.map((e) => (e ? CertifiedModel.toJSON(e) : undefined))
    } else {
      obj.certifiedModelList = []
    }
    if (message.revokedModelList) {
      obj.revokedModelList = message.revokedModelList.map((e) => (e ? RevokedModel.toJSON(e) : undefined))
    } else {
      obj.revokedModelList = []
    }
    if (message.provisionalModelList) {
      obj.provisionalModelList = message.provisionalModelList.map((e) => (e ? ProvisionalModel.toJSON(e) : undefined))
    } else {
      obj.provisionalModelList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.complianceInfoList = []
    message.certifiedModelList = []
    message.revokedModelList = []
    message.provisionalModelList = []
    if (object.complianceInfoList !== undefined && object.complianceInfoList !== null) {
      for (const e of object.complianceInfoList) {
        message.complianceInfoList.push(ComplianceInfo.fromPartial(e))
      }
    }
    if (object.certifiedModelList !== undefined && object.certifiedModelList !== null) {
      for (const e of object.certifiedModelList) {
        message.certifiedModelList.push(CertifiedModel.fromPartial(e))
      }
    }
    if (object.revokedModelList !== undefined && object.revokedModelList !== null) {
      for (const e of object.revokedModelList) {
        message.revokedModelList.push(RevokedModel.fromPartial(e))
      }
    }
    if (object.provisionalModelList !== undefined && object.provisionalModelList !== null) {
      for (const e of object.provisionalModelList) {
        message.provisionalModelList.push(ProvisionalModel.fromPartial(e))
      }
    }
    return message
  }
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
