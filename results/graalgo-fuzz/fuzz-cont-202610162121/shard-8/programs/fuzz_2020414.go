package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int; F2 uint }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int8; F1 bool }
func genFn0() uint64 {
    switch 0 {
    case 32768:
        var v6 bool = (false && false)
        _ = v6
    case -686:
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 string = "a"
            fmt.Println(v8)
            _ = fi7
            _ = v8
        }
    case 9223372036854775807:
        v9 := []int{438, -388}
        _ = v9
    default:
        var v10 string = "\n\t"
        _ = v10
    }
    var v11 uint = 349
    fmt.Println(v11)
    _ = v11
    return uint64(v11)
}

func main() {
    var v1 uint8 = 255
    var v2 float64 = 0.1
    v3 := make(map[int]int)
    var v4 int8 = -1
    var v5 bool = (0 < 1)
    v12 := []int{1}
    fmt.Println(v4, v5)
    var v13 int64 = 1
    v14 := GenS1{F0: (-1 < 2147483647)}
    for fi15 := 0; fi15 < 15; fi15++ {
        v5 = (42 >= 255)
        _ = fi15
    }
    switch -2147483648 {
    case 32767:
        v16 := GenS1{F0: false}
        _ = v16
    }
    var v17 bool = ("" < "a")
    switch "a" {
    case "hello":
        if ((!false) || ((v4 % ((v4) | 1)) >= -128)) {
            v18 := GenS1{F0: (v17 && v5)}
            _ = v18
        } else {
            fmt.Println(v2, v2)
        }
    default:
        var v19 uint8 = (v1 ^ v1)
        _ = v19
    }
    v20 := GenS2{F0: v1, F1: -186, F2: 65535}
    for fi21 := 0; fi21 < 7; fi21++ {
        v3[fi21] = fi21
        v2 = v2
        _ = fi21
    }
    switch "\n\t" {
    case "\n\t":
        v22 := GenS3{F0: (-(v13)), F1: 994}
        _ = v22
    default:
        v13 -= v13
    }
    delete(v3, 489)
    var v23 int16 = -32768
    var v24 uint64 = 763
    v25 := GenS4{F0: (v4 - int8(v1)), F1: v17}
    v26 := GenS1{F0: (true && true)}
    switch -438 {
    case 127:
        var v27 int32 = 2147483647
        _ = v27
    case 128:
        v2 = -2.25
    case 32767:
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 int64 = v13
            fmt.Println(v17, v5)
            _ = fi28
            _ = v29
        }
    default:
        if (0.0 >= 3.14159) {
            v2 -= v2
            switch (("" + "foo bar") + ("foo bar" + "")) {
            case "a":
                v30 := GenS3{F0: int64(v1), F1: 1}
                _ = v30
            case "\n\t":
                var v31 int8 = v4
                _ = v31
            }
            v32 := GenS3{F0: v13, F1: -125}
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
