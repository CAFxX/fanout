package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint16; F1 int; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 632
    var v3 uint8 = 0
    var v4 uint = 860
    var v5 float32 = 0.1
    if false {
        fmt.Println(v3, v4)
    }
    v2 += v2
    v4 -= v4
    v6 := GenS1{F0: -9223372036854775808}
    var v7 uint = 4294967295
    if true {
        if (1.5 == 1e10) {
            v3 -= (v3 - (v3 % ((v3) | 1)))
            v3 = (v3 ^ 5)
        }
        v7 = v7
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS2{F0: "\n\t", F1: 255}
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint = v4
                var v12 int8 = 37
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS3{F0: 2, F1: fi8, F2: 255}
                _ = fi13
                _ = v14
            }
            _ = fi8
            _ = v9
        }
    } else {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint8 = (^(v3))
            v16 = v16
            v17 := []int{-699, fi15, (fi15 + int(v7)), -341}
            _ = fi15
            _ = v16
            _ = v17
        }
        v1["foo bar"] = 771
    }
    v18 := GenS4{F0: 42}
    v19 := v1[""]
    var v20 float64 = float64(1e10)
    v2 += uint16(v7)
    var v21 uint = (v7 * v4)
    fmt.Println(v21)
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 int32 = 317
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 1}
    v25 := GenS5{F0: 127, F1: (("foo bar" + "hello") + "\n\t")}
    if (1 > 1) {
        for fi26 := 0; fi26 < 9; fi26++ {
            switch fi26 {
            case 9223372036854775807:
                fmt.Println(v3, v19)
            case 32768:
                if (false && (0 > -128)) {
                    v4 += (v21 - v4)
                    v27 := []int{358}
                    _ = v27
                } else {
                    v3 = v3
                    v21 = v7
                }
            default:
                v19 -= v19
            }
            _ = fi26
        }
        v28 := []int{0}
        _ = v28
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
