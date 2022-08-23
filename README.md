# Go - Random Date generator

Easy way to just import and use random **Date, Time, Date/Time and DOB (Date of birth)** with specified criteria.

> NOTICE: Using `"crypto/rand"` module to avoid randomisation with same values.
 
## Usage:

    go get github.com/duktig-solutions/go-random-date-generator

### Get random date between specified values

```go
package main

import "github.com/duktig-solutions/go-random-date-generator"

func main()  {
    
  randomDate, err := randomDataTime.GenerateDate("2020-08-01", "2022-08-01")

  if err != nil {
    panic(err) // just it...
  }

  fmt.Println(randomDate)
	
}
```

The result will look like:

    2022-01-22
    2021-04-09
    2020-12-06
    2022-05-07

### Get Random time

```go
package main

import "github.com/duktig-solutions/go-random-date-generator"

func main() {
	
  for i := 1; i < 5; i++ {
    randomTime, err := randomDataTime.GenerateTime(11, 19)

    if err != nil {
      panic(err) // just it ...
    }

    fmt.Println(randomTime)
  }
	
}
```

The result will look like:

    13:58
    16:44
    11:54
    18:55

### Get Random Date/Time

```go
package main

import "github.com/duktig-solutions/go-random-date-generator"

func main() {

  for i := 1; i < 5; i++ {
    randomDateTime, err := randomDataTime.GenerateDateTime("2015-06-15 16:45:00", "2019-09-21 10:43:58")

    if err != nil {
      panic(err) // just it ...
    }

    fmt.Println(randomDateTime)
  }
	
}
```

The result will look like:

    2017-05-18 17:22:13
    2019-08-26 11:00:01
    2019-08-28 18:24:36
    2018-12-14 16:41:51

### Get Random DOB (Date of birth)

This function is different from random date generation with possibility to quickly set age range in integers.

```go
package main

import "github.com/duktig-solutions/go-random-date-generator"

func main() {

  for i := 1; i < 5; i++ {
    randomDOB, err := randomDataTime.GenerateDOB(18, 75)

    if err != nil {
      panic(err) // just it ...
    }

    fmt.Println(randomDOB)
  }
}
```

The result will look like:

    1949-10-02
    1988-02-12
    1997-09-17
    1972-09-28

## Tests

> NOTE: Having output in test process to see results.

    go test -v

End of document