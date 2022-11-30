# imapchecker
Simple [IMAP/IMAPS](https://www.rfc-editor.org/rfc/rfc3501) checker package.

## Includes:
- authorization & inbox select
- proxy support

## Examples:

IMAP

```golang
email := "email@mail.domain"
password := "password"
address := "example.mail.domain:143"

// create a dialer that implements proxy.ContextDialer interface
dialer := &net.Dialer{}

// set a timeout for context
ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

// create a checker config
config := NewConfig(address)

// open a connection
conn, err := DialWithContextDialer(ctx, dialer, config)
if err != nil {
	log.Fatal(err)
}

// check if mail is valid 
err = conn.CheckInbox(ctx, email, password)
if err != nil {
	log.Fatal(err)
}

log.Printf("IMAP: account is valid with %s %s for %s", email, password, address)


```

IMAPS

```golang

email := "email@mail.domain"
password := "password"
address := "example.mail.domain:993"

// create a dialer that implements proxy.ContextDialer interface
dialer := &net.Dialer{}

// set a timeout for context
ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

// create tls config
tlsConfig := &tls.Config{
	InsecureSkipVerify: true,
}

// create a checker config with tls
config := NewTLSConfig(address, tlsConfig)

// open a connection
conn, err := DialWithContextDialer(ctx, dialer, config)
if err != nil {
	log.Fatal(err)
}

// check if mail is valid 
err = conn.CheckInbox(ctx, email, password)
if err != nil {
	log.Fatal(err)
}


log.Printf("IMAPS: account is valid with %s %s for %s", email, password, address)

```

Contributions are welcome
