package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 string; F1 uint16 }

func main() {
    var v1 int16 = -15
    var v2 int8 = 100
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 82
    var v5 int8 = 1
    v6 := []int{0, 32767, 127}
    if (1.5 < 1e10) {
        v7 := GenS1{F0: (v5 <= v2), F1: 522}
        v3 = (v3 | int64(v2))
        _ = v7
    }
    v4 -= (v4 - v4)
    v1 += 501
    for fi8 := 0; fi8 < 19; fi8++ {
        switch fi8 {
        case 42:
            if len(v6) > 0 {
                v6[0] = -2147483648
            }
        case 1:
            v9 := GenS1{F0: (!(255 != 260)), F1: fi8}
            _ = v9
        case 816:
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 uint16 = 2
                _ = fi10
                _ = v11
            }
        default:
            var v12 int16 = v1
            _ = v12
        }
        _ = fi8
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v5 = v2
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    v4 += (^(v4))
    switch 42 {
    case 127:
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(v1)
            v16 := GenS1{F0: (false || (false && (42 <= -858))), F1: fi15}
            var v17 int8 = v5
            _ = fi15
            _ = v16
            _ = v17
        }
    case 158:
        if false {
            if len(v6) > 0 {
                v6[0] = 42
            }
            var v18 uint8 = (^(v4))
            v2 = v2
            _ = v18
        }
    }
    var v19 bool = ((2 <= 516) || ((v4 ^ (v4 % ((v4) | 1))) != v4))
    fmt.Println(v19, v4)
    v4 = v4
    var v20 uint64 = 656
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 uint64 = (v20 + 0)
        if v19 {
            for fi23 := 0; fi23 < 8; fi23++ {
                v6 = append(v6, (fi23 * 185))
                v24 := GenS2{F0: "hello", F1: 0}
                v25 := GenS1{F0: v19, F1: (-(fi21))}
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v26 uint = 2
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            v27 := GenS2{F0: "a", F1: 1}
            _ = v27
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
