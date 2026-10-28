package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0(p0 uint16) int32 {
    v7 := GenS1{F0: 0, F1: "hello"}
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int, p1 uint32, p2 int64) int16 {
    var v9 bool = ((1.5 + (1.5 - ((1e10 * ((-2.25 * float32(3.14159)) + 0.0)) + -2.25))) == 1e10)
    for fi10 := 0; fi10 < 8; fi10++ {
        switch 0 {
        case 2147483647:
            var v11 uint32 = p1
            _ = v11
        case -9223372036854775808:
            if (1 != 2) {
                p0 -= fi10
                v12 := []int{p0, (fi10 | 9223372036854775807), (fi10 + (p0 >> ((18446744073709551615) & 63)))}
                var v13 uint16 = 255
                _ = v12
                _ = v13
            }
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    var v1 bool = (false || true)
    var v2 uint32 = 2147483647
    var v3 int = 32768
    v4 := make(map[string]bool)
    var v5 float64 = 3.14159
    var v6 uint16 = 270
    v14 := GenS1{F0: 1, F1: "hello"}
    for fi15 := 0; fi15 < 5; fi15++ {
        v1 = v1
        fmt.Println(v2, fi15, v5)
        v5 = v5
        _ = fi15
    }
    fmt.Println(v5)
    v2 = uint32(v6)
    if v1 {
        v4["hello"] = (346 < 9223372036854775807)
        v16 := []int{(^(v3)), v3, v3, (-(v3))}
        if (!false) {
            v1 = v1
            var v17 uint8 = 214
            _ = v17
        }
        _ = v16
    }
    if v1 {
        v18 := GenS2{F0: 1, F1: 42}
        _ = v18
    }
    v2 = v2
    v3 -= (v3 ^ v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
