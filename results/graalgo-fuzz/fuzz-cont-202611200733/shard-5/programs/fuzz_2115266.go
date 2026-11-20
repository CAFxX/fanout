package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    var v1 int16 = 166
    var v2 int8 = -128
    v3 := make(map[uint32]uint32)
    var v4 int8 = (v2 - (v2 + v2))
    v5 := make(map[uint32]int)
    v6 := make(map[int]uint32)
    v4 -= (v2 >> ((1) & 63))
    var v7 string = "\n\t"
    fmt.Println(v2, v4)
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v1)
        _ = fi8
    }
    v9 := GenS1{F0: 666}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int = fi10
        var v12 int = fi10
        v4 += v2
        _ = fi10
        _ = v11
        _ = v12
    }
    v4 = (v4 & -1)
    v7 = "a"
    var v13 int64 = int64(v1)
    switch "\n\t" {
    case "a":
        var v14 uint16 = 255
        _ = v14
    case "foo bar":
        var v15 uint = 65535
        _ = v15
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            if (((v13 * (v13 ^ v13)) > 1) && true) {
                v2 -= (v2 / (((-(v4))) | 1))
                v1 -= (v1 >> ((255) & 63))
            }
            var v17 uint8 = 127
            for fi18 := 0; fi18 < 12; fi18++ {
                if (false && false) {
                    if (("hello" + (v7 + ((v7 + "hello") + "a"))) > "\n\t") {
                        v19 := GenS2{F0: 135, F1: v2}
                        _ = v19
                    }
                } else {
                    v20 := []int{fi18}
                    v3[4294967295] = 42
                    if true {
                        v21 := GenS2{F0: 2, F1: int8(v17)}
                        v20 = append(v20, (fi16 << ((255) & 63)))
                        switch (fi18 * fi18) {
                        case 836:
                            v6[fi16] = 0
                        case 128:
                            v4 = (v2 >> ((42) & 63))
                        default:
                            v17 += (v17 | (v17 * v17))
                        }
                        _ = v21
                    } else {
                        fmt.Println(fi16, v4)
                        switch fi16 {
                        case 127:
                            fmt.Println(v4)
                        case 0:
                            v17 -= v17
                        default:
                            v1 = v1
                        }
                        var v22 uint32 = 949
                        _ = v22
                    }
                    _ = v20
                }
                fmt.Println(v7)
                var v23 uint16 = 857
                _ = fi18
                _ = v23
            }
            _ = fi16
            _ = v17
        }
    }
    if false {
        if (!(137 >= 4294967295)) {
            var v24 int16 = v1
            v25 := GenS3{F0: -2147483648, F1: 255}
            var v26 uint64 = 0
            _ = v24
            _ = v25
            _ = v26
        } else {
            var v27 float64 = ((0.0 * (-2.25 + 0.0)) + 1.5)
            var v28 int16 = -735
            var v29 int = 9223372036854775807
            _ = v27
            _ = v28
            _ = v29
        }
        var v30 float64 = (float64(-2.25) + float64(((((-2.25 - 1e10) - (-2.25 - -2.25)) * 3.14159) * (float32(3.14159) - -2.25))))
        var v31 bool = (!true)
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
