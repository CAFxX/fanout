package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 string }

func main() {
    var v1 uint8 = 233
    var v2 string = "a"
    var v3 int = -9223372036854775808
    var v4 int64 = 1
    v5 := make(map[uint32]int)
    v6 := make(map[int]int)
    v2 = v2
    _ = len(v5)
    var v7 uint = 42
    switch 685 {
    case 9223372036854775807:
        var v8 uint64 = 559
        _ = v8
    case 475:
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint = (v7 ^ 4294967295)
            _ = fi9
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 8; fi11++ {
            if true {
                var v12 string = ""
                _ = v12
            } else {
                v13 := v5[2]
                _ = v13
            }
            var v14 float64 = ((3.14159 - float64(0.1)) - 1e10)
            _ = fi11
            _ = v14
        }
    }
    v15 := GenS1{F0: 127}
    fmt.Println(v4, v1, v2)
    var v16 int16 = -879
    v3 += 0
    if false {
        v17 := GenS1{F0: int8(v7)}
        _ = v17
    } else {
        if (-128 >= 127) {
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := []int{(fi18 & v3), (-(fi18)), fi18, (v3 * fi18)}
                for fi20 := 0; fi20 < 14; fi20++ {
                    v21 := GenS2{F0: 1, F1: v7}
                    var v22 float64 = -2.25
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 int = v3
                _ = fi18
                _ = v19
                _ = v23
            }
            fmt.Println(v4)
            if (!false) {
                if (v16 != (v16 | (v16 >> (((v7 % (((v7 | 18446744073709551615)) | 1))) & 63)))) {
                    var v24 float64 = -2.25
                    v25 := []int{(v3 >> ((v7) & 63)), 865, (-(v3)), (^(v3))}
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 17; fi26++ {
                    v27 := GenS1{F0: 127}
                    _ = fi26
                    _ = v27
                }
            }
        } else {
            var v28 bool = ((false && false) && false)
            switch v3 {
            case -450:
                v29 := []int{}
                _ = v29
            case 32768:
                v30 := GenS3{F0: ""}
                _ = v30
            default:
                var v31 int8 = 42
                _ = v31
            }
            _ = v28
        }
        switch v3 {
        case 0:
            for fi32 := 0; fi32 < 9; fi32++ {
                v33 := []int{-1}
                _ = fi32
                _ = v33
            }
        }
        v16 += (v16 | (v16 + (v16 & v16)))
    }
    var v34 uint = (^(v7))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
