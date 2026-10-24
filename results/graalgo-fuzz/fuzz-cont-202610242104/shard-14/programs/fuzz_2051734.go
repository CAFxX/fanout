package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }
func genFn0(p0 uint32, p1 uint32, p2 string) uint8 {
    var v6 uint16 = 234
    var v7 int8 = 0
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint16 = (^(v6))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 131
}
func genFn1(p0 int16) int16 {
    switch "a" {
    case "\n\t":
        p0 = (p0 ^ (p0 % ((p0) | 1)))
    case "a":
        p0 = p0
    case "foo bar":
        v10 := GenS1{F0: 42, F1: 42}
        _ = v10
    default:
        v11 := []int{-9223372036854775808, 0, -9223372036854775808}
        _ = v11
    }
    p0 -= (p0 - (p0 + p0))
    p0 = 1
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int8 = 1
        fmt.Println(v13, fi12)
        v14 := GenS2{F0: ("foo bar" + "hello")}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    return -32768
}

func main() {
    var v1 float32 = (0.0 / 1e10)
    var v2 uint = 255
    v3 := []string{"a", (("\n\t" + "foo bar") + "a")}
    var v4 string = "\n\t"
    var v5 uint64 = 4294967295
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v5)
        v16 := GenS2{F0: ("foo bar" + v4)}
        _ = fi15
        _ = v16
    }
    var v17 int16 = -1
    if false {
        _ = len(v3)
    } else {
        if ((!true) || false) {
            if true {
                v2 += v2
            }
        } else {
            v5 = (v5 & (v5 >> ((v2) & 63)))
            var v18 uint16 = 65535
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v20 := GenS2{F0: v4}
        _ = fi19
        _ = v20
    }
    var v21 int32 = 1
    if (!((2147483647 <= 32767) && false)) {
        var v22 int8 = -128
        _ = v22
    }
    v2 += (^(v2))
    v4 = v4
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS3{F0: -146, F1: 0, F2: (!(true || false))}
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
