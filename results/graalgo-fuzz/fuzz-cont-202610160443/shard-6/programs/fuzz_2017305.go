package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
func genFn0(p0 int16, p1 bool) int32 {
    var v4 int16 = (p0 & -298)
    switch "a" {
    case "a":
        v5 := []int{32768, 2147483647, 917, 2147483647}
        _ = v5
    case "":
        v6 := []int{}
        _ = v6
    case "foo bar":
        if (!false) {
            v7 := []int{9223372036854775807, 128, int(p0), 4294967295}
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: (711 == 42), F1: -128}
                if (false && ((1 != -1) && p1)) {
                    if ((-128 >= -19) || (!p1)) {
                        fmt.Println(v4, fi8, v4)
                        var v10 int16 = (p0 >> ((1) & 63))
                        _ = v10
                    }
                    if len(v7) > 0 {
                        v7[0] = fi8
                    }
                    for fi11 := 0; fi11 < 16; fi11++ {
                        var v12 int16 = v4
                        var v13 string = "a"
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                fmt.Println(fi8)
                _ = fi8
                _ = v9
            }
            _ = v7
        }
    }
    v14 := []int{128}
    _ = p0
    _ = p1
    _ = v14
    _ = v4
    return 0
}

func main() {
    var v1 uint32 = 59
    var v2 uint32 = v1
    v3 := make(map[int]uint32)
    delete(v3, 127)
    fmt.Println(v1, v1)
    _ = len(v3)
    if true {
        for fi15 := 0; fi15 < 15; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := GenS1{F0: (!false), F1: -98}
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    }
    v18 := []int{-2147483648}
    if true {
        v19 := v3[32768]
        var v20 uint = 65535
        v19 += (v2 / ((v2) | 1))
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if (18446744073709551615 > 1) {
            v1 += (v2 / ((v1) | 1))
        } else {
            var v22 string = "foo bar"
            var v23 int16 = 520
            v1 += (^(v2))
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            _ = len(v3)
            _ = fi24
        }
        fmt.Println(v2, v2)
        _ = fi21
    }
    switch 127 {
    case -1:
        var v25 int8 = -1
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
}
