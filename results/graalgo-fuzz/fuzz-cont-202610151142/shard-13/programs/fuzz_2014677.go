package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int64 = -533
    var v2 uint16 = 255
    var v3 string = "hello"
    v4 := []uint32{2147483648, 2147483647}
    if true {
        var v5 uint64 = 2
        var v6 bool = (false || (!true))
        _ = v5
        _ = v6
    }
    switch v3 {
    case "foo bar":
        var v7 uint8 = 127
        _ = v7
    }
    v8 := GenS1{F0: 1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v9 int8 = 1
    if false {
        fmt.Println(v1, v2, v2)
        v1 -= (-(v1))
        if len(v4) > 0 {
            v4[0] = 483
        }
    }
    var v10 int16 = 1
    v3 = v3
    fmt.Println(v2, v3)
    var v11 int64 = (v1 | v1)
    if true {
        v12 := GenS1{F0: 470}
        _ = v12
    } else {
        v2 = (^(v2))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
