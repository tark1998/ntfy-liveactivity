# ntfy-liveactivity
This is a fork of [ntfy](https://github.com/tark1998/ntfy-liveactivity) with iOS Live activity.

See [ntfy-liveactivity-ios](https://github.com/tark1998/ntfy-liveactivity-ios) to build iOS app for the custom ntfy server.

**Caution**: This server works for only one Apple device, because I develop for my personal purpose.

# How to setup

1. Follow the [Firebase setup](https://docs.ntfy.sh/develop/#firebase-setup) and [APNS setup](https://docs.ntfy.sh/develop/#apple-setup) 
2. Save the firebase private key and Apple Authkey file under `/etc/ntfy/`
3. Edit `newapnsClient` function in `server/server_apns.go` matching with previous Apple Authkey and TeamID
3. Configure `/etc/ntfy/server.yml`. For example, see `server/server.yml`
  1. base-url: url and port that used for connection from outside of the proxy
  2. listen-http: url and port that used for connection from inside of the proxy
  3. firebase-key-file: location and file name of firbase private key
  4. behind-proxy: true

# How to publish

## Start live activity
```bash
curl https://mybaseurl.com/sometopic \
  -H "Activity: 1" \ 
  -d '{
    "aps": {
      "timestamp": 1747758082083,
      "event": "start",
      "content-state": {
        "emoji": "🍏🍏"
      },
      "attributes-type": "MywidgetAttributes",
      "attributes": {
        "name": "Apple"
      },
      "alert": {
        "title": "Hello",
        "body": "World",
        "sound": "chime.aiff"
      }
    }
  }'
```

## Update live activity

```bash
curl https://mybaseurl.com/sometopic \
  -H "Activity: 2" \ 
  -d '{
    "aps": {
      "timestamp": 1747758082083,
      "event": "start",
      "content-state": {
        "emoji": "🍏🍑"
      },
    }
  }'
```

## End live activity

```bash
curl https://mybaseurl.com/sometopic \
  -H "Activity: 3" \ 
  -d '{
    "aps": {
      "timestamp": 1747758082083,
      "event": "end",
      "content-state": {
        "emoji": "🍑🍑"
      },
      "dismissal-date": 1747758082093,
      "alert": {
        "title": "Hello",
        "body": "Update World",
        "sound": "chime.aiff"
      }
    }
  }'
```

