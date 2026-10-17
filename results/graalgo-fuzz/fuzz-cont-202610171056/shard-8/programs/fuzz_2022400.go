package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0(p0 int8) bool {
    fmt.Println(p0)
    v7 := []int{2147483647, 4294967295, 1}
    if (!true) {
        p0 = p0
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int16 = 1000
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi8
            _ = v9
        }
        p0 = 0
        for fi10 := 0; fi10 < 13; fi10++ {
            p0 = (p0 >> ((1) & 63))
            p0 = p0
            _ = fi10
        }
    }
    v11 := GenS1{F0: 476, F1: "hello", F2: 2147483647}
    _ = p0
    _ = v11
    _ = v7
    return true
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 232
    v3 := make(map[int]int)
    var v4 int32 = 1
    var v5 uint64 = v1
    var v6 int64 = 871
    var v12 uint32 = 2147483648
    fmt.Println(v5, v1)
    for fi13 := 0; fi13 < 11; fi13++ {
        v5 = 1
        var v14 uint64 = 156
        _ = fi13
        _ = v14
    }
    if (v6 != (v6 / (((v6 | (v6 + (v6 + -255)))) | 1))) {
        var v15 int = 4294967295
        var v16 int16 = int16(v15)
        var v17 string = "hello"
        _ = v15
        _ = v16
        _ = v17
    } else {
        switch 155 {
        case -139:
            var v18 uint = 42
            _ = v18
        case 9223372036854775807:
            v4 += v4
        default:
            var v19 uint = 469
            _ = v19
        }
    }
    v4 = 2147483647
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int = fi20
        v22 := []int{0, (^(fi20)), v21, int(v1)}
        for fi23 := 0; fi23 < 14; fi23++ {
            v21 = -1
            if false {
                var v24 uint16 = 248
                _ = v24
            } else {
                var v25 uint8 = v2
                fmt.Println(fi20, fi23, v2)
                _ = v25
            }
            _ = fi23
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    if true {
        var v26 int32 = int32(v5)
        v27 := []int{int(v2), 127, 127, 0}
        v3[524] = 2147483647
        _ = v26
        _ = v27
    }
    v28 := v3[-2147483648]
    if (-1 == -956) {
        v6 = (v6 ^ (v6 * 1))
    } else {
        var v29 int32 = v4
        switch v28 {
        case 1:
            if (-69 >= 18) {
                v30 := GenS1{F0: 669, F1: "hello", F2: 42}
                _ = v30
            } else {
                v31 := v3[v28]
                v32 := GenS2{F0: 255, F1: 255, F2: 9223372036854775807}
                v4 = (v29 % ((1) | 1))
                _ = v31
                _ = v32
            }
        }
        v5 = 2
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
