# scram-password -- Command-line utility for Postgres-compatible SCRAM-SHA-256 passwords

SCRAM-SHA-256 (see [RFC-7677](https://datatracker.ietf.org/doc/html/rfc7677), [Salted Challenge Response Authentication Mechanism](https://en.wikipedia.org/wiki/Salted_Challenge_Response_Authentication_Mechanism)) is a password based challenge-response authentication mechanism.

[Postgres 14 uses it](https://www.postgresql.org/docs/14/auth-password.html) to avoid needing to store or transmit plaintext passwords.

This repository contains a simple command-line utility to hash passwords into a Postgres-compatible format.
It may work with other SCRAM-using server software, but the actual storage format is not a standard.

```console
$ go install -v  eagain.net/go/scram-password@latest
[...]
$ scram-password jdoe </secrets/postgres-password-for-jdoe
SCRAM-SHA-256$4096:QmQ2A1cjD16nIqNIDV7h8zjEG1B2h3mc$Cg0t5o2dPlN7gjE4v023hrhGIegBF1aOLksORwBiTgA=:UEs8KN9wbs03QE6oyglm8egxWqNh6laUfNtvVoChtRM=
```

If you need similar helpers to easily generate the actual passwords, see https://github.com/tv42/entropy and <https://github.com/tv42/zbase32>:

```console
$ entropy 32 | zbase32-encode
rijwsgiuedt4bx86b5qsamxs1iyobbjdr7f9mieattztbgauxngo
```

# Acknowledgements

The library that does all actual work: https://github.com/xdg-go/scram

Configuration and usage advice: https://hacksoclock.blogspot.com/2018/10/how-to-set-up-scram-sha-256.html

More advice (including a Python script with unclear licensing that was *not* used to create this project): https://blog.crunchydata.com/blog/how-to-upgrade-postgresql-passwords-to-scram

Note in Postgres docker image documentation that says `POSTGRES_INITDB_ARGS=--auth-host=scram-sha-256` might be needed for SCRAM to work: https://github.com/docker-library/docs/commit/00ad08f4335b71b70cfed616ca81ab6dfc015f12
