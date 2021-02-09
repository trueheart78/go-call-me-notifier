# Go! Call Me Notifier

**No longer maintained due to workplace restrictions connecting to AWS Redis**

Local application that subscribes to the published items from the
[Go! Call Me (Maybe)!][go call me maybe] lambda.

Download your preferred executable from [the Releases page][releases] and extract it. That's it!

:warning: Only works on MacOS 10.9+

![taylor swift][drop everything now]

## Configuration

Setup a `~/.go-call-me.json` file with the following values:

```json
{
  "redis_url" : "https://url.to.redis:port-number",
  "redis_password" : "your password",
  "redis_channels" : {
    "emergency"    : "emergency",
    "nonemergent"  : "nonemergent"
  }
}
```

[go call me maybe]: https://github.com/trueheart78/go-call-me-maybe
[drop everything now]: assets/images/taylor-swift.gif
[releases]: https://github.com/trueheart78/go-call-me-notifier/releases
