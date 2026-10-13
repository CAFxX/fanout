package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint8 }
func genFn0(p0 uint64, p1 int, p2 int32) uint8 {
    var v4 int = (p1 | (p1 >> ((12) & 63)))
    switch (("" + "a") + "a") {
    case "foo bar":
        v5 := GenS1{F0: 0, F1: 255}
        _ = v5
    case "a":
        fmt.Println(p1, p1, p0)
    case "":
        fmt.Println(v4)
    default:
        v4 = v4
    }
    v6 := GenS2{F0: 2}
    p0 += p0
    if (!false) {
        if (2147483647 < 0) {
            p1 -= (p1 * 1)
            if (p2 < 2147483647) {
                v7 := []int{}
                var v8 int16 = -804
                _ = v7
                _ = v8
            } else {
                v9 := []int{}
                _ = v9
            }
            var v10 string = "hello"
            _ = v10
        }
    } else {
        v11 := GenS1{F0: 84, F1: 2}
        var v12 float32 = 3.14159
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 211
}

func main() {
    var v1 int32 = 379
    var v2 bool = (-1 <= 366)
    v3 := make(map[string]int)
    var v13 string = ("foo bar" + "")
    v2 = v2
    fmt.Println(v2, v13)
    var v14 int64 = -1
    v14 -= (-(v14))
    var v15 uint32 = 49
    var v16 float64 = 1.5
    v17 := GenS3{F0: "", F1: (v14 & (v14 * v14)), F2: 296}
    v18 := GenS4{F0: 2, F1: -800, F2: 2}
    var v19 bool = (117 > 16)
    fmt.Println(v16)
    for fi20 := 0; fi20 < 9; fi20++ {
        v14 += v14
        _ = fi20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v19 = v2
        var v22 int8 = 28
        _ = fi21
        _ = v22
    }
    fmt.Println(v13, v1, v14)
    v23 := []int{127, -840, 421}
    switch -62 {
    case 9223372036854775807:
        if v2 {
            var v24 int = 32767
            var v25 int64 = v14
            var v26 uint64 = 2
            _ = v24
            _ = v25
            _ = v26
        }
    case -84:
        var v27 string = v13
        _ = v27
    case 434:
        v13 = ("hello" + v13)
    }
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}
