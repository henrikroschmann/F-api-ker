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

Get a specific entry
GET: localhost:8080/content/{id}

Create new entry 
POST: localhost:8080/content/

Delete entry 
DELETE: localhost:8080/content/{id}

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