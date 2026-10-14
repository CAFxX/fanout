package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }

func main() {
    var v1 uint = 625
    var v2 float64 = (0.1 + 1e10)
    v3 := make(map[int]bool)
    var v4 int8 = 127
    v5 := []int{int(v1), 239}
    var v6 uint64 = 2
    var v7 int16 = 387
    switch 1 {
    case 32767:
        v8 := GenS1{F0: 32768, F1: -473}
        _ = v8
    case -918:
        for fi9 := 0; fi9 < 17; fi9++ {
            v7 -= v7
            _ = fi9
        }
    default:
        var v10 uint32 = 65535
        _ = v10
    }
    var v11 uint = 267
    if (!(42 == 42)) {
        v12 := GenS1{F0: 0, F1: -9223372036854775808}
        if ("hello" >= (("" + "foo bar") + "\n\t")) {
            var v13 float32 = ((3.14159 / 1.0) * -2.25)
            for fi14 := 0; fi14 < 2; fi14++ {
                fmt.Println(v1, v13, v2)
                v5 = append(v5, (-(fi14)))
                _ = fi14
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                _ = len(v3)
                _ = fi15
            }
            _ = v13
        }
        v2 = float64(0.1)
        _ = v12
    } else {
        var v16 uint = v1
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v5 = append(v5, (fi17 | fi17))
        fmt.Println(v7, v2, v2)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
