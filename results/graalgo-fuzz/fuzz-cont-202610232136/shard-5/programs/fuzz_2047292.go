package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int, p1 bool) uint {
    var v4 int32 = 2147483647
    var v5 uint8 = 127
    fmt.Println(p0, v4)
    v6 := GenS1{F0: v5, F1: p0, F2: p1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int8, p1 bool, p2 int64) uint32 {
    p2 = -143
    v7 := GenS2{F0: 65535}
    var v8 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 int8 = -1
    delete(v1, -1)
    switch 0 {
    case 32768:
        v9 := []int{4294967295, 127}
        _ = v9
    case -87:
        v10 := []int{-2147483648, -9223372036854775808, 32768, 2147483647}
        _ = v10
    case 42:
        v3 = (v3 << ((1) & 63))
    default:
        if (!(!false)) {
            for fi11 := 0; fi11 < 2; fi11++ {
                v3 += v3
                fmt.Println(v3, v3)
                fmt.Println(fi11)
                _ = fi11
            }
        } else {
            v3 = (v3 & (-(v3)))
            fmt.Println(v3)
        }
    }
    var v12 int = -2147483648
    v13 := []int{v12, (-(v12)), (v12 + (v12 - int(v3)))}
    fmt.Println(v3, v3)
    var v14 uint = 18446744073709551615
    var v15 uint64 = 2
    var v16 int = (v12 * (v12 / ((v12) | 1)))
    if (0 != int32(v3)) {
        v16 += v12
    } else {
        fmt.Println(v12, v3, v15)
        for fi17 := 0; fi17 < 19; fi17++ {
            v1[(v12 / (((^(fi17))) | 1))] = ("a" + "a")
            v3 = 71
            var v18 float32 = ((((0.0 / 1e10) + -2.25) * float32((1e10 - 0.1))) + 0.0)
            _ = fi17
            _ = v18
        }
    }
    _ = len(v2)
    var v19 uint = genFn0(-530, ((127 > 95) || false))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
