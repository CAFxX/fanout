package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 int32, p1 int32) int32 {
    var v5 int32 = (-(p1))
    v6 := GenS1{F0: 1, F1: 0, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 int, p1 uint64, p2 int64) bool {
    switch "\n\t" {
    case "a":
        var v7 int16 = 0
        _ = v7
    default:
        for fi8 := 0; fi8 < 16; fi8++ {
            fmt.Println(p1, p0, fi8)
            v9 := GenS1{F0: 521, F1: 1, F2: 1}
            _ = fi8
            _ = v9
        }
    }
    var v10 int16 = -32768
    switch "hello" {
    case "hello":
        v10 = (v10 - v10)
    case "\n\t":
        var v11 uint = 1
        _ = v11
    default:
        var v12 uint32 = 849
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 string = ("hello" + (("foo bar" + "") + "a"))
    var v4 int8 = 100
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint = 255
        v15 := []int{729}
        _ = fi13
        _ = v14
        _ = v15
    }
    v3 = ""
    var v16 int = 2147483647
    var v17 int32 = 0
    fmt.Println(v4)
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS2{F0: (-(v17)), F1: -2147483648}
        _ = fi18
        _ = v19
    }
    fmt.Println(v4, v17, v17)
    v4 = 1
    _ = len(v2)
    fmt.Println(v3, v16)
    if (((0.1 / 1.0) <= ((0.1 + float64(3.14159)) * float64(3.14159))) && true) {
        if (((^(v16)) >= v16) && (!(0 >= 42))) {
            var v20 uint32 = 2
            _ = v20
        }
        switch v3 {
        case "hello":
            v21 := []int{}
            _ = v21
        case "":
            v1[v3] = 42
        default:
            v22 := GenS2{F0: v17, F1: v17}
            _ = v22
        }
        var v23 int = 995
        _ = v23
    }
    fmt.Println(v4, v16, v17)
    fmt.Println(v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
