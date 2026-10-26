package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
func genFn0(p0 uint, p1 int32) bool {
    var v4 int = 904
    if (1 > -51) {
        var v5 uint32 = 42
        var v6 uint = (^(p0))
        _ = v5
        _ = v6
    } else {
        var v7 uint = (^(p0))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return (v4 >= -9223372036854775808)
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = -833
    v3 := make(map[uint32]int)
    var v8 int16 = 800
    v1 = (v1 - 18446744073709551615)
    var v9 int = 0
    if true {
        v10 := []int{128}
        _ = v10
    }
    v1 += 847
    var v11 float64 = float64(1.5)
    for fi12 := 0; fi12 < 12; fi12++ {
        v3[932] = -141
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := []int{int(v1), fi13, (fi12 ^ -26), 2147483647}
            fmt.Println(v9, v11)
            _ = fi13
            _ = v14
        }
        v2 -= v2
        _ = fi12
    }
    var v15 int64 = 943
    if (127 < 113) {
        var v16 string = ("a" + "a")
        v8 += 42
        _ = v16
    } else {
        v17 := []int{(v9 & (^(v9))), 32767, (v9 - (v9 & 9223372036854775807)), (-(v9))}
        _ = v17
    }
    fmt.Println(v8, v9, v15)
    v3[249] = 271
    v18 := GenS1{F0: (v15 ^ 314), F1: "hello", F2: 457}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
