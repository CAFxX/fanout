package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int; F1 string }
func genFn0(p0 uint64) int8 {
    p0 += p0
    p0 -= 335
    var v7 string = "hello"
    v8 := GenS1{F0: 153}
    fmt.Println(v7, v7)
    _ = p0
    _ = v7
    _ = v8
    return -86
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 bool = true
    v4 := make(map[uint32]bool)
    v5 := make(map[int]string)
    var v6 uint8 = 150
    v9 := []int{-893, -396, 128}
    for fi10 := 0; fi10 < 6; fi10++ {
        if (!((0.1 / 0.1) >= (((-2.25 / 3.14159) * 1e10) * 1.5))) {
            v11 := GenS2{F0: -1, F1: 42, F2: 1}
            for fi12 := 0; fi12 < 20; fi12++ {
                v3 = ((0.0 * (0.0 / 1.0)) < 3.14159)
                v13 := GenS3{F0: (fi10 % ((fi10) | 1)), F1: ("foo bar" + ("foo bar" + "hello"))}
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                if v3 {
                    v6 -= (v6 + 128)
                    v3 = (true || ((((0.1 - 1e10) * 0.1) != ((1e10 / 1.0) / -2.25)) && v3))
                    var v15 uint8 = v6
                    _ = v15
                } else {
                    v16 := GenS3{F0: (fi14 >> ((490) & 63)), F1: (("foo bar" + "hello") + ("" + "a"))}
                    v6 -= 204
                    if v3 {
                        v17 := GenS2{F0: 98, F1: 128, F2: 255}
                        var v18 int32 = 2147483647
                        fmt.Println(v3, fi14)
                        _ = v17
                        _ = v18
                    }
                    _ = v16
                }
                v9 = append(v9, (^(fi14)))
                _ = fi14
            }
            _ = v11
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 float64 = 0.1
            switch "" {
            case "a":
                v20 -= v20
            }
            _ = fi19
            _ = v20
        }
        _ = fi10
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v6, v6)
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS1{F0: 2}
        v3 = (true || v3)
        _ = fi21
        _ = v22
    }
    v6 = (v6 | 1)
    var v23 string = ("" + ("" + "hello"))
    var v24 uint64 = 18446744073709551615
    for fi25 := 0; fi25 < 1; fi25++ {
        v6 += v6
        _ = fi25
    }
    if true {
        if len(v9) > 0 {
            v9[0] = 32768
        }
        var v26 int16 = 1
        if v3 {
            var v27 uint16 = 65535
            _ = v27
        }
        _ = v26
    }
    switch 4294967295 {
    case 107:
        fmt.Println(v23, v3)
    case -922:
        v3 = (1000 == -32768)
    default:
        fmt.Println(v24, v6, v23)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
