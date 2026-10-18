package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
func genFn0() uint32 {
    var v7 int16 = 1000
    if true {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 bool = true
            _ = fi8
            _ = v9
        }
    }
    v10 := []int{2147483647}
    _ = v10
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 uint64, p2 string) int64 {
    var v11 uint = uint(p1)
    p0 = (!p0)
    var v12 string = (p2 + (((p2 + p2) + (p2 + "a")) + "a"))
    var v13 int = 0
    var v14 int32 = -565
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -128
    var v3 int32 = 1
    v4 := []int{}
    var v5 uint16 = 42
    v6 := []int{-325, 4294967295, 128, 128}
    v2 += v2
    v6 = append(v6, 991)
    var v15 int = 0
    fmt.Println(v15, v2)
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 float64 = 1.5
            if len(v4) > 0 {
                v4[0] = v15
            }
            var v19 string = (("hello" + "hello") + ("" + (("foo bar" + "hello") + "hello")))
            switch fi17 {
            case 9223372036854775807:
                v20 := GenS1{F0: 255, F1: 1, F2: (v3 + (v3 & (v3 | (v3 >> ((4294967295) & 63)))))}
                _ = v20
            default:
                if true {
                    switch v19 {
                    case "hello":
                        v18 -= v18
                    }
                    _ = len(v6)
                    v2 -= (v2 | -128)
                }
                v6 = append(v6, fi17)
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        v21 := GenS1{F0: 439, F1: 9223372036854775807, F2: 0}
        var v22 uint8 = 127
        _ = fi16
        _ = v21
        _ = v22
    }
    v15 = 724
    fmt.Println(v3)
    var v23 string = ("foo bar" + "")
    var v24 string = ""
    v1[v15] = 500
    fmt.Println(v23, v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
