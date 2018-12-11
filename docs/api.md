### /auth
---
##### ***GET***
**Summary:** Authenticate user with valid.

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Token | [Token](#token) |
| 500 | GenericError | [GenericError](#genericerror) |

### Models
---

### GenericError  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string | Error massage. | No |

### Token  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| token | string | JWT Token. | No |

### User  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer (uint64) |  | No |
| password | string | User's password | Yes |
| username | string | User's username | Yes |