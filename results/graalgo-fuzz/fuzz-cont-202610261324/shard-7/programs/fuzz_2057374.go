package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 127
    var v3 uint8 = 42
    var v4 float64 = float64(1.5)
    v5 := make(map[int]string)
    v6 := make(map[string]int)
    switch 4294967295 {
    case 0:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int8 = 29
            switch fi7 {
            case 127:
                v9 := GenS1{F0: 1, F1: -1}
                _ = v9
            case -891:
                v4 = v4
            default:
                var v10 string = "foo bar"
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    default:
        var v11 string = ("" + "\n\t")
        _ = v11
    }
    var v12 bool = true
    switch -1 {
    case -316:
        v12 = (v4 > 1.5)
    default:
        v13 := v5[127]
        _ = v13
    }
    switch 1 {
    case -562:
        var v14 uint8 = 251
        _ = v14
    case -1:
        switch "\n\t" {
        case "a":
            var v15 int64 = -621
            _ = v15
        }
    default:
        v16 := []int{909, 223, 0, -573}
        _ = v16
    }
    if v12 {
        switch 9223372036854775807 {
        case 128:
            v17 := GenS1{F0: -54, F1: v2}
            _ = v17
        case 42:
            var v18 float64 = 1.5
            _ = v18
        default:
            v19 := GenS1{F0: -20, F1: (v2 | (v2 * v2))}
            _ = v19
        }
        v12 = v12
        fmt.Println(v4)
    } else {
        fmt.Println(v12, v2)
        if false {
            v20 := GenS2{F0: 32767, F1: -2147483648, F2: v12}
            _ = v20
        }
    }
    fmt.Println(v2, v12)
    switch "a" {
    case "foo bar":
        for fi21 := 0; fi21 < 3; fi21++ {
            fmt.Println(v2, fi21, v4)
            if (v2 < v2) {
                var v22 float32 = float32((((((v4 / v4) * 0.0) + float64(0.0)) - (-2.25 - (v4 * ((v4 / v4) - 3.14159)))) + (0.1 / (float64(-2.25) * (v4 / v4)))))
                fmt.Println(v4, v4)
                v23 := []int{(fi21 / ((-9223372036854775808) | 1)), fi21, (fi21 / ((fi21) | 1)), fi21}
                _ = v22
                _ = v23
            } else {
                fmt.Println(v3, v3)
            }
            _ = fi21
        }
    case "":
        switch -274 {
        case -1:
            for fi24 := 0; fi24 < 11; fi24++ {
                var v25 string = "a"
                v26 := []int{}
                v27 := GenS2{F0: 32767, F1: fi24, F2: false}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        }
    case "a":
        var v28 uint32 = 352
        _ = v28
    }
    var v29 uint32 = 143
    var v30 float64 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
