package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 int32 = 0
    var v5 uint16 = 42
    delete(v2, ((("\n\t" + "hello") + "foo bar") + "a"))
    v4 = (v4 % ((v4) | 1))
    v6 := GenS1{F0: 65535, F1: 2147483647}
    _ = len(v3)
    v7 := GenS2{F0: uint64(v5)}
    var v8 int8 = -119
    v8 = 0
    v9 := v1[1]
    for fi10 := 0; fi10 < 15; fi10++ {
        if false {
            switch ((("foo bar" + "") + ((("hello" + "a") + (("" + "a") + "a")) + "")) + "a") {
            case "hello":
                var v11 int8 = (-(v8))
                _ = v11
            default:
                v4 = (v4 ^ 42)
            }
            fmt.Println(v9, v9)
            v4 += (v4 / ((v4) | 1))
        } else {
            fmt.Println(v9)
        }
        _ = fi10
    }
    v4 = (v4 - (v4 - (-(v4))))
    v12 := GenS1{F0: 18446744073709551615, F1: (v9 % ((v9) | 1))}
    if true {
        var v13 int = 611
        _ = v13
    } else {
        var v14 int8 = v8
        _ = v14
    }
    _ = len(v3)
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS1{F0: 336, F1: v9}
        var v17 int = (fi15 - (fi15 - fi15))
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
