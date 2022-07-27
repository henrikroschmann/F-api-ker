# F-api-ker (faker api)

### Purpose:
The reason for this tool is to help out when mocking API data while building for example frontends or other services. 

### Gettings started 
```
docker-compose up 
```

### Endpoints
Get all entries
GET: localhost:8080/content/

![image](https://user-images.githubusercontent.com/17333/181211658-bbacce39-e1b7-4382-b1dd-b3ae15ad427a.png)


Get a specific entry
GET: localhost:8080/content/{id}

![image](https://user-images.githubusercontent.com/17333/181211628-0a30e28f-f528-4bd3-b3e3-b13063e5aaa8.png)


Create new entry 
POST: localhost:8080/content/

![image](https://user-images.githubusercontent.com/17333/181211542-2f76ea77-17dc-46df-be13-fd9a3caeaef3.png)


Delete entry 
DELETE: localhost:8080/content/{id}

![image](https://user-images.githubusercontent.com/17333/181211694-0b525770-9fd4-40ed-be40-47e125e2d0d9.png)


### Example payload

```json
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}
```
Data taken from JSON Example

You can throw any valid json data on this one and it will parse correctly. 
