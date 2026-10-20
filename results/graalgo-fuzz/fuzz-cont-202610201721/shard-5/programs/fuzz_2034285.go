package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint) int64 {
    v4 := GenS1{F0: 63}
    var v5 int8 = 42
    var v6 float64 = float64((1e10 + 1e10))
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -93
}
func genFn1() string {
    var v7 int8 = 42
    switch 32767 {
    case 32767:
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 bool = (fi8 != 2147483647)
            _ = fi8
            _ = v9
        }
    case 128:
        v10 := []int{-2147483648}
        _ = v10
    case 9223372036854775807:
        v7 = -113
    default:
        v7 = (^(v7))
    }
    v11 := []int{401, 1}
    var v12 uint64 = 609
    _ = v11
    _ = v12
    _ = v7
    return "a"
}

func main() {
    var v1 string = ""
    var v2 uint32 = 4294967295
    var v3 string = v1
    v13 := GenS1{F0: 42}
    fmt.Println(v1, v3, v2)
    v3 = v3
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(fi14)
        _ = fi14
    }
    if true {
        var v15 string = v3
        var v16 uint64 = 277
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 2}
    switch v3 {
    case "foo bar":
        v18 := GenS1{F0: 4294967295}
        _ = v18
    case "a":
        var v19 float64 = (0.1 * (-2.25 + (0.0 / -2.25)))
        _ = v19
    case "\n\t":
        var v20 float32 = (((-2.25 + 0.1) / 1e10) + (1e10 - 3.14159))
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 float64 = 1e10
        v3 = v3
        if false {
            v3 = ((v1 + ("a" + v3)) + "a")
        }
        _ = fi21
        _ = v22
    }
    v2 += (v2 << ((42) & 63))
    v23 := []int{127}
    var v24 uint16 = 779
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
