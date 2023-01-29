# PAC

Role-based access control implemented with SQL's relations.

This was made completely only for playing around with [sqlc](https://sqlc.dev) + [postgresql](https://postgresql.org) and docker compose.

This is not intended to be actually used as there might be _some_ security flaws. Please note that this software wasn't tested for security in any other way than manually checking if a resource is limited only to a specific group.



## Setup

I don't know why you would want to setup this, but oh boy it's easy (I'm just extremely lazy and don't want to reproduce things constantly).

```sh
git clone git@github.com:romeq/pac.git
# or alternatively via http:
# git clone https://github.com/romeq/pac.git
docker-compose up -d
make clean-db run
```



## Architecture

Architecture for this is clearly rather simple. Therefore I'll just write it down.

### Database

| Table name                                    | Description                                                         |
| --------------------------------------------- | ------------------------------------------------------------------- |
| [`role`](#table-role)                         | Roles and their private information (such as roles' names)          |
| [`account`](#table-account)                   | Accounts and account related information                            |
| [`account_to_role`](#table-account-to-role)   | Account roles                                                       |
| [`resource`](#table-resource)                 | Resources's uuids and contents                                      |
| [`resource_to_role`](#table-resource-to-role) | Keeps track of what roles are allowed to access a specific resource |

#### Tables

##### <a id="table-role">role</a>

| Column name | Description                                                 |
| ----------- | ----------------------------------------------------------- |
| `role_uuid` | Unique identifier for role, used to refer a specific record |
| `name`      | Name of the role                                            |

##### <a id="table-account">account</a>

| Column name    | Description                                                  |
| -------------- | ------------------------------------------------------------ |
| `account_uuid` | Unique identifier for account, used to refer a specific record |
| `name`         | Username of the account                                      |

##### <a id="table-account-to-role">account_to_role</a>

| Column name    | Description                                                  |
| -------------- | ------------------------------------------------------------ |
| `account_uuid` | References [`account`](#table-account)-table's `role_uuid`-column |
| `role_uuid`    | References [`role`](#table-role)-table's `role_uuid`-column  |

##### <a id="table-resource">resource</a>

| Column name     | Description                                                  |
| --------------- | ------------------------------------------------------------ |
| `resource_uuid` | Unique identifier for resource, primarily used to refer a specific record |
| `content`       | Resource content (string)                                    |

##### <a id="table-resource-to-role">resource_to_role</a>

| Column name     | Description                                                  |
| --------------- | ------------------------------------------------------------ |
| `resource_uuid` | References [`resource`](#table-resource)-table's `role_uuid`-column |
| `role_uuid`     | References [`role`](#table-role)-table's `role_uuid`-column  |
