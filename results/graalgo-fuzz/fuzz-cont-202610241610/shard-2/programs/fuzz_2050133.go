package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int64; F1 int; F2 uint8 }
func genFn0(p0 uint, p1 uint8) string {
    v6 := GenS1{F0: 1}
    var v7 bool = false
    if (32767 == 585) {
        p1 = p1
        p1 = (^(p1))
    }
    switch 767 {
    case -2147483648:
        var v8 float64 = (1e10 - (-2.25 * 0.1))
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 string = "foo bar"
    var v2 float64 = (1e10 * float64(3.14159))
    v3 := make(map[uint32]string)
    var v4 uint16 = 44
    var v5 int8 = -72
    v9 := []int{-2147483648, int(v5), 32768, 4294967295}
    v9 = append(v9, 0)
    switch v1 {
    case "hello":
        var v10 int = -9223372036854775808
        _ = v10
    case "foo bar":
        v11 := GenS2{F0: v1, F1: 2147483647}
        _ = v11
    }
    fmt.Println(v4, v4, v1)
    var v12 int32 = 628
    v13 := GenS2{F0: ((v1 + "\n\t") + "foo bar"), F1: -9223372036854775808}
    v12 += v12
    if (237 <= 0) {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int = fi14
            _ = fi14
            _ = v15
        }
        if (true || true) {
            _ = len(v3)
            var v16 uint8 = 128
            _ = v16
        }
        fmt.Println(v12, v2)
    } else {
        v17 := GenS3{F0: 1, F1: -1, F2: 2}
        _ = v17
    }
    var v18 uint16 = 2
    _ = len(v3)
    fmt.Println(v5, v18, v2)
    for fi19 := 0; fi19 < 1; fi19++ {
        if false {
            v4 -= (v4 & (^(v18)))
            var v20 uint16 = 42
            _ = v20
        } else {
            v18 = v4
            if false {
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 float32 = 1.5
                    _ = fi21
                    _ = v22
                }
            }
        }
        var v23 int16 = 32767
        v12 = 0
        _ = fi19
        _ = v23
    }
    if len(v9) > 0 {
        v9[0] = -779
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 float64 = (v2 - (float64((-2.25 + float32(1.5))) - (v2 / float64(0.1))))
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int16 = 354
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
