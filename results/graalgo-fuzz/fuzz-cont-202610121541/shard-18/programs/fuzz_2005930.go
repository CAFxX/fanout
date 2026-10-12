package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }

func main() {
    var v1 int8 = 42
    v2 := []int{-395, 2147483647, 32768}
    v3 := make(map[string]uint32)
    switch "" {
    case "\n\t":
        switch 42 {
        case 32767:
            v4 := v3["\n\t"]
            _ = v4
        case 4294967295:
            _ = len(v2)
        }
    default:
        v5 := GenS1{F0: (v1 ^ (v1 & (-(v1)))), F1: 255}
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            if true {
                v1 += (v1 % (((v1 ^ 42)) | 1))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch "a" {
                case "\n\t":
                    v7 := GenS1{F0: v1, F1: 2}
                    _ = v7
                case "foo bar":
                    var v8 int = fi6
                    _ = v8
                }
            }
        }
        v9 := GenS1{F0: (^(v1)), F1: 350}
        _ = fi6
        _ = v9
    }
    var v10 uint8 = 83
    fmt.Println(v10)
    v1 = (v1 ^ (v1 | (v1 - (v1 % (((v1 & (v1 | v1))) | 1)))))
    for fi11 := 0; fi11 < 12; fi11++ {
        v1 += (v1 % ((v1) | 1))
        v1 += 42
        _ = fi11
    }
    v2 = append(v2, 32767)
    if len(v2) > 0 {
        v2[0] = 32768
    }
    var v12 uint64 = 153
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 9; fi14++ {
        if (fi14 != (fi14 * fi14)) {
            fmt.Println(v10, v10, v10)
        }
        _ = fi14
    }
    switch 32767 {
    case 2147483647:
        v10 -= v10
    case 42:
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                v1 += v1
                _ = fi16
            }
            var v17 int8 = v1
            var v18 int = (fi15 ^ (fi15 / ((fi15) | 1)))
            _ = fi15
            _ = v17
            _ = v18
        }
    case 0:
        v19 := GenS1{F0: 0, F1: 255}
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v2 = append(v2, fi20)
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
