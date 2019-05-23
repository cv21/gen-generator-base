# gen-generator-base


#### Generator Parameters
| Parameter Key | Example | Description |
| --- | --- | --- |
|generator_name|Mock||
|module_repository|github.com/cv21/gen-generator-mock||
|module_query|1.0.0||
|params_structure_name|generatorParams||

#### Config Example:

```json
	{
		"files": [
			{
				"path": "...",
				"generators": [
				{
					"repository": "github.com/cv21/gen-generator-base",
					"version": "master",
					"params": {
						"generator_name":"Mock",
						"module_repository":"github.com/cv21/gen-generator-mock",
						"module_query":"1.0.0",
						"params_structure_name":"generatorParams"
					}
				}
			]
			}
		]
	}

```

