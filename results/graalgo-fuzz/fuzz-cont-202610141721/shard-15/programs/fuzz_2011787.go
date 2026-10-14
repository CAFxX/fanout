package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint8; F1 uint; F2 int16 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]uint32)
    var v3 uint16 = 2
    var v4 uint16 = (v3 << ((2) & 63))
    var v5 uint = 255
    if false {
        fmt.Println(v1)
    } else {
        var v6 uint16 = v3
        for fi7 := 0; fi7 < 5; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                switch "a" {
                case "foo bar":
                    if (0 != 10) {
                        _ = len(v2)
                        var v9 int64 = -1
                        switch fi7 {
                        case 586:
                            var v10 uint32 = 4294967295
                            _ = v10
                        default:
                            v11 := []int{fi7}
                            _ = v11
                        }
                        _ = v9
                    } else {
                        v6 = v4
                        v12 := GenS1{F0: 2}
                        v5 = (^(v5))
                        _ = v12
                    }
                case "a":
                    if true {
                        v13 := GenS2{F0: 4294967295, F1: (v1 + v1)}
                        var v14 int16 = -138
                        fmt.Println(v5, v4)
                        _ = v13
                        _ = v14
                    }
                }
                v15 := v2[v1]
                _ = fi8
                _ = v15
            }
            for fi16 := 0; fi16 < 15; fi16++ {
                v6 += (v3 ^ (v6 % ((318) | 1)))
                _ = len(v2)
                _ = fi16
            }
            v17 := []int{fi7, -145, fi7, (^(fi7))}
            _ = fi7
            _ = v17
        }
        v4 = (v4 << ((v5) & 63))
        _ = v6
    }
    v18 := GenS1{F0: 65535}
    fmt.Println(v4, v5)
    v19 := GenS1{F0: 208}
    _ = len(v2)
    var v20 float32 = 1.5
    v21 := GenS2{F0: 1, F1: v1}
    if (((227 >= 4294967295) && (202 >= 127)) && false) {
        v22 := GenS2{F0: 0, F1: v1}
        if false {
            v1 = "hello"
            v20 = v20
        } else {
            v5 = 18446744073709551615
            switch 53 {
            case -964:
                _ = len(v2)
            }
            v23 := GenS2{F0: 709, F1: ("" + v1)}
            _ = v23
        }
        _ = v22
    }
    if (255 > 2) {
        v1 = "hello"
        var v24 float64 = (0.0 * float64((float32(-2.25) - 0.1)))
        v5 = 431
        _ = v24
    } else {
        v25 := GenS3{F0: 2, F1: v5, F2: -1}
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 int64 = -9223372036854775808
            v1 = "a"
            _ = fi26
            _ = v27
        }
        v28 := []int{127, -418, 32767, 2147483647}
        _ = v25
        _ = v28
    }
    v2[("foo bar" + (("a" + v1) + (((v1 + v1) + "a") + (v1 + v1))))] = 4294967295
    v4 += v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
