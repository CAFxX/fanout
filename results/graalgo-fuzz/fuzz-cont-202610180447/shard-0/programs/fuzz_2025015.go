package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint8) int16 {
    p0 -= (p0 & 427)
    var v4 uint64 = 0
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1(p0 int) uint {
    p0 = p0
    p0 = 42
    _ = p0
    return 0
}

func main() {
    var v1 float32 = (1e10 * 0.1)
    var v2 uint32 = 42
    var v3 int8 = 127
    var v5 float64 = ((-2.25 + 1.5) + 0.1)
    v6 := []int{42, 9223372036854775807, 786, 134}
    v7 := GenS1{F0: 255, F1: (false || (("" != "foo bar") || (true || (false || false)))), F2: ((("" + ("a" + "hello")) + (("hello" + "\n\t") + "foo bar")) + "foo bar")}
    for fi8 := 0; fi8 < 13; fi8++ {
        v2 += v2
        if true {
            v5 = (float64(1e10) * float64(-2.25))
            var v9 bool = (false && (!false))
            fmt.Println(v9, v9)
            _ = v9
        } else {
            v2 -= 42
            v6 = append(v6, 646)
            var v10 int32 = 1
            _ = v10
        }
        var v11 float32 = (float32(0.0) + 1.5)
        _ = fi8
        _ = v11
    }
    switch ("" + "a") {
    case "a":
        v6 = append(v6, 0)
    case "hello":
        v1 = (v1 + float32(3.14159))
    default:
        fmt.Println(v1)
    }
    fmt.Println(v1, v5, v2)
    var v12 float32 = v1
    v3 = (^(v3))
    if ((v12 - v1) >= v12) {
        v6 = append(v6, 127)
        v13 := GenS2{F0: 834}
        var v14 uint32 = (v2 * (v2 << ((65535) & 63)))
        _ = v13
        _ = v14
    }
    var v15 uint8 = 2
    v5 = 1e10
    var v16 float64 = 1e10
    v2 -= (v2 ^ v2)
    if (true && false) {
        fmt.Println(v3)
    } else {
        v17 := GenS3{F0: uint16(v15)}
        _ = v17
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
