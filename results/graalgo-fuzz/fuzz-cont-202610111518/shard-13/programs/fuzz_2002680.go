package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }

func main() {
    var v1 uint16 = 42
    var v2 int16 = -32768
    var v3 uint = 0
    var v4 int32 = -2147483648
    v5 := make(map[string]int)
    v2 = (v2 * int16(v4))
    var v6 uint16 = (^(v1))
    switch "a" {
    case "foo bar":
        var v7 float32 = (float32(0.0) * 3.14159)
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 string = ("\n\t" + "hello")
                var v11 bool = true
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    default:
        v12 := GenS1{F0: v3}
        _ = v12
    }
    v1 = 255
    v1 -= 32767
    switch 165 {
    case 32767:
        v13 := GenS2{F0: 0}
        _ = v13
    default:
        fmt.Println(v6, v6, v4)
    }
    if (-369 == 1) {
        if (!(false || false)) {
            v1 = (v6 - (^(v1)))
            var v14 float32 = -2.25
            _ = v14
        }
    }
    var v15 int64 = -9223372036854775808
    switch "hello" {
    case "hello":
        fmt.Println(v15, v1, v6)
    }
    switch int(v4) {
    case -591:
        var v16 int = 32767
        _ = v16
    default:
        fmt.Println(v1, v15)
    }
    v1 += (v1 ^ v1)
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int32 = v4
        _ = fi17
        _ = v18
    }
    fmt.Println(v15)
    v6 += v1
    v2 += (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
