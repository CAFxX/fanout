package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int16; F1 uint32; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 2147483647
    v3 := []int{42, 940, 684, 128}
    var v4 float32 = (1.5 / ((1e10 * float32(((0.0 / (((((((0.1 / (0.1 / float64(float32(0.1)))) / (-2.25 - 1e10)) - 1e10) * -2.25) / 1.5) * 0.1) * -2.25)) / (3.14159 + 1e10)))) + 0.1))
    var v5 int16 = 1
    var v6 int8 = 42
    v6 = (-(v6))
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v5, v4)
        _ = len(v1)
        v8 := GenS1{F0: (fi7 << ((65535) & 63))}
        _ = fi7
        _ = v8
    }
    switch -1 {
    case -1:
        var v9 bool = ((-(v6)) >= -30)
        _ = v9
    case 4294967295:
        v10 := GenS1{F0: 32768}
        _ = v10
    }
    switch 42 {
    case 666:
        if false {
            v11 := GenS2{F0: 1, F1: 255}
            v2 += (-(v2))
            _ = v11
        } else {
            var v12 int8 = -1
            if true {
                var v13 uint8 = 2
                v2 -= (-(v2))
                v14 := GenS3{F0: v5, F1: 2147483647, F2: -1}
                _ = v13
                _ = v14
            }
            v15 := GenS3{F0: -32768, F1: 2147483647, F2: -1}
            _ = v12
            _ = v15
        }
    }
    v2 = -1
    v4 = v4
    if false {
        v16 := GenS1{F0: 239}
        v17 := GenS1{F0: 1}
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 5; fi18++ {
            if (!false) {
                v5 = v5
                var v19 float64 = (0.0 / 1.5)
                _ = v19
            }
            fmt.Println(fi18)
            _ = fi18
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            if false {
                for fi21 := 0; fi21 < 11; fi21++ {
                    var v22 uint16 = 65535
                    _ = fi21
                    _ = v22
                }
                _ = len(v3)
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS1{F0: fi20}
                    v25 := GenS1{F0: fi23}
                    v6 = -114
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            switch ("a" + ("" + "hello")) {
            case "foo bar":
                _ = len(v3)
            case "\n\t":
                v26 := GenS1{F0: (fi20 * fi20)}
                _ = v26
            default:
                v5 -= v5
            }
            _ = fi20
        }
        var v27 uint8 = 0
        _ = v27
    }
    if true {
        v2 -= (v2 * 0)
        for fi28 := 0; fi28 < 19; fi28++ {
            if (32767 > 1) {
                var v29 string = "a"
                fmt.Println(v2, v4, v5)
                switch "\n\t" {
                case "":
                    var v30 float64 = 1e10
                    _ = v30
                case "\n\t":
                    var v31 float64 = (0.1 * 0.1)
                    _ = v31
                default:
                    if (32767 == 2) {
                        _ = len(v1)
                    } else {
                        v4 = v4
                        fmt.Println(v6, v2, fi28)
                    }
                }
                _ = v29
            } else {
                v4 = v4
            }
            _ = fi28
        }
        var v32 int8 = 127
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
