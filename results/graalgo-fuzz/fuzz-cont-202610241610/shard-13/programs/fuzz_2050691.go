package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }

func main() {
    var v1 int8 = -1
    var v2 bool = (-2.25 != 1.5)
    var v3 int64 = 9223372036854775807
    var v4 int = 1
    fmt.Println(v1, v1, v1)
    var v5 int16 = 42
    v4 = (v4 | v4)
    var v6 uint = 4294967295
    v7 := GenS1{F0: 4294967295}
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: (v6 - v6), F1: v6, F2: 1}
        var v11 string = "foo bar"
        _ = fi9
        _ = v10
        _ = v11
    }
    v6 -= v6
    switch ((("a" + "\n\t") + ("hello" + "hello")) + "") {
    case "\n\t":
        v3 = (v3 / ((v3) | 1))
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v3 = -474
        var v13 int8 = 0
        v3 -= (v3 * 1)
        _ = fi12
        _ = v13
    }
    var v14 int32 = 1
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(fi15)
        fmt.Println(v1, v14, v8)
        v8 = v8
        _ = fi15
    }
    switch "a" {
    case "":
        var v16 uint16 = 65535
        _ = v16
    case "a":
        v3 = (v3 % ((v3) | 1))
    default:
        switch (v4 / ((9223372036854775807) | 1)) {
        case -127:
            v17 := GenS1{F0: (v8 - v8)}
            _ = v17
        }
    }
    var v18 uint64 = v8
    v2 = v2
    v4 -= (v4 / ((2147483647) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
