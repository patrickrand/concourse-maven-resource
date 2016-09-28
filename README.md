# Maven Resource

A [Concourse](http://concourse.ci) resource for interacting with artifacts in a Maven repository. 

> WIP as of 9/28/16

## Resource Type Configuration

```yaml
resource_types:
- name: maven
  type: docker-image
  source:
    repository: patrickrand/concourse-maven-resource
    tag: latest
```

## Source Configuration

* `repository`: *Required.* The URI location of the Maven repository.

* `group_id`: *Required.* The group ID of the artifact.

* `artifact_id`: *Required.* The ID of the artifact.

* `username`: *Optional.* Username for interacting with Maven respository.

* `password`: *Optional.* Password for interacting with Maven respository.


### Example

Resource configuration for a private Maven repo:

``` yaml
resources:
- name: artifact
  type: maven
  source:
    repository: http://localhost/maven-repo
    group_id: com.my.company
    artifact_id: some-plugin
```

Fetching the latest version of the source artifact:

``` yaml
- get: artifact
```

## Behavior

### `check`: Check for new versions of the artifact.

### `in`: Download the latest version of an artifact.

#### Parameters

### `out`: Deploy artifact to a repository.

#### Parameters
