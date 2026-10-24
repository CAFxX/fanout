package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint8 }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 uint64 = 625
    v5 := GenS1{F0: 420, F1: 2147483647, F2: 1}
    var v6 uint = 2
    if false {
        var v7 uint32 = 577
        _ = v7
    }
    v8 := GenS1{F0: 610, F1: -2147483648, F2: 103}
    v6 += (v6 % ((65535) | 1))
    var v9 float64 = (v1 / ((1e10 * v1) + (((v1 / (-2.25 + ((v1 * v1) / -2.25))) - v1) + v1)))
    switch -393 {
    case 127:
        v4 -= (v4 >> (((v6 ^ v6)) & 63))
    case 669:
        var v10 string = ("" + "\n\t")
        _ = v10
    case -1:
        if (216 <= 780) {
            v11 := GenS1{F0: 0, F1: -863, F2: 2}
            _ = v11
        } else {
            if (42 != 74) {
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := []int{fi12, fi12}
                    _ = fi12
                    _ = v13
                }
            } else {
                v14 := []int{}
                fmt.Println(v6, v6)
                _ = v14
            }
            var v15 int64 = 644
            _ = v15
        }
    default:
        fmt.Println(v9, v9)
    }
    v16 := GenS2{F0: 821, F1: 0}
    v6 -= 65535
    v1 += (0.1 / -2.25)
    var v17 int32 = -1
    v18 := v2[2]
    v19 := v3["foo bar"]
    _ = len(v2)
    var v20 uint64 = (v4 ^ v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
