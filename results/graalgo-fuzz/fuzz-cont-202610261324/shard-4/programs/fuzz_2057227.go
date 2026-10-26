package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    v3 := []int{-535, 42}
    var v4 float64 = float64(0.0)
    for fi5 := 0; fi5 < 16; fi5++ {
        v4 -= v4
        var v6 int32 = 1
        var v7 uint64 = 296
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int32 = 248
    switch 32767 {
    case -85:
        fmt.Println(v8, v4)
    case -19:
        v8 = (v8 + (v8 * v8))
    default:
        fmt.Println(v8, v4)
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(v8, fi9)
        fmt.Println(fi9, v8)
        _ = fi9
    }
    var v10 string = ("foo bar" + "a")
    _ = len(v2)
    if ((!false) || true) {
        for fi11 := 0; fi11 < 15; fi11++ {
            if (928 == 0) {
                v4 = 3.14159
                for fi12 := 0; fi12 < 16; fi12++ {
                    v1[-232] = true
                    v13 := GenS1{F0: 397}
                    fmt.Println(v4, fi11, fi12)
                    v14 := v1[(fi12 + fi12)]
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                if len(v2) > 0 {
                    v2[0] = 177
                }
            }
            if true {
                v4 = (0.0 / (v4 - float64((0.1 * (1.5 / -2.25)))))
            }
            fmt.Println(v8)
            _ = fi11
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 -= float64(0.1)
    switch "\n\t" {
    case "foo bar":
        var v15 bool = (((((v10 + v10) + (((v10 + "") + v10) + v10)) + v10) + (v10 + v10)) >= v10)
        _ = v15
    case "\n\t":
        v16 := v1[9223372036854775807]
        _ = v16
    case "hello":
        v4 += 3.14159
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
