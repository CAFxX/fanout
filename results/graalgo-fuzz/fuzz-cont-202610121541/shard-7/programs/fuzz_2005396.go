package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64) int8 {
    p0 -= (^(p0))
    switch "a" {
    case "\n\t":
        v6 := []int{32768, 1, 9223372036854775807, 1}
        _ = v6
    }
    v7 := []int{127, 127, 128, 9223372036854775807}
    var v8 float32 = -2.25
    _ = p0
    _ = v7
    _ = v8
    return -10
}
func genFn1(p0 uint, p1 uint64) int {
    fmt.Println(p0)
    v9 := GenS1{F0: 4294967295, F1: 4294967295, F2: (918 < 2147483647)}
    if false {
        var v10 float32 = (0.0 * float32((float64(1e10) - 3.14159)))
        v11 := []int{-2147483648, 128, 2147483647, 9223372036854775807}
        _ = v10
        _ = v11
    }
    p1 = (p1 >> ((uint(p1)) & 63))
    _ = p0
    _ = p1
    _ = v9
    return 921
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := []string{"hello", "a", ""}
    var v4 int8 = -128
    v5 := make(map[string]string)
    var v12 int = 234
    if (false && true) {
        for fi13 := 0; fi13 < 18; fi13++ {
            switch -2147483648 {
            case 9223372036854775807:
                fmt.Println(v4)
            default:
                v14 := GenS1{F0: 2, F1: 18446744073709551615, F2: (0.1 != 0.1)}
                _ = v14
            }
            _ = fi13
        }
        fmt.Println(v12, v4)
    }
    v12 -= (v12 / ((genFn1(2, 0)) | 1))
    v15 := GenS2{F0: v12, F1: 42}
    fmt.Println(v12, v12)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int64 = 42
            var v19 string = ""
            fmt.Println(v18)
            _ = fi17
            _ = v18
            _ = v19
        }
        fmt.Println(fi16)
        _ = fi16
    }
    _ = len(v3)
    v4 = (^(v4))
    for fi20 := 0; fi20 < 19; fi20++ {
        fmt.Println(fi20, v12, v12)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi20
    }
    if false {
        var v21 int8 = (v4 ^ 100)
        var v22 float64 = (1e10 * (float64(0.1) * 0.1))
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v24 uint = 0
                if false {
                    var v25 float32 = (1e10 * (0.0 / (1e10 - ((-2.25 + (0.0 / 1e10)) * 1.5))))
                    _ = v25
                } else {
                    v26 := GenS3{F0: 42, F1: 232}
                    fmt.Println(v4, v4)
                    _ = v26
                }
                _ = v24
            }
            var v27 string = ("\n\t" + ("hello" + ("" + "foo bar")))
            fmt.Println(v4)
            _ = fi23
            _ = v27
        }
        _ = len(v5)
    }
    fmt.Println(v4, v4)
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 uint8 = 104
        if true {
            v30 := v5["a"]
            v29 += (v29 & v29)
            _ = v30
        }
        v1["a"] = 65535
        _ = fi28
        _ = v29
    }
    v12 -= 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
