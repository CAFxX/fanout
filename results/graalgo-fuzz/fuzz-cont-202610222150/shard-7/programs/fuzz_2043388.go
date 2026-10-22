package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 int32, p1 uint, p2 uint32) int {
    var v5 int64 = -9223372036854775808
    var v6 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 bool, p1 uint16) uint8 {
    fmt.Println(p1)
    var v7 bool = ((true && (32767 >= 867)) && true)
    v8 := []int{1, -502, -1000, 498}
    for fi9 := 0; fi9 < 17; fi9++ {
        v7 = p0
        fmt.Println(p1, v7)
        p0 = (42 <= 86)
        _ = fi9
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{483, 42}
    v3 := make(map[string]bool)
    var v4 int64 = -1
    var v10 int = -467
    v11 := GenS1{F0: 44, F1: 389}
    var v12 int = (v10 % ((genFn0(2147483647, 2, 4294967295)) | 1))
    v4 = 0
    v4 += (-(v4))
    if (2147483648 != 4294967295) {
        switch -9223372036854775808 {
        case -1:
            v13 := GenS1{F0: 255, F1: 835}
            _ = v13
        default:
            var v14 string = "a"
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            v10 = genFn0(2147483647, 1, 2147483648)
            _ = fi15
        }
        v1 = (-(v4))
    }
    v16 := GenS2{F0: 32767, F1: (417 != 0), F2: 89}
    if true {
        var v17 uint = 4294967295
        _ = v17
    }
    var v18 int16 = -1
    var v19 uint = 42
    switch "a" {
    case "hello":
        var v20 uint64 = 516
        _ = v20
    case "foo bar":
        v21 := GenS2{F0: 1, F1: true, F2: -128}
        _ = v21
    case "a":
        var v22 uint64 = 153
        _ = v22
    }
    if (!false) {
        v3[""] = (80 < 145)
        v3[""] = true
    } else {
        v18 -= v18
        fmt.Println(v10, v19, v12)
        for fi23 := 0; fi23 < 7; fi23++ {
            v10 += genFn0(1, (^(v19)), 4294967295)
            _ = fi23
        }
    }
    v24 := GenS3{F0: 64, F1: (v18 < v18)}
    v18 = v18
    v25 := GenS1{F0: 967, F1: 2}
    if false {
        v3[""] = true
    } else {
        for fi26 := 0; fi26 < 20; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 string = "a"
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi27
                _ = v28
            }
            v18 = v18
            var v29 bool = true
            _ = fi26
            _ = v29
        }
        v18 = 864
        v1 += v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
