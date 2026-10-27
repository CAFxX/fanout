package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0(p0 uint8, p1 bool) uint64 {
    switch ("a" + "") {
    case "":
        var v4 int8 = -1
        _ = v4
    case "foo bar":
        v5 := []int{-155, 42, 32767}
        _ = v5
    case "\n\t":
        var v6 uint64 = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint8 = p0
        _ = fi7
        _ = v8
    }
    p1 = true
    switch "foo bar" {
    case "":
        for fi9 := 0; fi9 < 2; fi9++ {
            p0 = 1
            _ = fi9
        }
    case "a":
        p0 -= 1
    case "foo bar":
        var v10 float64 = 3.14159
        _ = v10
    default:
        fmt.Println(p0, p0)
    }
    var v11 int16 = -32768
    _ = p0
    _ = p1
    _ = v11
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    v3 := make(map[int]string)
    v12 := []int{0, 4294967295, 2147483647, -9223372036854775808}
    var v13 int64 = 220
    for fi14 := 0; fi14 < 20; fi14++ {
        if (64 < 763) {
            v15 := GenS1{F0: 860, F1: -123}
            v1[fi14] = 234
            _ = v15
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 bool = false
                v12 = append(v12, fi16)
                _ = fi16
                _ = v17
            }
            var v18 uint32 = 42
            _ = v18
        }
        var v19 int64 = v13
        var v20 bool = false
        _ = fi14
        _ = v19
        _ = v20
    }
    var v21 uint = 543
    var v22 uint32 = 2147483647
    var v23 float32 = (1.5 / 3.14159)
    v24 := v1[1]
    v2 = (v2 - (v2 & (v2 ^ (v2 / (((v2 & v2)) | 1)))))
    v25 := GenS1{F0: v21, F1: 1}
    fmt.Println(v23)
    if len(v12) > 0 {
        v12[0] = 0
    }
    var v26 string = "foo bar"
    fmt.Println(v26)
    v21 -= v21
    v13 = v13
    for fi27 := 0; fi27 < 6; fi27++ {
        v22 = 2
        if true {
            var v28 int32 = (v2 + 42)
            fmt.Println(v21, fi27)
            _ = v28
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
