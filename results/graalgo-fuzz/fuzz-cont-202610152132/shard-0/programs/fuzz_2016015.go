package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int }
func genFn0(p0 bool, p1 uint16) uint8 {
    var v5 uint8 = 0
    var v6 uint = 1
    var v7 float32 = 1e10
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 int64, p1 int64) bool {
    var v8 uint16 = 338
    v8 = v8
    _ = p0
    _ = p1
    _ = v8
    return false
}

func main() {
    var v1 float32 = -2.25
    var v2 int8 = 52
    var v3 uint8 = 255
    v4 := make(map[int]bool)
    for fi9 := 0; fi9 < 20; fi9++ {
        switch "foo bar" {
        case "a":
            v10 := []int{fi9, (-(fi9))}
            _ = v10
        case "hello":
            v3 = 236
        default:
            var v11 int32 = 42
            _ = v11
        }
        fmt.Println(v2)
        delete(v4, (fi9 % (((fi9 | (fi9 >> ((65535) & 63)))) | 1)))
        _ = fi9
    }
    v1 -= float32(3.14159)
    v12 := GenS1{F0: 65535, F1: 0}
    var v13 float32 = ((0.1 + float32(0.0)) + (0.1 + (float32(3.14159) / (1.5 + v1))))
    var v14 string = "foo bar"
    var v15 int8 = 42
    v16 := GenS2{F0: 42, F1: (v3 * v3), F2: -9223372036854775808}
    v2 -= (v15 | v15)
    v13 = (v13 + 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
