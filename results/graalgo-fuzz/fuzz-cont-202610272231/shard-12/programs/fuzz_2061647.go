package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 bool; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint32, p1 int64, p2 int64) int8 {
    fmt.Println(p0, p1, p1)
    var v4 uint64 = 417
    var v5 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -128
}

func main() {
    var v1 uint8 = 156
    var v2 int = -1
    v3 := make(map[uint32]bool)
    var v6 uint64 = 326
    v7 := []int{v2, (v2 & (v2 & (v2 / (((-(v2))) | 1)))), v2}
    v1 -= 140
    v7 = append(v7, int(v1))
    v6 += 661
    if ((1 <= 2147483648) && true) {
        var v8 int = (v2 * 128)
        switch v8 {
        case 956:
            switch "\n\t" {
            case "\n\t":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case "hello":
                v9 := GenS1{F0: 1, F1: 42, F2: 32767}
                _ = v9
            case "a":
                v2 += int(v1)
            default:
                var v10 uint8 = v1
                _ = v10
            }
        case 317:
            var v11 uint = 255
            _ = v11
        default:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = v8
    } else {
        v12 := GenS2{F0: ""}
        var v13 int8 = 0
        if (v1 < (v1 % ((v1) | 1))) {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 uint8 = (v1 << ((18446744073709551615) & 63))
                _ = fi14
                _ = v15
            }
            v16 := GenS3{F0: 65535, F1: ((false || false) || false), F2: 255}
            if false {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := GenS4{F0: -1, F1: (v1 & 0)}
                    if (true && (!(!(!(42 <= 255))))) {
                        delete(v3, 42)
                        switch "hello" {
                        case "a":
                            if (32767 == -668) {
                                var v19 uint8 = 104
                                _ = v19
                            } else {
                                v2 -= (^(fi17))
                            }
                        case "":
                            if ((true || (-735 < -1)) || (!(42 <= 594))) {
                                v6 -= v6
                            }
                        default:
                            var v20 uint64 = (^(v6))
                            _ = v20
                        }
                    }
                    var v21 uint32 = 747
                    _ = fi17
                    _ = v18
                    _ = v21
                }
            } else {
                if len(v7) > 0 {
                    v7[0] = 32768
                }
            }
            _ = v16
        }
        _ = v12
        _ = v13
    }
    v2 = v2
    v1 += 42
    if len(v7) > 0 {
        v7[0] = 1
    }
    v22 := v3[1]
    v7 = append(v7, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v6
    _ = v7
}
