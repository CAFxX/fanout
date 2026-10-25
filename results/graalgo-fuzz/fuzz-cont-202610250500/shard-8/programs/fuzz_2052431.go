package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 uint64, p1 int8) bool {
    p0 -= (p0 ^ (p0 - p0))
    switch "\n\t" {
    case "":
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := []int{fi5, (fi5 & fi5), 0, fi5}
            _ = fi5
            _ = v6
        }
    case "hello":
        p0 = 867
    }
    _ = p0
    _ = p1
    return (true && false)
}
func genFn1(p0 int8, p1 uint64) int64 {
    if false {
        v7 := GenS1{F0: -887, F1: 184}
        for fi8 := 0; fi8 < 4; fi8++ {
            if false {
                v9 := GenS1{F0: int32(p0), F1: 342}
                p0 += 100
                _ = v9
            } else {
                var v10 int = 9223372036854775807
                _ = v10
            }
            var v11 int8 = 1
            p1 = (p1 & p1)
            _ = fi8
            _ = v11
        }
        _ = v7
    } else {
        p0 = p0
    }
    v12 := GenS1{F0: 42, F1: 1}
    switch -2147483648 {
    case 9223372036854775807:
        var v13 uint8 = 128
        _ = v13
    }
    v14 := GenS1{F0: -1, F1: 2147483647}
    _ = p0
    _ = p1
    _ = v12
    _ = v14
    return 0
}

func main() {
    var v1 int16 = -850
    v2 := []string{}
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    fmt.Println(v1)
    var v15 int32 = 42
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float64 = 0.1
        _ = fi16
        _ = v17
    }
    v18 := v3[("" + ("" + "hello"))]
    v15 -= (v15 & (v15 >> ((1) & 63)))
    v15 -= 2147483647
    v15 = v15
    var v19 string = ("foo bar" + "hello")
    var v20 uint32 = v18
    if true {
        if false {
            v21 := GenS2{F0: v19, F1: v18, F2: -1}
            _ = v21
        } else {
            v1 += v1
            var v22 int32 = (v15 + v15)
            var v23 uint = 255
            _ = v22
            _ = v23
        }
    }
    v3[(((v19 + "") + (v19 + v19)) + "\n\t")] = (v18 - (^(v20)))
    var v24 float32 = 1e10
    v20 -= 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
