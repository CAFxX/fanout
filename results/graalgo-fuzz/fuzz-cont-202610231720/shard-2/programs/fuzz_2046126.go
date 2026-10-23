package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint32) int16 {
    p0 = p0
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int8 = 127
        var v8 bool = (2 > 900)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    return 32767
}
func genFn1() uint8 {
    v9 := []int{127, 128, 9223372036854775807, 409}
    switch "hello" {
    case "foo bar":
        if (false || (67 != 1)) {
            switch 32767 {
            case -1:
                v9 = append(v9, 127)
            }
            var v10 int = 0
            var v11 string = (((("foo bar" + "\n\t") + ("" + "")) + "") + "a")
            _ = v10
            _ = v11
        }
    default:
        if ("foo bar" == (("foo bar" + "") + "")) {
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS1{F0: 128, F1: 11}
                _ = fi12
                _ = v13
            }
            var v14 uint16 = 986
            v14 = v14
            _ = v14
        }
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            switch "a" {
            case "foo bar":
                var v17 int64 = -1
                _ = v17
            default:
                var v18 float32 = 0.0
                _ = v18
            }
            var v19 float64 = (((3.14159 * ((-2.25 * 1e10) / 1.0)) - float64(-2.25)) / ((1e10 + -2.25) / 1e10))
            v19 = 0.0
            if ((65535 != 2) && true) {
                var v20 int16 = -161
                var v21 string = "hello"
                var v22 int = 42
                _ = v20
                _ = v21
                _ = v22
            }
            var v23 float64 = v19
            switch (fi15 + fi15) {
            case 903:
                var v24 uint32 = 2147483647
                _ = v24
            case -9223372036854775808:
                v25 := GenS2{F0: 9223372036854775807}
                _ = v25
            default:
                var v26 uint16 = 1
                _ = v26
            }
            _ = fi16
            _ = v19
            _ = v23
        }
        fmt.Println(fi15)
        _ = fi15
    }
    var v27 int8 = 0
    _ = v27
    _ = v9
    return 128
}

func main() {
    v1 := []uint32{}
    var v2 string = "hello"
    var v3 uint64 = 1
    var v4 int = 42
    v5 := make(map[int]string)
    fmt.Println(v4, v2)
    v3 = (v3 & (v3 & 754))
    _ = len(v5)
    var v28 uint8 = 42
    v29 := v5[(v4 + v4)]
    fmt.Println(v4, v3)
    v2 = (v29 + v2)
    v30 := GenS2{F0: 247}
    v28 -= 0
    _ = len(v1)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
