package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (float64(float32((0.0 / 1.0))) - 1.5)
    var v3 float64 = (v2 - v2)
    if true {
        v2 += ((v2 * float64((0.1 - -2.25))) / v3)
        var v4 uint8 = 15
        _ = v4
    } else {
        if true {
            v1["\n\t"] = 186
            var v5 uint16 = 1
            _ = v5
        }
        switch 4294967295 {
        case 9223372036854775807:
            v6 := []int{478, 128}
            _ = v6
        case -756:
            var v7 string = "a"
            _ = v7
        }
    }
    switch "hello" {
    case "a":
        v8 := []int{861, 583, 579}
        _ = v8
    case "hello":
        v1[("" + "a")] = 1
    case "":
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint8 = 1
            _ = fi9
            _ = v10
        }
    default:
        fmt.Println(v2, v3)
    }
    fmt.Println(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        switch "\n\t" {
        case "\n\t":
            var v12 uint32 = 401
            _ = v12
        case "":
            v3 = (0.1 + 0.0)
        }
        delete(v1, "\n\t")
        v3 -= (v3 / v3)
        _ = fi11
    }
    var v13 int64 = 1
    v14 := []int{435, 343, 1, 42}
    if (v3 <= 0.0) {
        var v15 float64 = (((float64(1.5) + (((v3 / (float64(1.5) * (v2 * (v2 * v2)))) * (v2 - (1.5 + -2.25))) + float64(0.1))) + (1e10 / v2)) / 0.1)
        var v16 int64 = 0
        v1[("" + "\n\t")] = 31
        _ = v15
        _ = v16
    } else {
        var v17 int = 42
        _ = v17
    }
    v18 := GenS1{F0: 1, F1: -9223372036854775808, F2: (v13 - v13)}
    var v19 uint64 = 168
    fmt.Println(v3, v13, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
