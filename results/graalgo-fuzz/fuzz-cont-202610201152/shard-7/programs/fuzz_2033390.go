package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
func genFn0() uint8 {
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 42}
    v7 = (v7 / (((v7 + (^(v7)))) | 1))
    v9 := GenS2{F0: 0, F1: 711}
    if true {
        v7 -= 42
    }
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 int32 = 42
    var v2 float64 = float64(-2.25)
    v3 := make(map[int]int)
    v4 := []uint32{2147483648, 4294967295, 869, 292}
    var v5 float64 = v2
    v6 := make(map[int]int)
    fmt.Println(v2, v2)
    if (-128 == -10) {
        v1 -= (v1 & v1)
        var v10 string = "hello"
        switch v10 {
        case "foo bar":
            var v11 int16 = -606
            _ = v11
        case "hello":
            v5 = v2
        case "":
            var v12 uint16 = 1
            _ = v12
        }
        _ = v10
    } else {
        var v13 uint = 4294967295
        v2 += v2
        _ = v13
    }
    v14 := GenS2{F0: 640, F1: 1}
    if false {
        v15 := v3[789]
        if true {
            var v16 float32 = 0.1
            _ = v16
        } else {
            v2 += float64(-2.25)
        }
        fmt.Println(v1, v5, v2)
        _ = v15
    }
    if (2 >= 2) {
        if true {
            for fi17 := 0; fi17 < 2; fi17++ {
                v5 = v2
                v18 := GenS1{F0: 200}
                _ = fi17
                _ = v18
            }
        }
        var v19 int = 832
        _ = v19
    } else {
        for fi20 := 0; fi20 < 4; fi20++ {
            v5 = 0.1
            _ = fi20
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            fmt.Println(fi21)
            switch fi21 {
            case 32767:
                for fi22 := 0; fi22 < 18; fi22++ {
                    v2 = ((1e10 / v5) + v2)
                    _ = fi22
                }
            case 9223372036854775807:
                v2 = 1.5
            default:
                if len(v4) > 0 {
                    v4[0] = 786
                }
            }
            fmt.Println(fi21, v2)
            _ = fi21
        }
        v23 := GenS1{F0: 640}
        _ = v23
    }
    v1 = v1
    var v24 int = 42
    v25 := GenS2{F0: 18446744073709551615, F1: 1}
    fmt.Println(v1, v5)
    v1 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
