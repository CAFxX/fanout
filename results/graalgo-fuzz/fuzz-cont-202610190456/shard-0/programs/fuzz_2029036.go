package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int32 }
func genFn0(p0 uint32) int {
    v7 := GenS1{F0: "hello", F1: 68}
    p0 -= (p0 ^ (^(p0)))
    _ = p0
    _ = v7
    return -629
}
func genFn1(p0 int, p1 uint32) uint8 {
    p1 += 42
    if false {
        p1 -= 476
        p0 = -402
    }
    v8 := GenS2{F0: 291, F1: 2}
    var v9 uint64 = 18446744073709551615
    var v10 uint16 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 128
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 743
    v3 := []int{v2, 4294967295, v2, 1}
    var v4 int32 = 0
    v5 := make(map[string]int)
    v6 := []int{(v2 / ((-503) | 1)), v2, (v2 ^ (-(v2))), 128}
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 float64 = (3.14159 * 3.14159)
        _ = fi11
        _ = v12
    }
    var v13 uint8 = genFn1(genFn0(42), 42)
    v6 = append(v6, 128)
    v2 = v2
    fmt.Println(v13, v4)
    switch "hello" {
    case "hello":
        if (0 <= 18446744073709551615) {
            v13 = (^(v13))
        }
    case "a":
        switch "hello" {
        case "hello":
            switch (v2 % ((v2) | 1)) {
            case 4294967295:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        case "\n\t":
            fmt.Println(v4, v2, v2)
        case "":
            var v14 uint16 = 1
            _ = v14
        }
    default:
        fmt.Println(v4)
    }
    var v15 string = "\n\t"
    v15 = (v15 + (v15 + "a"))
    var v16 int32 = v4
    var v17 uint64 = 649
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 string = ("a" + (v15 + "hello"))
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        if (4294967295 == (v17 - (v17 * 391))) {
            fmt.Println(v16, v16, v2)
        }
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := GenS3{F0: fi20, F1: 32767}
            _ = fi21
            _ = v22
        }
        v13 = (v13 ^ (v13 ^ (v13 + v13)))
        _ = fi20
    }
    v23 := GenS4{F0: 781, F1: (^(v17)), F2: -215}
    var v24 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
