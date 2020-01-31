# ec2-imds: Query the AWS Instance Metadata API

This is a simple executable that makes it easier to query the AWS Instance Metadata API.

* Uses the AWS Go SDK which has support for retries. The default is set to 3 and can be overridden with the global `--retries` flag.

* Supports [IMDSv2](https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service/)

* Support for specialized queries. Some endpoints from the API only return JSON. This allows to return more logical values such as for the `region` sub-command.

* Exit in error when value is not found.

### Usage

```
# Query the root of the meta-data sub-path
ec2-imds

# Query the ami-id meta-data property
ec2-imds ami-id

# Query the region
ecs-imds region

# Query the user data
ecs-imds user-data

# Query properties from the dynamic sub-path
ecs-imds dynamic instance-identity/signature

# With more retries
ecs-imds --retries=10 ami-id
```
