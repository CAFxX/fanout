package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 uint, p1 uint32, p2 int64) uint8 {
    v4 := []int{-9223372036854775808, 32767, 32767, 32767}
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5, p0)
        p1 -= (p1 + 42)
        _ = fi5
    }
    var v6 uint32 = (p1 / ((2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 127
}

func main() {
    var v1 bool = true
    var v2 float64 = (0.1 + 1e10)
    v3 := make(map[int]bool)
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        fmt.Println(v1, v1)
    default:
        switch -935 {
        case -75:
            var v7 bool = (true || v1)
            _ = v7
        default:
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int64 = -348
                var v10 uint32 = 42
                v11 := GenS1{F0: 1, F1: -165}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
    }
    v12 := []int{306, 4294967295, -2147483648, -410}
    v2 += ((v2 - v2) * v2)
    if (!(("\n\t" + ("hello" + ("\n\t" + ((((("foo bar" + "hello") + "a") + ("\n\t" + "")) + "hello") + "a")))) >= "hello")) {
        var v13 float64 = 3.14159
        _ = v13
    }
    v14 := v3[-622]
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: 1000, F1: 42}
        var v17 int16 = 1000
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 = v2
    v18 := GenS2{F0: 255, F1: 1}
    v19 := GenS2{F0: 189, F1: genFn0(1, 4294967295, 9223372036854775807)}
    if v14 {
        var v20 uint32 = 1
        _ = v20
    } else {
        if (439 > 2147483647) {
            var v21 uint32 = 4294967295
            switch 226 {
            case -9223372036854775808:
                v21 += 65535
            case -463:
                v1 = (127 <= -102)
            default:
                v22 := GenS2{F0: 42, F1: 128}
                _ = v22
            }
            _ = v21
        }
    }
    v23 := GenS2{F0: 4294967295, F1: 0}
    v12 = append(v12, 42)
    var v24 bool = (v14 || ((1000 > 42) && v1))
    if true {
        v25 := GenS3{F0: 18446744073709551615, F1: "hello"}
        var v26 uint8 = 128
        _ = v25
        _ = v26
    } else {
        var v27 float32 = float32((1.5 * 1e10))
        _ = v27
    }
    v28 := GenS1{F0: 0, F1: 1000}
    v29 := GenS3{F0: 42, F1: (("hello" + "hello") + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
}
