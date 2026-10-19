package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint16) int16 {
    var v5 int16 = -1
    var v6 uint32 = 0
    v7 := GenS1{F0: 1, F1: p0}
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(v6, v6, v6)
        _ = fi8
    }
    switch "a" {
    case "\n\t":
        v9 := []int{32768, 9223372036854775807, 9223372036854775807}
        _ = v9
    case "":
        fmt.Println(v6, v5, p0)
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v5
}
func genFn1(p0 uint8) bool {
    switch 127 {
    case 0:
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
    case 253:
        var v12 string = "hello"
        _ = v12
    case 32768:
        p0 = (p0 % ((p0) | 1))
    default:
        var v13 int16 = 42
        _ = v13
    }
    var v14 int16 = -86
    if false {
        v15 := GenS1{F0: -742, F1: 73}
        if (302 == -1) {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 float64 = (0.1 - (1e10 * 1.5))
                _ = fi16
                _ = v17
            }
            p0 += 1
        }
        _ = v15
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        switch fi18 {
        case 32767:
            var v19 int8 = -1
            _ = v19
        }
        _ = fi18
    }
    _ = p0
    _ = v14
    return false
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float32 = 1.5
    var v3 float64 = (0.0 - 0.0)
    var v4 float64 = v3
    v20 := []int{-254, 822, -162}
    switch -838 {
    case -9223372036854775808:
        fmt.Println(v1, v1, v3)
    case -620:
        var v21 string = ""
        _ = v21
    }
    if false {
        var v22 int = 1
        var v23 uint16 = 2
        _ = v22
        _ = v23
    }
    var v24 uint64 = v1
    var v25 uint16 = 255
    for fi26 := 0; fi26 < 8; fi26++ {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = len(v20)
        v3 -= v4
        _ = fi26
    }
    v27 := GenS1{F0: 42, F1: (v25 ^ v25)}
    fmt.Println(v3, v3, v3)
    fmt.Println(v25, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
}
