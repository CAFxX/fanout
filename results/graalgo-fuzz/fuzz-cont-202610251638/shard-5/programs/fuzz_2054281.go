package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int64 = 0
    var v2 uint8 = 124
    var v3 int8 = 127
    var v4 float32 = 1.5
    v5 := []string{}
    if (false || true) {
        v6 := GenS1{F0: (("foo bar" + "\n\t") + "foo bar")}
        _ = v6
    } else {
        var v7 int16 = 746
        _ = v7
    }
    if (!false) {
        v4 += -2.25
    }
    var v8 uint64 = uint64(v2)
    var v9 int16 = 42
    switch "" {
    case "\n\t":
        _ = len(v5)
    case "hello":
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    default:
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int16 = v9
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 int8 = 1
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    }
    if false {
        var v14 uint16 = 0
        if ((1 != 0) || false) {
            var v15 uint64 = (v8 & v8)
            _ = v15
        }
        v16 := GenS2{F0: v1, F1: "a"}
        _ = v14
        _ = v16
    } else {
        fmt.Println(v2, v2)
    }
    v17 := GenS3{F0: 291, F1: 2147483648}
    v9 += 0
    v9 += v9
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v18 uint64 = v8
    v19 := GenS3{F0: 2147483648, F1: 817}
    if false {
        var v20 int64 = (v1 * (v1 - v1))
        v4 = v4
        _ = v20
    }
    v8 -= v8
    v8 -= (v18 % ((v18) | 1))
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
