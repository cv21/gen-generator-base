

#### Generator Parameters
| Parameter Key | Type | Example | Description |
| --- | --- | --- | --- |
|`repository`|`string`|`github.com/cv21/gen-generator-mock@v1.0.0`|It is necessary for plugin registration. Also it is useful for building config example. It is better to use module queries if your generator supports versioning.|
|`params_structure_name`|`string`|`generatorParams`|It is name of params structure which holds all generator params.|

#### Config Example:

```yaml
files:
  - path: ...
    repository: github.com/cv21/gen-generator-base@v1.0.0
    params:
      repository: github.com/cv21/gen-generator-mock@v1.0.0
      params_structure_name: generatorParams
```

