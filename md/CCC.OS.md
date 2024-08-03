# CCC.OS: Object Storage

| Control Id | Service Taxonomy Id | Control |
| ---------- | ------------------- | ------- |
| CCC.OS.C1  | CCC-020115          | Prevent unencrypted requests to object storage bucket |
| CCC.OS.C2  | CCC-020114          | Ensure data encryption at rest |
| CCC.OS.C3  | CCC-020116          | Implement multi-factor authentication (MFA) for access |
| CCC.OS.C4  | CCC-020112          | Maintain immutable backups of data |
| CCC.OS.C5  | CCC-020118          | Log all access and changes to object storage |

---

## CCC.OS.C1: Prevent unencrypted requests to object storage bucket

- Corresponding Feature: CCC-020115
- NIST CSF: PR.DS-2
- MITRE ATT&CK TTP: T1573

### Objective

Prevent any unencrypted requests to the object storage bucket, ensuring that all communications are encrypted in transit to protect data integrity and confidentiality.

### Control Mappings

- CCM: IVS-09, DSI-03
- ISO_27001: 2013 A.13.1.1
- NIST_800_53: SC-8, SC-13

### Testing Requirements

The following validations must be performed against corresponding Control Implementation capabilities to ensure the Control Objective is thoroughly assessed:

1. **CCC.OS.C1.01**: All supported network data protocols must be running on secure channels.
1. **CCC.OS.C1.02**: All clear text channels should be disabled.
1. **CCC.OS.C1.03**: The cipher suite implemented for ensuring the integrity and confidentiality of data should conform with the latest suggested cipher suites. [NIST proposed latest standard cipher suites](<[#](https://csrc.nist.gov/pubs/sp/800/52/r2/final)>).

---

## CCC.OS.C2: Ensure data encryption at rest

- Corresponding Feature: CCC-020114
- NIST CSF: PR.DS-1
- MITRE ATT&CK TTP: T1486

### Objective

Ensure that all data stored within the object storage service is encrypted at rest to maintain confidentiality and integrity.

### Control Mappings

- ISO_27001: 2013 A.10.1.1
- NIST_800_53: SC-28
- CCM: DSI-01, DSI-02

### Testing Requirements

The following validations must be performed against corresponding Control Implementation capabilities to ensure the Control Objective is thoroughly assessed:

1. **CCC.OS.C2.01**: Verify that data stored in the object storage bucket is encrypted using industry-standard algorithms.
1. **CCC.OS.C2.02**: Ensure that encryption keys are managed securely and rotated periodically.
1. **CCC.OS.C2.03**: Confirm that decryption is only possible through authorized access mechanisms.

---

## CCC.OS.C3: Implement multi-factor authentication (MFA) for access

- Corresponding Feature: CCC-020116
- NIST CSF: PR.AC-7
- MITRE ATT&CK TTP: T1078

### Objective

Ensure that all human user access to object storage buckets requires multi-factor authentication (MFA), minimizing the risk of unauthorized access by enforcing strong authentication mechanisms.

### Control Mappings

- CCM: IAM-03, IAM-08
- ISO_27001: 2013 A.9.4.2
- NIST_800_53: IA-2

### Testing Requirements

The following validations must be performed against corresponding Control Implementation capabilities to ensure the Control Objective is thoroughly assessed:

1. **CCC.OS.C3.02**: Ensure that MFA is required for all administrative access to the storage management interface.
1. **CCC.OS.C3.03**: Confirm that users are unable to access the object storage bucket without completing MFA.
1. **CCC.OS.C3.01**: Verify that MFA is enforced for all access attempts to the object storage bucket.

---

## CCC.OS.C4: Maintain immutable backups of data

- Corresponding Feature: CCC-020112
- NIST CSF: PR.DS-1
- MITRE ATT&CK TTP: T1485

### Objective

Ensure that data stored in the object storage bucket is immutable for a defined period, preventing unauthorized modifications or deletions and thereby mitigating data destruction.

### Control Mappings

- CCM: DSI-05, DSI-07
- ISO_27001: 2013 A.12.3.1
- NIST_800_53: CP-9

### Testing Requirements

The following validations must be performed against corresponding Control Implementation capabilities to ensure the Control Objective is thoroughly assessed:

1. **CCC.OS.C4.01**: Verify that data in the object storage bucket is protected by immutability settings.
1. **CCC.OS.C4.02**: Ensure that attempts to modify or delete data within the immutability period are denied.
1. **CCC.OS.C4.03**: Confirm that immutable data remains unchanged throughout the defined retention period.

---

## CCC.OS.C5: Log all access and changes to object storage

- Corresponding Feature: CCC-020118
- NIST CSF: DE.AE-3
- MITRE ATT&CK TTP: T1530

### Objective

Ensure that all access and changes to the object storage bucket are logged to maintain a detailed audit trail for security and compliance purposes.

### Control Mappings

- CCM: DSI-06, STA-04
- ISO_27001: 2013 A.12.4.1
- NIST_800_53: AU-2, AU-3

### Testing Requirements

The following validations must be performed against corresponding Control Implementation capabilities to ensure the Control Objective is thoroughly assessed:

1. **CCC.OS.C5.01**: Verify that all access attempts to the object storage bucket are logged.
1. **CCC.OS.C5.02**: Ensure that all changes to the object storage bucket configurations are logged.
1. **CCC.OS.C5.03**: Confirm that logs are protected against unauthorized access and tampering.

---

