package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0() uint32 {
    if (float64(0.0) == float64(float32(0.0))) {
        var v7 string = "a"
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 float32 = 0.1
        for fi10 := 0; fi10 < 6; fi10++ {
            if true {
                if ((0 < 9223372036854775807) && false) {
                    var v11 uint64 = 4294967295
                    _ = v11
                }
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    return 4294967295
}
func genFn1(p0 int, p1 uint64, p2 uint32) int8 {
    var v14 bool = ((false && false) || (false || ((p2 + (p2 - (p2 | genFn0()))) == 742)))
    p1 -= (p1 ^ (p1 - 18446744073709551615))
    fmt.Println(p1, p2)
    for fi15 := 0; fi15 < 9; fi15++ {
        p1 = p1
        fmt.Println(fi15, fi15)
        switch 2147483647 {
        case -9223372036854775808:
            switch ("" + "hello") {
            case "\n\t":
                var v16 int = 2147483647
                _ = v16
            case "hello":
                v17 := []int{(p0 + (p0 << ((2) & 63))), fi15, p0, int(p2)}
                _ = v17
            case "":
                v18 := []int{127}
                _ = v18
            }
        case 9223372036854775807:
            p0 -= fi15
        case 32768:
            v14 = v14
        }
        _ = fi15
    }
    var v19 int = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v19
    return -74
}

func main() {
    var v1 uint16 = 2
    var v2 int8 = 0
    var v3 float32 = 0.1
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    if (295 <= 585) {
        var v20 float32 = (0.1 - v3)
        var v21 int = 32768
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: v1, F1: (v1 | v1)}
    delete(v6, 4294967295)
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 uint16 = 2
        for fi25 := 0; fi25 < 15; fi25++ {
            fmt.Println(fi25, fi25)
            _ = fi25
        }
        v2 += (v2 % ((v2) | 1))
        _ = fi23
        _ = v24
    }
    switch "hello" {
    case "a":
        if (v2 == (v2 / ((v2) | 1))) {
            var v26 int32 = 1
            fmt.Println(v1)
            _ = v26
        }
    case "hello":
        if true {
            v2 = v2
        }
    }
    fmt.Println(v3, v3, v1)
    v27 := []int{1, -9223372036854775808, 766, -2147483648}
    var v28 int64 = 42
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(v1, v1, v3)
        v30 := GenS1{F0: 612, F1: (v1 / ((v1) | 1))}
        v3 = v3
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
