package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint8 }
func genFn0(p0 bool, p1 int, p2 uint32) uint32 {
    if p0 {
        v4 := []int{p1, -369, p1, (p1 << ((255) & 63))}
        p1 -= (p1 >> ((0) & 63))
        v5 := GenS1{F0: 2, F1: 0}
        _ = v4
        _ = v5
    }
    v6 := []int{p1, 2147483647}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 443
}

func main() {
    var v1 float64 = 3.14159
    v2 := []int{-2147483648, -54}
    v3 := make(map[uint32]uint32)
    v1 += v1
    for fi7 := 0; fi7 < 17; fi7++ {
        v1 = v1
        if ((fi7 * (fi7 >> ((2) & 63))) >= fi7) {
            v1 = 0.0
            if true {
                v1 += float64((1e10 * 1e10))
                var v8 int64 = 0
                var v9 int16 = 0
                _ = v8
                _ = v9
            }
            if (!true) {
                v1 -= 0.1
            } else {
                var v10 int64 = -1
                v11 := GenS1{F0: 226, F1: 18446744073709551615}
                _ = v10
                _ = v11
            }
        }
        switch fi7 {
        case -2147483648:
            fmt.Println(fi7)
        default:
            switch "\n\t" {
            case "foo bar":
                v1 = (((0.1 / v1) + float64(3.14159)) / v1)
            }
        }
        _ = fi7
    }
    var v12 int16 = 32767
    var v13 uint = 952
    var v14 string = ("foo bar" + "hello")
    var v15 float32 = 0.0
    v13 -= (v13 | v13)
    v1 += (3.14159 * v1)
    v14 = "a"
    var v16 uint = v13
    v17 := GenS2{F0: 42}
    var v18 uint8 = 211
    var v19 uint64 = 454
    var v20 int = 2147483647
    var v21 int64 = -676
    fmt.Println(v18)
    switch v20 {
    case -740:
        v22 := GenS3{F0: (v12 >> ((v13) & 63)), F1: v18}
        _ = v22
    default:
        var v23 uint8 = (v18 ^ 53)
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
