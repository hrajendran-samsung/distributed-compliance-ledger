import { ApprovedCertificates } from '../pki/approved_certificates';
import { ProposedCertificate } from '../pki/proposed_certificate';
import { ChildCertificates } from '../pki/child_certificates';
import { ProposedCertificateRevocation } from '../pki/proposed_certificate_revocation';
import { RevokedCertificates } from '../pki/revoked_certificates';
import { UniqueCertificate } from '../pki/unique_certificate';
import { ApprovedRootCertificates } from '../pki/approved_root_certificates';
import { RevokedRootCertificates } from '../pki/revoked_root_certificates';
import { ApprovedCertificatesBySubject } from '../pki/approved_certificates_by_subject';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
/** GenesisState defines the pki module's genesis state. */
export interface GenesisState {
    approvedCertificatesList: ApprovedCertificates[];
    proposedCertificateList: ProposedCertificate[];
    childCertificatesList: ChildCertificates[];
    proposedCertificateRevocationList: ProposedCertificateRevocation[];
    revokedCertificatesList: RevokedCertificates[];
    uniqueCertificateList: UniqueCertificate[];
    approvedRootCertificates: ApprovedRootCertificates | undefined;
    revokedRootCertificates: RevokedRootCertificates | undefined;
    /** this line is used by starport scaffolding # genesis/proto/state */
    approvedCertificatesBySubjectList: ApprovedCertificatesBySubject[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
