package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int16; F1 uint8; F2 int }
func genFn0(p0 uint, p1 int64, p2 uint8) string {
    var v5 uint = (p0 ^ 42)
    switch 9223372036854775807 {
    case 128:
        v6 := GenS1{F0: 1, F1: (^(p2)), F2: (p1 << ((458) & 63))}
        _ = v6
    case 671:
        v7 := []int{-898, -9223372036854775808}
        _ = v7
    case -2147483648:
        fmt.Println(v5, p1, p0)
    default:
        var v8 uint8 = 128
        _ = v8
    }
    var v9 uint = (v5 << (((p0 - (^(v5)))) & 63))
    switch "hello" {
    case "a":
        v10 := []int{}
        _ = v10
    default:
        switch 9223372036854775807 {
        case 9223372036854775807:
            fmt.Println(v9, p2, v5)
        case -964:
            v5 += p0
        }
    }
    v11 := []int{-542, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    _ = v9
    return ("hello" + "foo bar")
}
func genFn1(p0 uint64) uint64 {
    v12 := []int{-9223372036854775808, 2147483647, 32767, 32768}
    v13 := GenS2{F0: 0, F1: 382}
    var v14 uint8 = 127
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 2
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]bool)
    var v3 int64 = -1
    var v4 uint64 = 42
    var v15 float32 = ((1.5 - (0.0 - float32((-2.25 + (3.14159 - ((0.1 / 1.5) / 1.5)))))) - (1e10 * 1.5))
    v16 := []int{503, 4294967295, 643}
    v17 := GenS3{F0: -2147483648, F1: 2}
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS4{F0: 127, F1: 42}
        v16 = append(v16, fi18)
        var v20 uint = v1
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int16 = 32767
    if (!(!((!((v4 - v4) <= (^(v4)))) || (42 > 1)))) {
        var v22 int32 = -11
        _ = v22
    }
    if (0.1 >= 0.1) {
        fmt.Println(v15, v3)
    }
    v15 -= 0.1
    v23 := GenS2{F0: 2147483647, F1: 1}
    v24 := GenS5{F0: (v21 + (v21 & (v21 & v21))), F1: 2, F2: 2147483647}
    v3 += 42
    var v25 int16 = (v21 | v21)
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
