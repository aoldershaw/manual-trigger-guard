# manual-trigger-guard

A Concourse resource type to fail a build when it is triggered by a user that
is not granted access (via GitHub username or team membership).

## Usage

```yaml
resource_types:
- name: manual-trigger-guard
  type: registry-image
  source: {repository: aoldershaw/manual-trigger-guard}

resources:
- name: allow-maintainers-guard
  type: manual-trigger-guard
  expose_build_created_by: true
  source:
    access_token: ((access_token))
    users:
    - some-user
    teams:
    - concourse/maintainers

jobs:
- name: some-job
  plan:
  - put: allow-maintainers-guard
  - ... # the rest of the job plan
```

With this configuration, `some-job` must be triggered manually, and the build
will fail if it is triggered by a user other than `some-user` or one of the
members of `concourse/maintainers`.

Note: the `access_token` must be specified if using the `teams` configuration,
and the access token must be granted the `read:org` permission. The access
token must belong to a user that has visibility to the `teams`.
