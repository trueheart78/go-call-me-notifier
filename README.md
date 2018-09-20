# Go! Call Me Notifier

Local application that subscribes to the published items from the
[Go! Call Me (Maybe)!][go call me maybe] lambda.

:warning: Only works on MacOS 10.9+

![taylor swift][drop everything now]

## Configuration

Setup a `~/.go-call-me.json` file with the following values:

```json
{
  "redis_url" : "https://url.to.redis",
  "redis_password" : "your password"
}
```

[go call me maybe]: https://github.com/trueheart78/go-call-me-maybe
[drop everything now]: assets/images/taylor-swift.gif
