package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8, p1 int16) string {
    var v5 int64 = -1
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = v5
    return ("a" + "\n\t")
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 2
    var v3 int32 = 42
    v4 := make(map[uint32]bool)
    switch 500 {
    case -358:
        var v6 int = 406
        _ = v6
    case -523:
        var v7 uint32 = (v2 ^ v2)
        _ = v7
    case 546:
        if ((false && false) || (!true)) {
            v1 = 42
            if (false && (-1 > -1)) {
                fmt.Println(v1, v1, v1)
            } else {
                var v8 int = 4294967295
                v9 := GenS1{F0: 0, F1: (^(v2)), F2: v1}
                var v10 float64 = -2.25
                _ = v10
                _ = v8
                _ = v9
            }
            var v11 int16 = 0
            _ = v11
        }
    default:
        v12 := v4[(v2 << ((65535) & 63))]
        _ = v12
    }
    v13 := []int{-9223372036854775808}
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 string = ("a" + ("a" + "foo bar"))
        v2 -= (v2 | 0)
        _ = fi14
        _ = v15
    }
    if len(v13) > 0 {
        v13[0] = 9223372036854775807
    }
    v16 := GenS2{F0: 2, F1: 4294967295}
    var v17 uint32 = 65535
    switch ("a" + "a") {
    case "\n\t":
        var v18 uint64 = 1
        _ = v18
    case "foo bar":
        v2 = v2
    case "":
        v19 := GenS3{F0: (!((!false) && (18446744073709551615 == uint(v2))))}
        _ = v19
    }
    v3 -= v3
    var v20 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
