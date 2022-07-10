---
title: Why should I use Go?
author : Mathis Joffre
styles:
    style: monokai
extensions:
    - image_ueberzug
---


> Why should I use Go?
>
>
> A quick introduction to Go and its language features.
>
>
>
> Presented by Mathis Joffre

---
# Agenda 

1. Who am I?
2. What is Go?
3. Why should you become a Gopher?
4. Quick tour of Go
    1. Hello World
    2. Features
       1. Structs and typed Methods
       2. Interfaces
       3. Goroutines
          1. Context
          2. Mutex and wait groups
       4. Generics
       5. Tests
5. Live-coding : Web Server



---

# Who am I?

- Students at ESIR (**Ecole Superieure d'Ingenieurs de Rennes**) in Networking and Internet Technology.


- I'm currently working as a **Software Engineer and SRE** at **OVHcloud**.


- Practicing Go since 1 year and I started with Python.


- Reach me on :
  - GitHub : **Joffref**
  - LinkedIn : **Mathis Joffre**
  - Discord : **Sugate#2563**

---

# What is Go?

- Go is a statically typed and compiled language designed @Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007.


- It is a **concurrent, garbage-collected, and memory-safe** language. It is designed to be **fast** and **memory-efficient**.


- Go1 : 2012-03-28


- Project built upon Go :

  
  - Docker (https://www.docker.com/)

  - Kubernetes (https://kubernetes.io/)

  - CockroachDB (https://www.cockroachlabs.com/)
---

# Why should you become a Gopher ?

- Easy to learn. Easy syntax. Easy to use. Easy to understand.
- Safe.
- Fast.
- Concurrent friendly.
- Use by a large number of devs and firms.

---
# Quick tour of Go

## Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

---
## Features
 
### Structs and typed Methods

```go
package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age: age,
    }
}

func (p *Person) Greet() {
    fmt.Printf("Hello, I'm %s\n", p.Name)
}

func main() {
    p := NewPerson("John", 30)
    p.Greet()
}
```

---
## Basics

### Interfaces

```go
package main

import "fmt"

type Gopher interface {
    Greet()
}

type Person struct {
    Name string
    Age int
}

type Robot struct {
    Name string
}

func (p *Person) Greet() {
    fmt.Printf("Hello, I'm %s\n", p.Name)
}

func (r *Robot) Greet() {
    fmt.Printf("Hello, I'm %s\n", r.Name)
}

func DoStuff(g Gopher) {
    g.Greet()
}

func main() {
    p := &Person{
        Name: "John",
        Age: 30,
    }
    r := &Robot{
        Name: "Robot",
    }
    DoStuff(p)
    DoStuff(r)
}
```

---
## Basics

### Goroutines

```go
package main

import "fmt"

func main() {
    c := make(chan int) 
    go func(c chan int) {
	    for i := 0; i < 10; i++ {
	        fmt.Println(i)
        }
        c <- 1
    }(c)
    <-c
}
```

---
## Basics
### Goroutines
#### Context
```go
package main

import (
	"context"
	"fmt"
)

func main() {
	main_ctx := context.Background()
	ctx, cancel := context.WithCancel(main_ctx)
	go func(ctx context.Context) {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		cancel()
	}(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("It's finish")
			return
		}
	}
}
```

---
## Basics
### Goroutines
#### Mutex and wait groups
```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Container struct {
	mu       sync.Mutex
	counters int
}

func (c *Container) inc(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < n; i++ {
		c.counters++
		fmt.Println(c.counters)
	}
	wg.Done()
}

func (c *Container) dec(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < n; i++ {
		c.counters--
		fmt.Println(c.counters)
	}
	wg.Done()
}

func main() {
	c := &Container{counters: 10}
	wg.Add(2)
	go c.inc(3)
	go c.dec(2)
	wg.Wait()
}
```
---
## Basics

### Generics

```go
package main

import "fmt"

func print[V any](v V) {
	fmt.Println(v)
}

func main() {
	print(1)
	print(1.0)
}
```
---
## Basics

### Tests

```go
package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

func TestIntToString(t *testing.T) {
	assert.Equal(t, "1", IntToString(1))
	assert.Equal(t, "2", IntToString(2))
	assert.Equal(t, "3", IntToString(3))
}

```
---
## Basics

### Errors handling

```go
package main

import (
	"fmt"
	"errors"
)

func createError() error {
    return errors.New("Error")
}

func main() {
    err := createError()
    if err != nil {
        fmt.Println(err)
    }
}
```

---
# Live-coding : Web Server


---
# Thank you for listening !
- Reach me on :
    - GitHub : **Joffref**
    - LinkedIn : **Mathis Joffre**
    - Discord : **Sugate#2563**





