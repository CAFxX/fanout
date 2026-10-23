package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 uint, p2 uint) int64 {
    var v6 int32 = 0
    v7 := []int{}
    var v8 uint = 874
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 820
}

func main() {
    var v1 float64 = 1e10
    var v2 int32 = 2147483647
    v3 := make(map[string]bool)
    var v4 int32 = v2
    var v5 uint = 240
    v1 = v1
    if ((-2.25 > (1.5 - v1)) || (genFn0(215, 1, (v5 - v5)) > 42)) {
        v9 := []int{-1}
        _ = v9
    } else {
        v5 += (v5 / ((255) | 1))
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: -103, F1: 65535, F2: 0}
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v2, v4, v4)
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := []int{(fi12 | (fi12 * 1))}
        _ = len(v3)
        v14 := GenS2{F0: 2}
        _ = fi12
        _ = v13
        _ = v14
    }
    if true {
        fmt.Println(v2)
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                v5 = v5
                switch "\n\t" {
                case "\n\t":
                    v2 -= 2147483647
                case "foo bar":
                    v5 -= 0
                case "":
                    v4 = v2
                }
                var v17 float64 = (0.1 / (1.5 + v1))
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                v19 := GenS1{F0: 42, F1: 418, F2: 4294967295}
                switch fi15 {
                case 127:
                    v4 = 772
                case -1:
                    v1 -= v1
                default:
                    v20 := GenS1{F0: 375, F1: 296, F2: 489}
                    _ = v20
                }
                v21 := v3[""]
                _ = len(v3)
                _ = fi18
                _ = v19
                _ = v21
            }
            _ = fi15
        }
    }
    var v22 int32 = v4
    v5 += 981
    v22 -= (v22 >> ((65535) & 63))
    v4 = 2147483647
    var v23 uint32 = 0
    v24 := GenS1{F0: 1000, F1: v23, F2: v23}
    v3["foo bar"] = (0 > 385)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
