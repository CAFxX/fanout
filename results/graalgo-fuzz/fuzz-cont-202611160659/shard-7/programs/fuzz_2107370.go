package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }

func main() {
    var v1 string = ("\n\t" + "foo bar")
    var v2 float32 = -2.25
    v3 := make(map[int]bool)
    var v4 float64 = (0.1 - float64(v2))
    var v5 int = 32767
    var v6 uint64 = 2
    var v7 float64 = (1.5 + (v4 * v4))
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int8 = 7
        if (1 == 4294967295) {
            var v10 uint64 = v6
            _ = v10
        } else {
            fmt.Println(v2, v9, v9)
        }
        v4 += float64(0.0)
        _ = fi8
        _ = v9
    }
    switch "" {
    case "\n\t":
        if (42 >= -1) {
            switch "foo bar" {
            case "hello":
                fmt.Println(v1)
            }
            switch (-(v5)) {
            case 32768:
                for fi11 := 0; fi11 < 16; fi11++ {
                    v5 = 575
                    var v12 int = (v5 ^ fi11)
                    _ = fi11
                    _ = v12
                }
            }
        } else {
            v13 := v3[9223372036854775807]
            v3[v5] = (4294967295 < 255)
            _ = v13
        }
    default:
        fmt.Println(v5, v2)
    }
    v14 := []int{(v5 >> ((4294967295) & 63))}
    for fi15 := 0; fi15 < 3; fi15++ {
        v2 -= float32((v4 - (v4 + (v4 / float64(1e10)))))
        if len(v14) > 0 {
            v14[0] = (-(v5))
        }
        _ = fi15
    }
    v1 = v1
    var v16 int32 = 42
    switch (^(v5)) {
    case -1:
        v14 = append(v14, v5)
    case 32768:
        fmt.Println(v16, v16)
    case 128:
        _ = len(v3)
    }
    v17 := GenS1{F0: -1, F1: true, F2: 2}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println(v1, v7)
    fmt.Println(v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
