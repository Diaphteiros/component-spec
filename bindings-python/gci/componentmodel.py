import dataclasses
import enum
import typing

import dacite

dc = dataclasses.dataclass


class SchemaVersion(enum.Enum):
    V1 = 'v1'
    V2 = 'v2'


class AccessType(enum.Enum):
    OCI_REGISTRY = 'ociRegistry'
    GITHUB = 'github'


class SourceType(enum.Enum):
    GIT = 'git'


class ResourceType(enum.Enum):
    OCI_IMAGE = 'ociImage'
    WEB_DEPENDENCY = 'web'
    GENERIC = 'generic'


@dc(frozen=True)
class ResourceAccess:
    type: AccessType


@dc(frozen=True)
class OciAccess(ResourceAccess):
    imageReference: str


@dc(frozen=True)
class GithubAccess(ResourceAccess):
    repoUrl: str
    ref: str


@dc(frozen=True)
class WebAccess(ResourceAccess):
    url: str


class Provider(enum.Enum):
    '''
    internal: from repositoryContext-owner
    external: from 3rd-party (not repositoryContext-owner)
    '''
    INTERNAL = 'internal'
    EXTERNAL = 'external'


@dc(frozen=True)
class Metadata:
    schemaVersion: SchemaVersion



@dc(frozen=True)
class ComponentReference:
    name: str
    version: str


@dc(frozen=True)
class Resource:
    name: str
    version: str
    type: ResourceType
    access: typing.Union[
        ResourceAccess,
        OciAccess,
        GithubAccess,
        WebAccess,
    ]


@dc
class RepositoryContext:
    baseUrl: str
    type: AccessType = AccessType.OCI_REGISTRY


@dc
class ComponentSource:
    name: str
    access: typing.Union[
        GithubAccess,
    ]
    type: SourceType = SourceType.GIT


@dc
class Component:
    name: str    # must be valid URL w/o schema
    version: str # relaxed semver

    repositoryContexts: typing.List[RepositoryContext]
    provider: Provider

    sources: typing.List[ComponentSource]
    componentReferences: typing.List[ComponentReference]
    localResources: typing.List[Resource]
    externalResources: typing.List[Resource]


@dc
class ComponentDescriptor:
    meta: Metadata
    component: Component

    @staticmethod
    def from_dict(component_descriptor_dict: dict):
        component_descriptor = dacite.from_dict(
            data_class=ComponentDescriptor,
            data=component_descriptor_dict,
            config=dacite.Config(
                cast=[
                    AccessType,
                    Provider,
                    ResourceType,
                    SchemaVersion,
                    SourceType,
                ]
            )
        )

        return component_descriptor