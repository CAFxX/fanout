package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int8 = -128
    var v2 uint64 = 955
    var v3 int8 = v1
    var v4 uint64 = (v2 * 170)
    v5 := []string{"", ""}
    var v6 int64 = -1
    var v7 int = 128
    v8 := GenS1{F0: 2, F1: false, F2: ("foo bar" + (("foo bar" + "foo bar") + "\n\t"))}
    var v9 bool = true
    var v10 uint64 = (^(v2))
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint32 = 269
        _ = fi11
        _ = v12
    }
    switch (v7 * (v7 % ((v7) | 1))) {
    case 165:
        var v13 uint = 255
        _ = v13
    case 32768:
        v6 = (v6 | (v6 + (v6 + (v6 / ((v6) | 1)))))
    }
    if (!(!(42 <= 2147483647))) {
        var v14 int8 = 100
        _ = v14
    } else {
        fmt.Println(v4)
    }
    var v15 string = "a"
    switch v7 {
    case 2147483647:
        var v16 bool = false
        _ = v16
    case 128:
        if ((!v9) && v9) {
            v10 += v10
        } else {
            v17 := GenS2{F0: 71, F1: 2147483648}
            v9 = ((-2.25 * (0.1 - (1e10 + -2.25))) != 0.0)
            v18 := GenS3{F0: 481}
            _ = v17
            _ = v18
        }
    case -402:
        var v19 float64 = (1e10 / (3.14159 * -2.25))
        _ = v19
    default:
        var v20 uint = 1
        _ = v20
    }
    v3 -= v1
    var v21 int64 = v6
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 int32 = 42
        var v24 int8 = (v1 >> ((0) & 63))
        _ = fi22
        _ = v23
        _ = v24
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 -= 42
    v7 -= 32767
    fmt.Println(v21, v3, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
