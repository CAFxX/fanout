package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 string }

func main() {
    v1 := []int{}
    var v2 int8 = -77
    var v3 int64 = -1
    v2 += 100
    var v4 uint8 = 127
    fmt.Println(v3, v3)
    var v5 string = "\n\t"
    if true {
        v6 := GenS1{F0: 2, F1: (v2 >> ((0) & 63))}
        if true {
            v3 = (v3 - 0)
            _ = len(v1)
        } else {
            var v7 uint32 = 137
            switch 265 {
            case 32768:
                switch "foo bar" {
                case "":
                    var v8 float64 = 0.1
                    _ = v8
                case "a":
                    v2 = v2
                case "foo bar":
                    _ = len(v1)
                default:
                    v9 := GenS2{F0: false, F1: 65535}
                    _ = v9
                }
            case 149:
                v7 = 665
            default:
                v10 := GenS1{F0: 255, F1: v2}
                _ = v10
            }
            _ = v7
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            if len(v1) > 0 {
                v1[0] = (fi11 + (fi11 % ((fi11) | 1)))
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                v3 -= v3
                if (!true) {
                    v13 := GenS3{F0: ""}
                    _ = v13
                } else {
                    fmt.Println(v5, v5, v4)
                }
                _ = fi12
            }
            _ = len(v1)
            _ = fi11
        }
        _ = v6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5)
    fmt.Println(v4, v5, v2)
    if len(v1) > 0 {
        v1[0] = 424
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        if (!((fi14 + fi14) < 9223372036854775807)) {
            if false {
                v4 += (v4 / ((v4) | 1))
            }
            var v15 int16 = 1000
            _ = v15
        }
        _ = fi14
    }
    fmt.Println(v3, v3, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
