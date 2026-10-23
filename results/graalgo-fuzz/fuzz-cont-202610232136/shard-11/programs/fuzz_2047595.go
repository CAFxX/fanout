package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int32 = 1
    var v2 int = 2147483647
    var v3 uint32 = 1
    var v4 float32 = (-2.25 / -2.25)
    v5 := []int{}
    var v6 uint = 255
    if (901 > 2) {
        v3 = 2147483648
        var v7 bool = true
        _ = v7
    } else {
        switch v2 {
        case 1:
            for fi8 := 0; fi8 < 5; fi8++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi8
            }
        case 2147483647:
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: (v6 / (((v6 >> ((v6) & 63))) | 1)), F1: 4294967295}
                v5 = append(v5, (-(v2)))
                var v11 uint16 = 836
                _ = fi9
                _ = v10
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 18; fi12++ {
                fmt.Println(v2)
                _ = fi12
            }
        }
        var v13 int8 = 0
        _ = v13
    }
    fmt.Println(v1, v6, v4)
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS1{F0: (v6 | (v6 | v6)), F1: 2}
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 uint64 = 42
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int64 = 9223372036854775807
            v2 = fi18
            if (true || (true || false)) {
                switch 128 {
                case -9223372036854775808:
                    fmt.Println(fi14, v3)
                case 854:
                    v20 := GenS2{F0: (^(v6)), F1: (v6 / ((v6) | 1)), F2: true}
                    _ = v20
                case 783:
                    v2 = -9223372036854775808
                default:
                    v3 = v3
                }
                var v21 uint = v6
                _ = v21
            } else {
                var v22 uint32 = v3
                v23 := GenS2{F0: v6, F1: v6, F2: false}
                var v24 float32 = 1.5
                _ = v22
                _ = v23
                _ = v24
            }
            v25 := GenS3{F0: 42, F1: -2}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi18
            _ = v19
            _ = v25
        }
        _ = fi14
        _ = v15
    }
    var v26 int8 = 1
    v3 = v3
    v2 += (v2 & v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
