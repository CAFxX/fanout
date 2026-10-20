package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []int{1, 9223372036854775807, 2147483647}
    var v2 float64 = (0.1 - float64(((float32(float64((0.1 - 1e10))) + -2.25) - (0.1 - (-2.25 + (0.1 * 1e10))))))
    var v3 int16 = 0
    if false {
        v4 := GenS1{F0: 100}
        v5 := GenS1{F0: 0}
        switch -1 {
        case -1:
            _ = len(v1)
        }
        _ = v4
        _ = v5
    } else {
        v2 = (1.5 + (float64((1.5 / 0.1)) + v2))
        var v6 int32 = -820
        _ = v6
    }
    var v7 string = ("a" + "")
    var v8 string = "foo bar"
    switch 467 {
    case 0:
        v7 = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS2{F0: 2, F1: v3}
        fmt.Println(v3)
        switch (v7 + v8) {
        case "":
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 uint = 18446744073709551615
                _ = fi11
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    v1 = append(v1, 2147483647)
    v13 := GenS2{F0: 0, F1: -302}
    v1 = append(v1, 9223372036854775807)
    if ("a" > (v7 + (v8 + ""))) {
        v3 += v3
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint64 = 289
            _ = fi14
            _ = v15
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
