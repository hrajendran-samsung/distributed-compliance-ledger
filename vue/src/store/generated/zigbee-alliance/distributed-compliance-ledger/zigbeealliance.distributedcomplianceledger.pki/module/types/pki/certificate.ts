/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface Certificate {
  pemCert: string
  serialNumber: string
  issuer: string
  authorityKeyId: string
  rootSubject: string
  rootSubjectKeyId: string
  isRoot: boolean
  owner: string
  subject: string
  subjectKeyId: string
}

const baseCertificate: object = {
  pemCert: '',
  serialNumber: '',
  issuer: '',
  authorityKeyId: '',
  rootSubject: '',
  rootSubjectKeyId: '',
  isRoot: false,
  owner: '',
  subject: '',
  subjectKeyId: ''
}

export const Certificate = {
  encode(message: Certificate, writer: Writer = Writer.create()): Writer {
    if (message.pemCert !== '') {
      writer.uint32(10).string(message.pemCert)
    }
    if (message.serialNumber !== '') {
      writer.uint32(18).string(message.serialNumber)
    }
    if (message.issuer !== '') {
      writer.uint32(26).string(message.issuer)
    }
    if (message.authorityKeyId !== '') {
      writer.uint32(34).string(message.authorityKeyId)
    }
    if (message.rootSubject !== '') {
      writer.uint32(42).string(message.rootSubject)
    }
    if (message.rootSubjectKeyId !== '') {
      writer.uint32(50).string(message.rootSubjectKeyId)
    }
    if (message.isRoot === true) {
      writer.uint32(56).bool(message.isRoot)
    }
    if (message.owner !== '') {
      writer.uint32(66).string(message.owner)
    }
    if (message.subject !== '') {
      writer.uint32(74).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(82).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Certificate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseCertificate } as Certificate
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pemCert = reader.string()
          break
        case 2:
          message.serialNumber = reader.string()
          break
        case 3:
          message.issuer = reader.string()
          break
        case 4:
          message.authorityKeyId = reader.string()
          break
        case 5:
          message.rootSubject = reader.string()
          break
        case 6:
          message.rootSubjectKeyId = reader.string()
          break
        case 7:
          message.isRoot = reader.bool()
          break
        case 8:
          message.owner = reader.string()
          break
        case 9:
          message.subject = reader.string()
          break
        case 10:
          message.subjectKeyId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Certificate {
    const message = { ...baseCertificate } as Certificate
    if (object.pemCert !== undefined && object.pemCert !== null) {
      message.pemCert = String(object.pemCert)
    } else {
      message.pemCert = ''
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = String(object.serialNumber)
    } else {
      message.serialNumber = ''
    }
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = String(object.issuer)
    } else {
      message.issuer = ''
    }
    if (object.authorityKeyId !== undefined && object.authorityKeyId !== null) {
      message.authorityKeyId = String(object.authorityKeyId)
    } else {
      message.authorityKeyId = ''
    }
    if (object.rootSubject !== undefined && object.rootSubject !== null) {
      message.rootSubject = String(object.rootSubject)
    } else {
      message.rootSubject = ''
    }
    if (object.rootSubjectKeyId !== undefined && object.rootSubjectKeyId !== null) {
      message.rootSubjectKeyId = String(object.rootSubjectKeyId)
    } else {
      message.rootSubjectKeyId = ''
    }
    if (object.isRoot !== undefined && object.isRoot !== null) {
      message.isRoot = Boolean(object.isRoot)
    } else {
      message.isRoot = false
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = String(object.subject)
    } else {
      message.subject = ''
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = String(object.subjectKeyId)
    } else {
      message.subjectKeyId = ''
    }
    return message
  },

  toJSON(message: Certificate): unknown {
    const obj: any = {}
    message.pemCert !== undefined && (obj.pemCert = message.pemCert)
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber)
    message.issuer !== undefined && (obj.issuer = message.issuer)
    message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId)
    message.rootSubject !== undefined && (obj.rootSubject = message.rootSubject)
    message.rootSubjectKeyId !== undefined && (obj.rootSubjectKeyId = message.rootSubjectKeyId)
    message.isRoot !== undefined && (obj.isRoot = message.isRoot)
    message.owner !== undefined && (obj.owner = message.owner)
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<Certificate>): Certificate {
    const message = { ...baseCertificate } as Certificate
    if (object.pemCert !== undefined && object.pemCert !== null) {
      message.pemCert = object.pemCert
    } else {
      message.pemCert = ''
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = object.serialNumber
    } else {
      message.serialNumber = ''
    }
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = object.issuer
    } else {
      message.issuer = ''
    }
    if (object.authorityKeyId !== undefined && object.authorityKeyId !== null) {
      message.authorityKeyId = object.authorityKeyId
    } else {
      message.authorityKeyId = ''
    }
    if (object.rootSubject !== undefined && object.rootSubject !== null) {
      message.rootSubject = object.rootSubject
    } else {
      message.rootSubject = ''
    }
    if (object.rootSubjectKeyId !== undefined && object.rootSubjectKeyId !== null) {
      message.rootSubjectKeyId = object.rootSubjectKeyId
    } else {
      message.rootSubjectKeyId = ''
    }
    if (object.isRoot !== undefined && object.isRoot !== null) {
      message.isRoot = object.isRoot
    } else {
      message.isRoot = false
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
    }
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = object.subject
    } else {
      message.subject = ''
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = object.subjectKeyId
    } else {
      message.subjectKeyId = ''
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
