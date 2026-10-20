package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0() int {
    v5 := []int{-1, 9223372036854775807, 127, -228}
    var v6 uint32 = 2
    v5 = append(v5, 9223372036854775807)
    v7 := GenS1{F0: (-1 != -98), F1: 1, F2: 0}
    _ = v5
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 string, p1 uint, p2 bool) uint {
    var v8 int16 = 32767
    var v9 uint = p1
    switch (p0 + p0) {
    case "":
        var v10 int16 = v8
        _ = v10
    }
    fmt.Println(v8)
    if (0 != 879) {
        v11 := GenS2{F0: 687}
        p2 = (!p2)
        var v12 uint64 = 2
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p1 | 4294967295)
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{1}
    var v3 uint8 = 154
    v4 := []int{4294967295}
    switch "foo bar" {
    case "":
        switch ("" + "") {
        case "":
            v13 := GenS2{F0: 1}
            _ = v13
        case "a":
            if (1e10 != 0.1) {
                v1["\n\t"] = (-1 > 40)
                v14 := v1["hello"]
                v15 := GenS1{F0: ((^(v3)) <= v3), F1: -187, F2: 1}
                _ = v14
                _ = v15
            }
        case "foo bar":
            v16 := GenS1{F0: (!false), F1: 822, F2: 165}
            _ = v16
        default:
            _ = len(v1)
        }
    case "hello":
        var v17 uint = 4294967295
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 int8 = 1
            _ = fi18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 17; fi20++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                v3 = v3
                _ = fi21
            }
            _ = fi20
        }
    }
    v3 = 70
    v22 := GenS1{F0: true, F1: 1, F2: 65535}
    v3 += (v3 * 127)
    var v23 uint32 = 2
    if (794 < 0) {
        fmt.Println(v3)
    }
    _ = len(v4)
    var v24 uint8 = v3
    v3 += v24
    fmt.Println(v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
