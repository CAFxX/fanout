package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0() uint64 {
    var v5 uint64 = 2
    var v6 uint = 567
    v7 := GenS1{F0: v5, F1: 0}
    v8 := []int{int(v6)}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 string = "\n\t"
    var v4 uint64 = 4294967295
    var v9 float64 = float64(3.14159)
    switch ((("a" + "foo bar") + v3) + ("a" + ("foo bar" + (v3 + (v3 + v3))))) {
    case "":
        var v10 string = (("\n\t" + (v3 + v3)) + v3)
        _ = v10
    case "foo bar":
        switch 289 {
        case 1:
            fmt.Println(v4, v4, v9)
        default:
            fmt.Println(v4, v4)
        }
    default:
        var v11 float32 = 0.0
        _ = v11
    }
    switch ("a" + v3) {
    case "foo bar":
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int64 = int64(v4)
            if false {
                v14 := []int{128, 24, 32768}
                _ = v14
            } else {
                var v15 float64 = float64(3.14159)
                v16 := []int{}
                var v17 uint64 = 384
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi12
            _ = v13
        }
    }
    v3 = ((v3 + "foo bar") + v3)
    v18 := GenS1{F0: (^(v4)), F1: -1}
    v19 := []int{-743, 128}
    if true {
        for fi20 := 0; fi20 < 1; fi20++ {
            if len(v19) > 0 {
                v19[0] = (fi20 ^ (fi20 % ((0) | 1)))
            }
            var v21 uint64 = genFn0()
            _ = fi20
            _ = v21
        }
    }
    if (!false) {
        v9 += (v9 / (0.0 - v9))
        fmt.Println(v4, v9, v4)
        v3 = "foo bar"
    } else {
        var v22 float32 = 3.14159
        _ = v22
    }
    v3 = (((((v3 + v3) + v3) + v3) + v3) + (v3 + ""))
    v23 := GenS1{F0: (^(v4)), F1: 720}
    for fi24 := 0; fi24 < 11; fi24++ {
        v4 = 491
        v25 := v2[(v3 + "\n\t")]
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
