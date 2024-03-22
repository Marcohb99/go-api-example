# GO API EXAMPLE

Api developed in Golang with Codely TV course

## Domain

The chosen domain is the [discogs](https://www.discogs.com/) API. The idea is to create a simple API that allows you to search for a release.
Being a release a music album, a single, a compilation, etc.
For now, the properties are:

- Id (uuid)
- Title
- Released: release date
- Resource url (in the discogs api)
- Uri (in the web)
- Year

## Endpoints

- POST /releases: Create a new release

Sample request:

```json
{
  "id": "BAD92BF5-9176-47BD-BCC6-8C38A5394A6E",
  "title": "ultra mono",
  "released": "2020-09-25",
  "resource_url": "https://api.discogs.com/releases/15951324",
  "uri": "https://www.discogs.com/release/15951324-Idles-Ultra-Mono",
  "year": "2020"
}
```

After this, you can search for the release in database:

```bash
docker exec -ti go-api-example-mysql-1 bash
mysql -umhb -pmhb
use mhb;
select * from releases;
```

## Installation

```bash
go mod download
```

## Run

```bash
docker compose up --build
```

## Test

```bash
make test
```