package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 bool }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    var v5 string = "hello"
    v6 := make(map[int]int)
    v5 = (("hello" + v5) + v5)
    var v7 uint8 = 0
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: v5, F1: 35, F2: true}
        v10 := GenS2{F0: true, F1: 54, F2: 671}
        var v11 int16 = -1
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v7 -= (v7 | 42)
    v12 := GenS2{F0: false, F1: 42, F2: -1}
    v13 := GenS2{F0: false, F1: 2, F2: 0}
    v14 := GenS3{F0: 1, F1: 1}
    var v15 float32 = 1.5
    switch 128 {
    case 32768:
        var v16 float32 = ((3.14159 - v15) / (v15 / (((v15 + (((v15 * 3.14159) * ((((0.0 / v15) + v15) * (0.0 - (1e10 / (v15 / (0.0 / v15))))) * (v15 - ((v15 / 3.14159) / v15)))) * (v15 + ((v15 + v15) + v15)))) / (1.5 - 1e10)) - v15)))
        _ = v16
    case 0:
        switch -2147483648 {
        case 4294967295:
            v17 := GenS2{F0: (1 > 4294967295), F1: 4294967295, F2: 566}
            _ = v17
        case 32768:
            fmt.Println(v7)
        case -208:
            for fi18 := 0; fi18 < 2; fi18++ {
                v2["a"] = fi18
                _ = fi18
            }
        }
    default:
        v19 := GenS1{F0: "hello", F1: 15, F2: ((!(!(1 < 42))) && (2147483647 == -1))}
        _ = v19
    }
    delete(v6, -9223372036854775808)
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := []int{}
        _ = fi20
        _ = v21
    }
    var v22 uint = 0
    _ = len(v1)
    v23 := []int{}
    v15 = ((0.1 + ((v15 - v15) / v15)) - v15)
    v5 = "a"
    if (-1 < -206) {
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 uint = 4294967295
            _ = fi24
            _ = v25
        }
        if false {
            fmt.Println(v15, v5, v7)
            var v26 bool = false
            _ = v26
        } else {
            var v27 uint = 18446744073709551615
            v28 := GenS2{F0: (!(v7 != (v7 * 2))), F1: v22, F2: 1000}
            _ = v27
            _ = v28
        }
    } else {
        var v29 float32 = ((v15 / 3.14159) * -2.25)
        v22 -= (v22 % ((v22) | 1))
        v7 += uint8(v22)
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
