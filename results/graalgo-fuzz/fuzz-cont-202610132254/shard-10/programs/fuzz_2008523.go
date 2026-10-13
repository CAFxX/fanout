package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 127
    var v3 float32 = 0.0
    v4 := v1[65535]
    var v5 int = 128
    if (3 == 127) {
        fmt.Println(v5)
        v6 := GenS1{F0: v4}
        _ = v6
    }
    fmt.Println(v3)
    v4 = 0
    switch v5 {
    case -174:
        if false {
            fmt.Println(v3, v2)
        } else {
            var v7 int16 = 1000
            _ = v7
        }
    default:
        var v8 uint64 = 2
        _ = v8
    }
    v3 = 3.14159
    fmt.Println(v4)
    switch "hello" {
    case "\n\t":
        v4 -= (v5 & (v4 ^ (^(v4))))
    default:
        var v9 uint16 = 32767
        _ = v9
    }
    v3 -= ((v3 - 1.5) / v3)
    v3 -= 0.0
    if (true && false) {
        v2 = 0
        v10 := GenS1{F0: v5}
        v11 := v1[0]
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS2{F0: 32767}
            _ = fi12
            _ = v13
        }
        var v14 int16 = 126
        if false {
            var v15 int = (v4 - v5)
            _ = v15
        }
        _ = v14
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v16 := GenS3{F0: 0, F1: 0}
        _ = v16
    case "a":
        var v17 int64 = -9223372036854775808
        _ = v17
    default:
        v18 := []int{v4, v5}
        _ = v18
    }
    v19 := []int{32768}
    v5 += (v4 * (v4 >> ((18446744073709551615) & 63)))
    var v20 int64 = 368
    var v21 uint = 173
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
