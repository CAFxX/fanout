package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint64 = 1
    var v2 int = 481
    v3 := []uint32{}
    var v4 uint8 = 1
    var v5 float64 = ((1e10 + 0.0) - (-2.25 + 0.0))
    var v6 float32 = 1.5
    if (true && (v4 < (v4 & 24))) {
        var v7 float32 = v6
        _ = v7
    } else {
        switch "\n\t" {
        case "a":
            v4 = (v4 << ((42) & 63))
        default:
            if (v2 > v2) {
                v2 -= (v2 & v2)
            } else {
                for fi8 := 0; fi8 < 15; fi8++ {
                    switch fi8 {
                    case -1:
                        v5 = 0.0
                    }
                    var v9 int32 = 1
                    v10 := GenS1{F0: v9}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
        }
    }
    switch v2 {
    case 32768:
        v6 = v6
    }
    v11 := GenS2{F0: uint16(v1), F1: 1}
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v12 := GenS1{F0: -2147483648}
    if (!true) {
        if true {
            var v13 uint64 = 722
            _ = v13
        }
    } else {
        var v14 int8 = -128
        _ = v14
    }
    switch 1 {
    case 400:
        if true {
            var v15 int32 = 740
            _ = v15
        } else {
            v6 = 0.1
            for fi16 := 0; fi16 < 8; fi16++ {
                if false {
                    v2 = (fi16 / (((v2 >> ((636) & 63))) | 1))
                    var v17 int64 = -9223372036854775808
                    _ = v17
                }
                _ = fi16
            }
            v18 := GenS1{F0: 2147483647}
            _ = v18
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
