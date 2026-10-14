package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
func genFn0(p0 int64, p1 uint32) uint64 {
    if true {
        if false {
            v5 := []int{0}
            _ = v5
        } else {
            v6 := GenS1{F0: 9223372036854775807, F1: ""}
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 int32 = 2147483647
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        if true {
            p0 = -757
        }
        p1 -= (p1 >> ((65535) & 63))
    }
    v9 := []int{0, -1, 32767}
    _ = p0
    _ = p1
    _ = v9
    return 4294967295
}

func main() {
    v1 := []uint32{361, 2147483647}
    var v2 uint32 = 895
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    if (true && (!true)) {
        v10 := GenS1{F0: -362, F1: "\n\t"}
        _ = v10
    }
    v11 := GenS1{F0: -9223372036854775808, F1: "hello"}
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS1{F0: 42, F1: "hello"}
        _ = fi12
        _ = v13
    }
    var v14 int32 = 105
    if (100 > 1) {
        fmt.Println(v2, v14)
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int = fi15
        v17 := GenS1{F0: -9223372036854775808, F1: ((("hello" + "a") + ("a" + ("" + "\n\t"))) + "\n\t")}
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v2, v14)
    var v18 int32 = v14
    _ = len(v4)
    if true {
        v14 = (v18 ^ (^(v14)))
    } else {
        v14 += (v18 & (v18 & (^(v18))))
    }
    v2 = 1
    v2 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
