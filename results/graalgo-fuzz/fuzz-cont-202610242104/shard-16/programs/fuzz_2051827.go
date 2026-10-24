package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 uint16) uint {
    var v4 int32 = 790
    if (1 >= 209) {
        var v5 int32 = int32(p0)
        for fi6 := 0; fi6 < 14; fi6++ {
            v4 += int32(fi6)
            _ = fi6
        }
        _ = v5
    } else {
        switch 1 {
        case -9223372036854775808:
            switch "hello" {
            case "":
                v4 += (v4 ^ v4)
            case "a":
                p0 = (^(p0))
            case "hello":
                var v7 int = 757
                _ = v7
            default:
                switch "foo bar" {
                case "hello":
                    p0 -= p0
                case "foo bar":
                    fmt.Println(p0, p0)
                }
            }
        case -1:
            for fi8 := 0; fi8 < 2; fi8++ {
                fmt.Println(v4, p0, fi8)
                if (0 == 203) {
                    v4 = (v4 << ((65535) & 63))
                    v9 := GenS1{F0: fi8, F1: 255}
                    for fi10 := 0; fi10 < 6; fi10++ {
                        if false {
                            var v11 uint = 878
                            _ = v11
                        }
                        if (!(208 >= 127)) {
                            var v12 int32 = (v4 / (((v4 * (v4 << ((4294967295) & 63)))) | 1))
                            v13 := GenS1{F0: 2147483647, F1: 1}
                            p0 -= (p0 | 919)
                            _ = v12
                            _ = v13
                        }
                        v4 += (^(v4))
                        _ = fi10
                    }
                    _ = v9
                } else {
                    p0 += 255
                }
                v14 := []int{0, (fi8 / (((^(fi8))) | 1))}
                _ = fi8
                _ = v14
            }
        case 32768:
            switch -857 {
            case -9223372036854775808:
                fmt.Println(v4)
            case -1:
                v4 += 2147483647
            case 42:
                v4 -= (v4 + v4)
            }
        }
        fmt.Println(p0)
    }
    v4 = v4
    for fi15 := 0; fi15 < 9; fi15++ {
        v4 -= v4
        _ = fi15
    }
    _ = p0
    _ = v4
    return 255
}
func genFn1(p0 uint32, p1 int) uint64 {
    v16 := []int{}
    for fi17 := 0; fi17 < 12; fi17++ {
        p0 += (p0 << ((255) & 63))
        var v18 int8 = -61
        _ = fi17
        _ = v18
    }
    switch "a" {
    case "":
        fmt.Println(p0)
    case "foo bar":
        fmt.Println(p1, p1)
    case "\n\t":
        p0 = p0
    default:
        if len(v16) > 0 {
            v16[0] = (-(p1))
        }
    }
    p1 = 42
    _ = p0
    _ = p1
    _ = v16
    return 613
}

func main() {
    var v1 float32 = float32(float64(0.0))
    var v2 uint8 = 128
    var v3 float32 = (-2.25 / ((v1 + (0.0 * float32((1.5 + 3.14159)))) - 1.5))
    v2 = (v2 >> ((4294967295) & 63))
    switch 127 {
    case 42:
        v19 := GenS2{F0: false, F1: 42}
        _ = v19
    default:
        var v20 uint8 = (v2 - v2)
        _ = v20
    }
    var v21 int64 = 79
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 uint = 42
        var v24 uint64 = 638
        if (!false) {
            v25 := GenS1{F0: (fi22 / ((fi22) | 1)), F1: v2}
            v24 += (v24 & v24)
            switch 89 {
            case 209:
                var v26 int32 = -581
                _ = v26
            case 32768:
                fmt.Println(v2)
            case 9223372036854775807:
                if true {
                    fmt.Println(v21, fi22, v1)
                    v2 -= (^(v2))
                    v1 = v1
                }
            }
            _ = v25
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    v27 := GenS3{F0: 0, F1: (921 != v21)}
    var v28 bool = true
    var v29 int32 = -1
    var v30 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
