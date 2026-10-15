package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool, p1 int32, p2 uint16) bool {
    for fi6 := 0; fi6 < 9; fi6++ {
        p0 = false
        p2 -= (^(p2))
        _ = fi6
    }
    v7 := GenS1{F0: 0, F1: "\n\t"}
    p1 = p1
    v8 := []int{-217, 32768}
    for fi9 := 0; fi9 < 14; fi9++ {
        if (p2 == p2) {
            p0 = (((p2 <= p2) || p0) && (!(!(1 != -128))))
            p0 = true
            v10 := GenS2{F0: p1}
            _ = v10
        } else {
            var v11 float32 = (float32(0.1) / (0.0 - (0.1 / 1.5)))
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 int64 = 1
    v2 := []string{"\n\t"}
    var v3 bool = false
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    var v12 int64 = 9223372036854775807
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint = 196
        _ = fi13
        _ = v14
    }
    v12 -= (v12 >> ((0) & 63))
    for fi15 := 0; fi15 < 8; fi15++ {
        if false {
            v1 = 420
            var v16 int64 = 0
            _ = v16
        }
        _ = fi15
    }
    v12 += (^(v12))
    fmt.Println(v1)
    v2 = append(v2, "")
    switch "foo bar" {
    case "a":
        var v17 uint = 4294967295
        _ = v17
    case "\n\t":
        v18 := GenS2{F0: -2147483648}
        _ = v18
    }
    if ((v1 - v12) <= v12) {
        _ = len(v4)
        var v19 uint8 = 15
        _ = v19
    } else {
        v20 := GenS1{F0: 881, F1: ""}
        var v21 int = -865
        v1 = -1
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if false {
            fmt.Println(v1, v12, v12)
        }
        var v23 int64 = v12
        var v24 int = (fi22 ^ (fi22 | -1))
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 string = "foo bar"
    v5[v25] = 628
    var v26 float64 = 0.0
    v27 := GenS1{F0: 2147483647, F1: v25}
    for fi28 := 0; fi28 < 15; fi28++ {
        v12 += (v1 + v12)
        fmt.Println(v12)
        v29 := GenS1{F0: fi28, F1: "hello"}
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
