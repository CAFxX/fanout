package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint, p1 int64) int16 {
    v7 := []int{2147483647}
    fmt.Println(p0, p0)
    var v8 uint64 = 478
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int, p1 string, p2 int) int16 {
    p1 = p1
    var v9 float32 = (0.0 - 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 32767
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    var v4 uint8 = 167
    v5 := []string{}
    v6 := make(map[int]int)
    v10 := GenS1{F0: -2147483648, F1: -20, F2: 51}
    v11 := GenS2{F0: 42}
    _ = len(v3)
    v1 = genFn1(0, "a", 32768)
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            if false {
                var v14 int64 = -494
                v14 += v14
                v1 = (v1 >> ((2) & 63))
                _ = v14
            }
            v4 = v4
            if true {
                _ = len(v5)
            } else {
                var v15 uint32 = 1
                v16 := GenS1{F0: -557, F1: -1, F2: 413}
                _ = v15
                _ = v16
            }
            var v17 int64 = 470
            v4 = (v4 | v4)
            _ = fi13
            _ = v17
        }
        var v18 uint8 = (^(v4))
        _ = fi12
        _ = v18
    }
    v1 = v1
    if ((v4 / ((v4) | 1)) <= v4) {
        var v19 int64 = -1
        _ = v19
    }
    if len(v5) > 0 {
        v5[0] = ("hello" + "foo bar")
    }
    v4 -= 255
    var v20 uint16 = 42
    if true {
        if ((^(v1)) <= (v1 >> ((4294967295) & 63))) {
            var v21 uint8 = 255
            var v22 uint = 0
            _ = v21
            _ = v22
        }
    } else {
        fmt.Println(v1)
    }
    v20 -= v20
    for fi23 := 0; fi23 < 11; fi23++ {
        v1 = v1
        _ = fi23
    }
    var v24 int64 = 1
    v1 = (v1 >> ((42) & 63))
    v3["hello"] = 42
    v25 := GenS3{F0: -9223372036854775808}
    v24 -= (v24 + v24)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
