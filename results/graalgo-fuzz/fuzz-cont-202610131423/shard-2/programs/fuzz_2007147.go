package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float32 = 1e10
    var v2 int64 = 42
    var v3 uint8 = 0
    var v4 string = "\n\t"
    v5 := make(map[string]int)
    var v6 uint8 = (v3 / (((v3 >> ((179) & 63))) | 1))
    v7 := v5["foo bar"]
    fmt.Println(v6, v4, v3)
    switch v7 {
    case 2147483647:
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint8 = (v6 ^ 255)
            v10 := []int{-9223372036854775808, fi8, v7}
            v7 += v7
            _ = fi8
            _ = v10
            _ = v9
        }
    case 32767:
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint16 = 32767
            _ = fi11
            _ = v12
        }
    case 382:
        v7 = 2147483647
    default:
        var v13 uint64 = 0
        _ = v13
    }
    fmt.Println(v7, v2)
    var v14 string = "a"
    v1 = ((v1 - v1) + (v1 + v1))
    v15 := []int{}
    switch (v7 ^ 128) {
    case 128:
        var v16 uint16 = 42
        _ = v16
    case 42:
        var v17 bool = false
        _ = v17
    case -9223372036854775808:
        v6 -= 0
    default:
        v6 = (v3 % ((128) | 1))
    }
    if (!false) {
        var v18 string = (v4 + v4)
        v2 = 397
        _ = v18
    } else {
        var v19 float64 = float64(v1)
        v1 = (-2.25 / v1)
        _ = v19
    }
    var v20 string = "foo bar"
    if (5 != 65535) {
        if (68 < 42) {
            var v21 uint16 = 255
            v6 = 25
            fmt.Println(v20, v14)
            _ = v21
        } else {
            v22 := GenS1{F0: (^(v7))}
            _ = v22
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            if true {
                v4 = v14
                v24 := GenS2{F0: 42}
                _ = v24
            }
            _ = fi23
        }
    } else {
        v15 = append(v15, v7)
        v25 := v5[(v14 + "a")]
        var v26 uint64 = 18446744073709551615
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: (v7 << ((2) & 63))}
    v2 = 42
    var v28 uint16 = 2
    v20 = (("a" + v20) + "")
    if false {
        v14 = (v4 + (v14 + "\n\t"))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
