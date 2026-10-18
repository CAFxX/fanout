package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 float32 = (float32(0.0) + 0.0)
    v5 := GenS1{F0: 1, F1: -1}
    var v6 float64 = -2.25
    if (3.14159 > (v4 + 0.0)) {
        switch ("foo bar" + "a") {
        case "foo bar":
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 int8 = -69
                switch "a" {
                case "foo bar":
                    v4 += 0.1
                case "\n\t":
                    if (!(440 <= 2147483648)) {
                        v8 = v8
                    }
                default:
                    var v9 uint16 = 42
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
        case "a":
            var v10 bool = (0 < -1)
            _ = v10
        }
        v6 += v6
        v11 := []int{}
        _ = v11
    }
    var v12 float64 = (v6 * (float64(v4) / 1e10))
    if true {
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := []int{(fi13 >> ((65535) & 63)), (fi13 + (fi13 / ((fi13) | 1)))}
            fmt.Println(fi13)
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 1, F1: 1}
        v12 -= (v12 / (float64(v4) - (((v12 - (((v12 - ((v12 - 0.0) / 0.1)) / float64((v4 / ((v4 + v4) * 0.1)))) * ((float64(0.0) * ((v6 / v12) + -2.25)) / -2.25))) - v12) * v12)))
        _ = v15
    }
    v4 -= (((v4 - v4) / (0.1 / float32(v12))) / (v4 / 3.14159))
    for fi16 := 0; fi16 < 16; fi16++ {
        if true {
            v6 = (-2.25 + v12)
            var v17 int = fi16
            _ = v17
        }
        fmt.Println(fi16)
        _ = fi16
    }
    var v18 uint = 419
    v19 := GenS2{F0: 42, F1: v18}
    if (1000 == 1) {
        var v20 uint64 = 4294967295
        _ = v20
    }
    var v21 string = ""
    var v22 uint64 = 651
    var v23 float32 = v4
    v24 := []int{}
    v25 := GenS1{F0: v18, F1: int64(v18)}
    if (-2.25 >= float64(float32(float64((float32(v12) + v4))))) {
        fmt.Println(v18)
        v23 += v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
