package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 string }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := []string{}
    var v4 uint = 255
    var v5 float64 = 1e10
    var v6 bool = false
    v4 += v4
    for fi7 := 0; fi7 < 19; fi7++ {
        switch "" {
        case "\n\t":
            var v8 int8 = -73
            _ = v8
        case "a":
            switch 32767 {
            case -2147483648:
                var v9 int16 = 618
                _ = v9
            default:
                var v10 int32 = -1
                _ = v10
            }
        }
        v11 := GenS1{F0: 0}
        _ = fi7
        _ = v11
    }
    switch "a" {
    case "\n\t":
        if v6 {
            switch 575 {
            case 42:
                v5 = float64(-2.25)
            case 128:
                v6 = v6
            }
        }
    }
    _ = len(v3)
    if ((v5 - 1.5) < (1.5 * v5)) {
        v6 = (!true)
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int32 = -813
        switch fi12 {
        case -1:
            fmt.Println(v5)
        case 0:
            var v14 float64 = 0.1
            _ = v14
        }
        switch -631 {
        case 659:
            v15 := GenS1{F0: -120}
            _ = v15
        case -198:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 42:
            for fi16 := 0; fi16 < 6; fi16++ {
                _ = len(v3)
                _ = fi16
            }
        default:
            v17 := GenS2{F0: 1, F1: "foo bar"}
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v18 uint64 = 0
    var v19 float32 = 0.0
    var v20 int64 = -1
    fmt.Println(v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
