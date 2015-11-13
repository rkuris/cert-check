# cert-check

Super simple check a list of certificates and output something
if the cert expires within the warning period.

This is intended to be stuck in a daily cron job to watch for
certificate expirations. If anything goes wrong or there is a cert
that is about to expire, this code tells you about it.

Kind of just starting with go and wanted something to whip up,
and so thought this might be useful later for someone.

Inspired by Kris Pruden's reference to this article:

http://www.theguardian.com/technology/2015/nov/12/apple-user-anger-mac-apps-break-security-certificate-lapse

## Configuring

Right now, everything is hardcoded. You probably want to change the
list of sites it checks and how many days before it warns you.

## Running

Easy to run if the go system is installed:
```sh
go run cert-check.go
```

You can also precompile it for a specific architecture and run cert-check
instead, which then doesn't require any go on the machine:

```sh
go build cert-check.go
```
