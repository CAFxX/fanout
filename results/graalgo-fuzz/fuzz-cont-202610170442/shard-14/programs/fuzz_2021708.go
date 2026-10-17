package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 float64 = float64(0.0)
    var v2 float64 = 3.14159
    var v3 string = "hello"
    var v4 uint = 2
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: 42}
        var v7 bool = (true && ((!true) || (false || (0 < 32767))))
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 42}
    v3 = "foo bar"
    var v9 uint64 = 18446744073709551615
    switch v3 {
    case "foo bar":
        v4 += v4
    case "a":
        fmt.Println(v9, v1)
    case "hello":
        var v10 int32 = 0
        _ = v10
    }
    v2 -= v1
    var v11 bool = (!true)
    v12 := GenS1{F0: 32767}
    var v13 uint8 = 255
    v3 = "hello"
    var v14 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
