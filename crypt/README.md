# Roki
The ultimate solution to sign and validate JSON Web Tokens (JWT) without breaking the ability to fully customize.

```go get github.com/z3ntl3/roki/crypt```

### Examples
##### Signing
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/z3ntl3/roki/crypt"
)

type MyCustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	*crypt.StandardClaims
}

func (c *MyCustomClaims) Valid() error {
	return nil
}

func main() {
	os.Setenv("SECRET", "root")

	c := &crypt.JWT{}
	{
		c.TokenStr = os.Getenv("token")
	}

	cl, err := c.Validate(&MyCustomClaims{}, func(t *crypt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := cl.(*MyCustomClaims)
	if !ok {
		log.Fatal("invalid token")
	}
	fmt.Println(claims.ExpiresAt, claims.Email, claims.Role)
}
```
##### Validating

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/z3ntl3/roki/crypt"
)

type MyCustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	*crypt.StandardClaims
}

func (c *MyCustomClaims) Valid() error {
	return nil
}

func main() {
	os.Setenv("SECRET", "root")

	c := &crypt.JWT{}
	{
		c.TokenStr = os.Getenv("token")
	}

	cl, err := c.Validate(&MyCustomClaims{}, func(t *crypt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := cl.(*MyCustomClaims)
	if !ok {
		log.Fatal("invalid token")
	}
	fmt.Println(claims.ExpiresAt, claims.Email, claims.Role)
}

```

##### Together
```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/z3ntl3/roki/crypt"
)

type MyCustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	*crypt.StandardClaims
}

func (c *MyCustomClaims) Valid() error {
	return nil
}

func main() {
	os.Setenv("SECRET", "root")

	c := &crypt.JWT{}
	{
		c.SecretEnv = "SECRET"
	}

	myclaims := &MyCustomClaims{
		Email: "efdal@gmail.com",
		Role:  "hoi",
		StandardClaims: &crypt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	if err := c.Sign(myclaims, crypt.HMAC_HS512); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.TokenStr)

	cl, err := c.Validate(&MyCustomClaims{}, func(t *crypt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := cl.(*MyCustomClaims)
	if !ok {
		log.Fatal("invalid token")
	}
	fmt.Println(claims.ExpiresAt, claims.Email, claims.Role)
}

```