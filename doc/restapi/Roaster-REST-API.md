FORMAT: 1A

HOST: https://roast.software/

# Roaster

The Roaster API let's people validate their code using statical code analysis.
Users can register and track their progress, also allowing friends to follow
your progress.

# Roaster API Root [/]
This resource returns the base `index.html`. The index contains JavaScript that
implements a SPA (Single Page Application) that is able to communicate with the
API.

## Retrieve the roast.software SPA [GET]
+ Response 200 (text/html)
	```html
	<!DOCTYPE html>
	<html [...]>
	[... The roast.software SPA ...]
	</html>
	```

## User [/user]
### Create User [POST]
+ Relation: user
+ Request: A user registration (application/json)
	+ Body
	```json
	{
		"email": "email@example.com",
		"username": "roastin_roger",
		"password": "MyTr3me3d0usPassw0rd!",
		"fullname": "Roger Roger"
	}
	```
	+ Schema
	```json
	{
		"type": "object",
		"properties": {
			"email": {
				"type": "string"
			},
			"username": {
				"type": "string"
			},
			"password": {
				"type": "string"
			},
			"fullname": {
				"type": "string",
				"required": false
			}
		}
	}
	```
+ Response: 200
	+ Headers
	```
	Set-Cookie: roaster_auth=AB32DEAC21A91DE123[...]; Expires=Wed, 21 Oct 2015 07:28:00 GMT;
	```

### View/Handle Specific User [/user/{username}]
+ Relation: user

#### Change User [PATCH]
+ Relation: user/{username}
+ Request: Change user (application/json)
	+ Body
	```json
	{
		"email": "email@example.com",
		"password": "My3venM0reTr3mend0usP@ssw0rd!",
		"fullname": "Sir. Roger Roger"
	}
	```
	+ Schema
	```json
	{
		"type": "object",
		"properties": {
			"email": {
				"type": "string",
				"required": false
			},
			"password": {
				"type": "string",
				"required": false
			},
			"fullname": {
				"type": "string",
				"required": false
			}
		}
	}
	```
+ Response: 200

#### Retrieve User Information [GET]
+ Relation: user/{username}
+ Request: Retrieve user information
+ Response: 200 (application/json)
	+ Body
	```json
	{
		"email": "email@example.com",
		"username": "roastin_roger",
		"fullname": "Roger Roger"
	}
	```
	+ Schema
	```json
	{
		"type": "object",
		"properties": {
			"email": {
				"type": "string"
			},
			"username": {
				"type": "string"
			},
			"fullname": {
				"type": "string"
			}
		}
	}
	```

#### Remove User [DELETE]
+ Relation: user/{username}
+ Request: Remove user
	+ Headers
	```
	Cookie: roaster_auth=AB32DEAC21A91DE123[...]
	```
+ Response: 200
	+ Headers
	```
	Set-Cookie: roaster_auth=deleted; Expires=Thu, 01 Jan 1970 00:00:00 GMT;
	```

## Session [/session]
### Authenticate for New Session (sign in) [POST]
+ Relation: session
+ Request: A session authentication (application/json)
	+ Body
	```json
	{
		"username": "roastin_roger",
		"password": "MyTr3me3d0usPassw0rd!",
	}
	```
	+ Schema
	```json
	{
		"type": "object",
		"properties": {
			"username": {
				"type": "string"
			},
			"password": {
				"type": "string"
			}

		}
	}
	```
+ Response: 200
	+ Headers
	```
	Set-Cookie: roaster_auth=AB32DEAC21A91DE123[...]; Expires=Wed, 21 Oct 2015 07:28:00 GMT;
	```

### Remove Current Session (sign out) [DELETE]
+ Relation: session
+ Request: Remove current session
	+ Headers
	```
	Cookie: roaster_auth=AB32DEAC21A91DE123[...]
	```
+ Response: 200
	+ Headers
	```
	Set-Cookie: roaster_auth=deleted; Expires=Thu, 01 Jan 1970 00:00:00 GMT;
	```
## Roast [/roast]
### Analyze code snippet [POST]
+ Relation: roast
+ Request: A static analysation of provided code snippet (application/json)
	+ Headers
	```
	Cookie: roaster_auth=AB32DEAC21A91DE123[...] (string, optional) - Should be sent if user has an active session.
	```
	+ Body
	```json
	{
		"language": "python3",
		"code": "print('I´m Roastin´ Roger... Roger.', invalid = 'Bönor är ändå rätt gött')",
	}
	```
	+ Schema
	```json
	{
		"type": "object",
		"properties": {
			"language": {
				"type": "string"
			},
			"code": {
				"type": "string"
			}

		}
	}
	```
+ Response: 200
	+ Body
	```json
	{
		"grade": 12,
		"warnings": [
			{
				"row": 1,
				"column": 48,
				"engine": "pycodestyle",
				"name": "E251",
				"description": "no spaces around keyword / parameter equals"
			}
		],
		"errors": [
			{
				"row": 1,
				"column": 41,
				"engine": "pyflakes",
				"name": "TypeError",
				"description": "print() got an unexpected keyword argument 'invalid'"
			}
		]
	}
	```
	+ Schema
	```json
	{
		"type": "object",
		"properties": {
			"grade": {
				"type": "number"
			},
			"warning": {
				"type": "array",
				"items": {
					"type": "object",
					"properties": {
						"row": {
							"type": "number"
						},
						"column": {
							"type": "number"
						},
						"engine": {
							"type": "string"
						},
						"name": {
							"type": "string"
						},
						"description": {
							"type": "string"
						}
					}
				}
			},
			"errors": {
				"type": "array",
				"items": {
					"type": "object",
					"properties": {
						"row": {
							"type": "number"
						},
						"column": {
							"type": "number"
						},
						"engine": {
							"type": "string"
						},
						"name": {
							"type": "string"
						},
						"description": {
							"type": "string"
						}
					}
				}
			}
		}
	}
	```