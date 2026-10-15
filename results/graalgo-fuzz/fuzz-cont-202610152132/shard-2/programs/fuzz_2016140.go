package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int8 = -76
    var v2 uint16 = 646
    v3 := make(map[int]int)
    v1 = -128
    var v4 int8 = (v1 / ((114) | 1))
    v3[2147483647] = 30
    var v5 int64 = 680
    v6 := GenS1{F0: (-(v5))}
    if false {
        if true {
            fmt.Println(v5, v1, v1)
            switch 9223372036854775807 {
            case 0:
                v7 := []int{32767, 950, 32767, 740}
                _ = v7
            case -748:
                var v8 uint8 = 196
                _ = v8
            default:
                switch "" {
                case "":
                    if (uint8(v2) < 42) {
                        v1 = -82
                        var v9 int = 0
                        _ = v9
                    } else {
                        fmt.Println(v4, v4, v4)
                    }
                default:
                    v10 := GenS1{F0: (v5 & v5)}
                    _ = v10
                }
            }
            var v11 bool = true
            _ = v11
        } else {
            fmt.Println(v2)
            switch "\n\t" {
            case "foo bar":
                v12 := GenS1{F0: -1}
                _ = v12
            case "hello":
                fmt.Println(v5)
            }
        }
        var v13 uint16 = v2
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v4 = v4
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 int16 = -1
            _ = fi15
            _ = v16
        }
        v5 = v5
        _ = fi14
    }
    v2 += 32767
    v17 := GenS1{F0: -678}
    var v18 float64 = (-2.25 * 0.0)
    var v19 float32 = 1.5
    var v20 int64 = (v5 - -1)
    v21 := v3[-1]
    switch (v21 * v21) {
    case 859:
        _ = len(v3)
    case 127:
        fmt.Println(v20, v21)
    default:
        if (!false) {
            v5 = v20
            v22 := []int{}
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
