package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint8 }
type GenS2 struct { F0 int32; F1 int; F2 uint16 }
func genFn0(p0 int8, p1 uint64, p2 int16) int32 {
    v4 := GenS1{F0: 542, F1: 2147483647, F2: 255}
    p2 += (-(p2))
    v5 := []int{128, 32767, 9223372036854775807, -539}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -937
}

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    var v3 float32 = 1.5
    v1 = v1
    v3 -= v3
    v6 := []int{-479, 32768}
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(fi7, fi7, fi7)
        switch 128 {
        case 1:
            if true {
                var v8 int8 = -37
                _ = v8
            } else {
                var v9 int = fi7
                var v10 float64 = 0.0
                _ = v10
                _ = v9
            }
        case 885:
            if v1 {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v11 uint64 = 2
                fmt.Println(v11, fi7, v11)
                _ = v11
            }
        case -1:
            var v12 uint = 2
            _ = v12
        }
        _ = fi7
    }
    _ = len(v2)
    if v1 {
        switch ("foo bar" + "hello") {
        case "hello":
            if true {
                var v13 int32 = 2147483647
                _ = v13
            }
        case "\n\t":
            var v14 int = 2147483647
            _ = v14
        default:
            if true {
                v1 = (18446744073709551615 < 2)
                var v15 int = 476
                v1 = (!false)
                _ = v15
            }
        }
        v16 := v2[588]
        _ = v16
    } else {
        if v1 {
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        }
        if v1 {
            v6 = append(v6, 771)
        } else {
            v3 = (0.0 + v3)
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 uint64 = 827
                _ = fi17
                _ = v18
            }
            fmt.Println(v1, v1)
        }
    }
    var v19 float64 = 3.14159
    v3 += (v3 / 1.0)
    for fi20 := 0; fi20 < 12; fi20++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(v19, v19, v19)
            var v22 float64 = (float64(v3) - (v19 - 0.0))
            _ = fi21
            _ = v22
        }
        var v23 string = "foo bar"
        _ = fi20
        _ = v23
    }
    v24 := GenS2{F0: 2147483647, F1: 128, F2: 336}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v6
}
