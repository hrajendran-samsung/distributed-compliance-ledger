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

import { Injectable } from '@angular/core';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { HttpParams } from '@angular/common/http';
import { TxService } from '../tx/tx.service';
import { ProposedCertificates } from '../../shared/models/pki/proposed-certificates';
import { ApprovedCertificates } from 'src/app/shared/models/pki/approved-certificates';
import { ProposedCertificate } from 'src/app/shared/models/pki/proposed-certificate';
import { MessageApproveAddX509RootCert } from 'src/app/shared/models/pki/message-approve-add-x509-root-certificate';

@Injectable({
  providedIn: 'root'
})
export class PkiService {

  private baseUrl = 'pki/certs';

  constructor(
    private http: HttpExtensionsService,
    private txService: TxService) {
  }

  getAllX509RootCerts(skip: number = 0, take: number = 0): Observable<ApprovedCertificates> {
    const params = new HttpParams()
      .append('skip', skip.toString())
      .append('take', take.toString());

    return this.http.get(`${this.baseUrl}/root`, ApprovedCertificates, params);
  }

  getAllX509Certs(skip: number = 0, take: number = 0): Observable<ApprovedCertificates> {
    const params = new HttpParams()
      .append('skip', skip.toString())
      .append('take', take.toString());

    return this.http.get(`${this.baseUrl}`, ApprovedCertificates, params);
  }

  getAllSubjectX509Certs(subject: string, skip: number = 0, take: number = 0): Observable<ApprovedCertificates> {
    const params = new HttpParams()
      .append('skip', skip.toString())
      .append('take', take.toString());

    return this.http.get(`${this.baseUrl}/${subject}`, ApprovedCertificates, params);
  }

  getX509Cert(subject: string, subjectKeyId: string): Observable<ApprovedCertificates> {
    const params = new HttpParams().append('prev_height', 'true');
    return this.http.get(`${this.baseUrl}/${subject}/${subjectKeyId}`, ApprovedCertificates, params);
  }

  getAllProposedX509RootCerts(skip: number = 0, take: number = 0): Observable<ProposedCertificates> {
    const params = new HttpParams()
      .append('skip', skip.toString())
      .append('take', take.toString());

    return this.http.get(`${this.baseUrl}/proposed/root`, ProposedCertificates, params);
  }

  getProposedX509RootCert(subject: string, subjectKeyId: string): Observable<ProposedCertificate> {
    const params = new HttpParams().append('prev_height', 'true');
    return this.http.get(`${this.baseUrl}/proposed/root/${subject}/${subjectKeyId}`, ProposedCertificate, params);
  }

  approveCertificate(item: ProposedCertificate) {
    const message = new MessageApproveAddX509RootCert(item.subject, item.subjectKeyId);
    this.txService.goPreview(message, '/proposed-certificates');
  }
}
