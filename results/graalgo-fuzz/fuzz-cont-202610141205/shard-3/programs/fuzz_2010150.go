package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int32 = 1
    var v2 uint = 2
    var v3 float64 = -2.25
    var v4 uint64 = 676
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: 1, F1: 1, F2: 1}
    v1 = v1
    v7 := GenS2{F0: 15, F1: 834}
    v8 := GenS2{F0: 0, F1: 838}
    if false {
        v9 := v5[4294967295]
        v4 += 1
        _ = v9
    }
    v10 := []int{127, -2147483648, -2147483648, 42}
    var v11 int = 788
    if false {
        v12 := GenS2{F0: (-(v11)), F1: 2147483648}
        _ = v12
    } else {
        if true {
            _ = len(v10)
        }
        switch v11 {
        case 1:
            v1 -= (v1 ^ -417)
        case 2147483647:
            if true {
                var v13 float32 = float32(((((1.5 - (0.0 * v3)) - v3) / float64(float32(v3))) * (v3 + -2.25)))
                _ = v13
            }
        default:
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                v5[14] = "foo bar"
            default:
                var v14 int64 = 1
                _ = v14
            }
        }
    }
    v11 += -1
    v15 := GenS2{F0: (^(v11)), F1: 0}
    v3 = (v3 / v3)
    for fi16 := 0; fi16 < 13; fi16++ {
        if ((v1 != (-(v1))) || false) {
            fmt.Println(v11)
        }
        v17 := GenS3{F0: 1}
        _ = fi16
        _ = v17
    }
    var v18 uint = v2
    var v19 uint16 = 437
    if true {
        for fi20 := 0; fi20 < 20; fi20++ {
            _ = len(v5)
            v2 += v18
            _ = fi20
        }
    }
    var v21 uint16 = v19
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
