package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint64) uint {
    v4 := []int{2147483647, -9223372036854775808}
    p0 -= (p0 & p0)
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            p0 = 42
            _ = fi6
        }
        _ = fi5
    }
    v4 = append(v4, -9223372036854775808)
    p0 -= p0
    _ = p0
    _ = v4
    return 870
}

func main() {
    var v1 float32 = (1.5 + (-2.25 / 1.0))
    v2 := make(map[string]bool)
    var v3 int16 = -1
    switch "foo bar" {
    case "foo bar":
        for fi7 := 0; fi7 < 12; fi7++ {
            v2["a"] = false
            _ = fi7
        }
    case "":
        v8 := GenS1{F0: 27, F1: (1.5 >= 0.1)}
        _ = v8
    case "\n\t":
        if (("a" + (("a" + "\n\t") + "a")) <= "") {
            v9 := []int{612}
            var v10 uint = 525
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 uint8 = 127
                v13 := v2["foo bar"]
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
            _ = v9
        }
    default:
        v1 += v1
    }
    switch "foo bar" {
    case "\n\t":
        var v14 int32 = -1
        _ = v14
    case "hello":
        switch "a" {
        case "a":
            var v15 bool = true
            _ = v15
        }
    case "foo bar":
        v1 -= v1
    }
    if true {
        var v16 uint = genFn0(0)
        v17 := []int{2147483647, -832, 23, -115}
        var v18 float32 = float32((float64((3.14159 / v1)) * ((3.14159 - 1.5) / 1.0)))
        _ = v16
        _ = v17
        _ = v18
    } else {
        v19 := GenS2{F0: 42}
        _ = v19
    }
    if true {
        delete(v2, "\n\t")
        v20 := []int{}
        _ = v20
    }
    var v21 uint16 = 114
    var v22 uint32 = 920
    v22 -= v22
    if (-543 > 4294967295) {
        var v23 uint16 = 2
        _ = v23
    }
    v24 := []int{155}
    var v25 uint16 = (v21 << ((1) & 63))
    var v26 int32 = -526
    v21 = v25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
