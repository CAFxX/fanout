package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 uint32; F1 string; F2 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint; F2 uint32 }
func genFn0(p0 uint32, p1 uint32) uint32 {
    p1 -= p1
    v5 := GenS1{F0: -459, F1: 255, F2: 143}
    v6 := []int{-1, 243, 2147483647}
    v7 := GenS1{F0: -375, F1: 2, F2: 1}
    v6 = append(v6, 2147483647)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 int) bool {
    p0 = (p0 - p0)
    p0 = (p0 >> ((0) & 63))
    v8 := GenS1{F0: 1, F1: 89, F2: 1}
    if (p0 > 4294967295) {
        p0 -= p0
    } else {
        p0 = p0
        p0 += 459
    }
    _ = p0
    _ = v8
    return false
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 uint64 = 1
    fmt.Println(v4)
    v9 := GenS2{F0: v4}
    v4 = 1
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(fi10, v4)
        _ = fi10
    }
    fmt.Println(v4)
    var v11 uint16 = 308
    v12 := GenS3{F0: 1, F1: 2147483647, F2: true}
    var v13 int64 = 1
    var v14 uint64 = (v4 - v4)
    switch 32768 {
    case -1:
        var v15 uint64 = 18446744073709551615
        _ = v15
    }
    var v16 int = -791
    v11 = 1
    var v17 float32 = (0.0 * 3.14159)
    switch -927 {
    case -328:
        v18 := GenS1{F0: -9223372036854775808, F1: 128, F2: 2}
        _ = v18
    case 32768:
        switch 0 {
        case 32767:
            v16 = (v16 << ((0) & 63))
        case 2147483647:
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 string = "hello"
                _ = fi19
                _ = v20
            }
        case 128:
            if (42 >= 677) {
                var v21 int = (v16 >> ((42) & 63))
                _ = v21
            } else {
                if false {
                    v22 := GenS4{F0: 2147483648, F1: "hello", F2: 127}
                    v23 := GenS2{F0: (v4 & (v14 & (^(v4))))}
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS3{F0: 2, F1: 32768, F2: false}
                    v4 += v4
                    v11 = 65535
                    _ = v24
                }
                v25 := GenS1{F0: (-(v13)), F1: 1, F2: 255}
                _ = v25
            }
        default:
            v26 := GenS5{F0: v13}
            _ = v26
        }
    default:
        var v27 uint8 = 127
        _ = v27
    }
    v4 -= 4294967295
    for fi28 := 0; fi28 < 1; fi28++ {
        v29 := []int{-1}
        v30 := GenS6{F0: 117, F1: 65535, F2: 676}
        fmt.Println(v13, v16)
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
