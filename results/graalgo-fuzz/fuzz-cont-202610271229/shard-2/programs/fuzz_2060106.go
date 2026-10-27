package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint = 589
    v2 := make(map[int]int)
    var v3 int16 = 1000
    v4 := make(map[uint32]bool)
    var v5 int64 = -1
    for fi6 := 0; fi6 < 13; fi6++ {
        if ((3.14159 / (-2.25 + (float64((1e10 / float32(1.5))) - 1.5))) == -2.25) {
            v7 := v4[918]
            _ = v7
        } else {
            v3 = v3
            v8 := []int{fi6, fi6}
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: 9223372036854775807, F1: 609}
    v3 = v3
    v10 := GenS1{F0: 32767, F1: 1}
    v11 := []int{}
    v12 := GenS1{F0: 481, F1: 619}
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int8 = -1
        if (0.1 >= -2.25) {
            if false {
                fmt.Println(v1, v14, v5)
                v5 = v5
                v15 := GenS1{F0: (fi13 / ((fi13) | 1)), F1: 1}
                _ = v15
            }
            var v16 uint = (v1 + (^(v1)))
            _ = v16
        } else {
            v1 = v1
            v5 = (v5 - v5)
            v17 := GenS2{F0: (v5 / ((v5) | 1)), F1: 1}
            _ = v17
        }
        v3 -= (-(v3))
        _ = fi13
        _ = v14
    }
    switch 385 {
    case 42:
        switch "" {
        case "\n\t":
            var v18 uint16 = 255
            _ = v18
        case "foo bar":
            if false {
                switch -127 {
                case 42:
                    fmt.Println(v5)
                default:
                    v11 = append(v11, -1)
                }
                if (((true && true) && (false || true)) && true) {
                    var v19 uint32 = 42
                    v3 -= -282
                    _ = v19
                }
                var v20 uint = v1
                _ = v20
            }
        default:
            v21 := GenS3{F0: 127, F1: (v5 << ((v1) & 63)), F2: (!true)}
            _ = v21
        }
    default:
        var v22 float32 = 0.0
        _ = v22
    }
    v23 := GenS4{F0: 2, F1: false}
    v1 -= (v1 - (v1 / (((^(v1))) | 1)))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
