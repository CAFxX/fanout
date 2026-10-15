package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int16 }
func genFn0(p0 int8, p1 uint16, p2 int8) uint16 {
    var v5 int32 = 1
    fmt.Println(p0, p2, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 665
}
func genFn1(p0 uint, p1 uint64) int16 {
    fmt.Println(p0)
    fmt.Println(p1, p1)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return -32768
}

func main() {
    var v1 int16 = 42
    var v2 uint8 = 1
    v3 := make(map[int]uint32)
    var v4 uint = 27
    var v6 int16 = genFn1(65535, 379)
    var v7 uint8 = (v2 / ((v2) | 1))
    var v8 uint8 = 42
    for fi9 := 0; fi9 < 6; fi9++ {
        switch "\n\t" {
        case "a":
            v2 += v7
        }
        var v10 uint16 = 2
        var v11 int = (fi9 | (fi9 + 127))
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v4 += 159
        var v13 uint32 = 2
        var v14 uint = v4
        _ = fi12
        _ = v13
        _ = v14
    }
    v4 = 255
    fmt.Println(v1)
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int32 = -1
        _ = fi15
        _ = v16
    }
    v1 = genFn1(v4, 32)
    v17 := GenS1{F0: "a", F1: -9223372036854775808, F2: (v1 << ((v4) & 63))}
    if false {
        switch 807 {
        case 0:
            var v18 float64 = (((0.0 - ((0.0 / 1.0) / 1.0)) - (0.1 - 3.14159)) / 3.14159)
            _ = v18
        }
        switch -9223372036854775808 {
        case 42:
            v2 -= 66
        default:
            var v19 string = ("" + "hello")
            _ = v19
        }
        v20 := GenS1{F0: "hello", F1: 760, F2: 556}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
