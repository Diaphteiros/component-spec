## Topics for Component-Spec-Meeting

- localOciBlob is a bad name, localBlob is much better

- Repository Context in CD
  - current schema only allows OCI-Ref, this needs to be extensible
  - Should OCI ref be part of this spec?
  
- CTF:
  - Should be included in OCM?
  
- What are the next step to specify?
  - particular types
  - Transport, CTF
  - Signing 
  - etc.

## Todo/Discuss
- Terminology section: beside others must describe landscape, component version, snapshot 
- complete component descriptor example
- change technical artifacts to software artifacts 
- references in the text to the CD spec for particular entries
- Some chapter with terminology (see https://github.com/openservicebrokerapi/servicebroker/blob/v2.16/spec.md#notations-and-terminology)
- Is Component Version Descriptor better than Component Descriptor because it describes a Component Version?
- Tables for presenting field descriptions (see https://github.com/openservicebrokerapi/servicebroker/blob/v2.16/spec.md)
- Tables for input export params of functions (see https://github.com/openservicebrokerapi/servicebroker/blob/v2.16/spec.md) 
- List with all errors and their semantics
- Use-Cases / Scenarios with diagrams for Component Repos
- Clarify that components are not allowed to have two version with the same number but different formats, e.g. v0.1.0 and 0.1.0
