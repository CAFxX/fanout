package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
func genFn0(p0 uint16, p1 uint, p2 uint16) int32 {
    p1 = p1
    v4 := []int{1, 1, 0}
    switch "foo bar" {
    case "a":
        fmt.Println(p2, p2, p1)
    case "":
        v5 := GenS1{F0: 2, F1: "foo bar"}
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: 42, F1: "\n\t"}
        _ = v6
    }
    p0 -= (p0 - p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v2[("foo bar" + "hello")] = 2
    if (-2.25 == float32((3.14159 * -2.25))) {
        var v7 int8 = -56
        _ = v7
    } else {
        var v8 bool = (false && (255 < 32767))
        var v9 bool = true
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: genFn0(1, 0, 1), F1: 319, F2: 2147483648}
    if true {
        v11 := []int{2147483647}
        var v12 float32 = (-2.25 - 1e10)
        if false {
            v12 += v12
            var v13 int = 128
            fmt.Println(v12, v13)
            _ = v13
        }
        _ = v11
        _ = v12
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint16 = 259
            fmt.Println(v15, v15)
            var v16 int16 = 32767
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int32 = genFn0(42, 754, 2)
        v17 = (v17 - (v17 + (v17 ^ genFn0(2, 0, 2))))
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int = (fi18 & fi18)
        var v20 string = "foo bar"
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 float32 = -2.25
    _ = len(v1)
    v21 = float32((-2.25 - 0.0))
    var v22 int32 = 86
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
