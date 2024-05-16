
# EPL API

This is REST API that shows list of clubs that play in EPL, and players of those clubs.
Base URL - https://golangepl.onrender.com/api/v1




## API Endpoints
###  Authorization and Authentication
#### Registration
```http
  POST /users
```
| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` |  User's name |
| `email`      | `string` |  User's email |
| `password`      | `string` |  User's password |

#### Activation
```http
 PUT /users/activated
```
| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` |  User's activation token |

#### Authentication
```http
 POST /tokens/authentication
```
| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` |  User's name |
| `email`      | `string` |  User's email |

### Clubs
#### Get all clubs

```http
  GET /clubs
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `clubname` | `string` | Search for matches in the club name |
| `clubcity` | `string` | Search for matches in the club city |
| `page` | `int` | Display individual page|
| `page_size` | `int` | Limit page size |
| `sort` | `string` | Sort data based on available options|

Sorting options: "clubid", "clubname", "clubcity", "leagueplacement", "-clubid", "-clubname", "-clubcity", "-leagueplacement"

#### Get specific club

```http
  GET /clubs/{id}
```
 User must be Activated


| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Add new club

```http
  POST /clubs
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `clubname`      | `string` |  Name of the club |
| `clubcity`      | `string` |  City of the club |
| `position`      | `int` |  Placement of the club |
| `points`      | `int` |  Points of the club |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Update club

```http
  PUT /clubs/{id}
```
User must have permission to that action

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `clubname`      | `string` | **Required**. Name of the club |
| `clubcity`      | `string` | **Required**. City of the club |
| `position`      | `int` | **Required**. Placement of the club |
| `points`      | `int` | **Required**. Points of the club |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Delete club

```http
  DELETE /clubs/{id}
```
User must have permission to that action

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to delete |

### Players
#### Get all players

```http
  GET /players
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `firstname` | `string` | Search for matches in the first name |
| `lastname` | `string` | Search for matches in the last name |
| `position` | `string` | Search for matches in the position |
| `nation` | `string` | Search for matches in the nationality |
| `age` | `int` | Search for matches in the age |
| `number` | `int` | Search for matches in the number |
| `page` | `int` | Display individual page|
| `page_size` | `int` | Limit page size |
| `sort` | `string` | Sort data based on available options|

Sorting options: "playerid", "playerclubid", "playerage", "playernumber", "playerposition", "playernationality", "-playerid", "-playerclubid", "-playerage", "-playernumber", "-playerposition", "-playernationality"

#### Get specific player

```http
  GET /players/{id}
```
 User must be Activated


| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Add new player

```http
  POST /players
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `playerclubid`      | `int` |  Club id of the player |
| `firstname`      | `string` |  First name of the player |
| `lastname`      | `string` |  Last name of the player |
| `age`      | `int` |  Age of the player |
| `number`      | `int` |  Number of the player |
| `position`      | `string` |  Position of the player |
| `nation`      | `string` |  National of the player |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Update player

```http
  PUT /players/{id}
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `playerclubid`      | `int` | **Required**. Club id of the player |
| `firstname`      | `string` | **Required**. First name of the player |
| `lastname`      | `string` | **Required**. Last name of the player |
| `age`      | `int` | **Required**. Age of the player |
| `number`      | `int` | **Required**. Number of the player |
| `position`      | `string` | **Required**. Position of the player |
| `nation`      | `string` | **Required**. National of the player |

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Delete player

```http
  DELETE /players/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to delete |

## Authors

# Жексенбиев Арсен Бахитжанович 22B030357
[Telegram](https://t.me/LightNext)

