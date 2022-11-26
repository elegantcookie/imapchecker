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

dialer := &net.Dialer{Timeout: 5 * time.Second}

// create a connection using dialer
conn, err := imapchecker.DialWithDialer(dialer, address)
if err != nil {
  log.Printf("failed to connect the server: %v", err)
  return
}

// check if credentials are valid and you have a permission to check mail in one function
err = conn.CheckInbox(email, password)
if err != nil {
  log.Printf("failed to login/select inbox")
  return
}

// or do it separately
err = conn.Authenticate(email, password)
if err != nil {
  log.Printf("failed to login")
  return
}

err = conn.OpenInbox()
if err != nil {
  log.Printf("have no permission to check mail")
  return
}

log.Printf("IMAP: account is valid with %s %s for %s", email, password, address)


```

IMAPS

```golang

email := "email@mail.domain"
password := "password"
address := "example.mail.domain:993"

dialer := &net.Dialer{Timeout: 5 * time.Second}

config := &tls.Config{
			InsecureSkipVerify: true,
		}

// create a tls connection using dialer
conn, err := imapchecker.DialWithDialerTLS(dialer, address, config)
if err != nil {
  log.Printf("failed to connect the server: %v", err)
  return
}

err = conn.CheckInbox(email, password)
if err != nil {
  log.Printf("failed to login/select inbox")
  return
}

err = conn.Authenticate(email, password)
if err != nil {
  log.Printf("failed to login")
  return
}

err = conn.OpenInbox()
if err != nil {
  log.Printf("have no permission to check mail")
  return
}

log.Printf("IMAPS: account is valid with %s %s for %s", email, password, address)

```

Contributions are welcome
