package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int; F2 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 255
    var v3 int8 = -125
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 65535}
    v3 += (-(v3))
    var v6 int16 = 514
    var v7 uint8 = 150
    if ((false && (true && true)) && ((v7 - 115) < 127)) {
        var v8 int16 = (v6 % ((v6) | 1))
        _ = v8
    } else {
        if false {
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 float64 = 1.5
                switch fi9 {
                case 9223372036854775807:
                    v11 := []int{}
                    _ = v11
                case 2147483647:
                    v3 = (v3 << ((17) & 63))
                case 0:
                    _ = len(v1)
                }
                var v12 bool = (65535 != 2)
                _ = fi9
                _ = v10
                _ = v12
            }
            var v13 uint8 = (^(v7))
            _ = v13
        }
        if (0 < 1) {
            v3 = int8(v7)
            var v14 bool = true
            _ = v14
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                for fi16 := 0; fi16 < 12; fi16++ {
                    v3 = (v3 << ((1) & 63))
                    _ = fi16
                }
                var v17 int = (fi15 * (-(fi15)))
                _ = fi15
                _ = v17
            }
        }
        v18 := GenS1{F0: 410}
        _ = v18
    }
    var v19 uint = 4294967295
    switch 333 {
    case 626:
        v6 = (v6 / ((int16(v3)) | 1))
    case -554:
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 float64 = 1.5
            var v22 int64 = -329
            _ = fi20
            _ = v21
            _ = v22
        }
    case -795:
        v23 := GenS2{F0: true, F1: 113, F2: 42}
        _ = v23
    default:
        fmt.Println(v6, v6, v3)
    }
    v3 -= 1
    fmt.Println(v19, v19, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
