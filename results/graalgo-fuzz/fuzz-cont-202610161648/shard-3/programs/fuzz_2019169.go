package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint, p1 int16, p2 int64) uint64 {
    fmt.Println(p0, p1, p2)
    var v7 uint8 = 28
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 994
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 0
    var v3 bool = false
    var v4 float32 = 1e10
    var v5 int64 = 42
    var v6 float32 = v4
    var v8 float32 = (((v6 + v4) * 0.0) + 0.1)
    if (v3 && v3) {
        v2 = (v2 % (((v2 / ((v2) | 1))) | 1))
        var v9 uint = 4294967295
        _ = v9
    }
    v2 -= -2147483648
    if (!((-2.25 + v8) == (-2.25 + v4))) {
        switch -2147483648 {
        case 0:
            var v10 int16 = 668
            _ = v10
        }
        v11 := v1[376]
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint32 = 1
        v5 -= (v5 ^ v5)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: v2}
    if false {
        _ = len(v1)
        v8 += float32(1.5)
    } else {
        v6 = (v4 / v4)
    }
    var v15 uint64 = 1
    v4 = v6
    v16 := GenS2{F0: 93, F1: 4294967295}
    switch "hello" {
    case "hello":
        v17 := []int{1, 127, 9223372036854775807, int(v5)}
        _ = v17
    case "":
        v18 := GenS3{F0: -1}
        _ = v18
    }
    var v19 int8 = 99
    switch 4294967295 {
    case 42:
        var v20 uint8 = 255
        _ = v20
    case 9223372036854775807:
        var v21 uint32 = 318
        _ = v21
    default:
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
