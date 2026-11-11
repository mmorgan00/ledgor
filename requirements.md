# Ledger overview

This project is a small 'production-thinking' ledger service written in Go, as a means to learn the language. This is not intended to be a product, a startup gamble, or other "sellable" angles, but something I could take into my day job and extend to fit into our stack as proof of competence.

## Non-negotiables

1. Double entry ledger
2. Idempotency
3. Balance enforcement
4. Audit Trail
5. Deterministic state under concurrency


## Data Model

- Postgres because in the spirit of "learning", I used MariaSQL in college and then went to NoSQL and now we're a Google Spanner shop. Postgres is something I should at least *look* at once

- Money rule: integer cents. No floats

TODO: Schema overview

## API

- I am very tempted by gRPC, because having something actually built with that beyond just the prototypes I've done at work would be nice, and I'm intending this to be deployed 'inside' a system, not something standalone that's client facing. gRPC fits that nicely, and then forces me to do more than just a `go run .` run step. However, that's probably more tooling than I want to commit to at this point, so initially it'll be a boring REST API. Something like:

```
POST /accounts
GET /accounts
GET /accounts/:id/entries


POST transactions

POST /transactions/:id/settle
POST /transacitons/:id/reverse
GET /transactions/:id
```

