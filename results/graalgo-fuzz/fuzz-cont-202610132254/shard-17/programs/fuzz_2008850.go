package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int8; F2 int16 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 int32 = 96
    var v2 int16 = 654
    var v3 uint16 = 1
    var v4 int64 = -9223372036854775808
    v5 := make(map[string]bool)
    v3 -= v3
    if true {
        fmt.Println(v4, v4, v2)
    }
    v6 := GenS1{F0: -39, F1: (v3 * (v3 & 209))}
    delete(v5, "a")
    v7 := GenS2{F0: v3, F1: -32, F2: 0}
    v2 -= v2
    if false {
        switch "a" {
        case "\n\t":
            var v8 uint64 = 0
            _ = v8
        case "a":
            v9 := []int{}
            _ = v9
        case "foo bar":
            v10 := GenS3{F0: 0, F1: 792}
            _ = v10
        default:
            v11 := GenS4{F0: -1}
            _ = v11
        }
    } else {
        v12 := GenS4{F0: 9}
        v13 := []int{0, 128, -138}
        var v14 int8 = 27
        _ = v12
        _ = v13
        _ = v14
    }
    v2 = v2
    switch 209 {
    case 4294967295:
        v15 := []int{int(v3), 4294967295, 1, -565}
        _ = v15
    case 42:
        var v16 uint = 18446744073709551615
        _ = v16
    default:
        var v17 float64 = 1.5
        _ = v17
    }
    v5["hello"] = false
    fmt.Println(v2, v2, v2)
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS2{F0: v3, F1: 0, F2: -45}
        var v20 uint8 = 2
        v5[("\n\t" + "\n\t")] = false
        _ = fi18
        _ = v19
        _ = v20
    }
    v3 -= 568
    for fi21 := 0; fi21 < 12; fi21++ {
        fmt.Println(v2, fi21)
        fmt.Println(v1)
        _ = fi21
    }
    if true {
        if true {
            v22 := []int{}
            v1 = -1
            _ = v22
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            v1 = 0
            _ = fi23
        }
    }
    var v24 uint = 216
    switch "a" {
    case "foo bar":
        var v25 uint = 0
        _ = v25
    case "a":
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 uint64 = 4294967295
            if true {
                v28 := GenS4{F0: 100}
                _ = v28
            } else {
                v29 := GenS5{F0: "\n\t", F1: 1, F2: (^(v2))}
                _ = v29
            }
            _ = fi26
            _ = v27
        }
    default:
        v30 := GenS6{F0: (-(v4))}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
