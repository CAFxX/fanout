package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 string }
type GenS3 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint64 = 42
    var v2 int16 = 1
    v3 := make(map[int]bool)
    v4 := []uint32{65535}
    var v5 uint32 = 602
    if true {
        switch 32767 {
        case 1:
            fmt.Println(v5)
        }
    } else {
        switch "" {
        case "hello":
            var v6 uint = 4294967295
            _ = v6
        case "":
            for fi7 := 0; fi7 < 6; fi7++ {
                v5 = (v5 - v5)
                v1 += (v1 & v1)
                var v8 int32 = 2147483647
                _ = fi7
                _ = v8
            }
        default:
            v5 = 2147483647
        }
        delete(v3, int(v5))
    }
    v2 += (v2 | v2)
    fmt.Println(v2, v1)
    v5 -= 2147483647
    if len(v4) > 0 {
        v4[0] = (v5 % ((v5) | 1))
    }
    v1 += (v1 % ((1) | 1))
    v9 := GenS1{F0: -91, F1: 0, F2: 2}
    _ = len(v3)
    v10 := v3[2147483647]
    var v11 int32 = 1
    fmt.Println(v5)
    var v12 int8 = 68
    switch int(v12) {
    case 9223372036854775807:
        switch "foo bar" {
        case "hello":
            v13 := GenS2{F0: 2, F1: (v1 * v1), F2: (("a" + "hello") + (("hello" + "\n\t") + "a"))}
            _ = v13
        case "\n\t":
            _ = len(v4)
        }
    case 562:
        _ = len(v4)
    case 511:
        if (-2.25 != (1.5 - 3.14159)) {
            fmt.Println(v10, v2)
            if false {
                v2 = (v2 | -59)
                v3[128] = ((1e10 > 1e10) || v10)
            }
        } else {
            v14 := GenS3{F0: (v1 * (v1 / (((^(v1))) | 1))), F1: -387}
            _ = v14
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
