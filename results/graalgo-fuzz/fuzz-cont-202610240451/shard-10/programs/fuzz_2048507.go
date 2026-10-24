package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0(p0 uint16) int {
    p0 = (^(p0))
    v4 := []int{32768, 2147483647, 814}
    _ = p0
    _ = v4
    return 127
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]bool)
    var v3 int8 = 43
    v5 := GenS1{F0: 2147483647}
    switch "foo bar" {
    case "foo bar":
        var v6 float32 = -2.25
        _ = v6
    default:
        for fi7 := 0; fi7 < 1; fi7++ {
            v3 -= (v3 >> ((4294967295) & 63))
            _ = fi7
        }
    }
    var v8 uint32 = 111
    v3 = 61
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float64 = ((((1.5 - 0.0) + -2.25) + (1e10 + 1.5)) / 1.5)
        _ = fi9
        _ = v10
    }
    v1 -= 79
    var v11 uint32 = (^(v8))
    v12 := GenS2{F0: 128}
    if false {
        var v13 uint64 = 750
        v2["\n\t"] = false
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint64 = v13
            v16 := v2["hello"]
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    v17 := []int{0}
    v11 -= (v8 | v8)
    v17 = append(v17, 280)
    switch 929 {
    case 647:
        v8 = (v8 >> ((42) & 63))
    default:
        var v18 int16 = -661
        _ = v18
    }
    v3 = 23
    v19 := v2["\n\t"]
    switch 128 {
    case -36:
        fmt.Println(v3)
    default:
        switch 127 {
        case 796:
            var v20 uint32 = 4294967295
            _ = v20
        default:
            v21 := GenS3{F0: 4294967295, F1: v3}
            _ = v21
        }
    }
    var v22 int8 = (-(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v5
    _ = v8
}
