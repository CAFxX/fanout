package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }

func main() {
    var v1 uint = 157
    v2 := make(map[int]int)
    var v3 uint8 = 127
    var v4 int8 = 83
    v5 := []int{-397, 464, -2147483648, 0}
    for fi6 := 0; fi6 < 18; fi6++ {
        v1 += 519
        v2[128] = (fi6 + -9223372036854775808)
        _ = fi6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint = v1
        switch fi7 {
        case 42:
            v9 := GenS1{F0: 2147483647}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v10 int8 = -128
    v5 = append(v5, 9223372036854775807)
    switch int(v1) {
    case 2147483647:
        _ = len(v5)
    case -2147483648:
        v10 = v4
    }
    var v11 bool = true
    var v12 int8 = (-(v4))
    v2[-222] = -9223372036854775808
    var v13 int16 = -679
    var v14 uint64 = 2
    var v15 int64 = -883
    var v16 int = -1
    v12 = (v12 + v12)
    v4 = v10
    _ = len(v5)
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            v1 = 65535
            switch fi18 {
            case 127:
                v16 -= v16
            default:
                var v19 float64 = 1.5
                _ = v19
            }
            if (v15 == (-(v15))) {
                var v20 string = "\n\t"
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v10 -= 0
                delete(v2, -2147483648)
                _ = v20
            } else {
                v3 = 230
                v21 := GenS2{F0: v16, F1: v11, F2: 1}
                _ = v21
            }
            _ = fi18
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            v2[fi17] = (fi22 % (((^(fi22))) | 1))
            _ = fi22
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            v12 = v12
            v10 -= (^(v10))
            var v24 bool = (!v11)
            _ = fi23
            _ = v24
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
