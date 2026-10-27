package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int32; F2 bool }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 float64 = 3.14159
    var v5 int16 = 816
    if (197 == 65535) {
        switch 0 {
        case 849:
            if (true || (!true)) {
                var v6 int8 = 42
                fmt.Println(v4)
                var v7 int = 32767
                _ = v6
                _ = v7
            } else {
                v4 -= float64((0.1 / 1.0))
            }
        case 0:
            var v8 float64 = v4
            _ = v8
        }
        v4 -= v4
    } else {
        var v9 uint = 42
        switch 559 {
        case -1:
            v5 = (-(v5))
        case 2147483647:
            v10 := []int{32767, -684, 1}
            _ = v10
        default:
            fmt.Println(v9)
        }
        var v11 bool = true
        _ = v11
        _ = v9
    }
    var v12 uint16 = 830
    v4 -= (1e10 / (v4 * (0.1 + 1.5)))
    v4 = (float64((1.5 + 1e10)) - 0.0)
    var v13 float64 = v4
    v14 := v2[(("a" + "hello") + "hello")]
    var v15 uint16 = v12
    v14 = (127 > 1)
    switch "" {
    case "\n\t":
        var v16 int8 = -128
        _ = v16
    case "a":
        v17 := GenS1{F0: 255}
        _ = v17
    case "foo bar":
        if v14 {
            v18 := []int{574, 0}
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 uint = 1
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 uint8 = 2
                    _ = fi21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        }
    }
    v5 = (v5 * -1)
    v23 := GenS2{F0: -1, F1: -1, F2: (!v14)}
    var v24 bool = false
    var v25 string = "\n\t"
    fmt.Println(v12, v5, v5)
    v26 := GenS3{F0: 1}
    if true {
        fmt.Println(v24)
    }
    v27 := GenS1{F0: 255}
    v28 := v1[301]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
