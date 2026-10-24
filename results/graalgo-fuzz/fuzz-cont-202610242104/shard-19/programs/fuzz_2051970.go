package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 int; F2 uint32 }
func genFn0(p0 uint32, p1 uint16) int8 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{9223372036854775807}
        p1 = (p1 & 1)
        _ = fi5
        _ = v6
    }
    var v7 uint = 2
    v8 := []int{128}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -128
}

func main() {
    v1 := []int{}
    var v2 uint64 = 135
    var v3 string = ("foo bar" + "hello")
    var v4 int64 = -531
    switch 2147483647 {
    case -522:
        var v9 uint = 255
        _ = v9
    case -520:
        var v10 uint64 = (v2 | v2)
        _ = v10
    case 32767:
        var v11 uint64 = (v2 ^ v2)
        _ = v11
    default:
        var v12 int32 = 1
        _ = v12
    }
    var v13 int8 = genFn0(2, 255)
    v14 := GenS1{F0: -1, F1: 89, F2: ""}
    switch 2147483647 {
    case 1:
        var v15 uint64 = (v2 / ((v2) | 1))
        _ = v15
    case 32767:
        var v16 int16 = -32768
        _ = v16
    case 32768:
        var v17 uint = 624
        _ = v17
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v2 += 737
        _ = fi18
    }
    if (!true) {
        for fi19 := 0; fi19 < 5; fi19++ {
            fmt.Println(fi19, v4)
            _ = fi19
        }
        v2 += v2
    }
    fmt.Println(v4, v13)
    fmt.Println(v2, v13)
    v3 = v3
    fmt.Println(v13)
    v4 -= v4
    v1 = append(v1, 337)
    var v20 uint64 = v2
    v4 += (v4 + v4)
    v21 := GenS2{F0: 12, F1: 32768, F2: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
