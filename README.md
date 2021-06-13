# AppStore Review Dumper

Dumps all recent Apple AppStore reviews for an app.

## Example

```
$ go build -o dumper .

$ ./dumper dump https://apps.apple.com/de/app/covpass/id1566140352

1 New Review

2021-13-06		5/5		JohnDoe
-------------------------------------------------------------------------------
Love the app design
This is probably the best looking app I've ever seen on the AppStore. So simple
and elegant.
```

## Available Commands

###### ./dumper dump <app_store_url>

Requests all recent reviews for the given AppStore URL, saves them and displays all new reviews.

###### ./dumper list <app_store_url>

List all saved reviews for given AppStore URL.