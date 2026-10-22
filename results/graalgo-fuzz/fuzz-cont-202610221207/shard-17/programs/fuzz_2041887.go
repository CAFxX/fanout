package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 string }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.0
    var v3 uint = 772
    v4 := make(map[string]int)
    var v5 int = -1
    fmt.Println(v2)
    if ((false || true) || false) {
        if (18446744073709551615 <= 4294967295) {
            var v6 bool = (127 == -128)
            var v7 int = -198
            _ = v6
            _ = v7
        }
        var v8 string = "a"
        _ = v8
    } else {
        if false {
            var v9 float64 = (3.14159 / 1.5)
            _ = v9
        } else {
            if (2 != 1) {
                v4[(("foo bar" + "foo bar") + "hello")] = (v5 - 32767)
            } else {
                v2 += (float32(3.14159) + 1.5)
            }
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(v5, v3)
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := v1[fi10]
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    fmt.Println(v5)
    if (!true) {
        var v13 uint = (v3 << ((2) & 63))
        _ = v13
    } else {
        v14 := GenS1{F0: 64, F1: -117, F2: ("a" + "a")}
        fmt.Println(v3)
        fmt.Println(v2)
        _ = v14
    }
    var v15 uint = v3
    var v16 int32 = -149
    var v17 int8 = 42
    switch ("a" + ("" + "foo bar")) {
    case "hello":
        v17 -= (v17 ^ (^(v17)))
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 int16 = 125
            v21 := GenS1{F0: v17, F1: v17, F2: "\n\t"}
            v22 := []int{v5, (^(v5)), v5, (fi19 ^ fi19)}
            v3 = v15
            fmt.Println(fi18, v16)
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            fmt.Println(v17, v16)
            v17 -= 42
            _ = fi23
        }
        _ = len(v1)
        _ = fi18
    }
    fmt.Println(v16)
    v3 = v3
    _ = len(v4)
    v24 := []int{v5}
    switch "a" {
    case "foo bar":
        v25 := v1[(v5 / ((v5) | 1))]
        _ = v25
    case "a":
        var v26 int = (v5 ^ 108)
        _ = v26
    default:
        v3 = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
