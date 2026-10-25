package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 uint8, p1 uint32, p2 uint64) string {
    v4 := GenS1{F0: 878}
    fmt.Println(p1)
    for fi5 := 0; fi5 < 18; fi5++ {
        p1 = (p1 * p1)
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "hello"
}
func genFn1() string {
    var v6 uint = 517
    var v7 float32 = float32(1e10)
    v6 -= (^(v6))
    v8 := GenS2{F0: 2147483647, F1: -1}
    v9 := GenS3{F0: v6, F1: (v6 << ((v6) & 63))}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 string = "hello"
    var v2 int = -150
    var v3 uint16 = 481
    if (((v3 & (v3 >> ((1) & 63))) > (^(v3))) || true) {
        v2 += v2
        v1 = v1
        v10 := []int{(v2 / ((-9223372036854775808) | 1)), v2}
        _ = v10
    }
    var v11 uint64 = 125
    v12 := []int{(-(v2)), (^(v2)), v2}
    var v13 uint8 = 42
    if (42 >= 1) {
        v13 = v13
        fmt.Println(v2)
    }
    fmt.Println(v3, v2)
    if len(v12) > 0 {
        v12[0] = (v2 + (v2 - (v2 - v2)))
    }
    switch ("foo bar" + ("hello" + v1)) {
    case "a":
        fmt.Println(v3, v2)
    case "":
        v14 := GenS2{F0: 1, F1: -197}
        _ = v14
    case "hello":
        var v15 float64 = 0.0
        _ = v15
    }
    if (1 < -92) {
        _ = len(v12)
    }
    v16 := GenS2{F0: 1, F1: 1}
    v11 = 18446744073709551615
    v13 = (v13 >> ((65535) & 63))
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int16 = int16(v2)
        _ = fi17
        _ = v18
    }
    var v19 bool = true
    if len(v12) > 0 {
        v12[0] = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
