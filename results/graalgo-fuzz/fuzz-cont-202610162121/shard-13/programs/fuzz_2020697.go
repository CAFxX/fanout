package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint8 }

func main() {
    v1 := []string{"", "a"}
    v2 := make(map[string]uint32)
    v3 := []string{"hello"}
    var v4 int16 = -742
    fmt.Println(v4)
    switch "hello" {
    case "a":
        fmt.Println(v4)
    default:
        v5 := GenS1{F0: 127}
        _ = v5
    }
    v2[""] = 2
    if true {
        v4 += (-(v4))
        fmt.Println(v4)
        var v6 uint16 = 35
        _ = v6
    } else {
        fmt.Println(v4)
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        if true {
            v4 -= 76
        }
        v8 := GenS2{F0: -1, F1: -2147483648, F2: 210}
        var v9 uint32 = 65535
        _ = fi7
        _ = v8
        _ = v9
    }
    v2[""] = 2
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "a" {
        case "foo bar":
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 float32 = (((1.5 + float32(0.0)) + (-2.25 / 0.1)) + 0.1)
                v4 = (v4 - 1000)
                if true {
                    fmt.Println(v12)
                    v12 = v12
                } else {
                    var v13 float32 = float32(3.14159)
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
        case "\n\t":
            v4 = (v4 & (^(v4)))
        case "a":
            fmt.Println(v4)
        }
        v14 := GenS1{F0: 106}
        _ = fi10
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v2["hello"] = 2147483648
        v4 += (v4 % ((v4) | 1))
        v4 = (v4 >> ((2) & 63))
        _ = fi15
    }
    if false {
        fmt.Println(v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
