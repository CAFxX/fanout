package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 int64, p1 int64) int8 {
    v4 := GenS1{F0: 529, F1: (p1 % (((p0 | p1)) | 1)), F2: 0}
    for fi5 := 0; fi5 < 15; fi5++ {
        p0 += p1
        v6 := GenS1{F0: 42, F1: int64(fi5), F2: -1}
        _ = fi5
        _ = v6
    }
    v7 := []int{9223372036854775807, -633, 32767, 970}
    for fi8 := 0; fi8 < 5; fi8++ {
        if (false || ("a" >= ("hello" + "a"))) {
            fmt.Println(fi8, p1, p0)
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return -12
}
func genFn1(p0 uint) uint64 {
    fmt.Println(p0)
    v9 := GenS1{F0: 42, F1: -1, F2: 0}
    p0 = 18446744073709551615
    _ = p0
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    if (true && true) {
        var v10 int32 = 2147483647
        v10 = (v10 >> ((0) & 63))
        _ = v10
    }
    var v11 float64 = 0.1
    v11 = v11
    var v12 int8 = -107
    var v13 uint32 = 270
    var v14 uint8 = 54
    for fi15 := 0; fi15 < 17; fi15++ {
        if (42 == 42) {
            v12 = 1
            v12 -= (v12 + (v12 * (v12 / ((v12) | 1))))
        }
        v16 := GenS2{F0: fi15}
        switch "\n\t" {
        case "a":
            v17 := GenS1{F0: (v13 / (((^(v13))) | 1)), F1: 9223372036854775807, F2: int32(v14)}
            _ = v17
        case "hello":
            v18 := GenS1{F0: v13, F1: 0, F2: -264}
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v12 += genFn0(-9223372036854775808, 9223372036854775807)
    v19 := []int{}
    if true {
        v12 += (v12 / (((v12 + (v12 % (((v12 + genFn0(1, 42))) | 1)))) | 1))
    } else {
        v11 += 0.1
        var v20 float64 = 0.0
        var v21 uint64 = 2
        _ = v20
        _ = v21
    }
    fmt.Println(v11)
    switch 128 {
    case 2147483647:
        fmt.Println(v13, v12, v12)
    case -9223372036854775808:
        delete(v1, -758)
    case 821:
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 uint64 = 0
            fmt.Println(fi22, v13)
            v13 = uint32(v14)
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v12)
    if true {
        if false {
            if true {
                var v24 float32 = float32(float64(3.14159))
                _ = v24
            }
            var v25 int = 82
            v13 += v13
            _ = v25
        }
        v11 = 1.5
        fmt.Println(v11)
    } else {
        for fi26 := 0; fi26 < 10; fi26++ {
            for fi27 := 0; fi27 < 13; fi27++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                v28 := v2[fi26]
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
}
