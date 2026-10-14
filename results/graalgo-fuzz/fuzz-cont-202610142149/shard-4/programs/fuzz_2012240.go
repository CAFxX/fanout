package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    v3 := make(map[int]uint32)
    var v4 int64 = -870
    var v5 int = 657
    v4 -= (v4 | (v4 | (v4 << ((2) & 63))))
    if (0 > 2) {
        v2 = (v2 << ((0) & 63))
        v2 -= (v2 * (v2 % (((v2 - (v2 + v2))) | 1)))
        v6 := GenS1{F0: v4}
        _ = v6
    } else {
        if (!true) {
            v7 := GenS1{F0: v4}
            fmt.Println(v2, v2)
            v2 += v2
            _ = v7
        } else {
            for fi8 := 0; fi8 < 5; fi8++ {
                if false {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 float64 = -2.25
                        var v11 int8 = 42
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 int = fi12
                        v2 += (^(v2))
                        _ = fi12
                        _ = v13
                    }
                }
                v14 := GenS2{F0: -1, F1: 128}
                var v15 uint = 616
                _ = fi8
                _ = v14
                _ = v15
            }
            v5 -= (v5 / (((^(v5))) | 1))
            var v16 int64 = -1
            _ = v16
        }
    }
    delete(v1, v5)
    if (507 != 8) {
        fmt.Println(v2, v4)
        v17 := GenS2{F0: v5, F1: 161}
        v4 -= (^(v4))
        _ = v17
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v2)
    case "":
        var v18 uint = 122
        _ = v18
    }
    var v19 string = ""
    var v20 uint16 = 1
    var v21 int = -9223372036854775808
    v1[(v21 & v5)] = true
    for fi22 := 0; fi22 < 2; fi22++ {
        fmt.Println(v19)
        v21 = v21
        switch v21 {
        case -1:
            for fi23 := 0; fi23 < 12; fi23++ {
                if (false && (true && false)) {
                    fmt.Println(v21, fi22, v21)
                } else {
                    var v24 uint8 = 127
                    fmt.Println(v24, v21, fi22)
                    _ = v24
                }
                var v25 bool = ((0.0 != (0.0 + 0.1)) || (v2 != 911))
                _ = fi23
                _ = v25
            }
        default:
            if (!false) {
                delete(v3, (^(v21)))
            }
        }
        _ = fi22
    }
    v20 = (v20 + (v20 ^ (v20 ^ 255)))
    fmt.Println(v21)
    v26 := GenS3{F0: 0, F1: 174}
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
